package controllers

import (
	"encoding/json"
	"httpServer/common"
	"httpServer/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = user.CheckPassword()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	common.WriteJSON(w, http.StatusOK, "login successful")

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = user.AddNew(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	common.WriteJSON(w, http.StatusOK, "register succesful")

}

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateInfoHandler(w http.ResponseWriter, r *http.Request) {

}
