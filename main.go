package main

import (
     "github.com/baselrabia/book-microservice/routes"
)

func main() {
    r := routes.SetupRouter()
    r.Run(":8080") // Run on port 8080
}
