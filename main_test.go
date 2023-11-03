package main

import "testing"

func TestNew(t *testing.T) {
	hashset := New()
	if hashset.size != InitialHashsetSize {
		t.Errorf("Expected size %d, got %d", InitialHashsetSize, hashset.size)
	}
	if len(hashset.elements) != int(InitialHashsetSize) {
		t.Errorf("Expected length %d, got %d", InitialHashsetSize, len(hashset.elements))
	}
}

func TestHashFunction(t *testing.T) {
	testCases := []struct {
		input    string
		expected uint32
	}{
		{"test", 4}, // Insert expected hash value for "test" after calculating it
		// Add more test cases for hash function here
	}

	hashset := New()

	for _, tc := range testCases {
		actual := hashset.hash(tc.input)
		if actual != tc.expected {
			t.Errorf("For input %s, expected %d but got %d", tc.input, tc.expected, actual)
		}
	}
}

func TestAddAndHas(t *testing.T) {
	hashset := New()

	// Add test cases to validate adding elements and checking existence
	testAdd := "example"
	hashset.Add(testAdd)

	if !hashset.Has(testAdd) {
		t.Errorf("Expected %s to be present in the hashset, but it's not", testAdd)
	}

	testNotAdded := "non-existent"
	if hashset.Has(testNotAdded) {
		t.Errorf("Expected %s to not be present in the hashset, but it is", testNotAdded)
	}
}

func TestMainFunction(t *testing.T) {
	// This test covers the main function for basic functionality
	// It doesn't have direct access to internal variables for verification
	// It verifies if the main function executes without errors.

	// Testing main() function
	main()
	// If the function completes without a runtime error, it's considered successful.
}
