package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type People struct {
	ID       string `json:"id"`
	LastName string `json:"title"`
}

var DATA []People

func Init()  {
	DATA= make([]People, 0)
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(DATA)
}

func CreateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj People
	_ = json.NewDecoder(r.Body).Decode(&obj)

	for _, item := range DATA {
		if item.ID == obj.ID {
			res := map[string]string{"error": "Object with id " + obj.ID + " already exists"}
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	DATA = append(DATA, obj)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&obj)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range DATA {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	res := map[string]string{"error": "Object with id " + params["id"] + " not found"}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range DATA {
		if item.ID == params["id"] {
			DATA = append(DATA[:index], DATA[index+1:]...)
			var obj People
			_ = json.NewDecoder(r.Body).Decode(&obj)
			obj.ID = params["id"]
			DATA = append(DATA, obj)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&obj)
			return
		}
	}
	res := map[string]string{"error": "Object with id " + params["id"] + " not found"}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range DATA {
		if item.ID == params["id"] {
			DATA = append(DATA[:index], DATA[index+1:]...)
			res := map[string]string{"msg": "Object with id " + params["id"] + " successfully deleted"}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	res := map[string]string{"error": "Object with id " + params["id"] + " not found"}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}
