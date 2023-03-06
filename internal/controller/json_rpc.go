package controller

import (
	"dev/lamoda_test/internal/model"
	"dev/lamoda_test/internal/service"
)

// RpcApi Api structure
type RpcApi struct {
	services service.Stocker
}

func New(services service.Stocker) *RpcApi {
	return &RpcApi{services: services}
}

func (r *RpcApi) Reserve(args *model.Args, result *string) error {
	err := r.services.Reserve(args.Ids)
	if err != nil {
		return err
	}
	*result = "Success"
	return nil
}

func (r *RpcApi) ReserveRelease(args *model.Args, result *string) error {
	err := r.services.ReserveRelease(args.Ids)
	if err != nil {
		return err
	}
	*result = "Success"
	return nil
}

func (r *RpcApi) GetAmount(args *model.Args, result *[]model.Products) error {
	amount, err := r.services.GetAmount(args.Store)
	if err != nil {
		return err
	}
	*result = amount
	return nil
}
