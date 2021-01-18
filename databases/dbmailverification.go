package databases

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// CheckVerificationExists func
func CheckVerificationExists(mail string) bool {
	if err := VerificationCollection.FindOne(ctx, bson.M{"email": mail}); err != nil {
		return true
	}
	return false
}

// CreateNewVerificationProfile func
func CreateNewVerificationProfile(profile bson.D) {
	_, err := VerificationCollection.InsertOne(ctx, profile)
	if err != nil {
		log.Println(err)
	}
}

// GetVerificationKey func
func GetVerificationKey(mail string) string {
	var document bson.M
	if err := VerificationCollection.FindOne(ctx, bson.M{"mail": mail}).Decode(&document); err != nil {
		log.Println(err)
	}
	var key string
	for k, v := range document {
		if k == "generatedKey" {
			key = v.(string)
			break
		}
	}

	return key
}

// DeleteVerificationDocument func
func DeleteVerificationDocument(mail string) {
	_, err := VerificationCollection.DeleteOne(ctx, bson.M{"mail": mail})
	if err != nil {
		log.Println(err)
	}
}

// GetAllVerificationInformation func
func GetAllVerificationInformation(mail string) bson.M {
	var user bson.M
	if err := VerificationCollection.FindOne(ctx, bson.M{"email": mail}).Decode(&user); err != nil {
		log.Println(err)
		return nil
	}
	DeleteVerificationDocument(mail)
	return user
}
