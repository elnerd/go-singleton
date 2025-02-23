# Singleton Go Package

**Singleton** is a design pattern that ensures a class has a single instance with global access. Learn more about
it [here](https://refactoring.guru/design-patterns/singleton). This pattern ensures the application has **one** instance
of a class, accessible globally.

> **⚠ Warning:** In most cases, avoid using this package. Implement the pattern yourself if possible.

## Why use this package?

This Go package provides:

- Type-safe instance retrieval.
- Thread-safe singleton implementation.
- Management of unique instances by assigned names.

These features enable easy retrieval of global instances and their assignment to correctly-typed variables.

This reduces tight coupling in your app by avoiding the need to pass instances as function arguments.

### Getter

```go
package somepackage

import (
	"database/sql"
	"github.com/elnerd/go-singleton/pkg/singleton"
)

func YourFunction() {
	var dbConn *db.Conn
	if err := singleton.GetInto("database-client", &dbConn); err != nil {
		// handle error
	}
}
```

### Setter

```go
package anotherpackage

import (
	"database/sql"
	"github.com/elnerd/go-singleton/pkg/singleton"
)

func main() {
	connectionString := "user:password@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		// handle error
	}
	singleton.Store("database-client", &db)
}
```

## Installation

To include this module, use:

```sh
go get github.com/elnerd/go-singleton
```

## API Reference

### `Store`

``` go
func Store(name string, instance interface{})
```

- **Description**: Saves an instance in the singleton container under a name.
- **Parameters**:
  - `name` (string): Unique identifier.
  - `instance` (interface{}): The instance to store.

### `Get`

``` go
func Get(name string) (interface{}, error)
```

- **Description**: Retrieves the instance associated with `name`.
- **Returns**:
  - `interface{}`: The instance.
  - `error`: If the instance is not found.

> It’s recommended to use `GetInto` as it provides a cleaner interface.

### `GetInto`

``` go
func GetInto(name string, into interface{}) error
```

- **Description**: Retrieves a named singleton instance and assigns it to a pointer, ensuring type safety.
- **Parameters**:
  - `name` (string): Unique name.
  - `into` (interface{}): A non-nil pointer to assign the instance.

- **Returns**:
  - `error`: If retrieval fails, `into` is not a pointer, or there’s a type mismatch.

### `Delete`

``` go
func Delete(name string)
```

- **Description**: Removes the instance associated with `name`.