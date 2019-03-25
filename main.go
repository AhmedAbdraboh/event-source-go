package main

import (
	"fmt"
	"net/http"
	"roosh-app/helpers"
	rooshRoutes "roosh-app/routes"
	"strconv"
)

func main() {
	fmt.Println("Hello world Roosh")
	fmt.Println("Server started listening on PORT=" +strconv.Itoa(8080))

	routes := rooshRoutes.SetupRoutes()
	err := http.ListenAndServe(":8080", routes)
	helpers.CheckPanicError(err)
}
