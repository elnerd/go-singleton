package singleton

import (
	"fmt"
	"sync"
)

// containerInstance is a singleton instance of the container struct that stores and retrieves named instances.
var containerInstance *container

var once sync.Once

// getContainerInstance returns a singleton instance of the containerInterface to ensure a single shared container.
// It initializes the container with a thread-safe mechanism if not already created.
func getContainerInstance() containerInterface {
	once.Do(func() {
		containerInstance = &container{
			instances: make(map[string]interface{}),
		}
	})
	return containerInstance
}

// containerInterface defines methods for storing and retrieving named singleton instances.
type containerInterface interface {
	get(name string) (interface{}, error)
	store(name string, instance interface{})
	del(name string)
}

// container is a struct that holds a map of singleton instances identified by string keys.
type container struct {
	instances map[string]interface{}
	mutex     sync.RWMutex
}

// del removes an instance from the container's map using the provided name as the key.
func (s *container) del(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.instances, name)
}

// get retrieves a singleton instance by its name from the container or returns an error if not found.
func (s *container) get(name string) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	instance, ok := s.instances[name]
	if !ok {
		return nil, fmt.Errorf("no singleton instance named '%s'", name)
	}
	return instance, nil
}

// store adds a new instance to the container with the given name as the key.
func (s *container) store(name string, instance interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.instances[name] = instance
}
