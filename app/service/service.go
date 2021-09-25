package service

import (
	"context"
	"log"
	"milosavljevicoa/gradebook/app/registry"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve(initializeServer func(r *gin.Engine), reg registry.Registration, port string) error {
	if err := registry.RegistryService(reg); err != nil {
		return err
	}

	r := gin.New()
	initializeServer(r)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		srv.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	if err := registry.ShutdownService(reg.ServiceURL); err != nil {
		return err
	}

	return nil
}
