package validation

import (
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/nbutton23/zxcvbn-go"
)

func RegisterAll() {
	govalidator.CustomTypeTagMap.Set("bool", BooleanValidator())
	govalidator.CustomTypeTagMap.Set("unicodeWord", UnicodeWordValidator())
	govalidator.CustomTypeTagMap.Set("unicodeWordsAndSpaces", UnicodeWordsAndSpacesValidator())
	govalidator.CustomTypeTagMap.Set("password", PasswordValidator())
}

type PasswordContainer interface {
	GetUserInputs() []string
}

func PasswordValidator() govalidator.CustomTypeValidator {
	const minComplexity = 2
	return func(i interface{}, context interface{}) bool {
		userInputs := []string{}
		switch o := context.(type) {
		case PasswordContainer:
			userInputs = o.GetUserInputs()
		default:
			return false
		}
		var password string
		switch v := i.(type) {
		case string:
			password = v
		case []byte:
			password = string(v)
		}
		measurements := zxcvbn.PasswordStrength(password, userInputs)
		return measurements.Score > minComplexity
	}
}

func UnicodeWordValidator() govalidator.CustomTypeValidator {
	unicodeWord := regexp.MustCompile(`^\p{L}+$`)
	return func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case string:
			return unicodeWord.Match([]byte(v))
		case []byte:
			return unicodeWord.Match(v)
		}
		return false
	}
}

func UnicodeWordsAndSpacesValidator() govalidator.CustomTypeValidator {
	unicodeWord := regexp.MustCompile(`^[\p{L}\s]+$`)
	return func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case string:
			return unicodeWord.Match([]byte(v))
		case []byte:
			return unicodeWord.Match(v)
		}
		return false
	}
}

func BooleanValidator() govalidator.CustomTypeValidator {
	boolean := regexp.MustCompile(`^true|false$`)
	return func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case string:
			return boolean.Match([]byte(v))
		case []byte:
			return boolean.Match(v)
		}
		return false
	}
}
