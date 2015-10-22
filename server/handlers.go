package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func NewAccount(w http.ResponseWriter, r *http.Request) {
	var ac Account

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &ac); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			panic(err)
		}
	}

	ac = RepoCreateAccount(ac)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}
