package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andhikasamudra/medium-go-di/handler"
	"github.com/andhikasamudra/medium-go-di/repo"
	"github.com/andhikasamudra/medium-go-di/service"
)

func main() {
	rp := repo.NewRepo()
	s := service.NewService(
		service.Dependency{
			Repo: rp,
		},
	)
	h := handler.NewHandler(
		handler.Dependency{
			Service: s,
		},
	)

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", h.CreateSomething)
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", ":8000"),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()
}
