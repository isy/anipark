package main

import "github.com/isy/ownedia/src/be/router"

func main() {
	r := router.GetRouter()

	r.Run(":3001")
}
