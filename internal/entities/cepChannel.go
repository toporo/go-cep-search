package entities

import "github.com/joomcode/errorx"

type CepChannel struct {
	Cep  CepDto
	Errx *errorx.Error
}
