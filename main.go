package main

import (
	"github.com/jclaudiotomasjr/go-live/api-go-gin-quiz/database"
	"github.com/jclaudiotomasjr/go-live/api-go-gin-quiz/routes"
)

func main() {
	database.ConectaComBanco()
	routes.HandleRequests()
}
