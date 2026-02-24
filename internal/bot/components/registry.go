package components

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

	Client     *bot.Client
	Components []types.Event `group:"components"`
	Cfg        *config.Config
}

func NewRegistry(p Params) {
	componentMap := make(map[reflect.Type][]types.Event)

	for _, h := range p.Components {
		eventType := reflect.TypeOf(h.Event())
		componentMap[eventType] = append(componentMap[eventType], h)
	}

	p.Client.EventManager.AddEventListeners(
		bot.NewListenerFunc(func(event bot.Event) {
			eventType := reflect.TypeOf(event)

			if components, ok := componentMap[eventType]; ok {
				for _, c := range components {
					if err := c.Handle(event); err != nil {
						logger.Error("event handler error:", zap.Error(err))
					}
				}
			}
		}),
	)
}
