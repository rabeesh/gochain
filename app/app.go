package app

import (
	"log"
	"net/http"
	"time"

	"github.com/rabeesh/gochain/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

type App struct {
	*negroni.Negroni
	Router *httprouter.Router
}

// New Creates a application Instance
func New() *App {
	ngApp := negroni.New()
	return &App{
		ngApp,
		httprouter.New(),
	}
}

func (ai *App) Run() {
	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()

	// use route as middleware
	ai.UseHandler(ai.Router)

	log.Println("Server http://localhost:" + config.Addr)
	http.ListenAndServe(":"+config.Addr, ai)
}
