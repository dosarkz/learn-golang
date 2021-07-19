package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is empty")
	}
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
