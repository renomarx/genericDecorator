package main

import "log"

type loggingDecorator[Input any, Output []any] struct {
	serviceName string
	handler     func(input Input) (output Output)
}

func (d loggingDecorator[Input, Output]) Handle(input Input) (output Output) {
	log.Printf("Before service %s call", d.serviceName)
	defer func() {
		for _, something := range output {
			err, ok := something.(error)
			if ok && err != nil {
				log.Printf("Error: %s", err)
			}
		}
		log.Printf("After service %s call", d.serviceName)
	}()
	return d.handler(input)
}
