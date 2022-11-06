package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/toporo/go-cep-search/internal/api"
)

func main() {
	e := echo.New()

	cepApi := api.NewCepApi()

	cepApi.Register(e)

	err := e.Start(":8080")
	fmt.Println(err.Error())
}
