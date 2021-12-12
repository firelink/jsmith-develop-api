package user

import (
	"encoding/json"
	"net/http"

	"github.com/firelink/api.jsmith-develop.com/response"
	"github.com/gorilla/mux"
)

func GetUserTroller(repo User) *UserTroller {
	return &UserTroller{repo: repo}
}

type UserTroller struct {
	repo User
}

func (user *UserTroller) Route(m *mux.Router) {
	//Failures
	m.NotFoundHandler = user
	//What we catch for
	m.HandleFunc("/", user.Create).Methods("POST")
	m.HandleFunc("/{id}/", user.Get).Methods("GET")
	m.HandleFunc("/{id}/", user.Update).Methods("PUT")
	m.HandleFunc("/{id}/", user.Delete).Methods("DELETE")
}

func (user *UserTroller) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	id := mux.Vars(r)["id"]

	userData, err := user.repo.GetUserByID(id)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = userData
	} else {
		finalModel.Items = userData
	}

	json.NewEncoder(w).Encode(finalModel)
}

func (user *UserTroller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Decode
	newUser := UserManager{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(newUser)

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	err := user.repo.SaveUser(newUser)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = ""
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)
}

func (user *UserTroller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Decode
	decoder := json.NewDecoder(r.Body)
	tempModel := UserManager{}
	decoder.Decode(tempModel)

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	id := mux.Vars(r)["id"]

	err := user.repo.UpdateUser(id, tempModel)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = ""
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)
}

func (user *UserTroller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	id := mux.Vars(r)["id"]

	err := user.repo.DeleteUser(id)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Message = "Deleted"
		finalModel.Items = nil
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)
}

func (user *UserTroller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	finalModel := response.BaseResponse{
		Status:  "error",
		Code:    404,
		Message: "Not a valid API call",
	}

	json.NewEncoder(w).Encode(finalModel)
}
