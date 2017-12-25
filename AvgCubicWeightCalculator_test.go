package main

import (
	"math"
	"testing"
)

// Testing function signature:  func avgCubicWeight(weightTotal, productTotal float64) float64 {...}
func TestAvgCubicWeight(t *testing.T) {
	F64_TOLERANCE := 0.000001	// tolerance parameter to check floating points results within

	// testcase for positive total weight and zero product total - must return error
	val1, err := avgCubicWeight(200, 0)
	if err == nil {
		t.Errorf(`avgCubicWeight(200, 0) did not return an error (return value %f) `, val1)
	}

	// testcase for zero total weight and zero product total - must return zero and nil error
	val2, err := avgCubicWeight(0, 0)

	if err != nil {
		t.Errorf(`avgCubicWeight(0, 0) returned an error: %v`, err)
	}

	if val2 != 0.0 {
		t.Errorf(`avgCubicWeight(0, 0) must return zero (returned %f)`, val2)
	}

	// testcase for negative total weight and negative product total - must return error
	val3, err := avgCubicWeight(-1, -1)

	if err == nil {
		t.Errorf(`avgCubicWeight(-1, -1) did not return an error (return value %f) `, val3)
	}

	// testcase for positive total weight and positive product total - returns non-decimal result
	val4, err := avgCubicWeight(200, 2)

	if err != nil {
		t.Errorf(`avgCubicWeight(200, 2) returned an error: %v`, err)
	}

	if val4 != 100 {
		t.Errorf(`avgCubicWeight(200, 2) != 100 (got %f)`, val1)
	}

	// testcase for positive total weight and positive product total - returns non-decimal result
	val5, err := avgCubicWeight(350, 35)

	if err != nil {
		t.Errorf(`avgCubicWeight(350, 35) returned an error: %v`, err)
	}

	if val5 != 10 {
		t.Errorf(`avgCubicWeight(350, 35) != 10 (got %f)`, val5)
	}

	// testcase for positive total weight and positive product total - returns decimal result, and is tested for being within the tolerance parameter
	val6, err := avgCubicWeight(3216841, 354)

	if err != nil {
		t.Errorf(`avgCubicWeight(3216841, 354) returned an error: %v`, err)
	}

	if math.Abs(val6 - 9087.121468927) > F64_TOLERANCE {
		t.Errorf(`avgCubicWeight(3216841, 354) - 9087.121468927 > %f (got %f)`, F64_TOLERANCE, val6)
	}

	// testcase for positive total weight and positive product total - returns decimal result, and is tested for being within the tolerance parameter
	val7, err := avgCubicWeight(321986454, 321825)

	if err != nil {
		t.Errorf(`avgCubicWeight(321986454, 321825) returned an error: %v`, err)
	}

	if math.Abs(val7 - 1000.501682591) > F64_TOLERANCE {
		t.Errorf(`avgCubicWeight(321986454, 321825) - 1000.501682591 > %f (got %f)`, F64_TOLERANCE, val7)
	}
}

// Testing function signature: func cubicWeight(pz *ProductSize, cubicWeightConversionFactor float64) (float64, error) {...}
func TestCubicWeight(t *testing.T) {
	var cubicWeightConversionFactor float64
	cubicWeightConversionFactor = 250
	F64_TOLERANCE := 0.000001

	val1, err := cubicWeight(&ProductSize{1.0, 1.0, 1.0}, cubicWeightConversionFactor)
	if err != nil {
		t.Errorf(`cubicWeight(&ProductSize{1.0, 1.0, 1.0}, %f) returned error %v`, err)
	}

	if math.Abs(val1 - 0.00025) > F64_TOLERANCE {
		t.Errorf(`cubicWeight(&ProductSize{1.0, 1.0, 1.0}, %f) - 1000.501682591 > %f (got %f)`, cubicWeightConversionFactor, F64_TOLERANCE, val1)
	}

	_, err = cubicWeight(&ProductSize{-20.5, 30.0, 40.0}, cubicWeightConversionFactor)
	if err == nil {
		t.Errorf(`cubicWeight(&ProductSize{-20.5, 30.0, 40.0}, %f) did not return an error`)
	}

	_, err = cubicWeight(&ProductSize{0, 0, 0}, cubicWeightConversionFactor)
	if err == nil {
		t.Errorf(`cubicWeight(&ProductSize{0, 0, 0}, %f) did not return an error`)
	}

	val2, err := cubicWeight(&ProductSize{40, 20, 30}, cubicWeightConversionFactor)
	if err != nil {
		t.Errorf(`cubicWeight(&ProductSize{40, 20, 30}, %f) returned error %v`, err)
	}

	if math.Abs(val2 - 6.0) > F64_TOLERANCE {
		t.Errorf(`cubicWeight(&ProductSize{40, 20, 30}, %f) - 6.0 > %f (got %f)`, cubicWeightConversionFactor, F64_TOLERANCE, val2)
	}
}
