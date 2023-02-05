package main

import "fmt"

type Service1Handler interface {
	SayHello()
}

type Service1 struct {
}

func (s *Service1) SayHello() {
	fmt.Println("Hello!")
}

type Service1Logged struct {
	d loggingDecorator[any, []any]
}

func NewService1Logged(service1 Service1Handler) *Service1Logged {
	return &Service1Logged{
		d: loggingDecorator[any, []any]{
			serviceName: "service1",
			handler: func(any) (output []any) {
				service1.SayHello()
				return output
			},
		},
	}
}

func (s Service1Logged) SayHello() {
	s.d.Handle(nil)
}
