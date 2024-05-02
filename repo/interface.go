package repo

import (
	"context"

	"github.com/andhikasamudra/medium-go-di/config"
)

type Interface interface{
	CreateSomething(ctx context.Context, dbAdapter config.ExampleAdapter, data Something) (*Something, error)
}