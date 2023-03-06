package app

import (
	"dev/lamoda_test/internal/config"
	"dev/lamoda_test/internal/controller"
	"dev/lamoda_test/internal/driver"
	"dev/lamoda_test/internal/repository"
	"dev/lamoda_test/internal/server"
	"dev/lamoda_test/internal/service"
	"github.com/rs/zerolog/log"
	"net/rpc"
)

func Run(cfg config.Config) error {
	db, err := driver.NewPostgresPool(cfg.PSQL)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
		log.Info().Msg("Connection with db close")
	}()

	repo := repository.NewStockPostgres(db)
	services := service.NewStock(repo)
	api := controller.New(services)

	serverRpc := rpc.NewServer()
	serverRpc.Register(api)

	srv := new(server.Listener)

	srv.New(cfg.Server.Port, serverRpc)

	return nil
}
