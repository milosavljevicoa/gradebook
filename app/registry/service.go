package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const ServerPort = ":3000"
const ServerURL = "http://localhost" + ServerPort
const ServiceRegistrationUrl = ServerURL + "/services"

func AddService(c *gin.Context) {
	log.Print("Request for adding service received")

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

func RemoveService(c *gin.Context) {
	log.Print("Request for removing service received")
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	url := buf.String()
	log.Printf("Removing service at URL: %v", url)

	if err := reg.remove(url); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
}

type registry struct {
	registrations []Registration
	mutex         *sync.RWMutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

func (r *registry) remove(url string) error {
	for i := 0; i < len(r.registrations); i++ {
		if r.registrations[i].ServiceURL == url {
			r.mutex.Lock()
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}

	return fmt.Errorf("Service url: %v not found", url)
}

func (r *registry) sendRequiredServices(reg Registration) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var p patch
	for _, serviceReg := range r.registrations {
		for _, reqService := range reg.RequiredServices {
			if serviceReg.ServiceName == reqService {
				p.Added = append(p.Added, patchEntry{
					Name: serviceReg.ServiceName,
					URL:  serviceReg.ServiceURL,
				})
			}
		}
	}

	err := r.sendPatch(p, req.ServiceUpdateURL)
	if err != nil {
		return err
	}
	return nil
}

func (r *registry) sendPatch(p patch, url string) err {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}

	if _, err = http.Post(url, "application/json", bytes.NewBuffer(d)); err != nil {
		return err
	}
	return nil
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.RWMutex),
}

//type RegistryService struct{}
