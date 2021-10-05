package main

import (
	"fmt"
	stlog "log"
	"milosavljevicoa/gradebook/app/grades"
	"milosavljevicoa/gradebook/app/registry"
	"milosavljevicoa/gradebook/app/service"
)

func main() {
	host, port := "localhost", "6000"
	reg := registry.Registration{
		ServiceName: registry.GradeService,
		ServiceURL:  fmt.Sprintf("http://%v:%v", host, port),
	}
	if err := service.Serve(grades.Initialize, reg, ":"+port); err != nil {
		stlog.Printf("%v", err)
	}
}
