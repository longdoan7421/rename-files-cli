package main

import (
	"reflect"
	"testing"
)

func TestValidatePathFlag(t *testing.T) {
	invalidPath := ""
	expected := false

	actual, _ := validateFlags(invalidPath, "title", 10)
	if actual != expected {
		t.Fatalf(`Expected %t, but got %t`, expected, actual)
	}
}

func TestValidateCaseFlag(t *testing.T) {
	invalidCase := "some-unknown-case"
	expected := false

	actual, _ := validateFlags("/Users/abc/xyz", invalidCase, 10)
	if actual != expected {
		t.Fatalf(`Expected %t, but got %t`, expected, actual)
	}
}

func TestValidateDepthFlag(t *testing.T) {
	invalidDepth := -1
	expected := false

	actual, _ := validateFlags("/Users/abc/xyz", "title", invalidDepth)
	if actual != expected {
		t.Fatalf(`Expected %t, but got %t`, expected, actual)
	}
}

func TestSplitNameWithoutKeepingUpper(t *testing.T) {
	inputs := []string{"This-is_an_random-Name", "This is - an Random Name", "this_is-an_random-Name", "This_is_an-Random_Name", "This is AN random, name"}
	expected := []string{"this", "is", "an", "random", "name"}

	for _, input := range inputs {
		actual := splitNameParts(input, false)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Expected %v, but got %v`, expected, actual)
		}
	}
}

func TestSplitNameWithKeepingUpper(t *testing.T) {
	inputs := []string{"this-is_an_random-NAME", "This is - an Random NAME", "this_is-an_random-NAME", "This_is_an-Random_NAME", "This is an random, NAME"}
	expected := []string{"this", "is", "an", "random", "NAME"}

	for _, input := range inputs {
		actual := splitNameParts(input, true)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Expected %v, but got %v`, expected, actual)
		}
	}
}

func TestCalculateDepth(t *testing.T) {
	rootPath := "/Users/abc/root-dir-depth-1"
	dirPath := rootPath + "/sub-dir-depth-2/sub--dir-depth-3"
	expected := 3

	actual := calculateDepthOfChildDirectory(rootPath, dirPath)
	if actual != expected {
		t.Fatalf(`Expected %d, but got %d`, expected, actual)
	}
}

func TestToTitleCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"This Is an Random Example",
		"An Example Starts With a Short",
	}

	for index, input := range inputs {
		actual := toTitleCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"ThisIsAnRandomExample",
		"AnExampleStartsWithAShort",
	}

	for index, input := range inputs {
		actual := toPascalCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"thisIsAnRandomExample",
		"anExampleStartsWithAShort",
	}

	for index, input := range inputs {
		actual := toCamelCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"this_is_an_random_example",
		"an_example_starts_with_a_short",
	}

	for index, input := range inputs {
		actual := toSnakeCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"this-is-an-random-example",
		"an-example-starts-with-a-short",
	}

	for index, input := range inputs {
		actual := toKebabCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}

func TestToPascalSnakeCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"This_Is_An_Random_Example",
		"An_Example_Starts_With_A_Short",
	}

	for index, input := range inputs {
		actual := toPascalSnakeCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}

func TestToPascalKebabCase(t *testing.T) {
	inputs := [][]string{
		{"this", "is", "an", "random", "example"},
		{"an", "example", "starts", "with", "a", "short"},
	}
	expectedOutputs := []string{
		"This-Is-An-Random-Example",
		"An-Example-Starts-With-A-Short",
	}

	for index, input := range inputs {
		actual := toPascalKebabCaseString(input)
		if actual != expectedOutputs[index] {
			t.Fatalf(`Expected %s, but got %s`, expectedOutputs[index], actual)
		}
	}
}
