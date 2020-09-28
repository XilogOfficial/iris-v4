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

var commands []gommand.CommandInterface // commands slice for other files to append commands to

func main() {
	// viper config
	viper.SetConfigName("conf") // name of config file (without extension)
	viper.AddConfigPath(".")    // look for config in current dir
	e := viper.ReadInConfig()   // find and read the config file
	if e != nil {               // error reading file
		panic(fmt.Errorf("Fatal error config file: %s \n", e))
	}

	// set all the commands in the slice
	for _, v := range commands {
		router.SetCommand(v)
	}

	// error handler
	router.AddErrorHandler(func(ctx *gommand.Context, err error) bool {
		switch err.(type) {
		case *gommand.CommandNotFound, *gommand.CommandBlank:
			// command no exist
			return true
		case *gommand.InvalidTransformation:
			_, _ = ctx.Reply(disgord.Embed{
				Title:       "Whoops!",
				Description: fmt.Sprintf("Invalid argument: %s", err.Error()),
				Color:       16711680,
			})
			return true
		case *gommand.IncorrectPermissions:
			_, _ = ctx.Reply(disgord.Embed{
				Title:       "Whoops!",
				Description: fmt.Sprintf("You don't have permission to do this! %s", err.Error()),
				Color:       16711680,
			})
			return true
		case *gommand.InvalidArgCount:
			_, _ = ctx.Reply(disgord.Embed{
				Title:       "Whoops!",
				Description: "Invalid amount of arguments.",
				Color:       16711680,
			})
			return true
		}

		// This was not handled here.
		return false
	})

	logrus.SetLevel(logrus.DebugLevel)

	s := disgord.New(disgord.Config{
		BotToken: viper.GetString("bot.token"),
		Logger:   logrus.New(),
		ShardConfig: disgord.ShardConfig{
			// this is a copy so u can't manipulate the config later on
			ShardIDs:   []uint{0},                        // this number must change for each instance. Try to automate this.
			ShardCount: uint(viper.GetInt("bot.shards")), // total of 5 shards, but this disgord instance only has one. AutoScaling is disabled - use OnScalingRequired.
		},
	})

	router.Hook(s)

	err := s.StayConnectedUntilInterrupted(context.Background())
	if err != nil {
		panic(err)
	}
}
