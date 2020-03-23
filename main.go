package main

import (
	"flag"
	"fmt"
	"github.com/OliCoder/RecSys/engine"
	router "github.com/OliCoder/RecSys/router"
	"github.com/OliCoder/RecSys/settings"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
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
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for s := range c {
			log.Info("Process Exit with %v", s)
			engine.ClientTransport.Close()
			os.Exit(0)
		}
	}()
}
