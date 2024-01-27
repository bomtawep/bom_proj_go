package configs

import (
	"context"
	"time"
)

func CtxWithTimout() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}
