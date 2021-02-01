package models

import (
	"../databases"
	"../utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// UserVerification struct
type UserVerification struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Lastname     string             `bson:"lastname,omitempty"`
	Email        string             `bson:"email,omitempty"`
	Password     string             `bson:"password,omitempty"`
	GeneratedKey string             `bson:"genratedKey,omitempty"`
	Verified     bool               `bson:"verified,omitempty"`
}

// GetVerifiedUserName func
func (verif *UserVerification) GetVerifiedUserName() string {
	return verif.Name
}

// GetVerifiedUserLastname func
func (verif *UserVerification) GetVerifiedUserLastname() string {
	return verif.Lastname
}

// GetVerifiedUserEmail func
func (verif *UserVerification) GetVerifiedUserEmail() string {
	return verif.Email
}

// GetVerifiedUserPassword func
func (verif *UserVerification) GetVerifiedUserPassword() string {
	return verif.Password
}

// GetVerifiedUserGeneratedKey func
func (verif *UserVerification) GetVerifiedUserGeneratedKey() string {
	return verif.GeneratedKey
}

// GetVerifiedUserVerified func
func (verif *UserVerification) GetVerifiedUserVerified() bool {
	return verif.Verified
}

// SetUserVerificationEmail func
func (verif *UserVerification) SetUserVerificationEmail(str string) {
	verif.Email = str
}

// SetUserVerificationName func
func (verif *UserVerification) SetUserVerificationName(str string) {
	verif.Name = str
}

// SetUserVerificationLastname func
func (verif *UserVerification) SetUserVerificationLastname(str string) {
	verif.Lastname = str
}

// SetUserVerificationGeneratedKey func
func (verif *UserVerification) SetUserVerificationGeneratedKey(str string) {
	verif.GeneratedKey = str
}

// SetUserVerificationPassword func
func (verif *UserVerification) SetUserVerificationPassword(str string) {
	verif.Password = str
}

// SetUserVerificationVerified func
func (verif *UserVerification) SetUserVerificationVerified(bo bool) {
	verif.Verified = bo
}

// CreateVerificationProfile func
func (verif *UserVerification) CreateVerificationProfile() {
	ok := databases.CheckVerificationExists(verif.Email)
	if !ok {
		return
	}

	pwd := verif.GetVerifiedUserPassword()
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	if err != nil {
		return
	}

	/* time := utils.CreateTimeStamp() */
	verificationDocument := bson.D{
		{Key: "name", Value: verif.GetVerifiedUserName()},
		{Key: "lastname", Value: verif.GetVerifiedUserLastname()},
		{Key: "email", Value: verif.GetVerifiedUserEmail()},
		{Key: "password", Value: string(hash)},
		{Key: "genratedKey", Value: verif.GetVerifiedUserGeneratedKey()},
		{Key: "verified", Value: verif.GetVerifiedUserVerified()},
	}
	databases.CreateNewVerificationProfile(verificationDocument)
	return
}

// CheckVerification func
func CheckVerification(key string) (string, bool) {
	return databases.GetVerificationKey(key)
}

// CreateNewUser func
func CreateNewUser(mail string) {
	document := databases.GetAllVerificationInformation(mail)

	user := BsonToUser(document)

	timestamp := utils.CreateTimeStamp()
	user.setUserCreatedAt(timestamp)
	user.setUserUpdatedAt(timestamp)
	user.CreateNewUser()

}
