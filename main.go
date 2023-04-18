package main

import (
	"log"
	"simple_rest_api/routers"
)

func main() {

	err := routers.StartServer().Run(":8080")
	if err != nil {
		log.Println(err.Error())
		return
	}
}
