package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	s := &http.Server{
		Addr:    ":1234",
		Handler: router,
	}

	log.Printf("listening :1234")
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
