package v1

import (
	"encoding/json"
	"net/http"
	"ted/internal/model"
	"ted/internal/repository"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type User struct {
	userRepo repository.User
	lg       *log.Entry
	renderer *render.Render
}

func NewUserController(userRepo repository.User, lg *log.Entry, renderer *render.Render) *User {
	return &User{
		userRepo: userRepo,
		lg:       lg,
		renderer: renderer,
	}
}

func (u *User) AddUser(writer http.ResponseWriter, request *http.Request) {
	var data model.User
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		u.lg.Error(err)
		err = u.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	if (data.Name == "") || (data.Password == "") {
		err = u.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	addedUser, err := u.userRepo.AddUser(request.Context(), data)
	if err != nil {
		u.lg.Error(err)
		err = u.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(addedUser)
	if err != nil {
		u.lg.Error(err)
		err = u.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	return
}

func (u *User) GetUserByID(writer http.ResponseWriter, request *http.Request) {
	var data model.User
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		u.lg.Error(err)
		err = u.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Bad request"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	if data.ID == "" {
		err = u.renderer.JSON(writer, http.StatusBadRequest, map[string]string{"Error": "Mandatory field is empty"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	user, err := u.userRepo.GetUserByID(request.Context(), data.ID)
	if err != nil {
		u.lg.Error(err)
		err = u.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		u.lg.Error(err)
		err = u.renderer.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
		if err != nil {
			u.lg.Error(err)
			return
		}
		return
	}
	return
}
