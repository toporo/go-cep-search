package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"

	"github.com/toporo/go-cep-search/internal/api"
)

func main() {
	e := echo.New()

	cepApi := api.NewCepApi()

	cepApi.Register(e)

	go func() {
		err := e.Start(":8080")
		fmt.Println(err.Error())
	}()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")

	<-done

	fmt.Println("exiting")

}
