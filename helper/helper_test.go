package helper_test

import (
	"testing"

	"github.com/plaub/go-test-project/helper"
)

func TestGetArrayValueSumWithInt64(t *testing.T) {
	expected := int64(100)
	actual := helper.GetArrayValueSum([]int64{10, 20, 30, 40})

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetArrayValueSumWithFloat64(t *testing.T) {
	expected := float64(100)
	actual := helper.GetArrayValueSum([]float64{10, 20, 30, 40})

	if expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestGetArrayValuesAverageWithInt64(t *testing.T) {
	expected := int64(25)
	actual := helper.GetArrayValuesAverage([]int64{10, 20, 30, 40})

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetArrayValuesAverageWithFloat64(t *testing.T) {
	expected := float64(25)
	actual := helper.GetArrayValuesAverage([]float64{10, 20, 30, 40})

	if expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}
