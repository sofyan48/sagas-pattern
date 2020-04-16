package registry

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

// ServiceRegistry ...
type ServiceRegistry struct {
}

// ServiceRegistryHandler ...
func ServiceRegistryHandler() *ServiceRegistry {
	return &ServiceRegistry{}
}

// ServiceRegistryInterface ...
type ServiceRegistryInterface interface {
	Health(path string)
}

// Health ...
func (registry *ServiceRegistry) Health(path string) {
	healthchecks := fmt.Sprintf("0.0.0.0:%s", os.Getenv("HEALTH_CHECK_PORT"))
	if healthchecks == "" {
		healthchecks = "0.0.0.0:8001"
	}
	http.HandleFunc(path, healthcheck)
	http.ListenAndServe(healthchecks, nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `user service is good`)
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "product"
	registration.Name = "product"
	address := hostname()
	registration.Address = address
	port, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = port
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

func port() string {
	p := os.Getenv("PRODUCT_SERVICE_PORT")
	h := os.Getenv("PRODUCT_SERVICE_HOST")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8100"
	}
	return fmt.Sprintf("%s:%s", h, p)
}

func hostname() string {
	// return os.Getenv("CONSUL_HTTP_ADDR")
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
