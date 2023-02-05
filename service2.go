package main

import "fmt"

type Service2Handler interface {
	SayMyName(firstname string, lastname string) error
}

type Service2 struct {
}

func (s *Service2) SayMyName(firstname string, lastname string) error {
	fmt.Printf("My name is %s, %s %s\n", lastname, firstname, lastname)
	return nil
}

type Service2Logged struct {
	d loggingDecorator[[]string, []any]
}

func NewService2Logged(service2 Service2Handler) *Service2Logged {
	return &Service2Logged{
		d: loggingDecorator[[]string, []any]{
			serviceName: "service2",
			handler: func(names []string) []any {
				err := service2.SayMyName(names[0], names[1])
				return []any{err}
			},
		},
	}
}

func (s Service2Logged) SayMyName(firstname string, lastname string) error {
	result := s.d.Handle([]string{firstname, lastname})
	res, ok := result[0].(error)
	if !ok {
		return nil
	}
	return res
}
