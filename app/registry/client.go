package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegistryService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(r); err != nil {
		return err
	}
	res, err := http.Post(ServiceRegistrationUrl, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry service responded with code %v", res.StatusCode)
	}
	return nil
}

func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(http.MethodDelete, ServiceRegistrationUrl, bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service. Registr service responded with code %v", res.StatusCode)
	}
	return nil
}
