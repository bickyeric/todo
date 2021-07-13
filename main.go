package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var todos = []Todo{}

func main() {
	router := echo.New()
	router.POST("/todo", PostTodo)

	s := &http.Server{
		Addr:    ":1234",
		Handler: router,
	}

	log.Printf("listening :1234")
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func PostTodo(c echo.Context) error {
	var todo Todo
	err := c.Bind(&todo)
	if err != nil {
		return c.String(500, "terjadi error.")
	}

	return c.String(201, "todo telah ditambahkan")
}
