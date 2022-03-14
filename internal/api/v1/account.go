package v1

import (
	"encoding/json"
	"net/http"
	"ted/internal/api/v1/dto"
	"ted/internal/model"
	"ted/internal/repository"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type Account struct {
	accountRepo repository.Account
	lg          *log.Entry
	renderer    *render.Render
}

func NewAccountController(accountRepo repository.Account, lg *log.Entry, renderer *render.Render) *Account {
	return &Account{
		accountRepo: accountRepo,
		lg:          lg,
		renderer:    renderer,
	}
}

func (a *Account) AddAccount(writer http.ResponseWriter, request *http.Request) {
	var data model.Account
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	if data.Owner == "" {
		err = a.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	addedAccount, err := a.accountRepo.AddAccount(request.Context(), data)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(addedAccount)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	return
}

func (a *Account) IsExist(writer http.ResponseWriter, request *http.Request) {
	var data model.Account
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	if data.ID == "" {
		err = a.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	ok, err := a.accountRepo.IsExist(request.Context(), data.ID)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	var reply = dto.AccountExistanceDTO{
		Exist: ok,
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(reply)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	return
}

func (a *Account) GetBalanceByID(writer http.ResponseWriter, request *http.Request) {
	var data model.Account
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	if data.ID == "" {
		err = a.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	balance, err := a.accountRepo.GetBalanceByID(request.Context(), data.ID)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	var reply = dto.AccountBalanceDTO{
		Balance: balance,
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(reply)
	if err != nil {
		a.lg.Error(err)
		err = a.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			a.lg.Error(err)
			return
		}
		return
	}
	return
}
