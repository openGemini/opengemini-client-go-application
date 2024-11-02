package boot

import (
	"github.com/libgox/envx"
	"github.com/libgox/slogsimple"
	"golang.org/x/exp/slog"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func NewTestClientApplication() *Application {
	logger := slog.Default()
	envx.GetStrOr("POD_NAME", "defaultName")
	logger.Info("Starting client application")
	return &Application{}
}

func (a *Application) Boot() {
	a.logger = slog.New(slogsimple.NewHandler(&slogsimple.Config{
		Output:   os.Stdout,
		MinLevel: slog.LevelInfo,
	}))
}
