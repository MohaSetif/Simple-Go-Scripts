package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func GetRandomName() string {
	names := []string{"Khalil", "Aymen", "Anes", ""}
	return names[rand.Intn(len(names))]
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("wrong name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
