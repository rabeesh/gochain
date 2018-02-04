package main

import (
	"github.com/rabeesh/gochain/app"
	api "github.com/rabeesh/gochain/chainapi"
	_ "github.com/rabeesh/gochain/config"
)

func main() {
	ai := app.New()
	api.Init(ai)
	ai.Run()
}
