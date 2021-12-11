package account_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/account/mocks"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	newAccountReq := account.AccountRegisterRequest{
		Email:     "user@example.com",
		Password:  "secret",
		FirstName: "test",
		LastName:  "test",
	}

	newAccountRes := account.Account{
		ID:        1,
		Email:     newAccountReq.Email,
		Password:  &newAccountReq.Password,
		FirstName: newAccountReq.FirstName,
		LastName:  newAccountReq.LastName,
		CreatedAt: time.Now(),
	}

	resp := response.Success(response.StatusCreated, newAccountRes)

	validate := validator.New()

	accountUsecase := new(mocks.AccountUsecase)
	accountUsecase.On("Register", mock.Anything, mock.AnythingOfType("account.AccountRegisterRequest")).Return(resp)

	newAccountRegisterRequestBuff, _ := json.Marshal(newAccountReq)

	accountHandler := account.AccountHandler{
		Validate: validate,
		Usecase:  accountUsecase,
	}

	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newAccountRegisterRequestBuff))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(accountHandler.Register)
	handler.ServeHTTP(recorder, r)

	rb := response.ResponseImpl{}
	if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.StatusCreated, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusCreated))
	assert.NotNil(t, rb.Data, "should not be nil")

	data, ok := rb.Data.(map[string]interface{})
	if !ok {
		t.Error("response data isn't a type of 'map[string]interface{}'")
		return
	}

	assert.Equal(t, newAccountRes.Email, data["email"], fmt.Sprintf("email should be '%s'", newAccountRes.Email))

	accountUsecase.AssertExpectations(t)
}
