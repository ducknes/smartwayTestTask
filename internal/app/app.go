package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"smartwayTestTAsk/internal/config"
	"smartwayTestTAsk/internal/database"
	"smartwayTestTAsk/internal/handler"
	"smartwayTestTAsk/internal/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Start() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("env")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("hello User")); err != nil {
			log.Println("fail write")
		}
		log.Println("handle success")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	cfg := config.NewStorageConfig()

	postgreSQLClient, errPSQL := database.NewClient(context.TODO(), cfg)
	if errPSQL != nil {
		panic(errPSQL)
	}

	repos := database.NewRepository(postgreSQLClient)
	services := service.NewService(&repos)
	handlers := handler.NewHandler(&services)
	handlers.Register(router)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("Server is running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Stopping server")
}
