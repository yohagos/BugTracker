package databases

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists func
func CheckUserExists(username string) bool {
	if err := UserCollection.FindOne(ctx, bson.M{"email": username}); err != nil {
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

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)

	for k, v := range result {
		if k == "password" {
			userHash = fmt.Sprintf("%v", v)
			break
		}
	}

	fmt.Println("\n", userHash)
	fmt.Println(string(hash))

	err := bcrypt.CompareHashAndPassword([]byte(userHash), []byte(hash))
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
