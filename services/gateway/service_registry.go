package main

import (
	"errors"
	"fmt"
	"log"
)

type registryItem struct {
	ServiceName    string
	ServiceBaseUrl string
	HealthUrl      string
	Health         float32
	TimeFrame      int
}

type ServiceRegistry struct {
	services map[string]registryItem
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string]registryItem),
	}
}

func (s *ServiceRegistry) AddService(item registryItem) {
	s.services[item.ServiceName] = item
	log.Printf("Service Registered | %s | %s", item.ServiceName, item.ServiceBaseUrl)
}

func (s *ServiceRegistry) GetService(serviceName string) (registryItem, error) {
	item, ok := s.services[serviceName]
	if !ok {
		return registryItem{}, errors.New(fmt.Sprintf("unregistered service %s", serviceName))
	}

	return item, nil
}
