package grpc

import (
	"context"

	v1 "testgrpc/gen/acme/weather/v1"
	v1connect "testgrpc/gen/acme/weather/v1/v1connect"

	connect "connectrpc.com/connect"
)

type WeatherServiceServer struct {
	v1connect.UnimplementedWeatherServiceHandler
}

func (s *WeatherServiceServer) Get(ctx context.Context, req *connect.Request[v1.Location]) (*connect.Response[v1.GetWeatherResponse], error) {
	return connect.NewResponse(&v1.GetWeatherResponse{}), nil
}

func (s *WeatherServiceServer) GetWeather(c context.Context, req *connect.Request[v1.Location]) (*connect.Response[v1.GetWeatherResponse], error) {
	return connect.NewResponse(&v1.GetWeatherResponse{
		Temperature: 2333.23,
	}), nil
}
