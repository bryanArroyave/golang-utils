package dtos

type MongoConnectionDTO struct {
	Host     string `json:"host" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	DBName   string `json:"dbname" validate:"required"`
	Port     string `json:"port,omitempty"`
}
