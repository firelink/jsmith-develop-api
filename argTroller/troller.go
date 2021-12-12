package argTroller

import (
	"encoding/json"
	"net/http"

	"github.com/firelink/api.jsmith-develop.com/argModel"
	"github.com/firelink/api.jsmith-develop.com/argStore"
	"github.com/firelink/api.jsmith-develop.com/response"
	"github.com/gorilla/mux"
)

type ArgTroller interface {
	Route(m *mux.Router)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func GetRouter(entity argModel.Entity, cache argStore.ArgDB) *GenericTroller {
	return &GenericTroller{entity: entity, cache: cache}
}

type GenericTroller struct {
	entity argModel.Entity
	cache  argStore.ArgDB
}

func (gen *GenericTroller) Route(m *mux.Router) {
	//Failures
	m.NotFoundHandler = gen
	//What we catch for
	m.HandleFunc("/", gen.Create).Methods("POST")
	m.HandleFunc("/{id}/", gen.Get).Methods("GET")
	m.HandleFunc("/{id}/", gen.Update).Methods("PUT")
	m.HandleFunc("/{id}/", gen.Delete).Methods("DELETE")
}

func (gen *GenericTroller) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	id := mux.Vars(r)["id"]

	tempModel := gen.entity.New()
	err := gen.cache.Get("key", id, tempModel)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = tempModel
	} else {
		finalModel.Items = tempModel
	}

	json.NewEncoder(w).Encode(finalModel)
}

func (gen *GenericTroller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	model := gen.entity.New()

	//Decode
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(model)

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	result, err := gen.cache.Write(model)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = result
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)
}

func (gen *GenericTroller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Decode
	decoder := json.NewDecoder(r.Body)
	tempModel := gen.entity.New()
	decoder.Decode(tempModel)

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	id := mux.Vars(r)["id"]
	err := gen.cache.Update(id, tempModel)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Message = "Updated"
		getModel := gen.entity.New()
		gen.cache.Get("key", id, getModel)
		finalModel.Items = getModel
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)
}

func (gen *GenericTroller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Make model
	finalModel := response.BaseResponse{Status: "ok", Code: 200}
	id := mux.Vars(r)["id"]

	err := gen.cache.Delete(id)

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

func (base *GenericTroller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	finalModel := response.BaseResponse{
		Status:  "error",
		Code:    404,
		Message: "Not a valid API call",
	}

	json.NewEncoder(w).Encode(finalModel)
}
