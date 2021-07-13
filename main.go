package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var todos = []Todo{
	{ID: 1, Name: "Buat Video"},
}

func main() {
	router := echo.New()
	router.GET("/todo", GetTodos)

	s := &http.Server{
		Addr:    ":1234",
		Handler: router,
	}

	log.Printf("listening :1234")
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}
