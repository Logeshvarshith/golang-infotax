package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "www.ivtlinfoview.com/infotax/infotax-backend/cmd/docs"
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/config"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/external/framework"

	"github.com/joho/godotenv"
)

// @title Info Tax
// @version 1.0
// @description User friendly tax deduction system
// @host localhost:5000
// @BasePath /
func main() {

	log.MakeLogger("./tmp/infotax.log", true)

	log.Logger.Println("Load app environment variables")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Logger.Fatalf("Failed to load environment variables. Reason: %v\n", err)
	}

	log.Logger.Println("Parse app configurations")
	conf := config.ParseConfig()

	router := framework.Handler(conf)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.AppPort),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("Error while initialzing server : %v\n", err.Error())
		}
	}()

	log.Logger.Printf("Server listening on port %v\n", server.Addr)

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	log.Logger.Printf("Shutting down server %v\n", server.Addr)

	if err := server.Shutdown(ctx); err != nil {
		log.Logger.Fatalf("Error while shutting down server : %v\n", err.Error())
	}

}
