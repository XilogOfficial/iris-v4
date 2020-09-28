package main

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var router = gommand.NewRouter(&gommand.RouterConfig{
	// the prefix function should be set here or it will be blank
	PrefixCheck: gommand.MultiplePrefixCheckers(gommand.StaticPrefix("="), gommand.MentionPrefix),
})

func init() {
	var coreCategory = &gommand.Category{
		Name:        "Core",
		Description: "Base Iris commands",
	}

	router.GetCommand("help").(*gommand.Command).Category = coreCategory
	router.GetCommand("help").(*gommand.Command).Aliases = []string{"h"}

	router.SetCommand(&gommand.Command{
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
				Color:       viper.GetInt("color"),
			})
			return nil
		},
		Aliases: []string{"p"},
	})
}

func main() {
	// viper config
	viper.SetConfigName("conf") // name of config file (without extension)
	viper.AddConfigPath(".")    // look for config in current dir
	e := viper.ReadInConfig()   // find and read the config file
	if e != nil {               // error reading file
		panic(fmt.Errorf("Fatal error config file: %s \n", e))
	}

	logrus.SetLevel(logrus.DebugLevel)

	s := disgord.New(disgord.Config{
		BotToken: viper.GetString("token"),
		Logger:   logrus.New(),
	})

	router.Hook(s)

	err := s.StayConnectedUntilInterrupted(context.Background())
	if err != nil {
		panic(err)
	}

	// error handler
	router.AddErrorHandler(func(ctx *gommand.Context, err error) bool {
		switch err.(type) {
		case *gommand.CommandNotFound, *gommand.CommandBlank:
			// command no exist
			return true
		case *gommand.InvalidTransformation:
			_, _ = ctx.Reply("Invalid argument:", err.Error())
			return true
		case *gommand.IncorrectPermissions:
			_, _ = ctx.Reply("You don't have permission to do that.", err.Error())
			return true
		case *gommand.InvalidArgCount:
			_, _ = ctx.Reply("Invalid argument count.")
			return true
		}

		// This was not handled here.
		return false
	})
}
