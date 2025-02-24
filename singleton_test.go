package singleton

import (
	"errors"
	"testing"
)

type TestStruct struct {
	Name string
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCreateSingleton(t *testing.T) {
	var testSingleton *TestStruct
	testSingleton = &TestStruct{Name: "test"}
	Store("testSingleton", testSingleton)
}

func TestGetNonExistingSingleton(t *testing.T) {
	instance, err := Get("test_getNonExistingSingleton")
	if !errors.Is(err, ErrSingletonNotFound) {
		t.Fatalf(`Get("test_getNonExistingSingleton") = %v, want error`, instance)
	}
}

func TestCreateGetSingleton(t *testing.T) {
	// Create
	var testSingleton *TestStruct
	testSingleton = &TestStruct{Name: "test"}
	Store("testSingleton", testSingleton)
	testSingleton.Name = "test-changed"
	// Get
	instance, err := Get("testSingleton")
	if err != nil {
		t.Fatalf(`Get("testSingleton") = %v`, testSingleton)
	}
	testSingleton, ok := instance.(*TestStruct)
	if !ok {
		t.Fatalf(`instance.(*TestStruct) = %v, want TestStruct`, testSingleton)
	}
	if testSingleton.Name != "test-changed" {
		t.Fatal(`testSingleton.Name != "test-changed"`)
	}
}

func TestGetInto(t *testing.T) {
	var testSingleton *TestStruct
	testSingleton = &TestStruct{Name: "test"}
	Store("test_getinto_singleton", testSingleton)
	var testSingleton2 *TestStruct
	err := GetInto("test_getinto_singleton", &testSingleton2)
	testSingleton2.Name = "test-changed"
	if err != nil {
		t.Fatalf(`GetInto("test_getinto_singleton", &testSingleton2) = %v`, testSingleton)
	}
	if testSingleton2.Name != "test-changed" {
		t.Fatal(`testSingleton2.Name != "test-changed"`)
	}
}

func TestGetIntoIncorrectType(t *testing.T) {
	var testSingleton *TestStruct
	testSingleton = &TestStruct{Name: "test"}
	Store("test_getinto_singleton_incorrect_type", testSingleton)
	var testSingleton2 int
	err := GetInto("test_getinto_singleton_incorrect_type", &testSingleton2)
	// type mismatch: cannot assign instance type *singleton.TestStruct to target type int
	if !errors.Is(err, ErrInvalidType) {
		t.Fatalf(`GetInto("test_getinto_singleton_incorrect_type", &testSingleton2) = %v, want error`, testSingleton)
	}
}
func TestGetIntoInvalidType(t *testing.T) {
	var testSingleton TestStruct
	var testSingleton2 TestStruct // should not be possible to GetInto a non-pointer type
	testSingleton = TestStruct{Name: "test"}
	Store("test_getinto_singleton_invalid_type", &testSingleton)

	err := GetInto("test_getinto_singleton_invalid_type", &testSingleton2)
	// type mismatch: cannot assign instance type *singleton.TestStruct to target type singleton.TestStruct
	if !errors.Is(err, ErrInvalidType) {
		t.Fatalf(`GetInto("test_getinto_singleton_invalid_type", &testSingleton2) = %v, want error`, testSingleton)
	}
}

func TestDelete(t *testing.T) {
	var testSingleton *TestStruct
	testSingleton = &TestStruct{Name: "test"}
	Store("testDeleteSingleton", testSingleton)
	Delete("testDeleteSingleton")
	instance, err := Get("testDeleteSingleton")
	if !errors.Is(err, ErrSingletonNotFound) {
		t.Fatalf(`Get("testDeleteSingleton") = %v, want error`, instance)
	}
}

func TestDeleteNonExistingSingleton(t *testing.T) {
	// Should never fail
	Delete("testSingleton")
}

func TestGetIntoWithErr(t *testing.T) {
	var testSingleton *TestStruct
	err := GetInto("testGetIntoNotExisting", &testSingleton)
	if !errors.Is(err, ErrSingletonNotFound) {
		t.Fatalf(`GetInto("testGetIntoNotExisting", &testSingleton2) = %v, want error`, testSingleton)
	}
}

func TestGetIntoWithNilTarget(t *testing.T) {
	var testSingleton TestStruct
	Store("testGetIntoWithNilTarget", &testSingleton)
	err := GetInto("testGetIntoWithNilTarget", nil)
	if !errors.Is(err, ErrInvalidType) {
		t.Fatalf(`GetInto("testGetIntoWithNilTarget", nil) = %v, want error`, testSingleton)
	}
}

func TestGetIntoWithNonPointerStruct(t *testing.T) {
	var testSingleton TestStruct
	Store("testGetIntoWithNonPointerStruct", &testSingleton)
	err := GetInto("testGetIntoWithNonPointerStruct", &testSingleton)
	if !errors.Is(err, ErrInvalidType) {
		t.Fatalf(`GetInto("testGetIntoWithNonPointerStruct", nil) = %v, want error`, testSingleton)
	}
}
