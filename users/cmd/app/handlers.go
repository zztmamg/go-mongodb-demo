package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// 获取所有用户
	users, err := app.users.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(users)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("查询所有用户完毕")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
