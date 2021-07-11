package services

import (
	"context"

	"github.com/Barben360/texas-holdem-game/errors"
	"github.com/Barben360/texas-holdem-game/services/config"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Config config.Config
}

func Init(ctx context.Context) (Services, error) {
	svcs := Services{}

	// Initializing config
	cfg, err := config.New(ctx)
	if err != nil {
		errMsg := "could not initialize config"
		logrus.WithError(err).Error(errMsg)
		return svcs, errors.ErrorfAppend(err, errMsg)
	}
	svcs.Config = cfg

	return svcs, nil
}
