package tests

import (
	"GoDB/sqlc"
	"GoDB/tests/utils"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type UserData struct {
	Email        string
	Username     string
	PasswordHash string
	FullName     string
	Role         int32
}

func ToUser(UserData UserData) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Email:        UserData.Email,
		Username:     UserData.Username,
		Passwordhash: UserData.PasswordHash,
		Fullname:     UserData.FullName,
		Createdate:   time.Now(),
		Role:         UserData.Role,
	}
}

func TestCreateUser(test *testing.T) {
	test.Log("Account Test - Starting CreateAccount test...")
	arg := utils.ToUser(utils.GenerateRandomUserData())
	fmt.Print(arg)
	account, err := testQueries.CreateUser(context.Background(), arg)
	// Assuming role is between 1 and 3
	require.NoError(test, err)
	require.NotEmpty(test, account)
	require.Equal(test, arg.Email, account.Email)
	require.Equal(test, arg.Username, account.Username)
	require.Equal(test, arg.Passwordhash, account.Passwordhash)
	require.Equal(test, arg.Fullname, account.Fullname)
	require.Equal(test, arg.Role, account.Role)

	require.NotZero(test, account.ID)
	require.NotZero(test, account.Createdate)

}
func TestGetUser(test *testing.T) {
	test.Log("Account Test - Starting GetUser test...")

	some_user_data := UserData{
		Email:        "jawhardjebbi@gmail.com",
		Username:     "jawhardjebbi",
		PasswordHash: "pushforward123",
		FullName:     "Jawhar Djebbi",
		Role:         1}

	User := ToUser(some_user_data)
	account, err := testQueries.CreateUser(context.Background(), User)
	require.NoError(test, err)

	some_user, err := testQueries.GetUser(context.Background(), account.ID)

	require.NoError(test, err)
	require.NotEmpty(test, some_user)
	require.Equal(test, account.ID, some_user.ID)
	require.Equal(test, account.Email, some_user.Email)
	require.Equal(test, account.Username, some_user.Username)
	require.Equal(test, account.Passwordhash, some_user.Passwordhash)
	require.Equal(test, account.Fullname, some_user.Fullname)
	require.Equal(test, account.Role, some_user.Role)
	require.WithinDuration(test, account.Createdate, some_user.Createdate, time.Second)

}

func TestDeleteUser(test *testing.T) {
	test.Log("Account Test - Starting DeleteAccount test...")

	some_user_data := UserData{
		Email:        "jawhardjebbi@gmail.com",
		Username:     "jawhardjebbi",
		PasswordHash: "pushforward123",
		FullName:     "Jawhar Djebbi",
		Role:         1,
	}

	User := ToUser(some_user_data)

	account, err := testQueries.CreateUser(context.Background(), User)
	require.NoError(test, err)
	err = testQueries.DeleteUser(context.Background(), account.ID)
	require.NoError(test, err)
	account, err = testQueries.GetUser(context.Background(), account.ID)
	require.Error(test, err)
	require.Empty(test, account)

}
