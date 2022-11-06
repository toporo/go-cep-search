package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type CepClientHandle struct {
}

func NewCepClientHandle() *CepClientHandle {
	return &CepClientHandle{}
}

func (cs *CepClientHandle) GetCep(cep string, target interface{}) error {
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetResult(target).
		ForceContentType("application/json").
		Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))

	return err
}
