package types

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type Event interface {
	Event() any
	Handle(event any) error
}

type Command interface {
	Name() string
	Command() discord.ApplicationCommandCreate
	Handle(discord.SlashCommandInteractionData, *handler.CommandEvent) error
	DevOnly() bool
}
