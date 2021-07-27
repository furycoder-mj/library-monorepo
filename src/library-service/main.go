package main

import (
	"fmt"
	"net/http"

	"epam.com/web-services/library-management/library-service/config"
	"epam.com/web-services/library-management/library-service/handlers"
)

func main() {
	err := http.ListenAndServe(":"+fmt.Sprint(config.Config.Port), http.HandlerFunc(handlers.Serve))
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() != nil {
			fmt.Println("found a panic")
		}
	}()
}
