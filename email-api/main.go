package main

import (
	"fmt"

	"github.com/eFico/email-api/config"
	"github.com/eFico/email-api/emails/web"
)

func main() {
	fmt.Println("Hello, World!")
	API_URL := "http://localhost:4080/es/email/_search"

	handler := web.NewEmailSearchHandler(API_URL)

	mux := config.Routes(handler)
	server := config.NewServer(mux)
	server.Run()
}
