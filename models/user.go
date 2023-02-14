package models

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	hashCost = 4

	//nolint:lll // Email regex so cannot be split
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	//nolint:gosec // This is just a regex
	passwordRegex = `^[a-zA-Z]\w{3,14}$`
)

var (
	ErrValidationFailed = errors.New("non valid data")
)

type User struct {
	gorm.Model
	Email    string `gorm:"UNIQUE;NOT NULL" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SetPassword will accept a string password and attempt to hash and salt it. Providing the
// hash is successful, the Password field of the User will be updated.
func (u *User) SetPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		return err
	}

	u.Password = string(pass)
	return nil
}

// ComparePassword will evaluate the provided password witht the one stored against the
// user by hashing it first. If the passwords match, true will be returned (with a nil,
// error) otherswise, false will be returned.
//
// If an error is returned, it means that the passwords weren't able to be compared, and so
// the result cannot be trusted.
func (u *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, err
}

// Validate used to validate the data in this struct, needs to be ran before password
// hashing.
func (u *User) Validate() ([]string, error) {
	result := []string{}

	emailMatch, err := regexp.MatchString(emailRegex, u.Email)
	if err != nil {
		return nil, err
	}

	if !emailMatch {
		result = append(result, "email")
	}

	if u.Username == "" {
		result = append(result, "username")
	}

	passMatch, err := regexp.MatchString(passwordRegex, u.Password)
	if err != nil {
		return nil, err
	}

	if !passMatch {
		result = append(result, "password")
	}

	return result, nil
}
