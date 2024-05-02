package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andhikasamudra/medium-go-di/dto"
	"github.com/andhikasamudra/medium-go-di/service"
)

type Dependency struct {
	Service service.Interface
}

type Handler struct {
	Service service.Interface
}

func NewHandler(d Dependency) *Handler {
	return &Handler{
		Service: d.Service,
	}
}

func (h *Handler) CreateSomething(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	request := dto.CreateSomethingRequest{}

	err := h.Service.CreateSomething(ctx, request)
	if err != nil {
		fmt.Println("something wrong")
	}

}
