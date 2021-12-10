package account

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
)

type AccountHandler struct {
	Validate *validator.Validate
	Usecase  AccountUsecase
}

func NewAccountHandler(router *mux.Router, validate *validator.Validate, usecase AccountUsecase) {
	handler := &AccountHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/accounts/registration", handler.Register).Methods(http.MethodPost)
}

func (handler *AccountHandler) Register(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var params AccountRegisterRequest

	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		res = response.Error(err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		res = response.Error(err)
		res.JSON(w)
		return
	}

	res = handler.Usecase.Register(ctx, params)
	res.JSON(w)

}
