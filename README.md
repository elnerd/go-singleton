# Singleton Module
## Overview
This module provides a thread-safe implementation of a singleton pattern for managing global instances in a Go application. It features a centralized container that allows you to store, retrieve, and manage unique instances by their assigned names. The module abstracts away the complexity of singleton management, ensuring type safety during retrieval while promoting code reuse and consistency.
## Features
- **Store**: Save an instance in a shared container with a unique name.
- **Retrieve**: Access any stored instance by its name.
- **Delete**: Remove a stored instance using its name.
- **Type-Safe Retrieval with `GetInto`**: Retrieve and assign a singleton instance to a provided non-nil pointer with type validation.

## Installation
To include this module in your Go project, add it to your dependencies using:
``` sh
go get github.com/elnerd/go-singleton
```
## Usage
### 1. **Storing an instance**
You can store an instance in the singleton container by providing a unique name:
``` go

import "github.com/elnerd/go-singleton/pkg/singleton"

type MyService struct {
    Name string
}

service := &MyService{Name: "ExampleService"}
singleton.Store("myService", service)
```

### 2. **Type-Safe Retrieval using `GetInto`**
To simplify type retrieval and validation, use `GetInto`:
``` go
var retrievedService *MyService
if err := singleton.GetInto("myService", &retrievedService); err != nil {
    log.Fatalf("Error retrieving and assigning instance: %v", err)
}
fmt.Println("Retrieved service name:", retrievedService.Name)
```
### 3. **Retrieving an instance using `Get`**
To retrieve the stored instance, use the `singleton.Get` function:
``` go
import "github.com/elnerd/go-singleton/pkg/singleton"

func demo() {
  instance, err := singleton.Get("myService")
  if err != nil {
      log.Fatalf("Error retrieving instance: %v", err)
  }
  service, ok := instance.(*MyService)
  if !ok {
      log.Fatalf("Instance type mismatch")
  }
}
```
### 4. **Deleting an instance**
To remove an instance from the container:
``` go
singleton.Delete("myService")
```
## API Reference
### `Store`
``` go
func Store(name string, instance interface{})
```
- **Description**: Saves an `instance` in the singleton container under a given `name`.
- **Parameters**:
    - `name` (string): A unique identifier for the instance.
    - `instance` (interface{}): The instance to store.

### `Get`
``` go
func Get(name string) (interface{}, error)
```
- **Description**: Retrieves the instance associated with the provided `name`.
- **Returns**:
    - `interface{}`: The retrieved instance.
    - `error`: An error if the instance is not found.

NB. The recommended way to retrieve an instance is using the GetInto function. See function below.

### `GetInto`
``` go
func GetInto(name string, into interface{}) error
```
- **Description**: Retrieves a named singleton instance and assigns it to the `into` pointer, validating type compatibility.
- **Parameters**:
    - `name` (string): Unique name of the instance.
    - `into` (interface{}): A non-nil pointer to which the retrieved instance will be assigned.

- **Returns**:
    - `error`: An error if the instance cannot be retrieved, `into` is not a pointer, or there is a type mismatch.

### `Delete`
``` go
func Delete(name string)
```
- **Description**: Removes the instance associated with the provided `name` from the container.

## Example Use Cases
The singleton module can be used for:
1. **Global Configuration**: Managing shared configurations across the application.
2. **Service Instances**: Accessing and sharing network clients, services, or database connections globally.
3. **Caching**: Implementing a caching mechanism with one instance for managing stored objects.

## Error Handling
- **Instance Not Found**: When trying to retrieve or delete a name that doesn't exist, appropriate errors are returned.
- **Invalid Type Assignment**: The `GetInto` function ensures that the type of the retrieved instance matches the type of the provided pointer.

## Thread Safety
Internally, the module ensures thread-safe operations for storing, retrieving, and deleting instances, preventing race conditions in concurrent environments.

