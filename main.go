package main

import (
	"flag"
	"fmt"
	router "github.com/OliCoder/RecSys/router"
	"github.com/OliCoder/RecSys/settings"
	"net/http"
)

func main() {
	flag.Parse()
	r := router.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", settings.GetConfigInstance().ServerPort),
		Handler:        r,
		ReadTimeout:    settings.GetConfigInstance().ServerReadTimeout,
		WriteTimeout:   settings.GetConfigInstance().ServerWriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
