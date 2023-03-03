package v1

import (
	"dev/lamoda_test/internal/entity"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strconv"
)

// reserve
// @Summary Reserve products
// @Tags Products
// @Description Reserve some products
// @Accept  json
// @Produce  json
// @Param input body entity.Ids true "products id"
// @Success 200
// @Failure 500
// @Router /v1/reserve [POST]
func (h *Handler) reserve(w http.ResponseWriter, req bunrouter.Request) error {
	body := req.Body
	defer req.Body.Close()

	var product entity.Ids
	if err := json.NewDecoder(body).Decode(&product); err != nil {
		return err
	}

	err := h.services.Reserve(product.Ids)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	return h.responseJSON(w, req, http.StatusOK, "product(s) was reserved")
}

// reserveRelease
// @Summary Release products
// @Tags Products
// @Description Release some products
// @Accept  json
// @Produce  json
// @Param input body entity.Ids true "products id"
// @Success 200
// @Failure 500
// @Router /v1/release [POST]
func (h *Handler) reserveRelease(w http.ResponseWriter, req bunrouter.Request) error {
	body := req.Body
	defer req.Body.Close()

	var product entity.Ids
	if err := json.NewDecoder(body).Decode(&product); err != nil {
		return err
	}

	err := h.services.ReserveRelease(product.Ids)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	return h.responseJSON(w, req, http.StatusOK, "product(s) was unreserved")
}

// amount
// @Summary Take amounts products
// @Tags Products
// @Description Take amounts products
// @Accept  json
// @Produce  json
// @Param storage path string true "storage id"
// @Success 200 {array} entity.Products
// @Failure 500
// @Router /v1/amount/{storage} [GET]
func (h *Handler) amount(w http.ResponseWriter, req bunrouter.Request) error {
	s := req.Params().ByName("storage")

	storage, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	result, err := h.services.GetAmount(storage)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	return h.responseJSON(w, req, 200, result)
}
