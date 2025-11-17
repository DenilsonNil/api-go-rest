package controller

import (
	"fmt"
	"net/http"
)

func Home(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(r, "Welcome to the Home Page!")
}
