// Copyright 2011 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ogletest

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

var someInt int = -17

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type equalsTestCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func checkTestCases(t *testing.T, matcher Matcher, cases []equalsTestCase) {
	for i, c := range cases {
		result, err := matcher.Matches(c.candidate)

		if result != c.expectedResult {
			t.Errorf(
				"Case %d (candidate %v): expected %v, got %v",
				i,
				c.candidate,
				c.expectedResult,
				result)
		}

		actualError := ""
		if err != nil {
			actualError = err.Error()
		}

		if actualError != c.expectedError {
			t.Errorf("Case %d: expected error %v, got %v", i, c.expectedError, err)
		}
	}
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func TestNegativeIntegerLiteral(t *testing.T) {
	// -2^30
	matcher := Equals(-1073741824)
	desc := matcher.Description()
	expectedDesc := "-1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -1073741824.
		equalsTestCase{-1073741824, MATCH_TRUE, ""},
		equalsTestCase{-1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{-1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegerLiteral(t *testing.T) {
	// 2^30
	matcher := Equals(1073741824)
	desc := matcher.Description()
	expectedDesc := "1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 1073741824.
		equalsTestCase{1073741824, MATCH_TRUE, ""},
		equalsTestCase{1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(1073741823), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// Floating point literals
////////////////////////////////////////////////////////////

func TestNegativeIntegralFloatingPointLiteral(t *testing.T) {
	// -2^30
	matcher := Equals(-1073741824.0)
	desc := matcher.Description()
	expectedDesc := "-1.073741824e+09"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -1073741824.
		equalsTestCase{-1073741824, MATCH_TRUE, ""},
		equalsTestCase{-1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{-1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloatingPointLiteral(t *testing.T) {
	// 2^30
	matcher := Equals(1073741824.0)
	desc := matcher.Description()
	expectedDesc := "1.073741824e+09"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 1073741824.
		equalsTestCase{1073741824, MATCH_TRUE, ""},
		equalsTestCase{1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(1073741824)), MATCH_TRUE, ""},

		// Values that would be 1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(1073741823), MATCH_FALSE, ""},
		equalsTestCase{uint(1073741823), MATCH_FALSE, ""},
		equalsTestCase{uint32(1073741823), MATCH_FALSE, ""},
		equalsTestCase{uint64(1073741823), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonIntegralFloatingPointLiteral(t *testing.T) {
	matcher := Equals(17.1)
	desc := matcher.Description()
	expectedDesc := "17.1"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 17.1.
		equalsTestCase{17.1, MATCH_TRUE, ""},
		equalsTestCase{17.1, MATCH_TRUE, ""},
		equalsTestCase{17.1 + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(17.1), MATCH_TRUE, ""},
		equalsTestCase{float64(17.1), MATCH_TRUE, ""},
		equalsTestCase{complex64(17.1), MATCH_TRUE, ""},
		equalsTestCase{complex128(17.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{17, MATCH_FALSE, ""},
		equalsTestCase{17.2, MATCH_FALSE, ""},
		equalsTestCase{18, MATCH_FALSE, ""},
		equalsTestCase{int(17), MATCH_FALSE, ""},
		equalsTestCase{int(18), MATCH_FALSE, ""},
		equalsTestCase{int32(17), MATCH_FALSE, ""},
		equalsTestCase{int64(17), MATCH_FALSE, ""},
		equalsTestCase{uint(17), MATCH_FALSE, ""},
		equalsTestCase{uint32(17), MATCH_FALSE, ""},
		equalsTestCase{uint64(17), MATCH_FALSE, ""},
		equalsTestCase{complex128(17.1 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// bool
////////////////////////////////////////////////////////////

func TestFalse(t *testing.T) {
	matcher := Equals(false)
	desc := matcher.Description()
	expectedDesc := "false"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// bools
		equalsTestCase{false, MATCH_TRUE, ""},
		equalsTestCase{bool(false), MATCH_TRUE, ""},

		equalsTestCase{true, MATCH_FALSE, ""},
		equalsTestCase{bool(true), MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a bool"},
	}

	checkTestCases(t, matcher, cases)
}

func TestTrue(t *testing.T) {
	matcher := Equals(true)
	desc := matcher.Description()
	expectedDesc := "true"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// bools
		equalsTestCase{true, MATCH_TRUE, ""},
		equalsTestCase{bool(true), MATCH_TRUE, ""},

		equalsTestCase{false, MATCH_FALSE, ""},
		equalsTestCase{bool(false), MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{int(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int8(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int16(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int32(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{int64(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint8(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint16(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint32(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uint64(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{uintptr(1), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a bool"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a bool"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int
////////////////////////////////////////////////////////////

func TestNegativeInt(t *testing.T) {
	// -2^30
	matcher := Equals(int(-1073741824))
	desc := matcher.Description()
	expectedDesc := "-1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -1073741824.
		equalsTestCase{-1073741824, MATCH_TRUE, ""},
		equalsTestCase{-1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{-1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt(t *testing.T) {
	// 2^30
	matcher := Equals(int(1073741824))
	desc := matcher.Description()
	expectedDesc := "1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 1073741824.
		equalsTestCase{1073741824, MATCH_TRUE, ""},
		equalsTestCase{1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(1073741823), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int8
////////////////////////////////////////////////////////////

func TestNegativeInt8(t *testing.T) {
	matcher := Equals(int8(-17))
	desc := matcher.Description()
	expectedDesc := "-17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -17.
		equalsTestCase{-17, MATCH_TRUE, ""},
		equalsTestCase{-17.0, MATCH_TRUE, ""},
		equalsTestCase{-17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(-17), MATCH_TRUE, ""},
		equalsTestCase{int8(-17), MATCH_TRUE, ""},
		equalsTestCase{int16(-17), MATCH_TRUE, ""},
		equalsTestCase{int32(-17), MATCH_TRUE, ""},
		equalsTestCase{int64(-17), MATCH_TRUE, ""},
		equalsTestCase{float32(-17), MATCH_TRUE, ""},
		equalsTestCase{float64(-17), MATCH_TRUE, ""},
		equalsTestCase{complex64(-17), MATCH_TRUE, ""},
		equalsTestCase{complex128(-17), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(-17)), MATCH_TRUE, ""},

		// Values that would be -17 in two's complement.
		equalsTestCase{uint((1 << 32) - 17), MATCH_FALSE, ""},
		equalsTestCase{uint8((1 << 8) - 17), MATCH_FALSE, ""},
		equalsTestCase{uint16((1 << 16) - 17), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 17), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 17), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-16), MATCH_FALSE, ""},
		equalsTestCase{int8(-16), MATCH_FALSE, ""},
		equalsTestCase{int16(-16), MATCH_FALSE, ""},
		equalsTestCase{int32(-16), MATCH_FALSE, ""},
		equalsTestCase{int64(-16), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float32(-17.1), MATCH_FALSE, ""},
		equalsTestCase{float32(-16.9), MATCH_FALSE, ""},
		equalsTestCase{complex64(-16), MATCH_FALSE, ""},
		equalsTestCase{complex64(-17 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr((1 << 32) - 17), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{-17}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{-17}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"-17", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroInt8(t *testing.T) {
	matcher := Equals(int8(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 0.
		equalsTestCase{0, MATCH_TRUE, ""},
		equalsTestCase{0.0, MATCH_TRUE, ""},
		equalsTestCase{0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(0), MATCH_TRUE, ""},
		equalsTestCase{int8(0), MATCH_TRUE, ""},
		equalsTestCase{int16(0), MATCH_TRUE, ""},
		equalsTestCase{int32(0), MATCH_TRUE, ""},
		equalsTestCase{int64(0), MATCH_TRUE, ""},
		equalsTestCase{float32(0), MATCH_TRUE, ""},
		equalsTestCase{float64(0), MATCH_TRUE, ""},
		equalsTestCase{complex64(0), MATCH_TRUE, ""},
		equalsTestCase{complex128(0), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(0)), MATCH_TRUE, ""},
		equalsTestCase{uint(0), MATCH_TRUE, ""},
		equalsTestCase{uint8(0), MATCH_TRUE, ""},
		equalsTestCase{uint16(0), MATCH_TRUE, ""},
		equalsTestCase{uint32(0), MATCH_TRUE, ""},
		equalsTestCase{uint64(0), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1), MATCH_FALSE, ""},
		equalsTestCase{int8(1), MATCH_FALSE, ""},
		equalsTestCase{int16(1), MATCH_FALSE, ""},
		equalsTestCase{int32(1), MATCH_FALSE, ""},
		equalsTestCase{int64(1), MATCH_FALSE, ""},
		equalsTestCase{float32(-0.1), MATCH_FALSE, ""},
		equalsTestCase{float32(0.1), MATCH_FALSE, ""},
		equalsTestCase{complex64(1), MATCH_FALSE, ""},
		equalsTestCase{complex64(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"0", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt8(t *testing.T) {
	matcher := Equals(int8(17))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 17.
		equalsTestCase{17, MATCH_TRUE, ""},
		equalsTestCase{17.0, MATCH_TRUE, ""},
		equalsTestCase{17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(17), MATCH_TRUE, ""},
		equalsTestCase{int8(17), MATCH_TRUE, ""},
		equalsTestCase{int16(17), MATCH_TRUE, ""},
		equalsTestCase{int32(17), MATCH_TRUE, ""},
		equalsTestCase{int64(17), MATCH_TRUE, ""},
		equalsTestCase{float32(17), MATCH_TRUE, ""},
		equalsTestCase{float64(17), MATCH_TRUE, ""},
		equalsTestCase{complex64(17), MATCH_TRUE, ""},
		equalsTestCase{complex128(17), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(17)), MATCH_TRUE, ""},
		equalsTestCase{uint(17), MATCH_TRUE, ""},
		equalsTestCase{uint8(17), MATCH_TRUE, ""},
		equalsTestCase{uint16(17), MATCH_TRUE, ""},
		equalsTestCase{uint32(17), MATCH_TRUE, ""},
		equalsTestCase{uint64(17), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(16), MATCH_FALSE, ""},
		equalsTestCase{int8(16), MATCH_FALSE, ""},
		equalsTestCase{int16(16), MATCH_FALSE, ""},
		equalsTestCase{int32(16), MATCH_FALSE, ""},
		equalsTestCase{int64(16), MATCH_FALSE, ""},
		equalsTestCase{float32(16.9), MATCH_FALSE, ""},
		equalsTestCase{float32(17.1), MATCH_FALSE, ""},
		equalsTestCase{complex64(16), MATCH_FALSE, ""},
		equalsTestCase{complex64(17 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(17), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{17}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{17}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"17", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int16
////////////////////////////////////////////////////////////

func TestNegativeInt16(t *testing.T) {
	matcher := Equals(int16(-32766))
	desc := matcher.Description()
	expectedDesc := "-32766"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -32766.
		equalsTestCase{-32766, MATCH_TRUE, ""},
		equalsTestCase{-32766.0, MATCH_TRUE, ""},
		equalsTestCase{-32766 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(-32766), MATCH_TRUE, ""},
		equalsTestCase{int16(-32766), MATCH_TRUE, ""},
		equalsTestCase{int32(-32766), MATCH_TRUE, ""},
		equalsTestCase{int64(-32766), MATCH_TRUE, ""},
		equalsTestCase{float32(-32766), MATCH_TRUE, ""},
		equalsTestCase{float64(-32766), MATCH_TRUE, ""},
		equalsTestCase{complex64(-32766), MATCH_TRUE, ""},
		equalsTestCase{complex128(-32766), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(-32766)), MATCH_TRUE, ""},

		// Values that would be -32766 in two's complement.
		equalsTestCase{uint((1 << 32) - 32766), MATCH_FALSE, ""},
		equalsTestCase{uint16((1 << 16) - 32766), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 32766), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 32766), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-16), MATCH_FALSE, ""},
		equalsTestCase{int8(-16), MATCH_FALSE, ""},
		equalsTestCase{int16(-16), MATCH_FALSE, ""},
		equalsTestCase{int32(-16), MATCH_FALSE, ""},
		equalsTestCase{int64(-16), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float32(-32766.1), MATCH_FALSE, ""},
		equalsTestCase{float32(-32765.9), MATCH_FALSE, ""},
		equalsTestCase{complex64(-32766.1), MATCH_FALSE, ""},
		equalsTestCase{complex64(-32766 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr((1 << 32) - 32766), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{-32766}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{-32766}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"-32766", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroInt16(t *testing.T) {
	matcher := Equals(int16(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 0.
		equalsTestCase{0, MATCH_TRUE, ""},
		equalsTestCase{0.0, MATCH_TRUE, ""},
		equalsTestCase{0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(0), MATCH_TRUE, ""},
		equalsTestCase{int8(0), MATCH_TRUE, ""},
		equalsTestCase{int16(0), MATCH_TRUE, ""},
		equalsTestCase{int32(0), MATCH_TRUE, ""},
		equalsTestCase{int64(0), MATCH_TRUE, ""},
		equalsTestCase{float32(0), MATCH_TRUE, ""},
		equalsTestCase{float64(0), MATCH_TRUE, ""},
		equalsTestCase{complex64(0), MATCH_TRUE, ""},
		equalsTestCase{complex128(0), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(0)), MATCH_TRUE, ""},
		equalsTestCase{uint(0), MATCH_TRUE, ""},
		equalsTestCase{uint8(0), MATCH_TRUE, ""},
		equalsTestCase{uint16(0), MATCH_TRUE, ""},
		equalsTestCase{uint32(0), MATCH_TRUE, ""},
		equalsTestCase{uint64(0), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1), MATCH_FALSE, ""},
		equalsTestCase{int8(1), MATCH_FALSE, ""},
		equalsTestCase{int16(1), MATCH_FALSE, ""},
		equalsTestCase{int32(1), MATCH_FALSE, ""},
		equalsTestCase{int64(1), MATCH_FALSE, ""},
		equalsTestCase{float32(-0.1), MATCH_FALSE, ""},
		equalsTestCase{float32(0.1), MATCH_FALSE, ""},
		equalsTestCase{complex64(1), MATCH_FALSE, ""},
		equalsTestCase{complex64(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"0", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt16(t *testing.T) {
	matcher := Equals(int16(32765))
	desc := matcher.Description()
	expectedDesc := "32765"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 32765.
		equalsTestCase{32765, MATCH_TRUE, ""},
		equalsTestCase{32765.0, MATCH_TRUE, ""},
		equalsTestCase{32765 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(32765), MATCH_TRUE, ""},
		equalsTestCase{int16(32765), MATCH_TRUE, ""},
		equalsTestCase{int32(32765), MATCH_TRUE, ""},
		equalsTestCase{int64(32765), MATCH_TRUE, ""},
		equalsTestCase{float32(32765), MATCH_TRUE, ""},
		equalsTestCase{float64(32765), MATCH_TRUE, ""},
		equalsTestCase{complex64(32765), MATCH_TRUE, ""},
		equalsTestCase{complex128(32765), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(32765)), MATCH_TRUE, ""},
		equalsTestCase{uint(32765), MATCH_TRUE, ""},
		equalsTestCase{uint16(32765), MATCH_TRUE, ""},
		equalsTestCase{uint32(32765), MATCH_TRUE, ""},
		equalsTestCase{uint64(32765), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(32764), MATCH_FALSE, ""},
		equalsTestCase{int16(32764), MATCH_FALSE, ""},
		equalsTestCase{int32(32764), MATCH_FALSE, ""},
		equalsTestCase{int64(32764), MATCH_FALSE, ""},
		equalsTestCase{float32(32764.9), MATCH_FALSE, ""},
		equalsTestCase{float32(32765.1), MATCH_FALSE, ""},
		equalsTestCase{complex64(32765.9), MATCH_FALSE, ""},
		equalsTestCase{complex64(32765 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(32765), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{32765}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{32765}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"32765", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int32
////////////////////////////////////////////////////////////

func TestNegativeInt32(t *testing.T) {
	// -2^30
	matcher := Equals(int32(-1073741824))
	desc := matcher.Description()
	expectedDesc := "-1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -1073741824.
		equalsTestCase{-1073741824, MATCH_TRUE, ""},
		equalsTestCase{-1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{-1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(-1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt32(t *testing.T) {
	// 2^30
	matcher := Equals(int32(1073741824))
	desc := matcher.Description()
	expectedDesc := "1073741824"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 1073741824.
		equalsTestCase{1073741824, MATCH_TRUE, ""},
		equalsTestCase{1073741824.0, MATCH_TRUE, ""},
		equalsTestCase{1073741824 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{int64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{uint64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float32(1073741824), MATCH_TRUE, ""},
		equalsTestCase{float64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex64(1073741824), MATCH_TRUE, ""},
		equalsTestCase{complex128(1073741824), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int32(1073741823), MATCH_FALSE, ""},
		equalsTestCase{int64(1073741823), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741824.1), MATCH_FALSE, ""},
		equalsTestCase{float64(1073741823.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741823), MATCH_FALSE, ""},
		equalsTestCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// int64
////////////////////////////////////////////////////////////

func TestNegativeInt64(t *testing.T) {
	// -2^40
	matcher := Equals(int64(-1099511627776))
	desc := matcher.Description()
	expectedDesc := "-1099511627776"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -1099511627776.
		equalsTestCase{-1099511627776.0, MATCH_TRUE, ""},
		equalsTestCase{-1099511627776 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int64(-1099511627776), MATCH_TRUE, ""},
		equalsTestCase{float32(-1099511627776), MATCH_TRUE, ""},
		equalsTestCase{float64(-1099511627776), MATCH_TRUE, ""},
		equalsTestCase{complex64(-1099511627776), MATCH_TRUE, ""},
		equalsTestCase{complex128(-1099511627776), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int64(-1099511627776)), MATCH_TRUE, ""},

		// Values that would be -1099511627776 in two's complement.
		equalsTestCase{uint64((1 << 64) - 1099511627776), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int64(-1099511627775), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1099511627776.1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1099511627775.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1099511627775), MATCH_FALSE, ""},
		equalsTestCase{complex128(-1099511627776 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveInt64(t *testing.T) {
	// 2^40
	matcher := Equals(int64(1099511627776))
	desc := matcher.Description()
	expectedDesc := "1099511627776"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 1099511627776.
		equalsTestCase{1099511627776.0, MATCH_TRUE, ""},
		equalsTestCase{1099511627776 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int64(1099511627776), MATCH_TRUE, ""},
		equalsTestCase{uint64(1099511627776), MATCH_TRUE, ""},
		equalsTestCase{float32(1099511627776), MATCH_TRUE, ""},
		equalsTestCase{float64(1099511627776), MATCH_TRUE, ""},
		equalsTestCase{complex64(1099511627776), MATCH_TRUE, ""},
		equalsTestCase{complex128(1099511627776), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int64(1099511627776)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(uint64(1099511627776)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1099511627775), MATCH_FALSE, ""},
		equalsTestCase{uint64(1099511627775), MATCH_FALSE, ""},
		equalsTestCase{float64(1099511627776.1), MATCH_FALSE, ""},
		equalsTestCase{float64(1099511627775.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(1099511627775), MATCH_FALSE, ""},
		equalsTestCase{complex128(1099511627776 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestInt64NotExactlyRepresentableBySinglePrecision(t *testing.T) {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := Equals(int64(kTwoTo25 + 1))
	desc := matcher.Description()
	expectedDesc := "33554433"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestInt64NotExactlyRepresentableByDoublePrecision(t *testing.T) {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := Equals(int64(kTwoTo54 + 1))
	desc := matcher.Description()
	expectedDesc := "18014398509481985"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo54 + 0), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo54 + 0), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// uint
////////////////////////////////////////////////////////////

func TestSmallUint(t *testing.T) {
	const kExpected = 17
	matcher := Equals(uint(kExpected))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{17, MATCH_TRUE, ""},
		equalsTestCase{17.0, MATCH_TRUE, ""},
		equalsTestCase{17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, MATCH_FALSE, ""},
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeUint(t *testing.T) {
	const kExpected = (1 << 16) + 17
	matcher := Equals(uint(kExpected))
	desc := matcher.Description()
	expectedDesc := "65553"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{65553, MATCH_TRUE, ""},
		equalsTestCase{65553.0, MATCH_TRUE, ""},
		equalsTestCase{65553 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int16(17), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(17), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestUintNotExactlyRepresentableBySinglePrecision(t *testing.T) {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := Equals(uint(kTwoTo25 + 1))
	desc := matcher.Description()
	expectedDesc := "33554433"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// uint8
////////////////////////////////////////////////////////////

func TestSmallUint8(t *testing.T) {
	const kExpected = 17
	matcher := Equals(uint8(kExpected))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{17, MATCH_TRUE, ""},
		equalsTestCase{17.0, MATCH_TRUE, ""},
		equalsTestCase{17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, MATCH_FALSE, ""},
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// uint16
////////////////////////////////////////////////////////////

func TestSmallUint16(t *testing.T) {
	const kExpected = 17
	matcher := Equals(uint16(kExpected))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{17, MATCH_TRUE, ""},
		equalsTestCase{17.0, MATCH_TRUE, ""},
		equalsTestCase{17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, MATCH_FALSE, ""},
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeUint16(t *testing.T) {
	const kExpected = (1 << 8) + 17
	matcher := Equals(uint16(kExpected))
	desc := matcher.Description()
	expectedDesc := "273"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{273, MATCH_TRUE, ""},
		equalsTestCase{273.0, MATCH_TRUE, ""},
		equalsTestCase{273 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int8(17), MATCH_FALSE, ""},
		equalsTestCase{int16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint8(17), MATCH_FALSE, ""},
		equalsTestCase{uint16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// uint32
////////////////////////////////////////////////////////////

func TestSmallUint32(t *testing.T) {
	const kExpected = 17
	matcher := Equals(uint32(kExpected))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{17, MATCH_TRUE, ""},
		equalsTestCase{17.0, MATCH_TRUE, ""},
		equalsTestCase{17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, MATCH_FALSE, ""},
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeUint32(t *testing.T) {
	const kExpected = (1 << 16) + 17
	matcher := Equals(uint32(kExpected))
	desc := matcher.Description()
	expectedDesc := "65553"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{65553, MATCH_TRUE, ""},
		equalsTestCase{65553.0, MATCH_TRUE, ""},
		equalsTestCase{65553 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int16(17), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(17), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestUint32NotExactlyRepresentableBySinglePrecision(t *testing.T) {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := Equals(uint32(kTwoTo25 + 1))
	desc := matcher.Description()
	expectedDesc := "33554433"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// uint64
////////////////////////////////////////////////////////////

func TestSmallUint64(t *testing.T) {
	const kExpected = 17
	matcher := Equals(uint64(kExpected))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{17, MATCH_TRUE, ""},
		equalsTestCase{17.0, MATCH_TRUE, ""},
		equalsTestCase{17 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint8(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint16(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, MATCH_FALSE, ""},
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint8(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint16(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeUint64(t *testing.T) {
	const kExpected = (1 << 32) + 17
	matcher := Equals(uint64(kExpected))
	desc := matcher.Description()
	expectedDesc := "4294967313"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{4294967313.0, MATCH_TRUE, ""},
		equalsTestCase{4294967313 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int(17), MATCH_FALSE, ""},
		equalsTestCase{int32(17), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(17), MATCH_FALSE, ""},
		equalsTestCase{uint32(17), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 1), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestUint64NotExactlyRepresentableBySinglePrecision(t *testing.T) {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := Equals(uint64(kTwoTo25 + 1))
	desc := matcher.Description()
	expectedDesc := "33554433"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestUint64NotExactlyRepresentableByDoublePrecision(t *testing.T) {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := Equals(uint64(kTwoTo54 + 1))
	desc := matcher.Description()
	expectedDesc := "18014398509481985"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo54 + 0), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo54 + 0), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// uintptr
////////////////////////////////////////////////////////////

func TestNilUintptr(t *testing.T) {
	var ptr1 uintptr
	var ptr2 uintptr

	matcher := Equals(ptr1)
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// uintptrs
		equalsTestCase{ptr1, MATCH_TRUE, ""},
		equalsTestCase{ptr2, MATCH_TRUE, ""},
		equalsTestCase{uintptr(0), MATCH_TRUE, ""},
		equalsTestCase{uintptr(17), MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a uintptr"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilUintptr(t *testing.T) {
	matcher := Equals(uintptr(17))
	desc := matcher.Description()
	expectedDesc := "17"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// uintptrs
		equalsTestCase{uintptr(17), MATCH_TRUE, ""},
		equalsTestCase{uintptr(16), MATCH_FALSE, ""},
		equalsTestCase{uintptr(0), MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a uintptr"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a uintptr"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// float32
////////////////////////////////////////////////////////////

func TestNegativeIntegralFloat32(t *testing.T) {
	matcher := Equals(float32(-32769))
	desc := matcher.Description()
	expectedDesc := "-32769"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -32769.
		equalsTestCase{-32769.0, MATCH_TRUE, ""},
		equalsTestCase{-32769 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int32(-32769), MATCH_TRUE, ""},
		equalsTestCase{int64(-32769), MATCH_TRUE, ""},
		equalsTestCase{float32(-32769), MATCH_TRUE, ""},
		equalsTestCase{float64(-32769), MATCH_TRUE, ""},
		equalsTestCase{complex64(-32769), MATCH_TRUE, ""},
		equalsTestCase{complex128(-32769), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float32(-32769)), MATCH_TRUE, ""},
		equalsTestCase{interface{}(int64(-32769)), MATCH_TRUE, ""},

		// Values that would be -32769 in two's complement.
		equalsTestCase{uint64((1 << 64) - 32769), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(-32770), MATCH_FALSE, ""},
		equalsTestCase{float32(-32769.1), MATCH_FALSE, ""},
		equalsTestCase{float32(-32768.9), MATCH_FALSE, ""},
		equalsTestCase{float64(-32769.1), MATCH_FALSE, ""},
		equalsTestCase{float64(-32768.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(-32768), MATCH_FALSE, ""},
		equalsTestCase{complex128(-32769 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralFloat32(t *testing.T) {
	matcher := Equals(float32(-32769.1))
	desc := matcher.Description()
	expectedDesc := "-32769.1"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of -32769.1.
		equalsTestCase{-32769.1, MATCH_TRUE, ""},
		equalsTestCase{-32769.1 + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(-32769.1), MATCH_TRUE, ""},
		equalsTestCase{float64(-32769.1), MATCH_TRUE, ""},
		equalsTestCase{complex64(-32769.1), MATCH_TRUE, ""},
		equalsTestCase{complex128(-32769.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int32(-32769), MATCH_FALSE, ""},
		equalsTestCase{int32(-32770), MATCH_FALSE, ""},
		equalsTestCase{int64(-32769), MATCH_FALSE, ""},
		equalsTestCase{int64(-32770), MATCH_FALSE, ""},
		equalsTestCase{float32(-32769.2), MATCH_FALSE, ""},
		equalsTestCase{float32(-32769.0), MATCH_FALSE, ""},
		equalsTestCase{float64(-32769.2), MATCH_FALSE, ""},
		equalsTestCase{complex128(-32769.1 + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeFloat32(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(float32(kExpected))
	desc := matcher.Description()
	expectedDesc := "-3.689349e+19"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float32(kExpected)
	castedInt := int64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroFloat32(t *testing.T) {
	matcher := Equals(float32(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of zero.
		equalsTestCase{0.0, MATCH_TRUE, ""},
		equalsTestCase{0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(0), MATCH_TRUE, ""},
		equalsTestCase{int8(0), MATCH_TRUE, ""},
		equalsTestCase{int16(0), MATCH_TRUE, ""},
		equalsTestCase{int32(0), MATCH_TRUE, ""},
		equalsTestCase{int64(0), MATCH_TRUE, ""},
		equalsTestCase{uint(0), MATCH_TRUE, ""},
		equalsTestCase{uint8(0), MATCH_TRUE, ""},
		equalsTestCase{uint16(0), MATCH_TRUE, ""},
		equalsTestCase{uint32(0), MATCH_TRUE, ""},
		equalsTestCase{uint64(0), MATCH_TRUE, ""},
		equalsTestCase{float32(0), MATCH_TRUE, ""},
		equalsTestCase{float64(0), MATCH_TRUE, ""},
		equalsTestCase{complex64(0), MATCH_TRUE, ""},
		equalsTestCase{complex128(0), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), MATCH_FALSE, ""},
		equalsTestCase{int64(-1), MATCH_FALSE, ""},
		equalsTestCase{float32(1), MATCH_FALSE, ""},
		equalsTestCase{float32(-1), MATCH_FALSE, ""},
		equalsTestCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloat32(t *testing.T) {
	matcher := Equals(float32(32769))
	desc := matcher.Description()
	expectedDesc := "32769"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 32769.
		equalsTestCase{32769.0, MATCH_TRUE, ""},
		equalsTestCase{32769 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(32769), MATCH_TRUE, ""},
		equalsTestCase{int32(32769), MATCH_TRUE, ""},
		equalsTestCase{int64(32769), MATCH_TRUE, ""},
		equalsTestCase{uint(32769), MATCH_TRUE, ""},
		equalsTestCase{uint32(32769), MATCH_TRUE, ""},
		equalsTestCase{uint64(32769), MATCH_TRUE, ""},
		equalsTestCase{float32(32769), MATCH_TRUE, ""},
		equalsTestCase{float64(32769), MATCH_TRUE, ""},
		equalsTestCase{complex64(32769), MATCH_TRUE, ""},
		equalsTestCase{complex128(32769), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float32(32769)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(32770), MATCH_FALSE, ""},
		equalsTestCase{uint64(32770), MATCH_FALSE, ""},
		equalsTestCase{float32(32769.1), MATCH_FALSE, ""},
		equalsTestCase{float32(32768.9), MATCH_FALSE, ""},
		equalsTestCase{float64(32769.1), MATCH_FALSE, ""},
		equalsTestCase{float64(32768.9), MATCH_FALSE, ""},
		equalsTestCase{complex128(32768), MATCH_FALSE, ""},
		equalsTestCase{complex128(32769 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralFloat32(t *testing.T) {
	matcher := Equals(float32(32769.1))
	desc := matcher.Description()
	expectedDesc := "32769.1"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 32769.1.
		equalsTestCase{32769.1, MATCH_TRUE, ""},
		equalsTestCase{32769.1 + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(32769.1), MATCH_TRUE, ""},
		equalsTestCase{float64(32769.1), MATCH_TRUE, ""},
		equalsTestCase{complex64(32769.1), MATCH_TRUE, ""},
		equalsTestCase{complex128(32769.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int32(32769), MATCH_FALSE, ""},
		equalsTestCase{int32(32770), MATCH_FALSE, ""},
		equalsTestCase{uint64(32769), MATCH_FALSE, ""},
		equalsTestCase{uint64(32770), MATCH_FALSE, ""},
		equalsTestCase{float32(32769.2), MATCH_FALSE, ""},
		equalsTestCase{float32(32769.0), MATCH_FALSE, ""},
		equalsTestCase{float64(32769.2), MATCH_FALSE, ""},
		equalsTestCase{complex128(32769.1 + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveFloat32(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(float32(kExpected))
	desc := matcher.Description()
	expectedDesc := "3.689349e+19"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float32(kExpected)
	castedInt := uint64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{uint64(0), MATCH_FALSE, ""},
		equalsTestCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestFloat32AboveExactIntegerRange(t *testing.T) {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := Equals(float32(kTwoTo25 + 1))
	desc := matcher.Description()
	expectedDesc := "3.3554432e+07"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// float64
////////////////////////////////////////////////////////////

func TestNegativeIntegralFloat64(t *testing.T) {
	const kExpected = -(1 << 50)
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "-1.125899906842624e+15"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{-1125899906842624.0, MATCH_TRUE, ""},
		equalsTestCase{-1125899906842624.0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Values that would be kExpected in two's complement.
		equalsTestCase{uint64((1 << 64) + kExpected), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralFloat64(t *testing.T) {
	const kTwoTo50 = 1 << 50
	const kExpected = -kTwoTo50 - 0.25

	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "-1.1258999068426242e+15"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(-kTwoTo50), MATCH_FALSE, ""},
		equalsTestCase{int64(-kTwoTo50 - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeFloat64(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "-3.6893488147419103e+19"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := int64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroFloat64(t *testing.T) {
	matcher := Equals(float64(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of zero.
		equalsTestCase{0.0, MATCH_TRUE, ""},
		equalsTestCase{0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(0), MATCH_TRUE, ""},
		equalsTestCase{int8(0), MATCH_TRUE, ""},
		equalsTestCase{int16(0), MATCH_TRUE, ""},
		equalsTestCase{int32(0), MATCH_TRUE, ""},
		equalsTestCase{int64(0), MATCH_TRUE, ""},
		equalsTestCase{uint(0), MATCH_TRUE, ""},
		equalsTestCase{uint8(0), MATCH_TRUE, ""},
		equalsTestCase{uint16(0), MATCH_TRUE, ""},
		equalsTestCase{uint32(0), MATCH_TRUE, ""},
		equalsTestCase{uint64(0), MATCH_TRUE, ""},
		equalsTestCase{float32(0), MATCH_TRUE, ""},
		equalsTestCase{float64(0), MATCH_TRUE, ""},
		equalsTestCase{complex64(0), MATCH_TRUE, ""},
		equalsTestCase{complex128(0), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), MATCH_FALSE, ""},
		equalsTestCase{int64(-1), MATCH_FALSE, ""},
		equalsTestCase{float32(1), MATCH_FALSE, ""},
		equalsTestCase{float32(-1), MATCH_FALSE, ""},
		equalsTestCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloat64(t *testing.T) {
	const kExpected = 1 << 50
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "1.125899906842624e+15"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 32769.
		equalsTestCase{1125899906842624.0, MATCH_TRUE, ""},
		equalsTestCase{1125899906842624.0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralFloat64(t *testing.T) {
	const kTwoTo50 = 1 << 50
	const kExpected = kTwoTo50 + 0.25
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "1.1258999068426242e+15"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kTwoTo50), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo50 - 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveFloat64(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "3.6893488147419103e+19"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := uint64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{uint64(0), MATCH_FALSE, ""},
		equalsTestCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestFloat64AboveExactIntegerRange(t *testing.T) {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := Equals(float64(kTwoTo54 + 1))
	desc := matcher.Description()
	expectedDesc := "1.8014398509481984e+16"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// complex64
////////////////////////////////////////////////////////////

func TestNegativeIntegralComplex64(t *testing.T) {
	const kExpected = -32769
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(-32769+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{-32769.0, MATCH_TRUE, ""},
		equalsTestCase{-32769.0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Values that would be kExpected in two's complement.
		equalsTestCase{uint32((1 << 32) + kExpected), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) + kExpected), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralComplex64(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = -kTwoTo20 - 0.25

	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(-1.0485762e+06+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(-kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int(-kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{int32(-kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int32(-kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{int64(-kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int64(-kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 0.75), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 0.75), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeComplex64(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(-3.689349e+19+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := int64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroComplex64(t *testing.T) {
	matcher := Equals(complex64(0))
	desc := matcher.Description()
	expectedDesc := "(0+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of zero.
		equalsTestCase{0.0, MATCH_TRUE, ""},
		equalsTestCase{0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(0), MATCH_TRUE, ""},
		equalsTestCase{int8(0), MATCH_TRUE, ""},
		equalsTestCase{int16(0), MATCH_TRUE, ""},
		equalsTestCase{int32(0), MATCH_TRUE, ""},
		equalsTestCase{int64(0), MATCH_TRUE, ""},
		equalsTestCase{uint(0), MATCH_TRUE, ""},
		equalsTestCase{uint8(0), MATCH_TRUE, ""},
		equalsTestCase{uint16(0), MATCH_TRUE, ""},
		equalsTestCase{uint32(0), MATCH_TRUE, ""},
		equalsTestCase{uint64(0), MATCH_TRUE, ""},
		equalsTestCase{float32(0), MATCH_TRUE, ""},
		equalsTestCase{float64(0), MATCH_TRUE, ""},
		equalsTestCase{complex64(0), MATCH_TRUE, ""},
		equalsTestCase{complex128(0), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), MATCH_FALSE, ""},
		equalsTestCase{int64(-1), MATCH_FALSE, ""},
		equalsTestCase{float32(1), MATCH_FALSE, ""},
		equalsTestCase{float32(-1), MATCH_FALSE, ""},
		equalsTestCase{float64(1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1), MATCH_FALSE, ""},
		equalsTestCase{complex64(0 + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralComplex64(t *testing.T) {
	const kExpected = 1 << 20
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(1.048576e+06+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 32769.
		equalsTestCase{1048576.0, MATCH_TRUE, ""},
		equalsTestCase{1048576.0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralComplex64(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = kTwoTo20 + 0.25
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(1.0485762e+06+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 1i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveComplex64(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(3.689349e+19+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := uint64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{uint64(0), MATCH_FALSE, ""},
		equalsTestCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestComplex64AboveExactIntegerRange(t *testing.T) {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := Equals(complex64(kTwoTo25 + 1))
	desc := matcher.Description()
	expectedDesc := "(3.3554432e+07+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo25 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo25 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestComplex64WithNonZeroImaginaryPart(t *testing.T) {
	const kRealPart = 17
	const kImagPart = 0.25i
	const kExpected = kRealPart + kImagPart
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "(17+0.25i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kRealPart + kImagPart, MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int8(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int16(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int32(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint8(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint16(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint32(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{float32(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{float64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{complex64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// complex128
////////////////////////////////////////////////////////////

func TestNegativeIntegralComplex128(t *testing.T) {
	const kExpected = -32769
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(-32769+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{-32769.0, MATCH_TRUE, ""},
		equalsTestCase{-32769.0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Values that would be kExpected in two's complement.
		equalsTestCase{uint32((1 << 32) + kExpected), MATCH_FALSE, ""},
		equalsTestCase{uint64((1 << 64) + kExpected), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralComplex128(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = -kTwoTo20 - 0.25

	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(-1.04857625e+06+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(-kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int(-kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{int32(-kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int32(-kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{int64(-kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int64(-kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 0.75), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 0.75), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeComplex128(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(-3.6893488147419103e+19+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := int64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroComplex128(t *testing.T) {
	matcher := Equals(complex128(0))
	desc := matcher.Description()
	expectedDesc := "(0+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of zero.
		equalsTestCase{0.0, MATCH_TRUE, ""},
		equalsTestCase{0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(0), MATCH_TRUE, ""},
		equalsTestCase{int8(0), MATCH_TRUE, ""},
		equalsTestCase{int16(0), MATCH_TRUE, ""},
		equalsTestCase{int32(0), MATCH_TRUE, ""},
		equalsTestCase{int64(0), MATCH_TRUE, ""},
		equalsTestCase{uint(0), MATCH_TRUE, ""},
		equalsTestCase{uint8(0), MATCH_TRUE, ""},
		equalsTestCase{uint16(0), MATCH_TRUE, ""},
		equalsTestCase{uint32(0), MATCH_TRUE, ""},
		equalsTestCase{uint64(0), MATCH_TRUE, ""},
		equalsTestCase{float32(0), MATCH_TRUE, ""},
		equalsTestCase{float64(0), MATCH_TRUE, ""},
		equalsTestCase{complex64(0), MATCH_TRUE, ""},
		equalsTestCase{complex128(0), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), MATCH_FALSE, ""},
		equalsTestCase{int64(-1), MATCH_FALSE, ""},
		equalsTestCase{float32(1), MATCH_FALSE, ""},
		equalsTestCase{float32(-1), MATCH_FALSE, ""},
		equalsTestCase{float64(1), MATCH_FALSE, ""},
		equalsTestCase{float64(-1), MATCH_FALSE, ""},
		equalsTestCase{complex64(0 + 2i), MATCH_FALSE, ""},
		equalsTestCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralComplex128(t *testing.T) {
	const kExpected = 1 << 20
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(1.048576e+06+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of 32769.
		equalsTestCase{1048576.0, MATCH_TRUE, ""},
		equalsTestCase{1048576.0 + 0i, MATCH_TRUE, ""},
		equalsTestCase{int(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{int64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{uint64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},
		equalsTestCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{int64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralComplex128(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = kTwoTo20 + 0.25
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(1.04857625e+06+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo20), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo20 - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected + 1), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex64(kExpected - 1i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected - 1i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveComplex128(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(3.6893488147419103e+19+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := uint64(floatExpected)

	cases := []equalsTestCase{
		// Equal values of numeric type.
		equalsTestCase{kExpected + 0i, MATCH_TRUE, ""},
		equalsTestCase{float32(kExpected), MATCH_TRUE, ""},
		equalsTestCase{float64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, MATCH_FALSE, ""},
		equalsTestCase{int64(0), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MinInt64), MATCH_FALSE, ""},
		equalsTestCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		equalsTestCase{uint64(0), MATCH_FALSE, ""},
		equalsTestCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		equalsTestCase{float32(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kExpected / 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestComplex128AboveExactIntegerRange(t *testing.T) {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := Equals(complex128(kTwoTo54 + 1))
	desc := matcher.Description()
	expectedDesc := "(1.8014398509481984e+16+0i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Integers.
		equalsTestCase{int64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		equalsTestCase{uint64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestComplex128WithNonZeroImaginaryPart(t *testing.T) {
	const kRealPart = 17
	const kImagPart = 0.25i
	const kExpected = kRealPart + kImagPart
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "(17+0.25i)"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Various types of the expected value.
		equalsTestCase{kExpected, MATCH_TRUE, ""},
		equalsTestCase{kRealPart + kImagPart, MATCH_TRUE, ""},
		equalsTestCase{complex64(kExpected), MATCH_TRUE, ""},
		equalsTestCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int8(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int16(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int32(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{int64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint8(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint16(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint32(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{uint64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{float32(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{float64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{complex64(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
		equalsTestCase{complex128(kRealPart), MATCH_FALSE, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// array
////////////////////////////////////////////////////////////

func TestArray(t *testing.T) {
	panicked := false

	defer func() {
		if !panicked {
			t.Errorf("Expected panic; got none.")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	var someArray [3]int
	Equals(someArray)
}

////////////////////////////////////////////////////////////
// chan
////////////////////////////////////////////////////////////

func TestNilChan(t *testing.T) {
	var nilChan1 chan int
	var nilChan2 chan int
	var nilChan3 chan uint
	var nonNilChan1 chan int = make(chan int)
	var nonNilChan2 chan uint = make(chan uint)

	matcher := Equals(nilChan1)
	desc := matcher.Description()
	expectedDesc := "0x0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// int channels
		equalsTestCase{nilChan1, MATCH_TRUE, ""},
		equalsTestCase{nilChan2, MATCH_TRUE, ""},
		equalsTestCase{nonNilChan1, MATCH_FALSE, ""},

		// uint channels
		equalsTestCase{nilChan3, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{nonNilChan2, MATCH_UNDEFINED, "which is not a chan int"},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a chan int"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilChan(t *testing.T) {
	var nilChan1 chan int
	var nilChan2 chan uint
	var nonNilChan1 chan int = make(chan int)
	var nonNilChan2 chan int = make(chan int)
	var nonNilChan3 chan uint = make(chan uint)

	matcher := Equals(nonNilChan1)
	desc := matcher.Description()
	expectedDesc := fmt.Sprintf("%v", nonNilChan1)

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// int channels
		equalsTestCase{nonNilChan1, MATCH_TRUE, ""},
		equalsTestCase{nonNilChan2, MATCH_FALSE, ""},
		equalsTestCase{nilChan1, MATCH_FALSE, ""},

		// uint channels
		equalsTestCase{nilChan2, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{nonNilChan3, MATCH_UNDEFINED, "which is not a chan int"},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a chan int"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a chan int"},
	}

	checkTestCases(t, matcher, cases)
}

func TestChanDirection(t *testing.T) {
	var chan1 chan<- int
	var chan2 <-chan int
	var chan3 chan int

	matcher := Equals(chan1)
	desc := matcher.Description()
	expectedDesc := fmt.Sprintf("%v", chan1)

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		equalsTestCase{chan1, MATCH_TRUE, ""},
		equalsTestCase{chan2, MATCH_UNDEFINED, "which is not a chan<- int"},
		equalsTestCase{chan3, MATCH_UNDEFINED, "which is not a chan<- int"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// func
////////////////////////////////////////////////////////////

func TestFunctions(t *testing.T) {
	func1 := func() {}
	func2 := func() {}
	func3 := func(x int) {}

	matcher := Equals(func1)
	desc := matcher.Description()
	expectedDesc := fmt.Sprintf("%v", func1)

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Functions.
		equalsTestCase{func1, MATCH_TRUE, ""},
		equalsTestCase{func2, MATCH_FALSE, ""},
		equalsTestCase{func3, MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a function"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a function"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// map
////////////////////////////////////////////////////////////

func TestNilMap(t *testing.T) {
	var nilMap1 map[int]int
	var nilMap2 map[int]int
	var nilMap3 map[int]uint
	var nonNilMap1 map[int]int = make(map[int]int)
	var nonNilMap2 map[int]uint = make(map[int]uint)

	matcher := Equals(nilMap1)
	desc := matcher.Description()
	expectedDesc := "map[]"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nilMap1, MATCH_TRUE, ""},
		equalsTestCase{nilMap2, MATCH_TRUE, ""},
		equalsTestCase{nilMap3, MATCH_TRUE, ""},
		equalsTestCase{nonNilMap1, MATCH_FALSE, ""},
		equalsTestCase{nonNilMap2, MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a map"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilMap(t *testing.T) {
	var nilMap1 map[int]int
	var nilMap2 map[int]uint
	var nonNilMap1 map[int]int = make(map[int]int)
	var nonNilMap2 map[int]int = make(map[int]int)
	var nonNilMap3 map[int]uint = make(map[int]uint)

	matcher := Equals(nonNilMap1)
	desc := matcher.Description()
	expectedDesc := "map[]"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nonNilMap1, MATCH_TRUE, ""},
		equalsTestCase{nonNilMap2, MATCH_FALSE, ""},
		equalsTestCase{nonNilMap3, MATCH_FALSE, ""},
		equalsTestCase{nilMap1, MATCH_FALSE, ""},
		equalsTestCase{nilMap2, MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a map"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a map"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// Pointers
////////////////////////////////////////////////////////////

func TestNilPointer(t *testing.T) {
	var someInt int = 17
	var someUint uint = 17

	var nilInt1 *int
	var nilInt2 *int
	var nilUint *uint
	var nonNilInt *int = &someInt
	var nonNilUint *uint = &someUint

	matcher := Equals(nilInt1)
	desc := matcher.Description()
	expectedDesc := "<nil>"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nilInt1, MATCH_TRUE, ""},
		equalsTestCase{nilInt2, MATCH_TRUE, ""},
		equalsTestCase{nonNilInt, MATCH_FALSE, ""},

		// Incorrect type.
		equalsTestCase{nilUint, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{nonNilUint, MATCH_UNDEFINED, "which is not a *int"},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a *int"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilPointer(t *testing.T) {
	var someInt int = 17
	var someOtherInt int = 17
	var someUint uint = 17

	var nilInt *int
	var nilUint *uint
	var nonNilInt1 *int = &someInt
	var nonNilInt2 *int = &someOtherInt
	var nonNilUint *uint = &someUint

	matcher := Equals(nonNilInt1)
	desc := matcher.Description()
	expectedDesc := fmt.Sprintf("%v", nonNilInt1)

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nonNilInt1, MATCH_TRUE, ""},
		equalsTestCase{nonNilInt2, MATCH_FALSE, ""},
		equalsTestCase{nilInt, MATCH_FALSE, ""},

		// Incorrect type.
		equalsTestCase{nilUint, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{nonNilUint, MATCH_UNDEFINED, "which is not a *int"},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a *int"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a *int"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// Slices
////////////////////////////////////////////////////////////

func TestNilSlice(t *testing.T) {
	var nilInt1 []int
	var nilInt2 []int
	var nilUint []uint

	var nonNilInt []int = make([]int, 0)
	var nonNilUint []uint = make([]uint, 0)

	matcher := Equals(nilInt1)
	desc := matcher.Description()
	expectedDesc := "[]"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nilInt1, MATCH_TRUE, ""},
		equalsTestCase{nilInt2, MATCH_TRUE, ""},
		equalsTestCase{nonNilInt, MATCH_FALSE, ""},

		// Incorrect type.
		equalsTestCase{nilUint, MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{nonNilUint, MATCH_UNDEFINED, "which is not a []int"},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a []int"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a []int"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilSlice(t *testing.T) {
	panicked := false

	defer func() {
		if !panicked {
			t.Errorf("Expected panic; got none.")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	nonNil := make([]int, 0)
	Equals(nonNil)
}

////////////////////////////////////////////////////////////
// string
////////////////////////////////////////////////////////////

func TestString(t *testing.T) {
	partial := "taco"
	expected := fmt.Sprintf("%s%d", partial, 1)

	matcher := Equals(expected)
	desc := matcher.Description()
	expectedDesc := "taco1"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{"taco1", MATCH_TRUE, ""},
		equalsTestCase{"taco" + "1", MATCH_TRUE, ""},
		equalsTestCase{expected, MATCH_TRUE, ""},

		equalsTestCase{"", MATCH_FALSE, ""},
		equalsTestCase{"taco", MATCH_FALSE, ""},
		equalsTestCase{"taco1\x00", MATCH_FALSE, ""},
		equalsTestCase{"taco2", MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a string"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a string"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// struct
////////////////////////////////////////////////////////////

func TestStruct(t *testing.T) {
	type someStruct struct{ foo uint }
	panicked := false

	defer func() {
		if !panicked {
			t.Errorf("Expected panic; got none.")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	Equals(someStruct{17})
}

////////////////////////////////////////////////////////////
// unsafe.Pointer
////////////////////////////////////////////////////////////

func TestNilUnsafePointer(t *testing.T) {
	someInt := int(17)

	var nilPtr1 unsafe.Pointer
	var nilPtr2 unsafe.Pointer
	var nonNilPtr unsafe.Pointer = unsafe.Pointer(&someInt)

	matcher := Equals(nilPtr1)
	desc := matcher.Description()
	expectedDesc := "0x0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nilPtr1, MATCH_TRUE, ""},
		equalsTestCase{nilPtr2, MATCH_TRUE, ""},
		equalsTestCase{nonNilPtr, MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilUnsafePointer(t *testing.T) {
	someInt := int(17)
	someOtherInt := int(17)

	var nilPtr unsafe.Pointer
	var nonNilPtr1 unsafe.Pointer = unsafe.Pointer(&someInt)
	var nonNilPtr2 unsafe.Pointer = unsafe.Pointer(&someOtherInt)

	matcher := Equals(nonNilPtr1)
	desc := matcher.Description()
	expectedDesc := fmt.Sprintf("%v", nonNilPtr1)

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Correct type.
		equalsTestCase{nonNilPtr1, MATCH_TRUE, ""},
		equalsTestCase{nonNilPtr2, MATCH_FALSE, ""},
		equalsTestCase{nilPtr, MATCH_FALSE, ""},

		// Other types.
		equalsTestCase{0, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{bool(false), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int8(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int16(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int32(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{int64(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint8(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint16(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint32(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uint64(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{uintptr(0), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{true, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{[...]int{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{make(chan int), MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{func() {}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{&someInt, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{[]int{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{"taco", MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
		equalsTestCase{equalsTestCase{}, MATCH_UNDEFINED, "which is not a unsafe.Pointer"},
	}

	checkTestCases(t, matcher, cases)
}
