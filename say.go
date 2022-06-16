
package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)


// Hello returns a greeting for the name person.
func Hello(name string) (string, error) {
	// if no  name was give, return an error with a message.

	if name == "" {
		return "", errors.New("Empty name is not allowed")
	}

	// If a name was received. return a  value that embeds the name
	// in greeting message
	message := fmt.Sprintf(randomFormat(), name)
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats

	formats := []string {
		"Hi!, %v. Welcome!",
		"Great to saee you, %v",
		"Hail, %v!, Well met!",
	}


	 // Return randomly selected message format by specifying
	 // a random index for the slice of formats

	 return formats[rand.Intn(len(formats))]
}
