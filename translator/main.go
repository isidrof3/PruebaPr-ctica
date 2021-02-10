package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"translator/endpoint"
	processTranslate "translator/process/translate"
	"syscall"

	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"github.com/go-kit/kit/log"
)

func main() {
	//Se crea la dependencia de para la lectura de configuracion (yaml y vault)
	config := configuration.GetInstance(configuration.NewDefaultSetting())
	//Se lee el puerto en el que correrá el servicio
	port := config.GetString("server.port")
	//Se crea la dependencia de para log
	logger := log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "time", log.DefaultTimestamp)
	//Se crea la dependencia de base de datos

	/*
		Se iniciliaza las depenendecinas
	*/

	//Se instancia el proceso de translate
	prcTranslate := processTranslate.NewProcessTranslate(config, logger)

	//se crea el servidor
	mux := http.NewServeMux()

	e := endpoint.EndPoints{
		Translate: endpoint.MakeTranslateEndpoint(prcTranslate),
	}

	mux.Handle("/translator/", endpoint.NewHttpHandler(e))

	//Se crea el canal para el manejo de error
	errs := make(chan error, 2)
	//Se crea la go routine para el servidor
	go func() {
		logger.Log("transport", "http", "address", port, "msg", "listening")
		errs <- http.ListenAndServe(":"+port, mux)
	}()
	//Se crea la go routine para cuando se termine la ejecución
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
