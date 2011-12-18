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
	"math"
	"testing"
)

var someInt int = -17

// TODO(jacobsa): interface
// TODO(jacobsa): map
// TODO(jacobsa): ptr
// TODO(jacobsa): slice
// TODO(jacobsa): string
// TODO(jacobsa): struct
// TODO(jacobsa): unsafe pointer

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type testCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func checkTestCases(t *testing.T, matcher Matcher, cases []testCase) {
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

		if err != c.expectedError {
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

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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
	expectedDesc := "-1073741824.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(float64(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloatingPointLiteral(t *testing.T) {
	// 2^30
	matcher := Equals(1073741824.0)
	desc := matcher.Description()
	expectedDesc := "1073741824.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(float64(1073741824)), MATCH_TRUE, ""},

		// Values that would be 1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{uint(1073741823), MATCH_FALSE, ""},
		testCase{uint32(1073741823), MATCH_FALSE, ""},
		testCase{uint64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonIntegralFloatingPointLiteral(t *testing.T) {
	matcher := Equals(17.5)
	desc := matcher.Description()
	expectedDesc := "17.5"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 17.1.
		testCase{17.1, MATCH_TRUE, ""},
		testCase{17.1, MATCH_TRUE, ""},
		testCase{17.1 + 0i, MATCH_TRUE, ""},
		testCase{float32(17.1), MATCH_TRUE, ""},
		testCase{float64(17.1), MATCH_TRUE, ""},
		testCase{complex64(17.1), MATCH_TRUE, ""},
		testCase{complex128(17.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{17, MATCH_FALSE, ""},
		testCase{17.2, MATCH_FALSE, ""},
		testCase{18, MATCH_FALSE, ""},
		testCase{int(17), MATCH_FALSE, ""},
		testCase{int(18), MATCH_FALSE, ""},
		testCase{int32(17), MATCH_FALSE, ""},
		testCase{int64(17), MATCH_FALSE, ""},
		testCase{uint(17), MATCH_FALSE, ""},
		testCase{uint32(17), MATCH_FALSE, ""},
		testCase{uint64(17), MATCH_FALSE, ""},
		testCase{complex128(17.1 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// bools
		testCase{false, MATCH_TRUE, ""},
		testCase{bool(false), MATCH_TRUE, ""},

		testCase{true, MATCH_FALSE, ""},
		testCase{bool(true), MATCH_FALSE, ""},

		// Other types.
		testCase{int(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int8(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int16(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int32(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int64(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint8(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint16(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint32(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint64(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not a bool"},
		testCase{true, MATCH_UNDEFINED, "which is not a bool"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not a bool"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a bool"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a bool"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a bool"},
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

	cases := []testCase{
		// bools
		testCase{true, MATCH_TRUE, ""},
		testCase{bool(true), MATCH_TRUE, ""},

		testCase{false, MATCH_FALSE, ""},
		testCase{bool(false), MATCH_FALSE, ""},

		// Other types.
		testCase{int(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int8(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int16(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int32(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{int64(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint8(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint16(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint32(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uint64(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{uintptr(1), MATCH_UNDEFINED, "which is not a bool"},
		testCase{true, MATCH_UNDEFINED, "which is not a bool"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not a bool"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a bool"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a bool"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a bool"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a bool"},
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

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of -17.
		testCase{-17, MATCH_TRUE, ""},
		testCase{-17.0, MATCH_TRUE, ""},
		testCase{-17 + 0i, MATCH_TRUE, ""},
		testCase{int(-17), MATCH_TRUE, ""},
		testCase{int8(-17), MATCH_TRUE, ""},
		testCase{int16(-17), MATCH_TRUE, ""},
		testCase{int32(-17), MATCH_TRUE, ""},
		testCase{int64(-17), MATCH_TRUE, ""},
		testCase{float32(-17), MATCH_TRUE, ""},
		testCase{float64(-17), MATCH_TRUE, ""},
		testCase{complex64(-17), MATCH_TRUE, ""},
		testCase{complex128(-17), MATCH_TRUE, ""},
		testCase{interface{}(int(-17)), MATCH_TRUE, ""},

		// Values that would be -17 in two's complement.
		testCase{uint((1 << 32) - 17), MATCH_FALSE, ""},
		testCase{uint8((1 << 8) - 17), MATCH_FALSE, ""},
		testCase{uint16((1 << 16) - 17), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 17), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 17), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-16), MATCH_FALSE, ""},
		testCase{int8(-16), MATCH_FALSE, ""},
		testCase{int16(-16), MATCH_FALSE, ""},
		testCase{int32(-16), MATCH_FALSE, ""},
		testCase{int64(-16), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float32(-17.1), MATCH_FALSE, ""},
		testCase{float32(-16.9), MATCH_FALSE, ""},
		testCase{complex64(-16), MATCH_FALSE, ""},
		testCase{complex64(-17 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr((1 << 32) - 17), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{-17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{-17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"-17", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 0.
		testCase{0, MATCH_TRUE, ""},
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(int(0)), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1), MATCH_FALSE, ""},
		testCase{int8(1), MATCH_FALSE, ""},
		testCase{int16(1), MATCH_FALSE, ""},
		testCase{int32(1), MATCH_FALSE, ""},
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{float32(-0.1), MATCH_FALSE, ""},
		testCase{float32(0.1), MATCH_FALSE, ""},
		testCase{complex64(1), MATCH_FALSE, ""},
		testCase{complex64(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"0", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 17.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(17), MATCH_TRUE, ""},
		testCase{int8(17), MATCH_TRUE, ""},
		testCase{int16(17), MATCH_TRUE, ""},
		testCase{int32(17), MATCH_TRUE, ""},
		testCase{int64(17), MATCH_TRUE, ""},
		testCase{float32(17), MATCH_TRUE, ""},
		testCase{float64(17), MATCH_TRUE, ""},
		testCase{complex64(17), MATCH_TRUE, ""},
		testCase{complex128(17), MATCH_TRUE, ""},
		testCase{interface{}(int(17)), MATCH_TRUE, ""},
		testCase{uint(17), MATCH_TRUE, ""},
		testCase{uint8(17), MATCH_TRUE, ""},
		testCase{uint16(17), MATCH_TRUE, ""},
		testCase{uint32(17), MATCH_TRUE, ""},
		testCase{uint64(17), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(16), MATCH_FALSE, ""},
		testCase{int8(16), MATCH_FALSE, ""},
		testCase{int16(16), MATCH_FALSE, ""},
		testCase{int32(16), MATCH_FALSE, ""},
		testCase{int64(16), MATCH_FALSE, ""},
		testCase{float32(16.9), MATCH_FALSE, ""},
		testCase{float32(17.1), MATCH_FALSE, ""},
		testCase{complex64(16), MATCH_FALSE, ""},
		testCase{complex64(17 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(17), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{17}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"17", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of -32766.
		testCase{-32766, MATCH_TRUE, ""},
		testCase{-32766.0, MATCH_TRUE, ""},
		testCase{-32766 + 0i, MATCH_TRUE, ""},
		testCase{int(-32766), MATCH_TRUE, ""},
		testCase{int16(-32766), MATCH_TRUE, ""},
		testCase{int32(-32766), MATCH_TRUE, ""},
		testCase{int64(-32766), MATCH_TRUE, ""},
		testCase{float32(-32766), MATCH_TRUE, ""},
		testCase{float64(-32766), MATCH_TRUE, ""},
		testCase{complex64(-32766), MATCH_TRUE, ""},
		testCase{complex128(-32766), MATCH_TRUE, ""},
		testCase{interface{}(int(-32766)), MATCH_TRUE, ""},

		// Values that would be -32766 in two's complement.
		testCase{uint((1 << 32) - 32766), MATCH_FALSE, ""},
		testCase{uint16((1 << 16) - 32766), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 32766), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 32766), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-16), MATCH_FALSE, ""},
		testCase{int8(-16), MATCH_FALSE, ""},
		testCase{int16(-16), MATCH_FALSE, ""},
		testCase{int32(-16), MATCH_FALSE, ""},
		testCase{int64(-16), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float32(-32766.1), MATCH_FALSE, ""},
		testCase{float32(-32765.9), MATCH_FALSE, ""},
		testCase{complex64(-32766.1), MATCH_FALSE, ""},
		testCase{complex64(-32766 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr((1 << 32) - 32766), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{-32766}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{-32766}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"-32766", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 0.
		testCase{0, MATCH_TRUE, ""},
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(int(0)), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1), MATCH_FALSE, ""},
		testCase{int8(1), MATCH_FALSE, ""},
		testCase{int16(1), MATCH_FALSE, ""},
		testCase{int32(1), MATCH_FALSE, ""},
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{float32(-0.1), MATCH_FALSE, ""},
		testCase{float32(0.1), MATCH_FALSE, ""},
		testCase{complex64(1), MATCH_FALSE, ""},
		testCase{complex64(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{0}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"0", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 32765.
		testCase{32765, MATCH_TRUE, ""},
		testCase{32765.0, MATCH_TRUE, ""},
		testCase{32765 + 0i, MATCH_TRUE, ""},
		testCase{int(32765), MATCH_TRUE, ""},
		testCase{int16(32765), MATCH_TRUE, ""},
		testCase{int32(32765), MATCH_TRUE, ""},
		testCase{int64(32765), MATCH_TRUE, ""},
		testCase{float32(32765), MATCH_TRUE, ""},
		testCase{float64(32765), MATCH_TRUE, ""},
		testCase{complex64(32765), MATCH_TRUE, ""},
		testCase{complex128(32765), MATCH_TRUE, ""},
		testCase{interface{}(int(32765)), MATCH_TRUE, ""},
		testCase{uint(32765), MATCH_TRUE, ""},
		testCase{uint16(32765), MATCH_TRUE, ""},
		testCase{uint32(32765), MATCH_TRUE, ""},
		testCase{uint64(32765), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(32764), MATCH_FALSE, ""},
		testCase{int16(32764), MATCH_FALSE, ""},
		testCase{int32(32764), MATCH_FALSE, ""},
		testCase{int64(32764), MATCH_FALSE, ""},
		testCase{float32(32764.9), MATCH_FALSE, ""},
		testCase{float32(32765.1), MATCH_FALSE, ""},
		testCase{complex64(32765.9), MATCH_FALSE, ""},
		testCase{complex64(32765 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(32765), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{32765}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{32765}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"32765", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of -1073741824.
		testCase{-1073741824, MATCH_TRUE, ""},
		testCase{-1073741824.0, MATCH_TRUE, ""},
		testCase{-1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(-1073741824), MATCH_TRUE, ""},
		testCase{int32(-1073741824), MATCH_TRUE, ""},
		testCase{int64(-1073741824), MATCH_TRUE, ""},
		testCase{float32(-1073741824), MATCH_TRUE, ""},
		testCase{float64(-1073741824), MATCH_TRUE, ""},
		testCase{complex64(-1073741824), MATCH_TRUE, ""},
		testCase{complex128(-1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(-1073741824)), MATCH_TRUE, ""},

		// Values that would be -1073741824 in two's complement.
		testCase{uint((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint32((1 << 32) - 1073741824), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) - 1073741824), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int(-1073741823), MATCH_FALSE, ""},
		testCase{int32(-1073741823), MATCH_FALSE, ""},
		testCase{int64(-1073741823), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1073741824.1), MATCH_FALSE, ""},
		testCase{float64(-1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(-1073741823), MATCH_FALSE, ""},
		testCase{complex128(-1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 1073741824.
		testCase{1073741824, MATCH_TRUE, ""},
		testCase{1073741824.0, MATCH_TRUE, ""},
		testCase{1073741824 + 0i, MATCH_TRUE, ""},
		testCase{int(1073741824), MATCH_TRUE, ""},
		testCase{uint(1073741824), MATCH_TRUE, ""},
		testCase{int32(1073741824), MATCH_TRUE, ""},
		testCase{int64(1073741824), MATCH_TRUE, ""},
		testCase{uint32(1073741824), MATCH_TRUE, ""},
		testCase{uint64(1073741824), MATCH_TRUE, ""},
		testCase{float32(1073741824), MATCH_TRUE, ""},
		testCase{float64(1073741824), MATCH_TRUE, ""},
		testCase{complex64(1073741824), MATCH_TRUE, ""},
		testCase{complex128(1073741824), MATCH_TRUE, ""},
		testCase{interface{}(int(1073741824)), MATCH_TRUE, ""},
		testCase{interface{}(uint(1073741824)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(1073741823), MATCH_FALSE, ""},
		testCase{int32(1073741823), MATCH_FALSE, ""},
		testCase{int64(1073741823), MATCH_FALSE, ""},
		testCase{float64(1073741824.1), MATCH_FALSE, ""},
		testCase{float64(1073741823.9), MATCH_FALSE, ""},
		testCase{complex128(1073741823), MATCH_FALSE, ""},
		testCase{complex128(1073741824 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of -1099511627776.
		testCase{-1099511627776.0, MATCH_TRUE, ""},
		testCase{-1099511627776 + 0i, MATCH_TRUE, ""},
		testCase{int64(-1099511627776), MATCH_TRUE, ""},
		testCase{float32(-1099511627776), MATCH_TRUE, ""},
		testCase{float64(-1099511627776), MATCH_TRUE, ""},
		testCase{complex64(-1099511627776), MATCH_TRUE, ""},
		testCase{complex128(-1099511627776), MATCH_TRUE, ""},
		testCase{interface{}(int64(-1099511627776)), MATCH_TRUE, ""},

		// Values that would be -1099511627776 in two's complement.
		testCase{uint64((1 << 64) - 1099511627776), MATCH_FALSE, ""},

		// Non-equal values of signed integer type.
		testCase{int64(-1099511627775), MATCH_FALSE, ""},

		// Non-equal values of other numeric types.
		testCase{float64(-1099511627776.1), MATCH_FALSE, ""},
		testCase{float64(-1099511627775.9), MATCH_FALSE, ""},
		testCase{complex128(-1099511627775), MATCH_FALSE, ""},
		testCase{complex128(-1099511627776 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 1099511627776.
		testCase{1099511627776.0, MATCH_TRUE, ""},
		testCase{1099511627776 + 0i, MATCH_TRUE, ""},
		testCase{int64(1099511627776), MATCH_TRUE, ""},
		testCase{uint64(1099511627776), MATCH_TRUE, ""},
		testCase{float32(1099511627776), MATCH_TRUE, ""},
		testCase{float64(1099511627776), MATCH_TRUE, ""},
		testCase{complex64(1099511627776), MATCH_TRUE, ""},
		testCase{complex128(1099511627776), MATCH_TRUE, ""},
		testCase{interface{}(int64(1099511627776)), MATCH_TRUE, ""},
		testCase{interface{}(uint64(1099511627776)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(1099511627775), MATCH_FALSE, ""},
		testCase{uint64(1099511627775), MATCH_FALSE, ""},
		testCase{float64(1099511627776.1), MATCH_FALSE, ""},
		testCase{float64(1099511627775.9), MATCH_FALSE, ""},
		testCase{complex128(1099511627775), MATCH_FALSE, ""},
		testCase{complex128(1099511627776 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		testCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo54 + 0), MATCH_FALSE, ""},
		testCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo54 + 0), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int8(kExpected), MATCH_TRUE, ""},
		testCase{int16(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint8(kExpected), MATCH_TRUE, ""},
		testCase{uint16(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{kExpected + 1, MATCH_FALSE, ""},
		testCase{int(kExpected + 1), MATCH_TRUE, ""},
		testCase{int8(kExpected + 1), MATCH_TRUE, ""},
		testCase{int16(kExpected + 1), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint8(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float32(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{65553, MATCH_TRUE, ""},
		testCase{65553.0, MATCH_TRUE, ""},
		testCase{65553 + 0i, MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{int16(17), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(17), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},
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

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		testCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int8(kExpected), MATCH_TRUE, ""},
		testCase{int16(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint8(kExpected), MATCH_TRUE, ""},
		testCase{uint16(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{kExpected + 1, MATCH_FALSE, ""},
		testCase{int(kExpected + 1), MATCH_TRUE, ""},
		testCase{int8(kExpected + 1), MATCH_TRUE, ""},
		testCase{int16(kExpected + 1), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint8(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float32(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int8(kExpected), MATCH_TRUE, ""},
		testCase{int16(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint8(kExpected), MATCH_TRUE, ""},
		testCase{uint16(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{kExpected + 1, MATCH_FALSE, ""},
		testCase{int(kExpected + 1), MATCH_TRUE, ""},
		testCase{int8(kExpected + 1), MATCH_TRUE, ""},
		testCase{int16(kExpected + 1), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint8(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float32(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{273, MATCH_TRUE, ""},
		testCase{273.0, MATCH_TRUE, ""},
		testCase{273 + 0i, MATCH_TRUE, ""},
		testCase{int16(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint16(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{int8(17), MATCH_TRUE, ""},
		testCase{int16(kExpected + 1), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint8(17), MATCH_TRUE, ""},
		testCase{uint16(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int8(kExpected), MATCH_TRUE, ""},
		testCase{int16(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint8(kExpected), MATCH_TRUE, ""},
		testCase{uint16(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{kExpected + 1, MATCH_FALSE, ""},
		testCase{int(kExpected + 1), MATCH_TRUE, ""},
		testCase{int8(kExpected + 1), MATCH_TRUE, ""},
		testCase{int16(kExpected + 1), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint8(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float32(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{65553, MATCH_TRUE, ""},
		testCase{65553.0, MATCH_TRUE, ""},
		testCase{65553 + 0i, MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{int16(17), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(17), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},
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

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		testCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{17, MATCH_TRUE, ""},
		testCase{17.0, MATCH_TRUE, ""},
		testCase{17 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int8(kExpected), MATCH_TRUE, ""},
		testCase{int16(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint8(kExpected), MATCH_TRUE, ""},
		testCase{uint16(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{kExpected + 1, MATCH_FALSE, ""},
		testCase{int(kExpected + 1), MATCH_TRUE, ""},
		testCase{int8(kExpected + 1), MATCH_TRUE, ""},
		testCase{int16(kExpected + 1), MATCH_TRUE, ""},
		testCase{int32(kExpected + 1), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint8(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint16(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint32(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float32(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of the expected value.
		testCase{4294967313.0, MATCH_TRUE, ""},
		testCase{4294967313 + 0i, MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric types.
		testCase{int(17), MATCH_TRUE, ""},
		testCase{int32(17), MATCH_TRUE, ""},
		testCase{int64(kExpected + 1), MATCH_TRUE, ""},
		testCase{uint(17), MATCH_TRUE, ""},
		testCase{uint32(17), MATCH_TRUE, ""},
		testCase{uint64(kExpected + 1), MATCH_TRUE, ""},
		testCase{float64(kExpected + 1), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 2i), MATCH_TRUE, ""},
		testCase{complex128(kExpected + 1), MATCH_TRUE, ""},
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

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Single-precision floating point.
		testCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo25 + 0), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 2), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo54 + 0), MATCH_FALSE, ""},
		testCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo54 + 0), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
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
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// uintptrs
		testCase{ptr1, MATCH_TRUE, ""},
		testCase{ptr2, MATCH_TRUE, ""},
		testCase{uintptr(0), MATCH_TRUE, ""},
		testCase{uintptr(17), MATCH_FALSE, ""},

		// Other types.
		testCase{0, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{bool(false), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{true, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a uintptr"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNonNilUintptr(t *testing.T) {
	matcher := Equals(uintptr(17))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// uintptrs
		testCase{uintptr(17), MATCH_TRUE, ""},
		testCase{uintptr(16), MATCH_FALSE, ""},
		testCase{uintptr(0), MATCH_FALSE, ""},

		// Other types.
		testCase{0, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{bool(false), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{int64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint8(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint16(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint32(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{uint64(0), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{true, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a uintptr"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a uintptr"},
	}

	checkTestCases(t, matcher, cases)
}

////////////////////////////////////////////////////////////
// float32
////////////////////////////////////////////////////////////

func TestNegativeIntegralFloat32(t *testing.T) {
	matcher := Equals(float32(-32769))
	desc := matcher.Description()
	expectedDesc := "-32769.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of -32769.
		testCase{-32769.0, MATCH_TRUE, ""},
		testCase{-32769 + 0i, MATCH_TRUE, ""},
		testCase{int32(-32769), MATCH_TRUE, ""},
		testCase{int64(-32769), MATCH_TRUE, ""},
		testCase{float32(-32769), MATCH_TRUE, ""},
		testCase{float64(-32769), MATCH_TRUE, ""},
		testCase{complex64(-32769), MATCH_TRUE, ""},
		testCase{complex128(-32769), MATCH_TRUE, ""},
		testCase{interface{}(float32(-32769)), MATCH_TRUE, ""},
		testCase{interface{}(int64(-32769)), MATCH_TRUE, ""},

		// Values that would be -32769 in two's complement.
		testCase{uint64((1 << 64) - 32769), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		testCase{int64(-32770), MATCH_FALSE, ""},
		testCase{float32(-32769.1), MATCH_FALSE, ""},
		testCase{float32(-32768.9), MATCH_FALSE, ""},
		testCase{float64(-32769.1), MATCH_FALSE, ""},
		testCase{float64(-32768.9), MATCH_FALSE, ""},
		testCase{complex128(-32768), MATCH_FALSE, ""},
		testCase{complex128(-32769 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of -32769.1.
		testCase{-32769.1, MATCH_TRUE, ""},
		testCase{-32769.1 + 0i, MATCH_TRUE, ""},
		testCase{float32(-32769.1), MATCH_TRUE, ""},
		testCase{float64(-32769.1), MATCH_TRUE, ""},
		testCase{complex64(-32769.1), MATCH_TRUE, ""},
		testCase{complex128(-32769.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int32(-32769), MATCH_FALSE, ""},
		testCase{int32(-32770), MATCH_FALSE, ""},
		testCase{int64(-32769), MATCH_FALSE, ""},
		testCase{int64(-32770), MATCH_FALSE, ""},
		testCase{float32(-32769.2), MATCH_FALSE, ""},
		testCase{float32(-32769.0), MATCH_FALSE, ""},
		testCase{float64(-32769.2), MATCH_FALSE, ""},
		testCase{complex128(-32769.1 + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeFloat32(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(float32(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float32(kExpected)
	castedInt := int64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Various types of zero.
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{int64(-1), MATCH_FALSE, ""},
		testCase{float32(1), MATCH_FALSE, ""},
		testCase{float32(-1), MATCH_FALSE, ""},
		testCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloat32(t *testing.T) {
	matcher := Equals(float32(32769))
	desc := matcher.Description()
	expectedDesc := "32769.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 32769.
		testCase{32769.0, MATCH_TRUE, ""},
		testCase{32769 + 0i, MATCH_TRUE, ""},
		testCase{int(32769), MATCH_TRUE, ""},
		testCase{int32(32769), MATCH_TRUE, ""},
		testCase{int64(32769), MATCH_TRUE, ""},
		testCase{uint(32769), MATCH_TRUE, ""},
		testCase{uint32(32769), MATCH_TRUE, ""},
		testCase{uint64(32769), MATCH_TRUE, ""},
		testCase{float32(32769), MATCH_TRUE, ""},
		testCase{float64(32769), MATCH_TRUE, ""},
		testCase{complex64(32769), MATCH_TRUE, ""},
		testCase{complex128(32769), MATCH_TRUE, ""},
		testCase{interface{}(float32(32769)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(32770), MATCH_FALSE, ""},
		testCase{uint64(32770), MATCH_FALSE, ""},
		testCase{float32(32769.1), MATCH_FALSE, ""},
		testCase{float32(32768.9), MATCH_FALSE, ""},
		testCase{float64(32769.1), MATCH_FALSE, ""},
		testCase{float64(32768.9), MATCH_FALSE, ""},
		testCase{complex128(32768), MATCH_FALSE, ""},
		testCase{complex128(32769 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
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

	cases := []testCase{
		// Various types of 32769.1.
		testCase{32769.1, MATCH_TRUE, ""},
		testCase{32769.1 + 0i, MATCH_TRUE, ""},
		testCase{float32(32769.1), MATCH_TRUE, ""},
		testCase{float64(32769.1), MATCH_TRUE, ""},
		testCase{complex64(32769.1), MATCH_TRUE, ""},
		testCase{complex128(32769.1), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int32(32769), MATCH_FALSE, ""},
		testCase{int32(32770), MATCH_FALSE, ""},
		testCase{uint64(32769), MATCH_FALSE, ""},
		testCase{uint64(32770), MATCH_FALSE, ""},
		testCase{float32(32769.2), MATCH_FALSE, ""},
		testCase{float32(32769.0), MATCH_FALSE, ""},
		testCase{float64(32769.2), MATCH_FALSE, ""},
		testCase{complex128(32769.1 + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveFloat32(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(float32(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float32(kExpected)
	castedInt := uint64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{uint64(0), MATCH_FALSE, ""},
		testCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
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
	expectedDesc := "33554432.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{int64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Single-precision floating point.
		testCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 3), MATCH_FALSE, ""},
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
	expectedDesc := "-1125899906842620"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{-1125899906842620.0, MATCH_TRUE, ""},
		testCase{-1125899906842620.0 + 0i, MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},
		testCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Values that would be kExpected in two's complement.
		testCase{uint64((1 << 64) + kExpected), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kExpected + 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralFloat64(t *testing.T) {
	const kTwoTo50 = 1 << 50
	const kExpected = -kTwoTo50 - 0.25

	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "-1125899906842620.25"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(-kTwoTo50), MATCH_FALSE, ""},
		testCase{int64(-kTwoTo50 - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeFloat64(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := int64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
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

	cases := []testCase{
		// Various types of zero.
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{int64(-1), MATCH_FALSE, ""},
		testCase{float32(1), MATCH_FALSE, ""},
		testCase{float32(-1), MATCH_FALSE, ""},
		testCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralFloat64(t *testing.T) {
	const kExpected = 1 << 50
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "1125899906842620"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 32769.
		testCase{1125899906842620.0, MATCH_TRUE, ""},
		testCase{1125899906842620.0 + 0i, MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},
		testCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralFloat64(t *testing.T) {
	const kTwoTo50 = 1 << 50
	const kExpected = kTwoTo50 + 0.25
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "1125899906842620.25"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kTwoTo50), MATCH_FALSE, ""},
		testCase{int64(kTwoTo50 - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveFloat64(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(float64(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := uint64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{uint64(0), MATCH_FALSE, ""},
		testCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
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
	expectedDesc := "18014398509481984.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
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
	expectedDesc := "-32769"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{-32769.0, MATCH_TRUE, ""},
		testCase{-32769.0 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},
		testCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Values that would be kExpected in two's complement.
		testCase{uint32((1 << 32) + kExpected), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) + kExpected), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kExpected + 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralComplex64(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = -kTwoTo20 - 0.25

	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "-1048576.25"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(-kTwoTo20), MATCH_FALSE, ""},
		testCase{int(-kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{int32(-kTwoTo20), MATCH_FALSE, ""},
		testCase{int32(-kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{int64(-kTwoTo20), MATCH_FALSE, ""},
		testCase{int64(-kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 0.75), MATCH_FALSE, ""},
		testCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 0.75), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeComplex64(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := int64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroComplex64(t *testing.T) {
	matcher := Equals(complex64(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of zero.
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{int64(-1), MATCH_FALSE, ""},
		testCase{float32(1), MATCH_FALSE, ""},
		testCase{float32(-1), MATCH_FALSE, ""},
		testCase{float64(1), MATCH_FALSE, ""},
		testCase{float64(-1), MATCH_FALSE, ""},
		testCase{complex64(0 + 2i), MATCH_FALSE, ""},
		testCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralComplex64(t *testing.T) {
	const kExpected = 1 << 20
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "1048576"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 32769.
		testCase{1048576.0, MATCH_TRUE, ""},
		testCase{1048576.0 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},
		testCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(kExpected + 1), MATCH_FALSE, ""},
		testCase{int32(kExpected + 1), MATCH_FALSE, ""},
		testCase{int64(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralComplex64(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = kTwoTo20 + 0.25
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "1048576.25"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kTwoTo20), MATCH_FALSE, ""},
		testCase{int64(kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo20), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected + 1), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 1i), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveComplex64(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := uint64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{uint64(0), MATCH_FALSE, ""},
		testCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
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
	expectedDesc := "33554433"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{int64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Single-precision floating point.
		testCase{float32(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float32(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo25 - 2), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo25 - 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 0), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 2), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestComplex64WithNonZeroImaginaryPart(t *testing.T) {
	const kRealPart = 17
	const kImagPart = 0.25i
	const kExpected = kRealPart + kImagPart
	matcher := Equals(complex64(kExpected))
	desc := matcher.Description()
	expectedDesc := "17 + 0.25i"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kRealPart + kImagPart, MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(kRealPart), MATCH_FALSE, ""},
		testCase{int8(kRealPart), MATCH_FALSE, ""},
		testCase{int16(kRealPart), MATCH_FALSE, ""},
		testCase{int32(kRealPart), MATCH_FALSE, ""},
		testCase{int64(kRealPart), MATCH_FALSE, ""},
		testCase{uint(kRealPart), MATCH_FALSE, ""},
		testCase{uint8(kRealPart), MATCH_FALSE, ""},
		testCase{uint16(kRealPart), MATCH_FALSE, ""},
		testCase{uint32(kRealPart), MATCH_FALSE, ""},
		testCase{uint64(kRealPart), MATCH_FALSE, ""},
		testCase{float32(kRealPart), MATCH_FALSE, ""},
		testCase{float64(kRealPart), MATCH_FALSE, ""},
		testCase{complex64(kRealPart), MATCH_FALSE, ""},
		testCase{complex64(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		testCase{complex64(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
		testCase{complex128(kRealPart), MATCH_FALSE, ""},
		testCase{complex128(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		testCase{complex128(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
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
	expectedDesc := "-32769"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{-32769.0, MATCH_TRUE, ""},
		testCase{-32769.0 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},
		testCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Values that would be kExpected in two's complement.
		testCase{uint32((1 << 32) + kExpected), MATCH_FALSE, ""},
		testCase{uint64((1 << 64) + kExpected), MATCH_FALSE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kExpected + 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestNegativeNonIntegralComplex128(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = -kTwoTo20 - 0.25

	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "-1048576.25"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(-kTwoTo20), MATCH_FALSE, ""},
		testCase{int(-kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{int32(-kTwoTo20), MATCH_FALSE, ""},
		testCase{int32(-kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{int64(-kTwoTo20), MATCH_FALSE, ""},
		testCase{int64(-kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 0.75), MATCH_FALSE, ""},
		testCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 0.75), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargeNegativeComplex128(t *testing.T) {
	const kExpected = -1 * (1 << 65)
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := int64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex64(kExpected + 2i), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestZeroComplex128(t *testing.T) {
	matcher := Equals(complex128(0))
	desc := matcher.Description()
	expectedDesc := "0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of zero.
		testCase{0.0, MATCH_TRUE, ""},
		testCase{0 + 0i, MATCH_TRUE, ""},
		testCase{int(0), MATCH_TRUE, ""},
		testCase{int8(0), MATCH_TRUE, ""},
		testCase{int16(0), MATCH_TRUE, ""},
		testCase{int32(0), MATCH_TRUE, ""},
		testCase{int64(0), MATCH_TRUE, ""},
		testCase{uint(0), MATCH_TRUE, ""},
		testCase{uint8(0), MATCH_TRUE, ""},
		testCase{uint16(0), MATCH_TRUE, ""},
		testCase{uint32(0), MATCH_TRUE, ""},
		testCase{uint64(0), MATCH_TRUE, ""},
		testCase{float32(0), MATCH_TRUE, ""},
		testCase{float64(0), MATCH_TRUE, ""},
		testCase{complex64(0), MATCH_TRUE, ""},
		testCase{complex128(0), MATCH_TRUE, ""},
		testCase{interface{}(float32(0)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(1), MATCH_FALSE, ""},
		testCase{int64(-1), MATCH_FALSE, ""},
		testCase{float32(1), MATCH_FALSE, ""},
		testCase{float32(-1), MATCH_FALSE, ""},
		testCase{float64(1), MATCH_FALSE, ""},
		testCase{float64(-1), MATCH_FALSE, ""},
		testCase{complex64(0 + 2i), MATCH_FALSE, ""},
		testCase{complex128(0 + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveIntegralComplex128(t *testing.T) {
	const kExpected = 1 << 20
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "1048576"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of 32769.
		testCase{1048576.0, MATCH_TRUE, ""},
		testCase{1048576.0 + 0i, MATCH_TRUE, ""},
		testCase{int(kExpected), MATCH_TRUE, ""},
		testCase{int32(kExpected), MATCH_TRUE, ""},
		testCase{int64(kExpected), MATCH_TRUE, ""},
		testCase{uint(kExpected), MATCH_TRUE, ""},
		testCase{uint32(kExpected), MATCH_TRUE, ""},
		testCase{uint64(kExpected), MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},
		testCase{interface{}(float64(kExpected)), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(kExpected + 1), MATCH_FALSE, ""},
		testCase{int32(kExpected + 1), MATCH_FALSE, ""},
		testCase{int64(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint32(kExpected + 1), MATCH_FALSE, ""},
		testCase{uint64(kExpected + 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - (1 << 30)), MATCH_FALSE, ""},
		testCase{float32(kExpected + (1 << 30)), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.5), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.5), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},

		// Non-numeric types.
		testCase{uintptr(0), MATCH_UNDEFINED, "which is not numeric"},
		testCase{true, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{make(chan int), MATCH_UNDEFINED, "which is not numeric"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not numeric"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not numeric"},
		testCase{"taco", MATCH_UNDEFINED, "which is not numeric"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not numeric"},
	}

	checkTestCases(t, matcher, cases)
}

func TestPositiveNonIntegralComplex128(t *testing.T) {
	const kTwoTo20 = 1 << 20
	const kExpected = kTwoTo20 + 0.25
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "1048576.25"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int64(kTwoTo20), MATCH_FALSE, ""},
		testCase{int64(kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo20), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo20 - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected - 1), MATCH_FALSE, ""},
		testCase{float32(kExpected + 1), MATCH_FALSE, ""},
		testCase{float64(kExpected - 0.25), MATCH_FALSE, ""},
		testCase{float64(kExpected + 0.25), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex64(kExpected - 1i), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1), MATCH_FALSE, ""},
		testCase{complex128(kExpected - 1i), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestLargePositiveComplex128(t *testing.T) {
	const kExpected = 1 << 65
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	floatExpected := float64(kExpected)
	castedInt := uint64(floatExpected)

	cases := []testCase{
		// Equal values of numeric type.
		testCase{kExpected + 0i, MATCH_TRUE, ""},
		testCase{float32(kExpected), MATCH_TRUE, ""},
		testCase{float64(kExpected), MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{castedInt, MATCH_FALSE, ""},
		testCase{int64(0), MATCH_FALSE, ""},
		testCase{int64(math.MinInt64), MATCH_FALSE, ""},
		testCase{int64(math.MaxInt64), MATCH_FALSE, ""},
		testCase{uint64(0), MATCH_FALSE, ""},
		testCase{uint64(math.MaxUint64), MATCH_FALSE, ""},
		testCase{float32(kExpected / 2), MATCH_FALSE, ""},
		testCase{float64(kExpected / 2), MATCH_FALSE, ""},
		testCase{complex128(kExpected + 2i), MATCH_FALSE, ""},
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
	expectedDesc := "18014398509481984.0"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Integers.
		testCase{int64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		testCase{uint64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Double-precision floating point.
		testCase{float64(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{float64(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},

		testCase{complex128(kTwoTo54 - 2), MATCH_FALSE, ""},
		testCase{complex128(kTwoTo54 - 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 0), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 1), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 2), MATCH_TRUE, ""},
		testCase{complex128(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	checkTestCases(t, matcher, cases)
}

func TestComplex128WithNonZeroImaginaryPart(t *testing.T) {
	const kRealPart = 17
	const kImagPart = 0.25i
	const kExpected = kRealPart + kImagPart
	matcher := Equals(complex128(kExpected))
	desc := matcher.Description()
	expectedDesc := "17 + 0.25i"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Various types of the expected value.
		testCase{kExpected, MATCH_TRUE, ""},
		testCase{kRealPart + kImagPart, MATCH_TRUE, ""},
		testCase{complex64(kExpected), MATCH_TRUE, ""},
		testCase{complex128(kExpected), MATCH_TRUE, ""},

		// Non-equal values of numeric type.
		testCase{int(kRealPart), MATCH_FALSE, ""},
		testCase{int8(kRealPart), MATCH_FALSE, ""},
		testCase{int16(kRealPart), MATCH_FALSE, ""},
		testCase{int32(kRealPart), MATCH_FALSE, ""},
		testCase{int64(kRealPart), MATCH_FALSE, ""},
		testCase{uint(kRealPart), MATCH_FALSE, ""},
		testCase{uint8(kRealPart), MATCH_FALSE, ""},
		testCase{uint16(kRealPart), MATCH_FALSE, ""},
		testCase{uint32(kRealPart), MATCH_FALSE, ""},
		testCase{uint64(kRealPart), MATCH_FALSE, ""},
		testCase{float32(kRealPart), MATCH_FALSE, ""},
		testCase{float64(kRealPart), MATCH_FALSE, ""},
		testCase{complex64(kRealPart), MATCH_FALSE, ""},
		testCase{complex64(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		testCase{complex64(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
		testCase{complex128(kRealPart), MATCH_FALSE, ""},
		testCase{complex128(kRealPart + kImagPart + 0.5), MATCH_FALSE, ""},
		testCase{complex128(kRealPart + kImagPart + 0.5i), MATCH_FALSE, ""},
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
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// int channels
		testCase{nilChan1, MATCH_TRUE, ""},
		testCase{nilChan2, MATCH_TRUE, ""},
		testCase{nonNilChan1, MATCH_FALSE, ""},

		// uint channels
		testCase{nilChan3, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{nonNilChan2, MATCH_UNDEFINED, "which is not a chan int"},

		// Other types.
		testCase{0, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{bool(false), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int8(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int16(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int32(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int64(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint8(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint16(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint32(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint64(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{true, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a chan int"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a chan int"},
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
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// int channels
		testCase{nonNilChan1, MATCH_TRUE, ""},
		testCase{nonNilChan2, MATCH_TRUE, ""},
		testCase{nilChan1, MATCH_FALSE, ""},

		// uint channels
		testCase{nilChan2, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{nonNilChan3, MATCH_UNDEFINED, "which is not a chan int"},

		// Other types.
		testCase{0, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{bool(false), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int8(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int16(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int32(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{int64(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint8(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint16(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint32(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{uint64(0), MATCH_UNDEFINED, "which is not a chan int"},
		testCase{true, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{func() {}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a chan int"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a chan int"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a chan int"},
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
	expectedDesc := "TODO"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []testCase{
		// Functions.
		testCase{func1, MATCH_TRUE, ""},
		testCase{func2, MATCH_FALSE, ""},
		testCase{func3, MATCH_FALSE, ""},

		// Other types.
		testCase{0, MATCH_UNDEFINED, "which is not a function"},
		testCase{bool(false), MATCH_UNDEFINED, "which is not a function"},
		testCase{int(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{int8(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{int16(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{int32(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{int64(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{uint(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{uint8(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{uint16(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{uint32(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{uint64(0), MATCH_UNDEFINED, "which is not a function"},
		testCase{true, MATCH_UNDEFINED, "which is not a function"},
		testCase{[...]int{}, MATCH_UNDEFINED, "which is not a function"},
		testCase{map[int]int{}, MATCH_UNDEFINED, "which is not a function"},
		testCase{&someInt, MATCH_UNDEFINED, "which is not a function"},
		testCase{[]int{}, MATCH_UNDEFINED, "which is not a function"},
		testCase{"taco", MATCH_UNDEFINED, "which is not a function"},
		testCase{testCase{}, MATCH_UNDEFINED, "which is not a function"},
	}

	checkTestCases(t, matcher, cases)
}
