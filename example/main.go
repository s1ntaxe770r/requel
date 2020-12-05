package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/requel"
)

func main() {
	router := mux.NewRouter()
	router.Use(requel.LogReq)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("consider leaving a star"))
	})

	http.ListenAndServe(":5050", router)

}
