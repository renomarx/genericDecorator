package main

import (
	"context"
	"fmt"
)

type Service3Handler interface {
	GetQuote(ctx context.Context, i int) (string, error)
}

type Service3 struct {
}

func (s *Service3) GetQuote(ctx context.Context, i int) (string, error) {
	quotes := []string{
		`The greatest glory in living lies not in never falling, but in rising every time we fall. -Nelson Mandela`,
		`The way to get started is to quit talking and begin doing. -Walt Disney`,
		`Your time is limited, so don't waste it living someone else's life. Don't be trapped by dogma – which is living with the results of other people's thinking. -Steve Jobs`,
		`If life were predictable it would cease to be life, and be without flavor. -Eleanor Roosevelt`,
		`If you look at what you have in life, you'll always have more. If you look at what you don't have in life, you'll never have enough. -Oprah Winfrey`,
		`If you set your goals ridiculously high and it's a failure, you will fail above everyone else's success. -James Cameron`,
		`Life is what happens when you're busy making other plans. -John Lennon`,
	}
	if i < 0 || i >= len(quotes) {
		return "", fmt.Errorf("%d out of range: [0,%d]", i, len(quotes))
	}
	return quotes[i], nil
}

type Service3Logged struct {
	d loggingDecorator[[]any, []any]
}

func NewService3Logged(service3 Service3Handler) *Service3Logged {
	return &Service3Logged{
		d: loggingDecorator[[]any, []any]{
			serviceName: "service3",
			handler: func(inputs []any) []any {
				quote, err := service3.GetQuote(inputs[0].(context.Context), inputs[1].(int))
				return []any{quote, err}
			},
		},
	}
}

func (s Service3Logged) GetQuote(ctx context.Context, i int) (string, error) {
	outputs := s.d.Handle([]any{ctx, i})
	res0, _ := outputs[0].(string)
	res1, ok := outputs[1].(error)
	if !ok {
		return res0, nil
	}
	return res0, res1
}
