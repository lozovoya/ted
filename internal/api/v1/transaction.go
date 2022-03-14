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

type Transaction struct {
	transactionRepo repository.Transaction
	lg              *log.Entry
	renderer        *render.Render
}

func NewTransactionController(transactionRepo repository.Transaction, lg *log.Entry, renderer *render.Render) *Transaction {
	return &Transaction{
		transactionRepo: transactionRepo,
		lg:              lg,
		renderer:        renderer,
	}
}

func (t *Transaction) Refill(writer http.ResponseWriter, request *http.Request) {
	var data model.Transaction
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	if (data.Dest == "") || (data.Amount == 0) {
		err = t.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	transaction, err := t.transactionRepo.Refill(request.Context(), data.Dest, data.Amount)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(transaction)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	return
}

func (t *Transaction) Transfer(writer http.ResponseWriter, request *http.Request) {
	var data model.Transaction
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	if (data.Source == "") || (data.Dest == "") || (data.Amount == 0) {
		err = t.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	transaction, err := t.transactionRepo.Transfer(request.Context(), data.Source, data.Dest, data.Amount)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(transaction)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	return
}

func (t *Transaction) Transactions(writer http.ResponseWriter, request *http.Request) {
	var data dto.TransactionsReqDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	if data.Account == "" {
		err = t.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	//var transactions = make([]model.Transaction, 0)
	transactions, err := t.transactionRepo.GetCurrentMonthTransactionsByAccountID(request.Context(), data.Account)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(transactions)
	if err != nil {
		t.lg.Error(err)
		err = t.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			t.lg.Error(err)
			return
		}
		return
	}
	return
}
