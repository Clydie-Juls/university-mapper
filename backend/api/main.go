package main

import (
	"api/api"
	"log"
)

func main() {
  apiServer := api.NewAPIServer(":8080")
  if err := apiServer.Run(); err != nil {
    log.Fatal(err)
  }
}
