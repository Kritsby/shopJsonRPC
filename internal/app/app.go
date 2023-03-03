package app

import (
	"context"
	"dev/lamoda_test/internal/config"
	v1 "dev/lamoda_test/internal/controller/v1"
	"dev/lamoda_test/internal/repository/postgresql"
	"dev/lamoda_test/internal/server"
	"dev/lamoda_test/internal/service"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg config.Config) error {
	db, err := postgresql.New(cfg.PSQL)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
		log.Info().Msg("Connection with db close")
	}()

	repo := postgresql.NewRepository(db)
	services := service.NewService(repo)
	handler := v1.New(services)

	srv := new(server.Server)
	defer func() {
		_ = srv.Shutdown(context.Background())
		log.Info().Msg("Server Stopped")
	}()

	go func() {
		if err = srv.Run(cfg.Server.Port, handler.InitRouter()); err != nil {
			log.Fatal().Msgf("error occurred while running the server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info().Msg("App Shutting Down")

	return nil
}
