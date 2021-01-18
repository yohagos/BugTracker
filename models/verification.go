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

func (verif *UserVerification) setUserVerificationEmail(str string) {
	verif.Email = str
}

func (verif *UserVerification) setUserVerificationName(str string) {
	verif.Name = str
}

func (verif *UserVerification) setUserVerificationLastname(str string) {
	verif.Lastname = str
}

func (verif *UserVerification) setUserVerificationGeneratedKey(str string) {
	verif.GeneratedKey = str
}

func (verif *UserVerification) setUserVerificationPassword(str string) {
	verif.Password = str
}

func (verif *UserVerification) setUserVerificationVerified(bo bool) {
	verif.Verified = bo
}

// CreateVerificationProfile func
func (verif *UserVerification) CreateVerificationProfile() {
	ok := databases.CheckVerificationExists(verif.Email)
	if ok {
		return
	}

	genKey := utils.GenerateVerificationKey()

	pwd := verif.GetVerifiedUserPassword()
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	if err != nil {
		return
	}

	/* time := utils.CreateTimeStamp() */
	verificationDocument := bson.D{
		{Key: "name", Value: verif.Name},
		{Key: "lastname", Value: verif.Lastname},
		{Key: "email", Value: verif.Email},
		{Key: "password", Value: string(hash)},
		{Key: "genratedKey", Value: genKey},
		{Key: "verified", Value: false},
	}
	databases.CreateNewVerificationProfile(verificationDocument)
	return
}
