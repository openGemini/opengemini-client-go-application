package boot

import (
	"github.com/libgox/envx"
	"golang.org/x/exp/slog"
)

type ClientApplication struct {
}

func NewClientApplication() *ClientApplication {
	logger := slog.Default()
	envx.GetStrOr("POD_NAME", "defaultName")
	logger.Info("Starting client application")
	return &ClientApplication{}
}
