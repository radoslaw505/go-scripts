package main

import (
	"fmt"
	"sync"
)

func main() {

	// Declare and initialize a map with string keys and int values.
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	// Make a map with string keys and string values.
	contacts := make(map[string]string)

	// Add key-value pairs to the contacts map.
	contacts["Alice"] = "alice@example.com"
	contacts["Bob"] = "bob@example.com"
	contacts["Charlie"] = "charlie@example.com"

	// Print the maps.
	fmt.Println("Ages:", ages)
	fmt.Println("Contacts:", contacts)

	// Access and print individual values using their keys.
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Bob's contact:", contacts["Bob"])

	// Update a value in the ages map.
	ages["Alice"] = 31
	fmt.Println("Updated Alice's age:", ages["Alice"])

	// Delete a key-value pair from the contacts map.
	delete(contacts, "Charlie")
	fmt.Println("Contacts after deleting Charlie:", contacts)

	// Check if a key exists in the ages map.
	age, exists := ages["Charlie"]
	if exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found in ages map.")
	}

	// Iterate over the ages map and print each key-value pair.
	fmt.Println("Iterating over ages map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Get the number of key-value pairs in the ages map.
	fmt.Println("Number of entries in ages map:", len(ages))

	// Nested map: map with string keys and map values.
	nestedMap := map[string]map[string]int{
		"Group1": {
			"Alice": 30,
			"Bob":   25,
		},
		"Group2": {
			"Charlie": 35,
			"David":   28,
		},
	}

	// Print the nested map.
	fmt.Println("Nested Map:", nestedMap)

	// Access a value in the nested map.
	fmt.Println("Alice's age in Group1:", nestedMap["Group1"]["Alice"])

	// Iterate over the nested map.
	fmt.Println("Iterating over nested map:")
	for group, members := range nestedMap {
		fmt.Println("Members of", group+":")
		for name, age := range members {
			fmt.Printf("  %s is %d years old.\n", name, age)
		}
	}

	// Compare maps (maps cannot be compared directly except to nil).
	// Here we demonstrate that two maps with the same content are not equal.
	map1 := map[string]int{"A": 1, "B": 2}
	map2 := map[string]int{"A": 1, "B": 2}
	fmt.Println("map1 and map2 are equal:", fmt.Sprintf("%v", map1) == fmt.Sprintf("%v", map2))

	// Clear function: reinitialize the map to an empty map.
	clear(ages)
	fmt.Println("Ages map after clearing:", ages)

	// Demonstrate the zero value behavior when accessing a non-existent key.
	_, unknown := contacts["Alice"]
	fmt.Println("Check if value is associated with a key 'Alice':", unknown)

	// Note: Maps are not safe for concurrent use. For concurrent access, consider using sync.Map or other synchronization mechanisms.
	// Example of using sync.Map:
	/*
		var m sync.Map
		m.Store("key", "value")
		value, ok := m.Load("key")
	*/

	var m sync.Map
	m.Store("Alice", "31")
	value, ok := m.Load("Alice")
	if ok {
		fmt.Println("Value from sync.Map:", value)
	} else {
		fmt.Println("Key not found in sync.Map")
	}

	// nille maps
	var nilMap map[string]int
	if nilMap == nil {
		fmt.Println("nilMap is nil")
	} else {
		fmt.Println("nilMap is not nil")
	}
	// Attempting to add a key-value pair to a nil map will cause a runtime panic.
	// Uncommenting the following line will result in a panic.
	// nilMap["NewKey"] = 100

	// To use a map, it must be initialized first.
	nilMap = make(map[string]int)
	nilMap["NewKey"] = 100
	fmt.Println("Initialized nil map after adding a key:", nilMap)

}
