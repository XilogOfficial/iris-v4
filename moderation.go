package main

import (
	"context"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/spf13/viper"
)

func init() {
	var modCategory = &gommand.Category{
		Name:        "Moderation",
		Description: "Moderation commands",
	}

	commands = append(commands,
		&gommand.Command{
			Name:        "kick",
			Description: "Kicks a user",
			Usage:       "<member> [reason]",
			Category:    modCategory,
			ArgTransformers: []gommand.ArgTransformer{
				{
					Function: gommand.MemberTransformer,
				},
				{
					Function:  gommand.StringTransformer,
					Remainder: true,
					Optional:  true,
				},
			},
			PermissionValidators: []gommand.PermissionValidator{
				gommand.KICK_MEMBERS(gommand.CheckMembersUserPermissions | gommand.CheckBotUserPermissions),
			},
			Function: func(ctx *gommand.Context) error {
				menu := gommand.NewEmbedMenu(&disgord.Embed{
					Title: "Are you sure?",
					Color: viper.GetInt("bot.color"),
				}, ctx)

				_ = menu.NewChildMenu(&gommand.ChildMenuOptions{
					Embed: &disgord.Embed{
						Title: "Kicked.",
						Color: viper.GetInt("bot.color"),
					},
					Button: &gommand.MenuButton{
						Emoji: "✅",
					},
					AfterAction: func() {
						reason, ok := ctx.Args[1].(string)
						if !ok {
							reason = ""
						}
						ctx.Session.KickMember(context.Background(), ctx.Message.GuildID, ctx.Args[0].(*disgord.Member).User.ID, reason)
					},
				})
				_ = menu.NewChildMenu(&gommand.ChildMenuOptions{
					Embed: &disgord.Embed{
						Title: "Cancelled.",
						Color: viper.GetInt("bot.color"),
					},
					Button: &gommand.MenuButton{
						Emoji: "❎",
					},
				})

				err := ctx.DisplayEmbedMenu(menu)
				if err != nil {
					return err
				}

				return nil
			},
		},
		&gommand.Command{
			Name:        "ban",
			Description: "Bans a user",
			Usage:       "<member> [reason]",
			Category:    modCategory,
			ArgTransformers: []gommand.ArgTransformer{
				{
					Function: gommand.MemberTransformer,
				},
				{
					Function:  gommand.StringTransformer,
					Remainder: true,
					Optional:  true,
				},
			},
			PermissionValidators: []gommand.PermissionValidator{
				gommand.BAN_MEMBERS(gommand.CheckMembersUserPermissions | gommand.CheckBotUserPermissions),
			},
			Function: func(ctx *gommand.Context) error {
				menu := gommand.NewEmbedMenu(&disgord.Embed{
					Title: "Are you sure?",
					Color: viper.GetInt("bot.color"),
				}, ctx)

				_ = menu.NewChildMenu(&gommand.ChildMenuOptions{
					Embed: &disgord.Embed{
						Title: "Banned.",
						Color: viper.GetInt("bot.color"),
					},
					Button: &gommand.MenuButton{
						Emoji: "✅",
					},
					AfterAction: func() {
						reason, ok := ctx.Args[1].(string)
						if !ok {
							reason = ""
						}
						ctx.Session.BanMember(context.Background(), ctx.Message.GuildID, ctx.Args[0].(*disgord.Member).User.ID, &disgord.BanMemberParams{Reason: reason, DeleteMessageDays: 0})
					},
				})
				_ = menu.NewChildMenu(&gommand.ChildMenuOptions{
					Embed: &disgord.Embed{
						Title: "Cancelled.",
						Color: viper.GetInt("bot.color"),
					},
					Button: &gommand.MenuButton{
						Emoji: "❎",
					},
				})

				err := ctx.DisplayEmbedMenu(menu)
				if err != nil {
					return err
				}

				return nil
			},
		},
	)
}
