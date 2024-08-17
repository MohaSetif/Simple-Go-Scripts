package greetings

import (
	"errors"
	"fmt"
	"reflect"
)

func Hello(name string) (string, error) {
	if name == "" || reflect.TypeOf(name).Kind() != reflect.String {
		return "", errors.New("wrong name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
