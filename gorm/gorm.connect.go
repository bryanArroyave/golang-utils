package gormconnect

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bryanArroyave/golang-utils/gorm/dtos"
	"github.com/bryanArroyave/golang-utils/gorm/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBManager struct {
	connectionData *dtos.GormConnectionDTO
	db             *gorm.DB
	mu             sync.Mutex
	reconnecting   bool
}

var (
	managerInstance *DBManager
	managerOnce     sync.Once
)

const (
	defaultMaxRetries      = 5
	maxRetryDelay          = 30 * time.Second
	baseRetryDelay         = time.Second
	defaultMaxIdleConns    = 10
	defaultMaxOpenConns    = 100
	defaultConnMaxLifetime = time.Hour
	defaultConnMaxIdleTime = 10 * time.Minute
)

func NewDBManager(connectionData *dtos.GormConnectionDTO) ports.IDBManager {
	managerOnce.Do(func() {
		managerInstance = &DBManager{connectionData: connectionData}

		if err := managerInstance.validate(); err != nil {
			log.Fatalf("❌ Failed to validate connection data: %v", err)
		}

		if err := managerInstance.connect(); err != nil {
			log.Fatalf("❌ Failed to connect to the database: %v", err)
		}
	})

	return managerInstance
}

func (dm *DBManager) validate() error {
	if dm.connectionData == nil {
		return fmt.Errorf("connectionData is required")
	}
	if dm.connectionData.URI == "" {
		return fmt.Errorf("URI is required")
	}
	if dm.connectionData.Dialector == nil {
		return fmt.Errorf("Dialector is required")
	}
	if dm.connectionData.Env == "" {
		dm.connectionData.Env = "LOCAL"
	}
	if dm.connectionData.MaxRetries <= 0 {
		dm.connectionData.MaxRetries = defaultMaxRetries
	}
	return nil
}

func (dm *DBManager) connect() error {
	dm.mu.Lock()
	if dm.reconnecting {
		dm.mu.Unlock()
		log.Println("connect | Already attempting to reconnect, skipping...")
		return fmt.Errorf("already attempting to reconnect")
	}
	dm.reconnecting = true
	dm.mu.Unlock()

	defer func() {
		dm.mu.Lock()
		dm.reconnecting = false
		dm.mu.Unlock()
	}()

	var lastErr error
	for i := 0; i < dm.connectionData.MaxRetries; i++ {
		dm.db, lastErr = gorm.Open(*dm.connectionData.Dialector, &gorm.Config{
			Logger: selectLevelLogger(dm.connectionData.Env),
		})
		if lastErr == nil {
			sqlDB, err := dm.db.DB()
			if err != nil {
				return fmt.Errorf("connect | Failed to get SQL DB instance: %w", err)
			}

			sqlDB.SetMaxIdleConns(defaultMaxIdleConns)
			sqlDB.SetMaxOpenConns(defaultMaxOpenConns)
			sqlDB.SetConnMaxLifetime(defaultConnMaxLifetime)
			sqlDB.SetConnMaxIdleTime(defaultConnMaxIdleTime)

			log.Printf("✅ Database connected successfully")
			return nil
		}

		delay := baseRetryDelay * (1 << i)
		if delay > maxRetryDelay {
			delay = maxRetryDelay
		}

		log.Printf("❌ Failed to connect to the database, retrying in %v... (%d/%d)", delay, i+1, dm.connectionData.MaxRetries)
		time.Sleep(delay)
	}

	return fmt.Errorf("failed to connect to the database after %d attempts: %w", dm.connectionData.MaxRetries, lastErr)
}

func (dm *DBManager) GetConnection() (*gorm.DB, error) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if dm.db == nil {
		return nil, fmt.Errorf("database connection is not established")
	}

	sqlDB, errDB := dm.db.DB()
	if errDB != nil {
		return nil, errDB
	}

	errPing := sqlDB.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("database connection is not active: %w", errPing)
	}

	return dm.db, nil
}

func (dm *DBManager) EnsureConnection() error {
	_, err := dm.GetConnection()
	if err != nil {
		log.Printf("EnsureConnection | Connection lost: %v. Attempting to reconnect...", err)
		if reconnectErr := dm.connect(); reconnectErr != nil {
			return fmt.Errorf("EnsureConnection | Failed to reconnect: %w", reconnectErr)
		}
	}
	return nil
}

func (dm *DBManager) Close() error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if dm.db == nil {
		log.Println("Close | No active database connection to close.")
		return nil
	}

	sqlDB, errDB := dm.db.DB()
	if errDB != nil {
		return errDB
	}

	errClose := sqlDB.Close()
	if errClose != nil {
		return errClose
	}

	dm.db = nil
	log.Println("✅ Database connection closed successfully")
	return nil
}

func selectLevelLogger(env string) logger.Interface {
	switch env {
	case "LOCAL":

		dbLogLevel := os.Getenv("DB_LOG_LEVEL")
		switch dbLogLevel {
		case "silent":
			return logger.Default.LogMode(logger.Silent)
		case "error":
			return logger.Default.LogMode(logger.Error)
		case "warn":
			return logger.Default.LogMode(logger.Warn)
		case "info":
			return logger.Default.LogMode(logger.Info)
		default:
			return logger.Default.LogMode(logger.Info)
		}

	case "DEV", "QA":
		return logger.Default.LogMode(logger.Info)
	case "PDN":
		return logger.Default.LogMode(logger.Error)
	default:
		return logger.Default.LogMode(logger.Info)
	}
}
