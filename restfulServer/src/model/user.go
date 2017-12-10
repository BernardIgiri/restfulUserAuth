package model

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Security Configuration
const (
	bcryptComplexity = 14
	saltLength       = 32
)

// UserStatus account status
type UserStatus int8

// User account status
const (
	UserUnverified UserStatus = iota
	UserPasswordReset
	UserActive
	UserSuspended
	UserDisabled
	UserDeleted
)

// UserLoginSessionKey session key for user login
const UserLoginSessionKey = "userLogin"

// Preferences of website users
type Preferences struct {
	SendSecurityAlerts bool
	SendNewsLetter     bool
	Enable2fa          bool
}

// User users of the website
type User struct {
	Login       string
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber string

	Password string
	Salt     string
	Version  int8

	LastLogin     int32
	LoginFailures int8
	LastLoginFail int32

	Status UserStatus

	Preferences
}

// NewUser creates new user
func NewUser(login, firstname, lastname, email, phonenumber, password string,
	sendSecurityAlerts, sendNewsLetter, enable2fa bool) (user User, err error) {
	user = User{
		Login:       login,
		Firstname:   firstname,
		Lastname:    lastname,
		Email:       email,
		PhoneNumber: phonenumber,
		Status:      UserUnverified,
		Preferences: Preferences{
			SendSecurityAlerts: sendSecurityAlerts,
			SendNewsLetter:     sendNewsLetter,
			Enable2fa:          enable2fa,
		},
	}
	err = user.SetPassword(password)
	return
}

// Authenticate checks user password
func (u *User) Authenticate(plaintextPassword string) (bool, error) {
	if u.Version != 1 {
		return false, errors.New("Invalid version")
	}
	hashBytes, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return false, err
	}
	saltBytes, err := base64.StdEncoding.DecodeString(u.Salt)
	if err != nil {
		return false, err
	}
	saltedPassword := append([]byte(plaintextPassword), saltBytes...)
	err = bcrypt.CompareHashAndPassword(hashBytes, saltedPassword)
	return err == nil, nil
}

// SetPassword sets user password
func (u *User) SetPassword(plaintextPassword string) (err error) {
	saltBytes := make([]byte, saltLength)
	_, err = io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		return
	}
	saltedPassword := append([]byte(plaintextPassword), saltBytes...)
	hashBytes, err := bcrypt.GenerateFromPassword(saltedPassword, bcryptComplexity)
	if err != nil {
		return
	}
	u.Password = base64.StdEncoding.EncodeToString(hashBytes)
	u.Salt = base64.StdEncoding.EncodeToString(saltBytes)
	u.Version = 1
	return
}
