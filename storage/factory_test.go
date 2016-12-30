package storage

import (
	"reflect"
	"testing"
)

func TestStorageFactory(t *testing.T) {
	testCases := []struct {
		storageType     int
		expectedStorage reflect.Type
	}{
		{FakeType, reflect.TypeOf(&Fake{})},
		{SQLiteType, reflect.TypeOf(&SQLite{})},
		{99, reflect.TypeOf(&Fake{})},
	}

	for _, tc := range testCases {
		actual := New(tc.storageType, "")
		typeOfActual := reflect.TypeOf(actual)
		if typeOfActual != tc.expectedStorage {
			t.Errorf("Expected %v, got %v", tc.expectedStorage, typeOfActual)
		}
	}
}
