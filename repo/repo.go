package repo

import (
	"context"
	"sync"

	"github.com/andhikasamudra/medium-go-di/config"
)

type Dependency struct{
	Mtx sync.Mutex 
}

type Repo struct{
	Mtx sync.Mutex
}

func NewRepo(d Dependency) *Repo{
	return &Repo{}
}

type Something struct {
	// any item
}

func (r *Repo) CreateSomething(ctx context.Context, dbAdapter config.ExampleAdapter, data Something) (*Something, error) {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()

	_, err := dbAdapter.Model(&data).Insert()
	if err != nil {
		return nil, err
	}

	return &data, nil
}