package cache

import (
	"testing"
)

func TestLiteStoreManager(t *testing.T) {

	/*
		in this example we are using anonymous structs, another of Go's
		features to implement dynamic unit tests.
		This pattern may be used to test multiple scenarios in which state
		may play a considerable role
	*/
	tests := []struct {
		expectError bool
		description string
		operation   func(*SimpleStore) error
		preset      func(*SimpleStore)
	}{
		{
			expectError: false,
			description: "Lite Store manager can store and load data",
			operation: func(lm *SimpleStore) error {
				_, err := lm.Load("existentKey")
				return err
			},
			preset: func(lm *SimpleStore) {
				lm.Store("existentKey", "somevalue")
			},
		},
		{
			expectError: true,
			description: "Lite Store Manager can't retrieve a value that is not stored",
			operation: func(lm *SimpleStore) error {
				_, err := lm.Load("nonexistentKey")
				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t2 *testing.T) {
			store := NewSimpleStore()
			if test.preset != nil {
				test.preset(store) // preset state if necessary
			}
			err := test.operation(store) // execute the operation
			switch test.expectError {
			case true:
				// if an error was EXPECTED, fail ONLY if there is NO ERROR
				if err == nil {
					t2.Errorf("Expected an error but got none\n")
				}

			// if errors are NOT EXPECTED, fail ONLY if an error IS RETURNED
			case false:
				if err != nil {
					t2.Errorf("Was not expecting error but got \"%v\"", err)
				}
			}
		})
	}
}
