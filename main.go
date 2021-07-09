package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"time"

	initApp "github.com/Barben360/texas-holdem-game/init"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initializing random seed
	rand.Seed(time.Now().UnixNano())

	// Initializing app
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app, err := initApp.Init(ctx)
	if err != nil {
		logrus.WithError(err).Error("could not initialize app, exiting")
		os.Exit(1)
	}

	// Waiting for sigterm
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Deferring stop
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := app.GracefulStop(ctx); err != nil {
			logrus.WithError(err).Error("could not gracefully stop app, force-exiting")
			os.Exit(1)
		}
	}()
}
