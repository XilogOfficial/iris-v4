package main

import (
	"fmt"
	"runtime"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/spf13/viper"
)

func init() {
	var coreCategory = &gommand.Category{
		Name:        "Core",
		Description: "Base Iris commands",
	}
	// change help command category and alias
	router.GetCommand("help").(*gommand.Command).Category = coreCategory
	router.GetCommand("help").(*gommand.Command).Aliases = []string{"h"}

	commands = append(commands,
		&gommand.Command{
			Name:        "ping",
			Description: "Responds with pong.",
			Category:    coreCategory,
			Function: func(ctx *gommand.Context) error {
				ping, err := ctx.Session.AvgHeartbeatLatency()
				if err != nil {
					return err
				}

				_, _ = ctx.Reply(disgord.Embed{
					Title:       "Ping",
					Description: fmt.Sprintf("Pong! %s", ping.String()),
					Color:       viper.GetInt("bot.color"),
				})
				return nil
			},
			Aliases: []string{"p"},
		},
		&gommand.Command{
			Name:        "invite",
			Description: "Invite Iris to a server.",
			Category:    coreCategory,
			Function: func(ctx *gommand.Context) error {
				_, _ = ctx.Reply(disgord.Embed{
					Title: ":link: Click here to invite me!",
					URL:   viper.GetString("bot.invite"),
					Color: viper.GetInt("bot.color"),
				})
				return nil
			},
			Aliases: []string{"inv"},
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
					Description: "Bot by [Xilog.](https://xilog.xyz) Thanks to [Alex](https://alex.lgbt/)",
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
					},
					Color: viper.GetInt("bot.color"),
				})
				return nil
			},
			Aliases: []string{"iris", "info"},
		},
	)
}
