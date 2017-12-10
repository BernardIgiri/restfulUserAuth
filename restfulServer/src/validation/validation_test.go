package validation_test

import (
	"testing"

	"github.com/bernardigiri/restfulUserAuth/validation"
	"github.com/stretchr/testify/assert"
)

func TestUnicodeWordValidator(t *testing.T) {
	assert.True(t, validation.UnicodeWordValidator()("Renee", nil))
	assert.True(t, validation.UnicodeWordValidator()("Reneé", nil))
	assert.True(t, validation.UnicodeWordValidator()("Rènéó", nil))
	assert.False(t, validation.UnicodeWordValidator()("Reneé ", nil))
	assert.False(t, validation.UnicodeWordValidator()("Renee ", nil))
	assert.False(t, validation.UnicodeWordValidator()("Rene©", nil))
}

func TestUnicodeWordsAndSpacesValidator(t *testing.T) {
	assert.True(t, validation.UnicodeWordsAndSpacesValidator()("Renee", nil))
	assert.True(t, validation.UnicodeWordsAndSpacesValidator()("Reneé", nil))
	assert.True(t, validation.UnicodeWordsAndSpacesValidator()("Rènéó", nil))
	assert.True(t, validation.UnicodeWordsAndSpacesValidator()("Reneé ", nil))
	assert.True(t, validation.UnicodeWordsAndSpacesValidator()("Renee ", nil))
	assert.True(t, validation.UnicodeWordsAndSpacesValidator()("Renee Rènéó", nil))
	assert.False(t, validation.UnicodeWordsAndSpacesValidator()("Rene©", nil))
}

func TestBooleanValidator(t *testing.T) {
	assert.True(t, validation.BooleanValidator()("true", nil))
	assert.True(t, validation.BooleanValidator()("false", nil))
	assert.False(t, validation.BooleanValidator()("", nil))
	assert.False(t, validation.BooleanValidator()("t", nil))
	assert.False(t, validation.BooleanValidator()("f", nil))
	assert.False(t, validation.BooleanValidator()("Ok", nil))
	assert.False(t, validation.BooleanValidator()("On", nil))
	assert.False(t, validation.BooleanValidator()("Off", nil))
}

type PasswordValue struct {
	Password string
}

func (p *PasswordValue) GetUserInputs() []string {
	return []string{
		"{xvNTZF5sYN9&MU>",
		"{xvNTZF5sYN9&MUə€ǒ©₌",
		"RxKXAUHfAv",
		"6KMxpKYp64",
		"TzUnZFHQsA",
		"hESSPgVV6a",
	}
}

func TestPasswordValidatorUserInputs(t *testing.T) {
	pass := &PasswordValue{Password: ""}
	pass.Password = "{xvNTZF5sYN9&MU>"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "{xvNTZF5sYN9&MUə€ǒ©₌"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "RxKXAUHfAv"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "6KMxpKYp64"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "TzUnZFHQsA"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "hESSPgVV6a"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "TQgSuEv79j"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
}

type PasswordValueDumb struct {
	Password string
}

func (p *PasswordValueDumb) GetUserInputs() []string {
	return []string{}
}

func TestPasswordValidatorNoUserInputs(t *testing.T) {
	pass := &PasswordValueDumb{Password: ""}
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = ""
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "password"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "password1234"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "1234567890"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "Bob"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "2°"
	assert.False(t, validation.PasswordValidator()(pass.Password, pass))

	pass.Password = "{xvNTZF5sYN9&MU>"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "{xvNTZF5sYN9&MUə€ǒ©₌"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "RxKXAUHfAv"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "6KMxpKYp64"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "TzUnZFHQsA"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "hESSPgVV6a"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
	pass.Password = "TQgSuEv79j"
	assert.True(t, validation.PasswordValidator()(pass.Password, pass))
}
