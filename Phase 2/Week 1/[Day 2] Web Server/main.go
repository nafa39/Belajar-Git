package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// //handler
	// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, World!")
	// }

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: handler,
	// }

	// port := ":8080"

	// fmt.Printf("Starting server at port %s\n", port)
	// log.Fatal(server.ListenAndServe())

	//httprouter

	router := httprouter.New()

	router.GET("/all", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		user := []user{
			{
				Name: "John Doe",
				Age:  20,
			},
			{
				Name: "Jane Doe",
				Age:  21,
			},
			{
				Name: "John Smith",
				Age:  22,
			},
		}
	})

}
