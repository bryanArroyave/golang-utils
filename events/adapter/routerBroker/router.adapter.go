package routerbroker

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

type Router struct {
	router *message.Router
}

func NewRouter() *Router {
	// ...
	logger := watermill.NewStdLogger(false, false)

	r, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	r.AddPlugin(plugin.SignalsHandler)

	r.AddMiddleware(
		middleware.CorrelationID,

		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		middleware.Recoverer,
	)

	return &Router{
		router: r,
	}
}

func (r *Router) AddHandler(
	name string,
	topic string,
	subscriber message.Subscriber,
	handlerFunc message.NoPublishHandlerFunc,
) {
	_ = r.router.AddNoPublisherHandler(
		name,
		topic,
		subscriber,
		handlerFunc,
	)
}

func (r *Router) Run(ctx context.Context) error {
	return r.router.Run(ctx)
}
