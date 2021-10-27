package main

import (
    "log"
    "net/http"
)

func main() {

    // create file server handler
    fs := http.FileServer( http.Dir( "/home/first_user/installer/public" ) )

    // start HTTP server with `fs` as the default handler
    log.Fatal(http.ListenAndServe( ":8081", fs ))

}
