package service

import (
	"context"

	"github.com/andhikasamudra/medium-go-di/dto"
)

type Interface interface{
	CreateSomething(ctx context.Context, request dto.CreateSomethingRequest) error
}