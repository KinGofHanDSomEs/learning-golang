package main

func DeleteLongKeys(m map[string]int) map[string]int {
	new_map := make(map[string]int)
	for key, value := range m {
		if len(key) >= 6 {
			new_map[key] = value
		}
	}
	return new_map
}
