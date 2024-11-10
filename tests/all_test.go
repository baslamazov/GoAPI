package tests

import (
	"GoAPI/protos/gen/proto"
	"GoAPI/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

func TestLogin_PositiveCase(t *testing.T) {
	st := suite.New(t)

	email := gofakeit.Email()
	pass := gofakeit.Password(true, true, true, true, false, 8)

	loginRequest := &proto.LoginRequest{
		Login:    email,
		Password: pass,
		AppId:    1,
	}

	//
	var callOptions []grpc.CallOption
	callOptions = append(callOptions, grpc.WaitForReady(true))

	response, err := st.AuthClient.Login(st.Context, loginRequest, callOptions...)

	require.NoError(t, err)
	assert.NotEmpty(t, response.GetToken())

	// TODO: Добавить проврку токена и времени его жизни
}
