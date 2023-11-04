package main

import (
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	hashset := New()
	if hashset.size != DefaultHashsetSize {
		t.Errorf("Expected size %d, got %d", DefaultHashsetSize, hashset.size)
	}
	if len(hashset.elements) != int(DefaultHashsetSize) {
		t.Errorf("Expected length %d, got %d", DefaultHashsetSize, len(hashset.elements))
	}
}

func TestHashFunction(t *testing.T) {
	testCases := []struct {
		input    string
		expected uint32
	}{
		{"test", 4}, // Insert expected hash value for "test" after calculating it
		// Insert more test cases for hash function here
	}

	hashset := New()

	for _, tc := range testCases {
		actual := hashset.hash(tc.input)
		if actual != tc.expected {
			t.Errorf("For input %s, expected %d but got %d", tc.input, tc.expected, actual)
		}
	}
}

func TestInsertAndHas(t *testing.T) {
	hashset := New()

	// Insert test cases to validate Inserting elements and checking existence
	testInsert := "example"
	hashset.Insert(testInsert)

	if !hashset.Has(testInsert) {
		t.Errorf("Expected %s to be present in the hashset, but it's not", testInsert)
	}

	testNotInserted := "non-existent"
	if hashset.Has(testNotInserted) {
		t.Errorf("Expected %s to not be present in the hashset, but it is", testNotInserted)
	}
}

func TestRemoveAndHas(t *testing.T) {
	hashset := New()

	// Insert elements into the hashset
	for i := 0; i < 5; i++ {
		el := strconv.Itoa(i)
		hashset.Insert(el)
	}

	// Remove elements from the hashset
	for i := 0; i < 3; i++ {
		el := strconv.Itoa(i)
		hashset.Remove(el)
	}

	// Test if removed elements don't exist in the hashset
	for i := 0; i < 3; i++ {
		el := strconv.Itoa(i)
		if hashset.Has(el) {
			t.Errorf("Expected %s to be removed from the hashset, but it's still present.", el)
		}
	}

	// Test if elements that were not removed are still present in the hashset
	for i := 3; i < 5; i++ {
		el := strconv.Itoa(i)
		if !hashset.Has(el) {
			t.Errorf("Expected %s to be in the hashset, but it's not.", el)
		}
	}
}

func TestCopy(t *testing.T) {
	hashsetX := New()

	// Insert elements into the hashset
	for i := 0; i < 5; i++ {
		el := strconv.Itoa(i)
		hashsetX.Insert(el)
	}

	// Create a copy of the hashset
	copyHashset := Copy(hashsetX)

	// Test if the copied hashset is equal to the original hashset
	for i := 0; i < 5; i++ {
		el := strconv.Itoa(i)
		if !copyHashset.Has(el) {
			t.Errorf("Expected %s to be in the copied hashset, but it's not.", el)
		}
	}
}

func TestRemoveNonExistentElement(t *testing.T) {
	hashset := New()

	// Insert elements into the hashset
	for i := 0; i < 5; i++ {
		el := strconv.Itoa(i)
		hashset.Insert(el)
	}

	// Remove an element that doesn't exist in the hashset
	nonExistentElement := "10"
	hashset.Remove(nonExistentElement)

	// Test if the non-existent element removal did not affect the hashset
	if hashset.len != 5 {
		t.Errorf("Expected hashset length to remain the same after removing a non-existent element.")
	}
}

func TestEmptyHashsetIntersection(t *testing.T) {
	hashsetX := New()
	hashsetY := New()

	// Test intersection of two empty hashsets
	intersection := hashsetX.Intersect(hashsetY)

	// Test if the intersection of two empty hashsets is also an empty hashset
	if intersection.len != 0 {
		t.Errorf("Expected the intersection of two empty hashsets to be empty, but it's not.")
	}
}

func TestEmptyHashsetUnion(t *testing.T) {
	hashsetX := New()
	hashsetY := New()

	// Test union of two empty hashsets
	union := hashsetX.Union(hashsetY)

	// Test if the union of two empty hashsets is also an empty hashset
	if union.len != 0 {
		t.Errorf("Expected the union of two empty hashsets to be empty, but it's not.")
	}
}

func TestEmptySubsetAndSuperset(t *testing.T) {
	hashsetX := New()
	hashsetY := New()

	// Test if an empty hashset is both a subset and a superset of another empty hashset
	if !hashsetX.Subset(hashsetY) || !hashsetY.Superset(hashsetX) {
		t.Errorf("Expected an empty hashset to be both a subset and a superset of another empty hashset, but it's not.")
	}
}

func TestEmptyHashsetHas(t *testing.T) {
	hashset := New()

	// Test if an element exists in an empty hashset
	if hashset.Has("0") {
		t.Errorf("Expected an empty hashset not to have any elements.")
	}
}


func TestHashCollision(t *testing.T) {
	hashset := New()

	// Force hash collision by inserting elements with the same hash
	// This test checks how the Hashset handles hash collisions
	hashset.Insert("a")
	hashset.Insert("b")
	hashset.Insert("aa") // This should hash to the same index as "a"

	if !hashset.Has("a") || !hashset.Has("b") || !hashset.Has("aa") {
		t.Errorf("Expected all elements to be present after hash collision, but some are missing.")
	}
}

func TestIntersect(t *testing.T) {
	hashsetX := New()
	hashsetY := New()

	// Insert elements into the hashsets
	for i := 0; i < 5; i++ {
		el := strconv.Itoa(i)
		hashsetX.Insert(el)
	}
	for i := 3; i < 8; i++ {
		el := strconv.Itoa(i)
		hashsetY.Insert(el)
	}

	// Get the intersection of hashsetX and hashsetY
	intersection := hashsetX.Intersect(hashsetY)

	// Test if the intersection contains elements that are common in both hashsets
	for i := 3; i < 5; i++ {
		el := strconv.Itoa(i)
		if !intersection.Has(el) {
			t.Errorf("Expected %s to be in the intersection, but it's not.", el)
		}
	}
}

func TestUnion(t *testing.T) {
	hashsetX := New()
	hashsetY := New()

	// Insert elements into the hashsets
	for i := 0; i < 5; i++ {
		el := strconv.Itoa(i)
		hashsetX.Insert(el)
	}
	for i := 3; i < 8; i++ {
		el := strconv.Itoa(i)
		hashsetY.Insert(el)
	}

	// Get the union of hashsetX and hashsetY
	union := hashsetX.Union(hashsetY)

	// Test if the union contains all elements from both hashsets
	for i := 0; i < 8; i++ {
		el := strconv.Itoa(i)
		if !union.Has(el) {
			t.Errorf("Expected %s to be in the union, but it's not.", el)
		}
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
