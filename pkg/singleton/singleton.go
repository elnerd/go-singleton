package singleton

import (
	"errors"
	"fmt"
	"reflect"
)

// Store saves an instance in a singleton container using a unique name as the key for future retrieval.
func Store[T any](name string, instance *T) {
	var container = getContainerInstance()
	container.store(name, instance)
}

// Get retrieves a singleton instance by its name from the container or returns an error if the instance is not found.
func Get(name string) (interface{}, error) {
	var container = getContainerInstance()
	return container.get(name)
}

// Delete removes a named singleton instance from the container shared by the application.
func Delete(name string) {
	var container = getContainerInstance()
	container.del(name)
}

// GetInto retrieves a singleton instance by name and assigns it to the provided non-nil pointer `into`.
// Returns an error if the instance cannot be retrieved, `into` is not a pointer, or there is a type mismatch.
func GetInto(name string, into interface{}) error {
	instance, err := Get(name)
	if err != nil {
		return fmt.Errorf("could not get singleton: %w", err)
	}

	intoVal := reflect.ValueOf(into)
	if intoVal.Kind() != reflect.Ptr || intoVal.IsNil() {
		return errors.New("into must be a non-nil pointer")
	}

	intoElem := intoVal.Elem()

	instanceVal := reflect.ValueOf(instance)
	if !instanceVal.Type().AssignableTo(intoElem.Type()) {
		return fmt.Errorf("type mismatch: cannot assign instance type %s to target type %s", instanceVal.Type(), intoElem.Type())
	}
	intoElem.Set(instanceVal)
	return nil
}
