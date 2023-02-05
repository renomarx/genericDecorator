package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type App struct {
	Service1 Service1Handler
	Service2 Service2Handler
	Service3 Service3Handler
}

func (a *App) Run() {
	a.Service1.SayHello()
	a.Service2.SayMyName("James", "Bond")
	quote, err := a.Service3.GetQuote(context.Background(), rand.Intn(14))
	if err != nil {
		fmt.Printf("Impossible to print quote: %s\n", err)
		return
	}
	fmt.Println(quote)
}

func main() {

	rand.Seed(time.Now().UnixMicro())
	// app := App{
	// 	Service1: &Service1{},
	// 	Service2: &Service2{},
	// 	Service3: &Service3{},
	// }
	app := App{
		Service1: NewService1Logged(&Service1{}),
		Service2: NewService2Logged(&Service2{}),
		Service3: NewService3Logged(&Service3{}),
	}
	app.Run()
}
