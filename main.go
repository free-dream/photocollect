package main

import (
	"photocollect/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}