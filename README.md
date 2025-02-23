# Singleton Go Package

**Singleton** is a design pattern that ensures a class has a single instance with global access. Learn more about
it [here](https://refactoring.guru/design-patterns/singleton). This pattern ensures the application has **one** instance
of a class, accessible globally.

> **⚠ Warning:** In most cases, avoid using this package. Implement the pattern yourself if possible.

Yes, the singleton is a design pattern and a design pattern is not meant to be replaced by a Go package.

## Why use this package?

So, why and when should you bother to use this package?

This package provides a simple way to manage singleton objects, identified by unique assigned names (e.g., `"database-client"` or `"app-config"`).

**Use Cases**
- Retrieve singleton instances anywhere in the application using a clean, type-safe interface.
- Assign singleton instances to local variables of the correct type by their associated name.
- Simplify the process of making any object a singleton.

Using this package is made to be thread safe, but this provide no guarantees that the instance you making a singleton is thread safe!


## Why should you NOT use the package?

If you only need for your package return a singleton, then simply do so.

# How to use this package

You use the `singleton.Store(name string, instance interface{})` to assign a singleton instance by name and
`singleton.GetInto(name string, into *interface{])` to load the instance into the variable *into.

See the example for storing and getting singleton instances below:


## Store(name, instance)

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

## GetInto(name, *instance)

This pattern show how you can retrieve the singleton by the assigned name "database-client".
Here we assign the instance to the local variable dbConn.

```go
package somepackage

import (
	sql "database/sql"
	"github.com/elnerd/go-singleton/pkg/singleton"
)

func YourFunction() {
	var dbConn *sql.DB
	if err := singleton.GetInto("database-client", &dbConn); err != nil {
		// handle error
	}
}
```


## API Reference

### `Store`

``` go
func Store[T Any](name string, instance *T)
```

- **Description**: Saves an instance in the singleton container under a name.
- **Parameters**:
  - `name` (string): Unique identifier.
  - `instance` (*T): The instance to store.

Note: The generic type T is used to provide error checking at compile time. You would always want a pointer.

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