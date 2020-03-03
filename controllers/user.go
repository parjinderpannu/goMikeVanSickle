package controllers

import "net/http"

type userController struct{}

//change func to method
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller!"))
}
