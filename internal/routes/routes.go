package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"van-life-api/internal/middlewares"
	db "van-life-api/internal/storage"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	myDir := http.Dir("public")
	myHandler := http.FileServer(myDir)

	mux.Handle("/", myHandler)
	mux.HandleFunc("/vans/", middlewares.LoggerMiddleware(getVans))
	mux.HandleFunc("/van/", middlewares.LoggerMiddleware(getVan))
	return mux
}


func getVans(w http.ResponseWriter, r *http.Request) {
	vans := db.LocalStrogaeInstance.GetAll()
	vansJson, err := json.Marshal(vans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(vansJson)
}

func getVan(w http.ResponseWriter, r *http.Request) {
	buff := strings.Split(r.URL.Path, "/")
	var part []string
	for _, v := range buff {
		if v != "" {
			part = append(part, v)
		}
	}
	if len(part) < 2 {
		http.NotFound(w, r)
		return
	}
	vanId, err := strconv.Atoi(part[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	van, err := db.LocalStrogaeInstance.Get(vanId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vanJson, err := json.Marshal(van)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(vanJson)
}
