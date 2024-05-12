package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/Aracelimartinez/email-platform-challenge/server/configs"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/router"
)

func main() {
	router := router.Generate()
	fmt.Printf("Listening in port :%s\n", configs.GlobalConfig.APIPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.GlobalConfig.APIPort), router))
}
