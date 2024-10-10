package main

import (
	"encoding/json"
	"fmt"
)

func processData(line string) {
	var parsedData map[string]interface{}
	err := json.Unmarshal([]byte(line), &parsedData)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	// Временное решение
	if request, ok := parsedData["request"].(map[string]interface{}); ok {
		if mountPoint, ok := request["mount_point"].(string); ok && mountPoint == "secret/" {
			// TODO: отправить в telegram
			fmt.Println("Получен запрос с mount_point 'secret/':", line)
		}
	}
}
