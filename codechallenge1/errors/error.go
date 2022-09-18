package errors

import (
	"errors"
	"log"
)

func Error(msg string) error {
	log.Print(msg)
	return errors.New(msg)
}
