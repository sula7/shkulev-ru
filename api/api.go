package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	API struct {
		Handlers
	}
)

func StartServer() {
	api := API{}

	http.HandleFunc("/ping", api.ping)
	http.HandleFunc("/api/v1/validate", api.validate)
	http.HandleFunc("/api/v1/fix", api.fix)

	srv := &http.Server{
		Addr: ":1323",
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("server started on port", srv.Addr)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server graceful shutdown failed with error:%+v", err)
	}

	log.Println("server gracefully shutdown")
}
