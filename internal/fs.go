package internal

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/viper"
)

func getFileName() string {
	return viper.GetString("FILE_PATH")
}

func readTasksFromDisk() ([]Task, error) {

	jsonFile, err := os.OpenFile(getFileName(), os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	return tasks, json.Unmarshal(byteValue, &tasks)
}

func writeTaskToDisk(tasks *[]Task) error {
	jsonFile, err := os.OpenFile(getFileName(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := json.Marshal(*tasks)

	if err != nil {
		return err
	}

	_, writeErr := jsonFile.Write(byteValue)

	return writeErr
}
