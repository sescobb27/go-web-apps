package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "os"
)

// START OMIT
func main() {
    router := httprouter.New()
    router.GET("/categories", Categories_Handler)
    router.GET("/locations", Locations_Handler)
    router.Handler("GET", "/css/*filename",
        http.StripPrefix("/css/",
            http.FileServer(http.Dir("resources/css"))))
    router.Handler("GET", "/js/*filename",
        http.StripPrefix("/js/",
            http.FileServer(http.Dir("resources/js"))))
    go func(router *httprouter.Router) {
        err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", router)
        if err != nil {
            panic(err)
        }
    }(router)
    http.ListenAndServe(":3000", router)
}

// END OMIT
