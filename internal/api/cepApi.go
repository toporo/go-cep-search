package api

import (
	"github.com/labstack/echo/v4"
	"github.com/toporo/go-cep-search/internal/service"
)

type CepApi struct {
	CepService service.CepServicer
}

func NewCepApi() *CepApi {
	return &CepApi{
		CepService: service.NewCepService(),
	}
}

func (ca *CepApi) Register(echo *echo.Echo) {
	v1 := echo.Group("/v1")

	v1.GET("/cep/:cep", ca.GetCep)
}

func (ca *CepApi) GetCep(c echo.Context) error {
	value := c.Param("cep")

	resp, errx := ca.CepService.GetCep(value)

	if errx != nil {
		return c.JSON(500, errx.Cause())
	}

	return c.JSON(200, resp)
}
