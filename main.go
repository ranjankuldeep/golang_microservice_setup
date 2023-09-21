package main

import (
	"context"
	"log"
)

func main() {
	println("Hello, World!")
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)
	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":8080"))

	_, err := svc.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
