package service

import (
	"context"

	"github.com/andhikasamudra/medium-go-di/config"
	"github.com/andhikasamudra/medium-go-di/dto"
	"github.com/andhikasamudra/medium-go-di/repo"
)

type Dependency struct {
	Repo      repo.Interface
	DbAdapter config.ExampleAdapter
}

type Service struct {
	Repo      repo.Interface
	DbAdapter config.ExampleAdapter
}

func NewService(d Dependency) *Service {
	return &Service{
		Repo:      d.Repo,
		DbAdapter: d.DbAdapter,
	}
}

func (s *Service) CreateSomething(ctx context.Context, request dto.CreateSomethingRequest) error {
	data := repo.Something{}
	_, err := s.Repo.CreateSomething(ctx, s.DbAdapter, data)
	if err != nil {
		return err
	}

	return nil
}
