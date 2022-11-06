package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/toporo/go-cep-search/internal/api"
	"github.com/toporo/go-cep-search/internal/entities"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	e := echo.New()

	cepApi := api.NewCepApi()

	cepApi.Register(e)

	err := e.Start(":8080")
	fmt.Println(err.Error())
}

func FindCep(cep string) entities.CepDto {
	cepObject := new(entities.CepDto)
	getCep(cep, cepObject)

	return *cepObject
}

func getCep(cep string, target interface{}) error {
	log.Info().Msg("findCep: " + cep)

	resp, err := httpClient.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		log.Error().Err(err).Msg("Error findCep: " + cep)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
