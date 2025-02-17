package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var result []Student
	for i := range jsonDataList {
		var students []Student
		err := json.Unmarshal(jsonDataList[i], &students)
		if err != nil {
			return nil, err
		}
		for j := range students {
			fmt.Println(students[j])
			result = append(result, students[j])
		}
	}
	bytesJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return bytesJson, nil
}
