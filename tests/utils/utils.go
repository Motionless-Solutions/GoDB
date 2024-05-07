package utils

import (
	"math/rand"
	"time"

	"GoDB/sqlc"

	"github.com/bxcodec/faker/v3"
)

type RandomUserData struct {
	Email        string `faker:"email"`
	Username     string `faker:"username"`
	PasswordHash string `faker:"password"`
	FullName     string `faker:"name"`
	Role         int32
}

func ToUser(randomUserData *RandomUserData) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Email:        randomUserData.Email,
		Username:     randomUserData.Username,
		Passwordhash: randomUserData.PasswordHash,
		Fullname:     randomUserData.FullName,
		Createdate:   time.Now(),
		Role:         randomUserData.Role,
	}
}

func GenerateRandomUserData() *RandomUserData {
	var randomUserData RandomUserData
	err := faker.FakeData(&randomUserData)
	if err != nil {
		panic(err)
	}

	randomUserData.Role = int32(rand.Intn(2))

	return &randomUserData
}
