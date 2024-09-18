package service

import (
	"be-golang-chapter-26/greeter"
	"errors"
)

type Service struct {
	Greeter *greeter.Greeter
}

func NewService(greeter *greeter.Greeter) (*Service, error) {
	if greeter.Error {
		return nil, errors.New("failed create service")
	} else {
		return &Service{Greeter: greeter}, nil
	}

}
