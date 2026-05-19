package main

import "fmt"

func main() {
	inventory := map[string]int{
		"apple":  5,
		"orange": 3,
	}

	// The comma-ok pattern distinguishes "key exists with zero value" from "key missing"
	// Without comma-ok: both missing keys and zero-value keys return 0
	fmt.Println("banana count:", inventory["banana"]) // 0 — but is it out of stock or never tracked?

	// With comma-ok: the second return value tells you if the key exists
	count, ok := inventory["apple"]
	fmt.Printf("apple: count=%d, exists=%t\n", count, ok)

	count, ok = inventory["banana"]
	fmt.Printf("banana: count=%d, exists=%t\n", count, ok)

	// Common pattern: conditional logic based on key existence
	if qty, exists := inventory["orange"]; exists {
		fmt.Printf("orange is in stock: %d units\n", qty)
	} else {
		fmt.Println("orange is not tracked")
	}

	if qty, exists := inventory["mango"]; exists {
		fmt.Printf("mango is in stock: %d units\n", qty)
	} else {
		fmt.Println("mango is not tracked")
	}

	// Why this matters: zero value vs absence
	inventory["grape"] = 0 // grape exists but is out of stock
	if _, exists := inventory["grape"]; exists {
		fmt.Println("grape is tracked (even though count is 0)")
	}
} 