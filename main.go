package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
})


	server := http.Server {
		Addr: "localhost:8888",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}