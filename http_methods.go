package main

import (
        "fmt"
        "net/http"
)

func method_matcher(method string, handler http.HandlerFunc) http.HandlerFunc {
        var err string
        return func(res http.ResponseWriter, req *http.Request) {
                if method != req.Method {
                        err = fmt.Sprintf("Expected %s Http Method, Got: %s Http Method", method, req.Method)
                        http.Error(res, err, http.StatusBadRequest)
                } else {
                        handler.ServeHTTP(res, req)
                }
        }
}

func Get(handler http.HandlerFunc) http.HandlerFunc {
        return method_matcher("GET", handler)
}

func Post(handler http.HandlerFunc) http.HandlerFunc {
        return method_matcher("POST", handler)
}

func Put(handler http.HandlerFunc) http.HandlerFunc {
        return method_matcher("PUT", handler)
}

func Delete(handler http.HandlerFunc) http.HandlerFunc {
        return method_matcher("DELETE", handler)
}

func Patch(handler http.HandlerFunc) http.HandlerFunc {
        return method_matcher("PATCH", handler)
}

func main() {
        server := http.NewServeMux()
        server.Handle("/", Get(Index_Handler))
        server.Handle("/users", Get(All_Handler))
        server.Handle("/user/new", Post(New_Handler))
        http.ListenAndServe(":8000", server)
}

func Index_Handler(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("Hello World"))
}

func New_Handler(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("Creating User"))
}

func All_Handler(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("All Users"))
}