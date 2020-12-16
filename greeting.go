package greeting

import (
	"errors"
	"fmt"
)

// Hello is receiveing name and return greeting message
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
