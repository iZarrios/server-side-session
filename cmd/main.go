package main

import (
	"www.github.com/iZarrios/server-side-session-golang/pkg/db"
	"www.github.com/iZarrios/server-side-session-golang/pkg/routes"
)

func main() {

	db.Setup()
	routes.Setup()
}
