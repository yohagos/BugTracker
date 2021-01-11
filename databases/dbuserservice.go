package databases

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists func
func CheckUserExists(username string) bool {
	var u bson.M
	if err := UserCollection.FindOne(ctx, bson.M{"email": username}).Decode(&u); err != nil {
		return false
	}
	return true
}

// CreateNewUser func
func CreateNewUser(user bson.D) {
	_, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
}

// AuthentificationUser func
func AuthentificationUser(bcryptCost int, username, password string) error {
	var result bson.M
	if err := UserCollection.FindOne(ctx, bson.M{"email": username}).Decode(&result); err != nil {
		log.Println(err)
		return err
	}

	var userHash string
	for k, v := range result {
		if k == "password" {
			userHash = fmt.Sprintf("%v", v)
			break
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(userHash), []byte(password))
	if err == nil {
		return nil
	}
	return bcrypt.ErrMismatchedHashAndPassword
}

// GetAllUserInformations func
func GetAllUserInformations(email string) (bson.M, error) {
	var result bson.M
	if err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&result); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
