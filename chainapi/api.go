package chainapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	"github.com/rabeesh/gochain/app"
)

func Init(ai *app.App) {
	router := ai.Router
	router.GET("/api/blocks", GetBlockHandler)
	router.POST("/api/blocks", PostBlockHandler)

	log.Println("Initialized api")
}

// write blockchain when we receive an http request
func GetBlockHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bytes, err := json.MarshalIndent(app.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// takes JSON payload as an input for heart rate (BPM)
func PostBlockHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var m app.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := app.GenerateBlock(m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if app.IsBlockValid(newBlock, app.Blockchain[len(app.Blockchain)-1]) {
		newBlockchain := append(app.Blockchain, newBlock)
		app.ReplaceChain(newBlockchain)
		spew.Dump(app.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
