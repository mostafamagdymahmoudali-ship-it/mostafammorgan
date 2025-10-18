package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Read the JSON file
	byteValue, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal into a slice of maps
	var data []map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("No data found in JSON.")
		return
	}

	// Prepare headers
	headers := []string{}
	for k := range data[0] {
		headers = append(headers, k)
	}

	// Create CSV file
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		fmt.Println("Error writing headers:", err)
		return
	}

	// Write data rows
	for _, row := range data {
		record := make([]string, len(headers))
		for i, h := range headers {
			if val, ok := row[h]; ok {
				record[i] = fmt.Sprintf("%v", val)
			}
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}

	fmt.Println("CSV file created: output.csv")
}
