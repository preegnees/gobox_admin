package main

import (
	sr "jwt/pkg/server"
	ss "jwt/pkg/services"
	se "jwt/pkg/storage"
	hs "jwt/pkg/handlers"
)

func main() {

	var storage se.IStorage = se.New()
	var services ss.IService = ss.New(storage)
	var handlers hs.IHandlers = hs.New(services)

	var server sr.IServer = sr.New("localhost:80", handlers)
	if err := server.Run(); err != nil {
		panic(err)
	}
}