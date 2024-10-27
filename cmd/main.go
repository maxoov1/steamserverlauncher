package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"steamserverlauncher/handler"
	"steamserverlauncher/sourcequery"
	"syscall"

	"github.com/cristalhq/aconfig"
)

type Configuration struct {
	TemplatesPattern string `default:"templates/*.html" env:"TEMPLATES_PATTERN"`
	Address          string `default:"127.0.0.1:8000" env:"ADDRESS"`
	ServerName       string `default:"Default Source Server" env:"SERVER_NAME"`
	ServerAddress    string `required:"true" env:"SERVER_ADDRESS"`
}

var GetConfiguration func() *Configuration

func init() {
	var config Configuration

	configLoader := aconfig.LoaderFor(
		&config, aconfig.Config{EnvPrefix: "LAUNCHER"},
	)
	if err := configLoader.Load(); err != nil {
		log.Fatalf("configLoader.Load: %s", err)
	}

	GetConfiguration = func() *Configuration {
		return &config
	}
}

func main() {
	client, err := sourcequery.New(
		GetConfiguration().ServerAddress,
	)
	if err != nil {
		log.Println("sourcequery.New:", err)
		return
	}

	handler, err := handler.New(
		client,
		GetConfiguration().TemplatesPattern,
		GetConfiguration().ServerAddress,
		GetConfiguration().ServerName,
	)
	if err != nil {
		log.Println("handler.New:", err)
		return
	}

	var exitChannel = make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting server at", GetConfiguration().Address)

		if err := http.ListenAndServe(
			GetConfiguration().Address, handler.RegisterRoutes(),
		); !errors.Is(err, http.ErrServerClosed) {
			log.Println("http.ListenAndServe:", err)
		}
	}()

	<-exitChannel

	if err := client.Close(); err != nil {
		log.Println("client.Close:", err)
	}
}
