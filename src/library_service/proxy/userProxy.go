package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"epam.com/web-services/library-management/library-service/client"
	"epam.com/web-services/library-management/library-service/config"

	"epam.com/web-services/library-management/library-service/models"
)

// Get all users from `users` table using `users-service`
func GetUsers() (uint, []models.UserDTO, error) {
	response, err := client.GetRequest(config.Config.UsersURI, "")
	if err != nil {
		return uint(response.StatusCode), nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), nil, err
	}
	users := make([]models.UserDTO, 0)
	unmarshallErr := json.Unmarshal(body, &users)
	if unmarshallErr != nil {
		return uint(response.StatusCode), nil, unmarshallErr
	}
	return uint(response.StatusCode), users, nil
}

// Get user with specific id from `users` table using `users-service`
func GetUser(id int64) (uint, models.UserDTO, error) {
	response, err := client.GetRequest(config.Config.UsersURI, fmt.Sprint(id))
	if err != nil {
		return uint(response.StatusCode), models.UserDTO{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), models.UserDTO{}, err
	}
	user := models.UserDTO{}
	unmarshallErr := json.Unmarshal(body, &user)
	if unmarshallErr != nil {
		return uint(response.StatusCode), models.UserDTO{}, unmarshallErr
	}
	return uint(response.StatusCode), user, nil
}

// Delete user with specific id from `users` table using `users-service`
func DeleteUser(id int64) (uint, string, error) {
	response, err := client.DeleteRequest(config.Config.UsersURI, fmt.Sprint(id))
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	return uint(response.StatusCode), string(body), nil
}

// Add user to `users` table using `users-service`
func AddUser(user models.UserDTO) (uint, string, error) {
	serializedUser, marshallErr := json.Marshal(user)
	if marshallErr != nil {
		return uint(http.StatusInternalServerError), "0", marshallErr
	}
	response, err := client.PostRequest(config.Config.UsersURI, "", bytes.NewReader(serializedUser))
	if err != nil {
		return uint(http.StatusInternalServerError), "0", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	return uint(response.StatusCode), string(body), nil
}

// Update user with specific id from `users` table using `users-service`
func UpdateUser(id int64, user models.UserDTO) (uint, string, error) {
	serializedUser, marshallErr := json.Marshal(user)
	if marshallErr != nil {
		return uint(http.StatusInternalServerError), "0", marshallErr
	}
	response, err := client.PutRequest(config.Config.UsersURI, fmt.Sprint(id), bytes.NewReader(serializedUser))
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	return uint(response.StatusCode), string(body), nil
}
