package conv

import (
	"reflect"
	"strings"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToJSONStr(t *testing.T) {
	t.Run("should marshal data into JSON string successfully", func(t *testing.T) {
		testData := TestStruct{Name: "Alice", Age: 25}
		expectedResult := `{"name":"Alice","age":25}`

		result, err := ToJSONStr(testData)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if result != expectedResult {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})

	t.Run("should return an error when unable to marshal data", func(t *testing.T) {
		testData := make(chan int) //channels cannot be serialized to JSON

		_, err := ToJSONStr(testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to marshal data") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})
}

func TestToMap(t *testing.T) {
	t.Run("should unmarshal data into map successfully", func(t *testing.T) {
		testData := TestStruct{Name: "Alice", Age: 25}
		expectedResult := map[string]interface{}{"name": "Alice", "age": float64(25)} // json.Unmarshal converts numbers to float64

		result, err := ToMap(testData)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})

	t.Run("should return an error when unable to marshal data", func(t *testing.T) {
		testData := make(chan int) //channels cannot be serialized to JSON

		_, err := ToMap(testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to marshal data") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})

	t.Run("should return an error when unable to unmarshal data into map", func(t *testing.T) {
		testData := "not a valid json string" // string cannot be unmarshalled to map

		_, err := ToMap(testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to unmarshal data into map") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})
}

func TestToArray(t *testing.T) {
	t.Run("should marshal and unmarshal data into array of struct successfully", func(t *testing.T) {
		testData := []TestStruct{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}}
		expectedResult := []TestStruct{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}}

		result, err := ToArray[TestStruct](testData)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})

	t.Run("should marshal and unmarshal data into array of ints successfully", func(t *testing.T) {
		testData := []int{1, 2, 3, 4, 5}

		result, err := ToArray[int](testData)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, testData) {
			t.Errorf("expected %v, got %v", testData, result)
		}
	})

	t.Run("should marshal and unmarshal data into array of strings successfully", func(t *testing.T) {
		testData := []string{"a", "b", "c"}

		result, err := ToArray[string](testData)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, testData) {
			t.Errorf("expected %v, got %v", testData, result)
		}
	})

	t.Run("should return an error when unable to marshal data", func(t *testing.T) {
		testData := make(chan int) //channels cannot be serialized to JSON

		_, err := ToArray[int](testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to marshal data") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})

	t.Run("should return an error when unable to unmarshal data into array", func(t *testing.T) {
		testData := "not an array" // string cannot be unmarshalled to array

		_, err := ToArray[int](testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to unmarshal data into array") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})
}

func TestToStruct(t *testing.T) {
	t.Run("should marshal and unmarshal data into struct successfully", func(t *testing.T) {
		testData := TestStruct{Name: "Alice", Age: 25}

		result, err := ToStruct[TestStruct](testData)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, &testData) {
			t.Errorf("expected %v, got %v", &testData, result)
		}
	})

	t.Run("should return an error when unable to marshal data", func(t *testing.T) {
		testData := make(chan int) //channels cannot be serialized to JSON

		_, err := ToStruct[TestStruct](testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to marshal data") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})

	t.Run("should return an error when unable to unmarshal data into target struct", func(t *testing.T) {
		testData := "not a valid json string" // string cannot be unmarshalled to struct

		_, err := ToStruct[TestStruct](testData)
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "unable to unmarshal data into target struct") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})
}
