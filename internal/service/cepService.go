package service

import (
	"github.com/joomcode/errorx"
	"github.com/toporo/go-cep-search/internal/client"
	"github.com/toporo/go-cep-search/internal/entities"
)

type CepServicer interface {
	GetCep(cep string) (entities.CepDto, *errorx.Error)
	GetCepInBatch(ceps []string) ([]entities.CepDto, *errorx.Error)
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

func (cs *CepService) GetCepInBatch(ceps []string) ([]entities.CepDto, *errorx.Error) {
	ch := make(chan entities.CepChannel, len(ceps))

	for i := 0; i < len(ceps); i++ {
		go cs.findCep(ceps[i], ch)
	}

	output := make([]entities.CepDto, len(ceps))

	for i := 0; i < len(ceps); i++ {
		chOuput := <-ch
		output[i] = chOuput.Cep
	}

	return output, nil
}

func (cs *CepService) findCep(cep string, cepCh chan entities.CepChannel) {
	cepDto := new(entities.CepDto)

	err := cs.CepClient.GetCep(cep, cepDto)

	if err != nil {
		cepCh <- entities.CepChannel{
			Cep: entities.CepDto{},
			Err: "failed to get CEP " + cep + "ERROR: " + err.Error(),
		}
	}

	cepCh <- entities.CepChannel{Cep: *cepDto, Err: ""}
}
