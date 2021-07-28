package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"epam.com/web-services/library-management/library-service/authorization"
	"epam.com/web-services/library-management/library-service/models"
	"epam.com/web-services/library-management/library-service/proxy"
	"epam.com/web-services/library-management/library-service/repository"
)

type UserId int64

// Get Users http FuncHandler
func getUsers(w http.ResponseWriter, r *http.Request) {
	statusCode, usersList, err := proxy.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(usersList)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(int(statusCode))
}

// Get User http FuncHandler
func (u UserId) getUser(w http.ResponseWriter, r *http.Request) {
	statusCode, user, err := proxy.GetUser(int64(u))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	list, err := repository.GetAssignedBooksForUser(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i := range list {
		user.AssociatedBooks = append(user.AssociatedBooks, list[i].AssignedBook)
	}
	j, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(int(statusCode))
}

// Add Users http FuncHandler
func addUser(w http.ResponseWriter, r *http.Request) {
	_, authErr := authorization.ValidateToken(w, r)
	if authErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token:", authErr)
		return
	}
	var user models.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	statusCode, body, err := proxy.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(body))
	w.WriteHeader(int(statusCode))
}

// Update User http FuncHandler
func (u UserId) updateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	statusCode, body, err := proxy.UpdateUser(int64(u), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(body))
	w.WriteHeader(int(statusCode))
}

// Delete User http FuncHandler
func (u UserId) deleteUser(w http.ResponseWriter, r *http.Request) {
	_, authErr := authorization.ValidateToken(w, r)
	if authErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token:", authErr)
		return
	}
	_, err := repository.DeleteByUser(int64(u))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	statusCode, body, err := proxy.DeleteUser(int64(u))
	if err != nil {
		w.WriteHeader(int(statusCode))
		return
	}
	w.Write([]byte(body))
	w.WriteHeader(int(statusCode))
}
