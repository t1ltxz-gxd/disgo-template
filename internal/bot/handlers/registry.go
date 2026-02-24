package handlers

import (
	"reflect"

	"github.com/disgoorg/disgo/bot"
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Client   *bot.Client
	Handlers []types.Event `group:"handlers"`
	Cfg      *config.Config
}

func NewRegistry(p Params) {
	handlerMap := make(map[reflect.Type][]types.Event)

	for _, h := range p.Handlers {
		eventType := reflect.TypeOf(h.Event())

		if eventType.Kind() == reflect.Pointer {
			eventType = eventType.Elem()
		}

		handlerMap[eventType] = append(handlerMap[eventType], h)
	}

	p.Client.EventManager.AddEventListeners(
		bot.NewListenerFunc(func(event bot.Event) {
			eventType := reflect.TypeOf(event)

			if eventType.Kind() == reflect.Pointer {
				eventType = eventType.Elem()
			}

			if handlers, ok := handlerMap[eventType]; ok {
				for _, h := range handlers {
					if err := h.Handle(event); err != nil {
						logger.Error("event handler error:", zap.Error(err))
					}
				}
			} else {
				logger.Debug("unknown event received",
					zap.String("event_type", reflect.TypeOf(event).String()),
				)
			}
		}),
	)
}
