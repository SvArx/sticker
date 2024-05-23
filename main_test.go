// main_test.go
package main

import (
	"reflect"
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSerilaizeAndDeserializeCart(t *testing.T) {
	cart := map[string]uint{
		"product1": 1,
		"product2": 2,
	}

	serialized := serialize_cart(cart)
	deserialized := deserialize_cart(serialized)

	if !reflect.DeepEqual(cart, deserialized) {
		t.Errorf("Maps are not equal. map1: %v, map2: %v", cart, deserialized)
	}
}
