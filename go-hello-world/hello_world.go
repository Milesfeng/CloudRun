package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// add middleware and routes
	// ...
	e.GET("/:id", func(c echo.Context) error {
		// return c.HTML(http.StatusOK, fmt.Sprintf("<p><h2>Hello World  !! </h2></p>"))
		id := c.Param("id")
		return c.String(http.StatusOK, "Hi V2 "+id)
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
