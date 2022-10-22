package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/toporo/go-search-cep/entities"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	cep := "01310940"
	cepObject := new(entities.CepDto)
	findCep(cep, cepObject)

	log.Info().Msg("# Found cep: " + cepObject.Logradouro)
}

func findCep(cep string, target interface{}) error {
	log.Info().Msg("## findCep: " + cep)

	resp, err := httpClient.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		log.Error().Err(err).Msg("Error findCep: " + cep)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
