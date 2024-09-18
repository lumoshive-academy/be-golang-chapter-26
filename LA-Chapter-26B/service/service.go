package service

import "be-golang-chapter-26/LA-Chapter-26B/greeter"

type Service struct {
	Greeter *greeter.Greeter
}

func NewService(greeter *greeter.Greeter) *Service {
	return &Service{Greeter: greeter}
}
