package main

import (
	"encoding/json"
)

type Student struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var students []Student

	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}
	classMap := make(map[string][]Student)
	for _, student := range students {
		classMap[student.Class] = append(classMap[student.Class], student)
	}
	result := make(map[string][]byte)
	for class, stds := range classMap {
		bytesStds, err := json.Marshal(stds)
		if err != nil {
			return nil, err
		}
		result[class] = bytesStds
	}
	return result, nil
}
