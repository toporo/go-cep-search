package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/toporo/go-cep-search/internal/api"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	e := echo.New()

	cepApi := api.NewCepApi()

	cepApi.Register(e)

	err := e.Start(":8080")
	fmt.Println(err.Error())
}
