package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		fmt.Println("unmarshall")
		return nil, err
	}
	for i := range students {
		students[i].Grade += 1
	}
	jsonBytes, err := json.Marshal(students)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
