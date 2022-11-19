package main

import (
	"context"
	"os"

	handlers "jwt/pkg/handlers"
	middlewares "jwt/pkg/middlewares"
	server "jwt/pkg/server"
	services "jwt/pkg/services"
	storage "jwt/pkg/storage"
	jwt "jwt/pkg/utils/jwt"
)

func main() {
	os.Setenv("SECRET", "TEST")
	jwt_, err := jwt.New()
	if err != nil {
		panic(err)
	}
	var storage_ storage.IStorage = storage.New()
	var services_ services.IService = services.New(storage_)
	var handlers_ handlers.IHandlers = handlers.New(services_, jwt_)
	var middlewares_ middlewares.IMiddleware = middlewares.New(jwt_)

	ctx := context.Background()
	var server server.IServer = server.New(ctx, "localhost:80", handlers_, middlewares_)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
