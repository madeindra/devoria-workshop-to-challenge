package article

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
)

type ArticleHandler struct {
	Validate *validator.Validate
	Usecase  ArticleUsecase
}

func NewArticleHandler(router *mux.Router, validate *validator.Validate, usecase ArticleUsecase) {
	handler := &ArticleHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/articles", handler.GetAllArticles).Methods(http.MethodGet)
	router.HandleFunc("/v1/articles/{id:[0-9]+}", handler.GetOneArticle).Methods(http.MethodGet)
	router.HandleFunc("/v1/articles", handler.CreateArticle).Methods(http.MethodPost)
	router.HandleFunc("/v1/articles", handler.UpdateArticle).Methods(http.MethodPut)
}

func (handler *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	var ctx = r.Context()

	resp = handler.Usecase.GetAllArticles(ctx)
	resp.JSON(w)
}

func (handler *ArticleHandler) GetOneArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	res := handler.Usecase.GetOneArticle(ctx, id)
	res.JSON(w)
}

func (handler *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var params CreateArticleRequest

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

	res = handler.Usecase.CreateArticle(ctx, params)
	res.JSON(w)
}

func (handler *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var params UpdateArticleRequest

	ctx := r.Context()

	vars := mux.Vars(r)
	params.ID, _ = strconv.ParseInt(vars["id"], 10, 64)

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

	res = handler.Usecase.UpdateArticle(ctx, params)
	res.JSON(w)
}
