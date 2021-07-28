package authentication

import (
	"log"

	"epam.com/web-services/library-management/login-service/models"
	"epam.com/web-services/library-management/login-service/repository"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	// fmt.Printf("hashed pass:")
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func IsUserAuthenticated(userCredential models.UserCredential) bool {
	storedCredential, err := repository.GetCredential(userCredential.UserName)
	if err != nil {
		return false
	}
	// fmt.Println((*storedCredential).Password, userCredential.Password)
	return comparePasswords((*storedCredential).Password, []byte(userCredential.Password))
}
