package config

import (
	"context"
	"os"
	"path"

	"github.com/Barben360/texas-holdem-game/errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	WorkersPlayerActions int `envconfig:"workers_player_actions" default:"1"`
	WorkersHandChecking  int `envconfig:"workers_hand_checking" default:"1"`
}

func New(ctx context.Context) (Config, error) {
	cfg := Config{}
	// Reading .env if required
	if os.Getenv("READ_ENV_FILE") == "1" {
		err := godotenv.Load(path.Join(os.Getenv("APP_ROOT_PATH"), ".env"))
		if err != nil {
			genErr := errors.Errorf(errors.ErrInternal, "could not load .env: %v", err)
			logrus.WithError(genErr).Error()
			return cfg, genErr
		}
	}
	// Parsing config
	err := envconfig.Process("", &cfg)
	if err != nil {
		genErr := errors.Errorf(errors.ErrInternal, "could not process config from environment: %v", err)
		logrus.WithError(genErr).Error()
		return cfg, genErr
	}
	return cfg, nil
}
