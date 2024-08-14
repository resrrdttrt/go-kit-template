package admin

import (
	"context"
)


type MathRepository interface {
	Add(ctx context.Context, numA, numB float32)(sum float32,err error)
}