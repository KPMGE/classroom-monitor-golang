package adapters

import (
	"github.com/gofiber/fiber/v2"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
	presentationprotocols "github.com/monitoring-go/src/presentation/presentation-protocols"
)

func FiberRouteAdapter(controller presentationprotocols.Controller) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		params := c.Route().Params

		if params == nil {
			request := httpprotocols.NewHttpRequest(nil, nil)
			httpResponse := controller.Handle(request)
			return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
		}

		paramName := params[0]
		paramValue := c.Params(paramName)
		request := httpprotocols.NewHttpRequest(paramValue, nil)
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	}
}
