package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

import "entityValidator.com/router"

import elastic "entityValidator.com/elastic"

func main() {
	fmt.Print("\033[H\033[2J") //clear terminal

	r := router.InitRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening on port", srv.Addr)

	_, err := elastic.NewClient(
		os.Getenv("ELASTICSEARCH_URL"), "",
	)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("ElasticSearch Client started successfully!")
	}

	go func() {
		printLogo()
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("FATAL: Failed to start server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("FATAL: Server forced to shudown: %v\n", err)
	}

}

func printLogo() {
	fmt.Println("")
	fmt.Println("  888                    888                            ")
	fmt.Println("  888   .op .ooooo.  ooo 888   .op  .ooooo.  oo. .oo.   ")
	fmt.Println("  888 .8P' d88   88b 888 888 .8P'  d88   88b 888P\"Y88b  ")
	fmt.Println("  888888.  888bod888 888 888888.   888bod888 888   888  ")
	fmt.Println("  888 `8b. 888    .o 888 888 `8b.  888    .o 888   888  ")
	fmt.Println("  888   `8b`Y8bod8P' 888 888   `8b.`Y8bod8P' 888   888  ")
	fmt.Println("")
	fmt.Println("  Validação de registros nacionais de profissionais cadastrados")
	fmt.Println("")
	fmt.Println("                                   made with ❤️by eduardo-moro")
	fmt.Println("")
}
