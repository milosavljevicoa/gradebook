package registry

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const ServerPort = ":3000"
const ServerURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

//type RegistryService struct{}

func ServeHttpPost(c *gin.Context) {
	log.Print("Request received")

	var r Registration
	if err := c.BindJSON(&r); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	log.Printf("Adding service: %v with URL: %v\n", r.ServiceName, r.ServiceURL)

	if err := reg.add(r); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
}
