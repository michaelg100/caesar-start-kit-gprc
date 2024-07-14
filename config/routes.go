package config

import (
	"testgrpc/app/grpc"
	"testgrpc/views/pages"

	v1connect "testgrpc/gen/acme/weather/v1/v1connect"

	caesar "github.com/caesar-rocks/core"
)

func authMiddleware(ctx *caesar.Context) error {
	return nil
}

func RegisterRoutes(
// Inject your controllers here...
) *caesar.Router {
	router := caesar.NewRouter()
	// Cant have generic '/' route when using grpc router
	router.Render("/home", pages.WelcomePage()).Use(authMiddleware)
	// Get GRPC Service and add to the
	path, handler := v1connect.NewWeatherServiceHandler(&grpc.WeatherServiceServer{})
	var extraHandler []caesar.ExtraHandler
	extraHandler = append(extraHandler, caesar.ExtraHandler{
		Path:    path,
		Handler: handler,
	})
	router.CustomHandlers = extraHandler
	return router
}
