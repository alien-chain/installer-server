package main

import (
    "log"
    "net/http"
)

func main() {

    // create file server handler
    fs := http.FileServer( http.Dir( "/home/first_user/installer" ) )

    // start HTTP server with `fs` as the default handler
    log.Fatal(http.ListenAndServe( ":80", fs ))

}
