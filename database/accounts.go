package database

import (
	"context"
	uuid "github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mo2/server/model"
)

//already check the validation in controller
//if add a newAccount success, return account info
func AddAccount(newAccount model.AddAccount) (account model.Account, err error) {
	collection := GetCollection("accounts")
	//collection.Distinct()
	model := []mongo.IndexModel{
		{
			Keys:    bson.D{{"username", 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{"email", 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err = collection.Indexes().CreateMany(context.TODO(), model)
	if err != nil {
		log.Fatal(err)
	}
	//var account model.Account
	// find the maxId in mongoDB
	account.ID = GetMaxID("accounts")
	account.Email = newAccount.Email
	account.UserName = newAccount.UserName
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(newAccount.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	account.HashedPwd = string(hashedPwd)
	account.UUID, err = uuid.NewV4()
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = collection.InsertOne(context.TODO(), account)
	return
}

//verify an account
func VerifyAccount(info model.LoginAccount) (account model.Account, err error) {

	//first use username, then use email to verify
	//var
	collection := GetCollection("accounts")
	userNameOrEmail := info.UserNameOrEmail
	err = collection.FindOne(context.TODO(), bson.D{{"username", userNameOrEmail}}).Decode(&account)

	if err != nil {
		//then verify email
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			err = collection.FindOne(context.TODO(), bson.D{{"email", userNameOrEmail}}).Decode(&account)
			if err != nil {
				// no chance
				return
			}
		}
		if err != mongo.ErrNoDocuments {
			log.Fatal(err)
			return
		}

	}
	password := info.Password
	hashedPassword := account.HashedPwd
	//judge hash with hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return
	}
	return

}
