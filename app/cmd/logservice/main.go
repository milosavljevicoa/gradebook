package main

import (
	"fmt"
	stlog "log"
	"milosavljevicoa/gradebook/app/log"
	"milosavljevicoa/gradebook/app/registry"
	"milosavljevicoa/gradebook/app/service"
)

func main() {
	host, port := "localhost", "4000"
	reg := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  fmt.Sprintf("http://%v:%v", host, port),
	}
	if err := service.Serve(log.Initialize, reg, ":"+port); err != nil {
		stlog.Printf("%v", err)
	}
}
