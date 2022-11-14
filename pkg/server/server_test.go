package server

import (
	"context"
	"testing"

	hsM "jwt/pkg/handlers/mock"
)

func TestXxx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2)
	defer cancel()
	var serverTest IServer = New(ctx, "localhost:80", &hsM.MockIHandlers{})
	if err := serverTest.Run(); err != nil {
		panic(err)
	}
}
