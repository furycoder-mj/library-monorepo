package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"epam.com/web-services/library-management/login-service/authentication"
	"epam.com/web-services/library-management/login-service/authorization"
	"epam.com/web-services/library-management/login-service/models"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userCredential models.UserCredential
		err := json.NewDecoder(r.Body).Decode(&userCredential)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// fmt.Println(authentication.HashAndSalt([]byte(userCredential.Password)))
		if !authentication.IsUserAuthenticated(userCredential) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		validToken, err := authorization.CreateToken(uint64(userCredential.Id))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("JWT-Token", validToken)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func SetupRoutes() {
	loginHandler := http.HandlerFunc(handleLogin)
	http.Handle("/login", loginHandler)
}
