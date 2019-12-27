package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/timojarv/orjabot/data"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	je := json.NewEncoder(w)
	je.Encode(map[string]string{
		"name":    "Orja",
		"tagline": "Koska tietojohtaminen on helppoa ja hauskaa!",
		"version": "never done",
	})
}

func messages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	je := json.NewEncoder(w)
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		je.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	messages, err := data.GetMessages(int64(id))
	if err != nil {
		je.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	je.Encode(map[string]interface{}{
		"chat":     id,
		"messages": *messages,
	})
}

// RunRouter starts the router
func RunRouter() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/messages/:id", messages)

	port := 8080
	log.Printf("API running on port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
