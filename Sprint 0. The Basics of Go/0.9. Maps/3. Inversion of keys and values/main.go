package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	new_map := make(map[string]string)
	for key, value := range m {
		new_map[value] = key
	}
	return new_map
}
