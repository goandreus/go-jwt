package main

import (
	"fmt"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndPoint)).Methods("GET")


	log.Println("Listen port 3030")
	log.Fatal(http.ListenAndServe(":3030",router))


}

func signup(w http.ResponseWriter, r *http.Request)  {
 fmt.Println("signup invoked")
}

func login(w http.ResponseWriter,  r *http.Request)  {
	fmt.Println("login invoked")
}

func protectedEndPoint(w http.HandlerFunc,r *http.Request){
	fmt.Println("protectedEndPoint invoked")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc{
	fmt.Println("TokenVerifyMiddleWare")
	return nil
}



