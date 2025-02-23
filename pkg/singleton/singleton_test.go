package singleton

import (
	"fmt"
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
	if err == nil {
		t.Fatalf(`Get("test_getNonExistingSingleton") = %v, want error`, instance)
	}
	_ = instance // ignore value
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
		t.Fatalf(`Get("testSingleton") = %v, want error`, testSingleton)
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
	fmt.Println(testSingleton2.Name)
	if err != nil {
		t.Fatalf(`GetInto("test_getinto_singleton", &testSingleton2) = %v, want error`, testSingleton)
	}
	if testSingleton2.Name != "test-changed" {
		t.Fatal(`testSingleton2.Name != "test"`)
	}
}
