package utility

import "strings"

func WrapTask(tasks []string) string {
	var task string
	for _, value := range tasks {
		task += value + " "
	}
	return strings.TrimSpace(task)
}