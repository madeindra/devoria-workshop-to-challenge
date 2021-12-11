package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/middleware"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
)

type AccountHandler struct {
	Validate *validator.Validate
	Usecase  AccountUsecase
}

func NewAccountHandler(router *mux.Router, basicAuthMiddleware middleware.RouteMiddleware, bearerAuthMiddleware middleware.RouteMiddlewareBearer, validate *validator.Validate, usecase AccountUsecase) {
	handler := &AccountHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/accounts/registration", basicAuthMiddleware.Verify(handler.Register)).Methods(http.MethodPost)
	router.HandleFunc("/v1/accounts/login", basicAuthMiddleware.Verify(handler.Login)).Methods(http.MethodPost)

	router.HandleFunc("/v1/accounts/{id:[0-9]+}", bearerAuthMiddleware.VerifyBearer(handler.GetAccount)).Methods(http.MethodGet)
}

func (handler *AccountHandler) Register(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var params AccountRegisterRequest

	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.Usecase.Register(ctx, params)
	res.JSON(w)
}

func (handler *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var params AccountLoginRequest

	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.Usecase.Login(ctx, params)
	res.JSON(w)
}

func (handler *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	res := handler.Usecase.GetAccount(ctx, id)
	res.JSON(w)
}
