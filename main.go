package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import router "entityValidator.com/router"

func main() {
	r := router.InitRouter()

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	go func() {
			log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("FATAL: Failed to start server: %v\n", err)
		}
	}()

	log.Println("Listening on port", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err:=srv.Shutdown(ctx); err != nil {
		log.Fatalf("FATAL: Server forced to shudown: %v\n", err)
	}
}


//made with ❤️ by eduardo-moro
