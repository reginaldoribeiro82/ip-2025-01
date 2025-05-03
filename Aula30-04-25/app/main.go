package main

import (
    "fmt"     
    "log"      
    "net/http" 
)

func main() {
    
    fileserver := http.FileServer(http.Dir("./static"))

    
    http.Handle("/", fileserver)

   
    fmt.Printf("port running on http://localhost:8081/\n") //  camimho para acessar o site do servidor local. para acessar só digitar no navegador

    
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err) 
    }
}
