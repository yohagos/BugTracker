package databases

import (
	"go.mongodb.org/mongo-driver/bson"
)

// CheckVerificationExists func
func CheckVerificationExists(mail string) bool {
	if err := VerificationCollection.FindOne(ctx, bson.M{"email": mail}); err != nil {
		return true
	}
	return false
}

func CreateNewVerificationProfile(profile bson.D) {

}
