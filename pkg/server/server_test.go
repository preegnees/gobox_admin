package server

import (
	"context"
	"testing"

	handlersMock "jwt/pkg/handlers/mock"
	middlewaresMock "jwt/pkg/middlewares/mock"
)

func TestXxx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 100)
	defer cancel()
	var serverTest IServer = New(ctx, "localhost:80", &handlersMock.MockIHandlers{}, &middlewaresMock.MockIMiddleware{})
	if err := serverTest.Run(); err != nil {
		panic(err)
	}
}
