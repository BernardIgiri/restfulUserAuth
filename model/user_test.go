package model_test

import (
	"testing"

	"github.com/bernardigiri/restfulUserAuth/model"
	"github.com/stretchr/testify/assert"
)

func TestSetPassword(t *testing.T) {
	user := model.User{}
	const password = "somePasswordValue123#.*"
	err := user.SetPassword(password)

	assert.Nil(t, err)
	assert.NotEqual(t, password, user.Password)
	assert.NotEqual(t, password, user.Salt)
	assert.NotEqual(t, user.Password, user.Salt)
	assert.NotEqual(t, len(password), len(user.Password))
	assert.NotZero(t, len(user.Password))
	assert.NotZero(t, len(user.Salt))

	oldPass := user.Password
	oldSalt := user.Salt
	err = user.SetPassword(password)

	assert.Nil(t, err)
	assert.NotEqual(t, oldPass, user.Password)
	assert.NotEqual(t, oldSalt, user.Salt)
}

func TestAuthenticate(t *testing.T) {
	user := model.User{
		Password: "JDJhJDE0JEZhVUsvZi5tcHhjQVJqQUhZYTIzVE9waVM1a0R5c1F5bjBTZmNVYVhwNHpEbjUwbDN2VEFL",
		Salt:     "WBIeCbvLdFm2f8zzO1MDic51teKxgC3ZiResyxpDKaI=",
		Version:  1,
	}
	const password = "somePasswordValue123#.*"

	match, err := user.Authenticate(password)

	assert.Nil(t, err)
	assert.True(t, match)
}

func TestSetAndAuthenticate(t *testing.T) {
	user := model.User{}
	const password = "SomeOtherPas(wo*d_=?9)"

	err := user.SetPassword(password)
	assert.Nil(t, err)
	match, err := user.Authenticate(password)

	assert.Nil(t, err)
	assert.True(t, match)
}
