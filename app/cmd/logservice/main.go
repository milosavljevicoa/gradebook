package main

import (
	"milosavljevicoa/gradebook/app/log"
	"milosavljevicoa/gradebook/app/service"
)

func main() {
	service.Serve(log.Initialize, ":4000")
}
