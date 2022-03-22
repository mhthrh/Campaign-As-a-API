package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/Campaign/Utilitys/ConfigUtil"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/DbPool"
	"github.com/mhthrh/Campaign/Utilitys/LogUtil"
	"github.com/mhthrh/Campaign/View"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	defer func() {
		if err := recover().(error); err != nil {
			fmt.Println(err)
		}
	}()
	//ConfigUtil.WriteConfig()
	cfg := ConfigUtil.ReadConfig("Files/ConfigCoded.json")
	if cfg == nil {
		log.Fatalln("Cant read Config, By.")
	}
	logger := LogUtil.New()
	sm := mux.NewRouter()
	db := DbPool.New(&DbPool.DbInfo{
		Host:            cfg.DB[0].Host,
		Port:            cfg.DB[0].Port,
		User:            cfg.DB[0].User.UserName,
		Pass:            cfg.DB[0].User.Password,
		Dbname:          cfg.DB[0].Dbname,
		Driver:          cfg.DB[0].Driver,
		ConnectionCount: 10,
		RefreshPeriod:   20,
	})
	if err := View.RunApiOnRouter(sm, logger, db, cfg); err != nil {
		panic(err)
	}

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port),
		Handler:      sm,
		ErrorLog:     log.New(LogUtil.LogrusErrorWriter{}, "", 0),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  180 * time.Second,
	}

	go func() {
		logger.Println("Starting server on  %s:%d", cfg.Server.IP, cfg.Server.Port)
		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	log.Println("Got signal:", <-c)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
