package api

import (
	"strings"

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
	v1.GET("/cep/batch", ca.GetBatchCep)
}

func (ca *CepApi) GetCep(c echo.Context) error {
	value := c.Param("cep")

	resp, errx := ca.CepService.GetCep(value)

	if errx != nil {
		return c.JSON(500, errx.Cause())
	}

	return c.JSON(200, resp)
}

func (ca *CepApi) GetBatchCep(c echo.Context) error {
	ceps := c.QueryParam("ceps")

	cepsArray := strings.Split(ceps, ",")

	resp, errx := ca.CepService.GetCepInBatch(cepsArray)

	if errx != nil {
		return c.JSON(500, errx.Cause())
	}

	return c.JSON(200, resp)
}
