package databases

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"../apperrors"

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
func AuthentificationUser(username, password string) error {
	var result bson.M
	if err := UserCollection.FindOne(ctx, bson.M{"email": username}).Decode(&result); err != nil {
		log.Println(err)
		return err
	}

	var userHash []byte

	for _, v := range result {
		if strings.Contains(v.(string), password) {
			userHash = v.([]byte)
			break
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(userHash), []byte(password))
	if err == nil {
		return nil
	}
	return apperrors.ErrorPasswordMismatch
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
