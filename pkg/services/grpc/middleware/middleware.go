package middleware

import (
	"google.golang.org/grpc"

	"github.com/SSH-Management/server/pkg/container"
)

func Register(c *container.Container) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			UnaryErrorHandler(
				c.GetTranslator(),
				c.GetDefaultLogger(),
			),
		),
	}
}
