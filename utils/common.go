package utils

import (
	"os"
)

func Map(arr []string, fn func(index int, value string) string) []string {
	var newArr []string
	for idx, element := range arr {
		newArr = append(newArr, fn(idx, element))
	}

	return newArr
}

func Filter(arr []string, fn func(index int, value string) bool) []string {
	var newArr []string
	for idx, element := range arr {
		if fn(idx, element) {
			newArr = append(newArr, element)
		}
	}

	return newArr
}

func Contains(arr []string, needCheckItem string) bool {
	for _, item := range arr {
		if item == needCheckItem {
			return true
		}
	}

	return false
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}
