package init

import (
	"context"
)

type App struct {
}

func Init(ctx context.Context) (App, error) {
	return App{}, nil
}

func (a *App) GracefulStop(ctx context.Context) error {
	return nil
}
