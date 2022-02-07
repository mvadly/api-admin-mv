package util

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetErrorJSON(err error) []string {
	var errors []string
	myErr := fmt.Sprintf("%v", err.Error())
	errors = strings.Split(myErr, "\n")
	return errors
}
func GetError(err error) map[string]string {
	myError := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		tag := err.ActualTag()
		if err.ActualTag() == "correct-format" {
			tag = "Correct Format"
		}
		myError[err.Field()] = fmt.Sprintf("(%s) is not a valid input, %s must be %s %s", err.Value(), err.Field(), tag, err.Param())
	}

	return myError
}

func ValidateCorrectInput(val validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9 &,.\\-\\/]+$")
	result := regex.MatchString(val.Field().String())
	return result

}
