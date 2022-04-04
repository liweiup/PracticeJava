package main

import (
	"bird/micro001/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  120 * time.Second,
		IdleTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("recieved terminate,graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	// http.ListenAndServe(":8081", sm)
}
