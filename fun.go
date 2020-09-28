package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/spf13/viper"
)

func init() {
	var coreCategory = &gommand.Category{
		Name:        "Fun",
		Description: "Fun commands! :partying_face:",
	}

	commands = append(commands,
		&gommand.Command{
			Name:        "gay",
			Description: "How gay are you?",
			Category:    coreCategory,
			Function: func(ctx *gommand.Context) error {
				gay := fmt.Sprint(rand.Intn(100))
				if ctx.Message.Author.ID == 251411894894592003 {
					gay = "999"
				}

				_, _ = ctx.Reply(disgord.Embed{
					Title:       ":rainbow_flag: Gay",
					Description: "You are " + gay + "% gay",
					Color:       viper.GetInt("bot.color"),
				})
				return nil
			},
			Aliases: []string{"homo", "homosexual"},
		},
		&gommand.Command{
			Name:        "furry",
			Description: "How furry are you?",
			Category:    coreCategory,
			Function: func(ctx *gommand.Context) error {
				_, _ = ctx.Reply(disgord.Embed{
					Title:       ":lion_face: Furry",
					Description: "You are " + fmt.Sprint(rand.Intn(100)) + "% furry",
					Color:       viper.GetInt("bot.color"),
				})
				return nil
			},
		},
		&gommand.Command{
			Name:        "about",
			Description: "About Iris.",
			Category:    coreCategory,
			Function: func(ctx *gommand.Context) error {
				var m runtime.MemStats
				runtime.ReadMemStats(&m)

				_, _ = ctx.Reply(disgord.Embed{
					Title:       "About Iris",
					Description: "Bot by [Xilog.](https://xilog.xyz) Thanks to [Alex](https://alex.lgbt/) and [Gommand](https://github.com/auttaja/gommand).",
					Fields: []*disgord.EmbedField{
						{
							Name:   "Go version",
							Value:  runtime.Version(),
							Inline: true,
						},
						{
							Name:   "Allocated mem",
							Value:  fmt.Sprintf("%v MiB", m.Alloc/1024/1024),
							Inline: true,
						},
						{
							Name:   "Disgord version",
							Value:  fmt.Sprintf(disgord.LibraryInfo()),
							Inline: true,
						},
						{
							Name: "Shard ID",
							// fmt.Sprint converts uint shard ID to the correct type of string
							Value:  fmt.Sprint(disgord.ShardID(ctx.Message.GuildID, uint(viper.GetInt("bot.shards")))),
							Inline: true,
						},
					},
					Color: viper.GetInt("bot.color"),
				})
				return nil
			},
			Aliases: []string{"iris", "info"},
		},
	)
}
