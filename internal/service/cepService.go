package service

import (
	"github.com/joomcode/errorx"
	"github.com/toporo/go-cep-search/internal/client"
	"github.com/toporo/go-cep-search/internal/entities"
)

type CepServicer interface {
	GetCep(cep string) (entities.CepDto, *errorx.Error)
}

type CepClientHandler interface {
	GetCep(cep string, target interface{}) error
}

type CepService struct {
	CepClient CepClientHandler
}

func NewCepService() *CepService {
	return &CepService{
		CepClient: client.NewCepClientHandle(),
	}
}

func (cs *CepService) GetCep(cep string) (entities.CepDto, *errorx.Error) {
	cepDto := new(entities.CepDto)

	err := cs.CepClient.GetCep(cep, cepDto)

	if err != nil {
		return entities.CepDto{}, errorx.Decorate(err, "failed to get CEP")
	}

	return *cepDto, nil
}
