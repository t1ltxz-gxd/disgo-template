package commands

import (
	"bytes"
	_ "embed"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
)

//go:embed thumbnail.jpg
var thumbnail []byte

type TestCommand struct{}

func NewTestCommand() types.Command {
	return &TestCommand{}
}

func (t *TestCommand) Name() string {
	return "test"
}

func (t *TestCommand) DevOnly() bool {
	return true
}

func (t *TestCommand) Command() discord.ApplicationCommandCreate {
	return discord.SlashCommandCreate{
		Name:        t.Name(),
		Description: "test",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionBool{
				Name:        "ephemeral",
				Description: "if the message should be ephemeral",
				Required:    false,
			},
		},
	}
}

func (t *TestCommand) Handle(data discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
	flags := discord.MessageFlagIsComponentsV2
	if ephemeral, ok := data.OptBool("ephemeral"); !ok || ephemeral {
		flags = flags.Add(discord.MessageFlagEphemeral)
	}

	return e.CreateMessage(discord.MessageCreate{
		Flags: flags,
		Components: []discord.LayoutComponent{
			discord.NewContainer(
				discord.NewSection(
					discord.NewTextDisplay("**Name: [Seeing Red](https://open.spotify.com/track/65qBr6ToDUjTD1RiE1H4Gl)**"),
					discord.NewTextDisplay("**Artist: [Architects](https://open.spotify.com/artist/3ZztVuWxHzNpl0THurTFCv)**"),
					discord.NewTextDisplay("**Album: [The Sky, The Earth & All Between](https://open.spotify.com/album/2W82VyyIFAXigJEiLm5TT1)**"),
				).WithAccessory(discord.NewThumbnail("attachment://thumbnail.png")),

				discord.NewTextDisplay("`0:08`/`3:40`"),
				discord.NewTextDisplay("[🔘▬▬▬▬▬▬▬▬▬]"),
				discord.NewSmallSeparator(),

				discord.NewActionRow(
					discord.NewPrimaryButton("", "player_previous").WithEmoji(discord.ComponentEmoji{Name: "⏮"}),
					discord.NewPrimaryButton("", "player_pause").WithEmoji(discord.ComponentEmoji{Name: "⏯"}),
					discord.NewPrimaryButton("", "player_next").WithEmoji(discord.ComponentEmoji{Name: "⏭"}),
					discord.NewDangerButton("", "player_stop").WithEmoji(discord.ComponentEmoji{Name: "⏹"}),
					discord.NewPrimaryButton("", "player_like").WithEmoji(discord.ComponentEmoji{Name: "❤️"}),
				),
			).WithAccentColor(0x5c5fea),
		},
		Files: []*discord.File{
			discord.NewFile("thumbnail.png", "", bytes.NewReader(thumbnail)),
		},
	})
}

var _ types.Command = (*TestCommand)(nil)
