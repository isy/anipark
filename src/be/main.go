package main

import (
	"github.com/isy/anipark/src/be/infra/datastore"
	"github.com/isy/anipark/src/be/infra/router"
)

func main() {
	datastore.Conn()
	defer datastore.Close()

	r := router.GetRouter()

	r.Run(":3001")
}
