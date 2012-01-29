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

package oglematchers_test

import (
	"fmt"
	. "github.com/jacobsa/oglematchers"
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
	expectedResult bool
	shouldBeFatal  bool
	expectedError  string
}

func checkTestCases(t *testing.T, matcher Matcher, cases []equalsTestCase) {
	for i, c := range cases {
		err := matcher.Matches(c.candidate)

		if (err == nil) != c.expectedResult {
			t.Errorf(
				"Case %d (candidate %v): expected %v, got error: %v",
				i,
				c.candidate,
				c.expectedResult,
				err)
		}

		if err == nil {
			continue
		}

		if _, isFatal := err.(*FatalError); isFatal != c.shouldBeFatal {
			t.Errorf(
				"Case %d (candidate %v): expected fatal %v, got fatal %v",
				i,
				c.candidate,
				c.shouldBeFatal,
				isFatal)
		}

		if err.Error() != c.expectedError {
			t.Errorf("Case %d: expected error %v, got %v", i, c.expectedError, err.Error())
		}
	}
}

////////////////////////////////////////////////////////////
// nil
////////////////////////////////////////////////////////////

func TestEqualsNil(t *testing.T) {
	matcher := Equals(nil)
	desc := matcher.Description()
	expectedDesc := "is nil"

	if desc != expectedDesc {
		t.Errorf("Expected description \"%s\", got \"%s\".", expectedDesc, desc)
	}

	cases := []equalsTestCase{
		// Legal types
		equalsTestCase{nil, true, false, ""},
		equalsTestCase{chan int(nil), true, false, ""},
		equalsTestCase{(func())(nil), true, false, ""},
		equalsTestCase{interface{}(nil), true, false, ""},
		equalsTestCase{map[int]int(nil), true, false, ""},
		equalsTestCase{(*int)(nil), true, false, ""},
		equalsTestCase{[]int(nil), true, false, ""},

		equalsTestCase{make(chan int), false, false, ""},
		equalsTestCase{func() {}, false, false, ""},
		equalsTestCase{map[int]int{}, false, false, ""},
		equalsTestCase{&someInt, false, false, ""},
		equalsTestCase{[]int{}, false, false, ""},

		// Illegal types
		equalsTestCase{17, false, true, "which cannot be compared to nil"},
		equalsTestCase{int8(17), false, true, "which cannot be compared to nil"},
		equalsTestCase{uintptr(17), false, true, "which cannot be compared to nil"},
		equalsTestCase{[...]int{}, false, true, "which cannot be compared to nil"},
		equalsTestCase{"taco", false, true, "which cannot be compared to nil"},
		equalsTestCase{equalsTestCase{}, false, true, "which cannot be compared to nil"},
		equalsTestCase{unsafe.Pointer(&someInt), false, true, "which cannot be compared to nil"},
	}

	checkTestCases(t, matcher, cases)
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
		equalsTestCase{-1073741824, true, false, ""},
		equalsTestCase{-1073741824.0, true, false, ""},
		equalsTestCase{-1073741824 + 0i, true, false, ""},
		equalsTestCase{int(-1073741824), true, false, ""},
		equalsTestCase{int32(-1073741824), true, false, ""},
		equalsTestCase{int64(-1073741824), true, false, ""},
		equalsTestCase{float32(-1073741824), true, false, ""},
		equalsTestCase{float64(-1073741824), true, false, ""},
		equalsTestCase{complex64(-1073741824), true, false, ""},
		equalsTestCase{complex128(-1073741824), true, false, ""},
		equalsTestCase{interface{}(int(-1073741824)), true, false, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), false, false, ""},
		equalsTestCase{int32(-1073741823), false, false, ""},
		equalsTestCase{int64(-1073741823), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), false, false, ""},
		equalsTestCase{float64(-1073741823.9), false, false, ""},
		equalsTestCase{complex128(-1073741823), false, false, ""},
		equalsTestCase{complex128(-1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1073741824, true, false, ""},
		equalsTestCase{1073741824.0, true, false, ""},
		equalsTestCase{1073741824 + 0i, true, false, ""},
		equalsTestCase{int(1073741824), true, false, ""},
		equalsTestCase{uint(1073741824), true, false, ""},
		equalsTestCase{int32(1073741824), true, false, ""},
		equalsTestCase{int64(1073741824), true, false, ""},
		equalsTestCase{uint32(1073741824), true, false, ""},
		equalsTestCase{uint64(1073741824), true, false, ""},
		equalsTestCase{float32(1073741824), true, false, ""},
		equalsTestCase{float64(1073741824), true, false, ""},
		equalsTestCase{complex64(1073741824), true, false, ""},
		equalsTestCase{complex128(1073741824), true, false, ""},
		equalsTestCase{interface{}(int(1073741824)), true, false, ""},
		equalsTestCase{interface{}(uint(1073741824)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), false, false, ""},
		equalsTestCase{int32(1073741823), false, false, ""},
		equalsTestCase{int64(1073741823), false, false, ""},
		equalsTestCase{float64(1073741824.1), false, false, ""},
		equalsTestCase{float64(1073741823.9), false, false, ""},
		equalsTestCase{complex128(1073741823), false, false, ""},
		equalsTestCase{complex128(1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{-1073741824, true, false, ""},
		equalsTestCase{-1073741824.0, true, false, ""},
		equalsTestCase{-1073741824 + 0i, true, false, ""},
		equalsTestCase{int(-1073741824), true, false, ""},
		equalsTestCase{int32(-1073741824), true, false, ""},
		equalsTestCase{int64(-1073741824), true, false, ""},
		equalsTestCase{float32(-1073741824), true, false, ""},
		equalsTestCase{float64(-1073741824), true, false, ""},
		equalsTestCase{complex64(-1073741824), true, false, ""},
		equalsTestCase{complex128(-1073741824), true, false, ""},
		equalsTestCase{interface{}(int(-1073741824)), true, false, ""},
		equalsTestCase{interface{}(float64(-1073741824)), true, false, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), false, false, ""},
		equalsTestCase{int32(-1073741823), false, false, ""},
		equalsTestCase{int64(-1073741823), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), false, false, ""},
		equalsTestCase{float64(-1073741823.9), false, false, ""},
		equalsTestCase{complex128(-1073741823), false, false, ""},
		equalsTestCase{complex128(-1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1073741824, true, false, ""},
		equalsTestCase{1073741824.0, true, false, ""},
		equalsTestCase{1073741824 + 0i, true, false, ""},
		equalsTestCase{int(1073741824), true, false, ""},
		equalsTestCase{int32(1073741824), true, false, ""},
		equalsTestCase{int64(1073741824), true, false, ""},
		equalsTestCase{uint(1073741824), true, false, ""},
		equalsTestCase{uint32(1073741824), true, false, ""},
		equalsTestCase{uint64(1073741824), true, false, ""},
		equalsTestCase{float32(1073741824), true, false, ""},
		equalsTestCase{float64(1073741824), true, false, ""},
		equalsTestCase{complex64(1073741824), true, false, ""},
		equalsTestCase{complex128(1073741824), true, false, ""},
		equalsTestCase{interface{}(int(1073741824)), true, false, ""},
		equalsTestCase{interface{}(float64(1073741824)), true, false, ""},

		// Values that would be 1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), false, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), false, false, ""},
		equalsTestCase{int32(1073741823), false, false, ""},
		equalsTestCase{int64(1073741823), false, false, ""},
		equalsTestCase{uint(1073741823), false, false, ""},
		equalsTestCase{uint32(1073741823), false, false, ""},
		equalsTestCase{uint64(1073741823), false, false, ""},
		equalsTestCase{float64(1073741824.1), false, false, ""},
		equalsTestCase{float64(1073741823.9), false, false, ""},
		equalsTestCase{complex128(1073741823), false, false, ""},
		equalsTestCase{complex128(1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{17.1, true, false, ""},
		equalsTestCase{17.1, true, false, ""},
		equalsTestCase{17.1 + 0i, true, false, ""},
		equalsTestCase{float32(17.1), true, false, ""},
		equalsTestCase{float64(17.1), true, false, ""},
		equalsTestCase{complex64(17.1), true, false, ""},
		equalsTestCase{complex128(17.1), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{17, false, false, ""},
		equalsTestCase{17.2, false, false, ""},
		equalsTestCase{18, false, false, ""},
		equalsTestCase{int(17), false, false, ""},
		equalsTestCase{int(18), false, false, ""},
		equalsTestCase{int32(17), false, false, ""},
		equalsTestCase{int64(17), false, false, ""},
		equalsTestCase{uint(17), false, false, ""},
		equalsTestCase{uint32(17), false, false, ""},
		equalsTestCase{uint64(17), false, false, ""},
		equalsTestCase{complex128(17.1 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{false, true, false, ""},
		equalsTestCase{bool(false), true, false, ""},

		equalsTestCase{true, false, false, ""},
		equalsTestCase{bool(true), false, false, ""},

		// Other types.
		equalsTestCase{int(0), false, true, "which is not a bool"},
		equalsTestCase{int8(0), false, true, "which is not a bool"},
		equalsTestCase{int16(0), false, true, "which is not a bool"},
		equalsTestCase{int32(0), false, true, "which is not a bool"},
		equalsTestCase{int64(0), false, true, "which is not a bool"},
		equalsTestCase{uint(0), false, true, "which is not a bool"},
		equalsTestCase{uint8(0), false, true, "which is not a bool"},
		equalsTestCase{uint16(0), false, true, "which is not a bool"},
		equalsTestCase{uint32(0), false, true, "which is not a bool"},
		equalsTestCase{uint64(0), false, true, "which is not a bool"},
		equalsTestCase{uintptr(0), false, true, "which is not a bool"},
		equalsTestCase{[...]int{}, false, true, "which is not a bool"},
		equalsTestCase{make(chan int), false, true, "which is not a bool"},
		equalsTestCase{func() {}, false, true, "which is not a bool"},
		equalsTestCase{map[int]int{}, false, true, "which is not a bool"},
		equalsTestCase{&someInt, false, true, "which is not a bool"},
		equalsTestCase{[]int{}, false, true, "which is not a bool"},
		equalsTestCase{"taco", false, true, "which is not a bool"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a bool"},
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
		equalsTestCase{true, true, false, ""},
		equalsTestCase{bool(true), true, false, ""},

		equalsTestCase{false, false, false, ""},
		equalsTestCase{bool(false), false, false, ""},

		// Other types.
		equalsTestCase{int(1), false, true, "which is not a bool"},
		equalsTestCase{int8(1), false, true, "which is not a bool"},
		equalsTestCase{int16(1), false, true, "which is not a bool"},
		equalsTestCase{int32(1), false, true, "which is not a bool"},
		equalsTestCase{int64(1), false, true, "which is not a bool"},
		equalsTestCase{uint(1), false, true, "which is not a bool"},
		equalsTestCase{uint8(1), false, true, "which is not a bool"},
		equalsTestCase{uint16(1), false, true, "which is not a bool"},
		equalsTestCase{uint32(1), false, true, "which is not a bool"},
		equalsTestCase{uint64(1), false, true, "which is not a bool"},
		equalsTestCase{uintptr(1), false, true, "which is not a bool"},
		equalsTestCase{[...]int{}, false, true, "which is not a bool"},
		equalsTestCase{make(chan int), false, true, "which is not a bool"},
		equalsTestCase{func() {}, false, true, "which is not a bool"},
		equalsTestCase{map[int]int{}, false, true, "which is not a bool"},
		equalsTestCase{&someInt, false, true, "which is not a bool"},
		equalsTestCase{[]int{}, false, true, "which is not a bool"},
		equalsTestCase{"taco", false, true, "which is not a bool"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a bool"},
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
		equalsTestCase{-1073741824, true, false, ""},
		equalsTestCase{-1073741824.0, true, false, ""},
		equalsTestCase{-1073741824 + 0i, true, false, ""},
		equalsTestCase{int(-1073741824), true, false, ""},
		equalsTestCase{int32(-1073741824), true, false, ""},
		equalsTestCase{int64(-1073741824), true, false, ""},
		equalsTestCase{float32(-1073741824), true, false, ""},
		equalsTestCase{float64(-1073741824), true, false, ""},
		equalsTestCase{complex64(-1073741824), true, false, ""},
		equalsTestCase{complex128(-1073741824), true, false, ""},
		equalsTestCase{interface{}(int(-1073741824)), true, false, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), false, false, ""},
		equalsTestCase{int32(-1073741823), false, false, ""},
		equalsTestCase{int64(-1073741823), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), false, false, ""},
		equalsTestCase{float64(-1073741823.9), false, false, ""},
		equalsTestCase{complex128(-1073741823), false, false, ""},
		equalsTestCase{complex128(-1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1073741824, true, false, ""},
		equalsTestCase{1073741824.0, true, false, ""},
		equalsTestCase{1073741824 + 0i, true, false, ""},
		equalsTestCase{int(1073741824), true, false, ""},
		equalsTestCase{uint(1073741824), true, false, ""},
		equalsTestCase{int32(1073741824), true, false, ""},
		equalsTestCase{int64(1073741824), true, false, ""},
		equalsTestCase{uint32(1073741824), true, false, ""},
		equalsTestCase{uint64(1073741824), true, false, ""},
		equalsTestCase{float32(1073741824), true, false, ""},
		equalsTestCase{float64(1073741824), true, false, ""},
		equalsTestCase{complex64(1073741824), true, false, ""},
		equalsTestCase{complex128(1073741824), true, false, ""},
		equalsTestCase{interface{}(int(1073741824)), true, false, ""},
		equalsTestCase{interface{}(uint(1073741824)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), false, false, ""},
		equalsTestCase{int32(1073741823), false, false, ""},
		equalsTestCase{int64(1073741823), false, false, ""},
		equalsTestCase{float64(1073741824.1), false, false, ""},
		equalsTestCase{float64(1073741823.9), false, false, ""},
		equalsTestCase{complex128(1073741823), false, false, ""},
		equalsTestCase{complex128(1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{-17, true, false, ""},
		equalsTestCase{-17.0, true, false, ""},
		equalsTestCase{-17 + 0i, true, false, ""},
		equalsTestCase{int(-17), true, false, ""},
		equalsTestCase{int8(-17), true, false, ""},
		equalsTestCase{int16(-17), true, false, ""},
		equalsTestCase{int32(-17), true, false, ""},
		equalsTestCase{int64(-17), true, false, ""},
		equalsTestCase{float32(-17), true, false, ""},
		equalsTestCase{float64(-17), true, false, ""},
		equalsTestCase{complex64(-17), true, false, ""},
		equalsTestCase{complex128(-17), true, false, ""},
		equalsTestCase{interface{}(int(-17)), true, false, ""},

		// Values that would be -17 in two's complement.
		equalsTestCase{uint((1 << 32) - 17), false, false, ""},
		equalsTestCase{uint8((1 << 8) - 17), false, false, ""},
		equalsTestCase{uint16((1 << 16) - 17), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 17), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 17), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-16), false, false, ""},
		equalsTestCase{int8(-16), false, false, ""},
		equalsTestCase{int16(-16), false, false, ""},
		equalsTestCase{int32(-16), false, false, ""},
		equalsTestCase{int64(-16), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float32(-17.1), false, false, ""},
		equalsTestCase{float32(-16.9), false, false, ""},
		equalsTestCase{complex64(-16), false, false, ""},
		equalsTestCase{complex64(-17 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr((1 << 32) - 17), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{-17}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{-17}, false, true, "which is not numeric"},
		equalsTestCase{"-17", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{0, true, false, ""},
		equalsTestCase{0.0, true, false, ""},
		equalsTestCase{0 + 0i, true, false, ""},
		equalsTestCase{int(0), true, false, ""},
		equalsTestCase{int8(0), true, false, ""},
		equalsTestCase{int16(0), true, false, ""},
		equalsTestCase{int32(0), true, false, ""},
		equalsTestCase{int64(0), true, false, ""},
		equalsTestCase{float32(0), true, false, ""},
		equalsTestCase{float64(0), true, false, ""},
		equalsTestCase{complex64(0), true, false, ""},
		equalsTestCase{complex128(0), true, false, ""},
		equalsTestCase{interface{}(int(0)), true, false, ""},
		equalsTestCase{uint(0), true, false, ""},
		equalsTestCase{uint8(0), true, false, ""},
		equalsTestCase{uint16(0), true, false, ""},
		equalsTestCase{uint32(0), true, false, ""},
		equalsTestCase{uint64(0), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1), false, false, ""},
		equalsTestCase{int8(1), false, false, ""},
		equalsTestCase{int16(1), false, false, ""},
		equalsTestCase{int32(1), false, false, ""},
		equalsTestCase{int64(1), false, false, ""},
		equalsTestCase{float32(-0.1), false, false, ""},
		equalsTestCase{float32(0.1), false, false, ""},
		equalsTestCase{complex64(1), false, false, ""},
		equalsTestCase{complex64(0 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{0}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{0}, false, true, "which is not numeric"},
		equalsTestCase{"0", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{17, true, false, ""},
		equalsTestCase{17.0, true, false, ""},
		equalsTestCase{17 + 0i, true, false, ""},
		equalsTestCase{int(17), true, false, ""},
		equalsTestCase{int8(17), true, false, ""},
		equalsTestCase{int16(17), true, false, ""},
		equalsTestCase{int32(17), true, false, ""},
		equalsTestCase{int64(17), true, false, ""},
		equalsTestCase{float32(17), true, false, ""},
		equalsTestCase{float64(17), true, false, ""},
		equalsTestCase{complex64(17), true, false, ""},
		equalsTestCase{complex128(17), true, false, ""},
		equalsTestCase{interface{}(int(17)), true, false, ""},
		equalsTestCase{uint(17), true, false, ""},
		equalsTestCase{uint8(17), true, false, ""},
		equalsTestCase{uint16(17), true, false, ""},
		equalsTestCase{uint32(17), true, false, ""},
		equalsTestCase{uint64(17), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(16), false, false, ""},
		equalsTestCase{int8(16), false, false, ""},
		equalsTestCase{int16(16), false, false, ""},
		equalsTestCase{int32(16), false, false, ""},
		equalsTestCase{int64(16), false, false, ""},
		equalsTestCase{float32(16.9), false, false, ""},
		equalsTestCase{float32(17.1), false, false, ""},
		equalsTestCase{complex64(16), false, false, ""},
		equalsTestCase{complex64(17 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(17), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{17}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{17}, false, true, "which is not numeric"},
		equalsTestCase{"17", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{-32766, true, false, ""},
		equalsTestCase{-32766.0, true, false, ""},
		equalsTestCase{-32766 + 0i, true, false, ""},
		equalsTestCase{int(-32766), true, false, ""},
		equalsTestCase{int16(-32766), true, false, ""},
		equalsTestCase{int32(-32766), true, false, ""},
		equalsTestCase{int64(-32766), true, false, ""},
		equalsTestCase{float32(-32766), true, false, ""},
		equalsTestCase{float64(-32766), true, false, ""},
		equalsTestCase{complex64(-32766), true, false, ""},
		equalsTestCase{complex128(-32766), true, false, ""},
		equalsTestCase{interface{}(int(-32766)), true, false, ""},

		// Values that would be -32766 in two's complement.
		equalsTestCase{uint((1 << 32) - 32766), false, false, ""},
		equalsTestCase{uint16((1 << 16) - 32766), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 32766), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 32766), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-16), false, false, ""},
		equalsTestCase{int8(-16), false, false, ""},
		equalsTestCase{int16(-16), false, false, ""},
		equalsTestCase{int32(-16), false, false, ""},
		equalsTestCase{int64(-16), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float32(-32766.1), false, false, ""},
		equalsTestCase{float32(-32765.9), false, false, ""},
		equalsTestCase{complex64(-32766.1), false, false, ""},
		equalsTestCase{complex64(-32766 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr((1 << 32) - 32766), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{-32766}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{-32766}, false, true, "which is not numeric"},
		equalsTestCase{"-32766", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{0, true, false, ""},
		equalsTestCase{0.0, true, false, ""},
		equalsTestCase{0 + 0i, true, false, ""},
		equalsTestCase{int(0), true, false, ""},
		equalsTestCase{int8(0), true, false, ""},
		equalsTestCase{int16(0), true, false, ""},
		equalsTestCase{int32(0), true, false, ""},
		equalsTestCase{int64(0), true, false, ""},
		equalsTestCase{float32(0), true, false, ""},
		equalsTestCase{float64(0), true, false, ""},
		equalsTestCase{complex64(0), true, false, ""},
		equalsTestCase{complex128(0), true, false, ""},
		equalsTestCase{interface{}(int(0)), true, false, ""},
		equalsTestCase{uint(0), true, false, ""},
		equalsTestCase{uint8(0), true, false, ""},
		equalsTestCase{uint16(0), true, false, ""},
		equalsTestCase{uint32(0), true, false, ""},
		equalsTestCase{uint64(0), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1), false, false, ""},
		equalsTestCase{int8(1), false, false, ""},
		equalsTestCase{int16(1), false, false, ""},
		equalsTestCase{int32(1), false, false, ""},
		equalsTestCase{int64(1), false, false, ""},
		equalsTestCase{float32(-0.1), false, false, ""},
		equalsTestCase{float32(0.1), false, false, ""},
		equalsTestCase{complex64(1), false, false, ""},
		equalsTestCase{complex64(0 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{0}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{0}, false, true, "which is not numeric"},
		equalsTestCase{"0", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{32765, true, false, ""},
		equalsTestCase{32765.0, true, false, ""},
		equalsTestCase{32765 + 0i, true, false, ""},
		equalsTestCase{int(32765), true, false, ""},
		equalsTestCase{int16(32765), true, false, ""},
		equalsTestCase{int32(32765), true, false, ""},
		equalsTestCase{int64(32765), true, false, ""},
		equalsTestCase{float32(32765), true, false, ""},
		equalsTestCase{float64(32765), true, false, ""},
		equalsTestCase{complex64(32765), true, false, ""},
		equalsTestCase{complex128(32765), true, false, ""},
		equalsTestCase{interface{}(int(32765)), true, false, ""},
		equalsTestCase{uint(32765), true, false, ""},
		equalsTestCase{uint16(32765), true, false, ""},
		equalsTestCase{uint32(32765), true, false, ""},
		equalsTestCase{uint64(32765), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(32764), false, false, ""},
		equalsTestCase{int16(32764), false, false, ""},
		equalsTestCase{int32(32764), false, false, ""},
		equalsTestCase{int64(32764), false, false, ""},
		equalsTestCase{float32(32764.9), false, false, ""},
		equalsTestCase{float32(32765.1), false, false, ""},
		equalsTestCase{complex64(32765.9), false, false, ""},
		equalsTestCase{complex64(32765 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(32765), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{32765}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{32765}, false, true, "which is not numeric"},
		equalsTestCase{"32765", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{-1073741824, true, false, ""},
		equalsTestCase{-1073741824.0, true, false, ""},
		equalsTestCase{-1073741824 + 0i, true, false, ""},
		equalsTestCase{int(-1073741824), true, false, ""},
		equalsTestCase{int32(-1073741824), true, false, ""},
		equalsTestCase{int64(-1073741824), true, false, ""},
		equalsTestCase{float32(-1073741824), true, false, ""},
		equalsTestCase{float64(-1073741824), true, false, ""},
		equalsTestCase{complex64(-1073741824), true, false, ""},
		equalsTestCase{complex128(-1073741824), true, false, ""},
		equalsTestCase{interface{}(int(-1073741824)), true, false, ""},

		// Values that would be -1073741824 in two's complement.
		equalsTestCase{uint((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint32((1 << 32) - 1073741824), false, false, ""},
		equalsTestCase{uint64((1 << 64) - 1073741824), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int(-1073741823), false, false, ""},
		equalsTestCase{int32(-1073741823), false, false, ""},
		equalsTestCase{int64(-1073741823), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1073741824.1), false, false, ""},
		equalsTestCase{float64(-1073741823.9), false, false, ""},
		equalsTestCase{complex128(-1073741823), false, false, ""},
		equalsTestCase{complex128(-1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1073741824, true, false, ""},
		equalsTestCase{1073741824.0, true, false, ""},
		equalsTestCase{1073741824 + 0i, true, false, ""},
		equalsTestCase{int(1073741824), true, false, ""},
		equalsTestCase{uint(1073741824), true, false, ""},
		equalsTestCase{int32(1073741824), true, false, ""},
		equalsTestCase{int64(1073741824), true, false, ""},
		equalsTestCase{uint32(1073741824), true, false, ""},
		equalsTestCase{uint64(1073741824), true, false, ""},
		equalsTestCase{float32(1073741824), true, false, ""},
		equalsTestCase{float64(1073741824), true, false, ""},
		equalsTestCase{complex64(1073741824), true, false, ""},
		equalsTestCase{complex128(1073741824), true, false, ""},
		equalsTestCase{interface{}(int(1073741824)), true, false, ""},
		equalsTestCase{interface{}(uint(1073741824)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(1073741823), false, false, ""},
		equalsTestCase{int32(1073741823), false, false, ""},
		equalsTestCase{int64(1073741823), false, false, ""},
		equalsTestCase{float64(1073741824.1), false, false, ""},
		equalsTestCase{float64(1073741823.9), false, false, ""},
		equalsTestCase{complex128(1073741823), false, false, ""},
		equalsTestCase{complex128(1073741824 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{-1099511627776.0, true, false, ""},
		equalsTestCase{-1099511627776 + 0i, true, false, ""},
		equalsTestCase{int64(-1099511627776), true, false, ""},
		equalsTestCase{float32(-1099511627776), true, false, ""},
		equalsTestCase{float64(-1099511627776), true, false, ""},
		equalsTestCase{complex64(-1099511627776), true, false, ""},
		equalsTestCase{complex128(-1099511627776), true, false, ""},
		equalsTestCase{interface{}(int64(-1099511627776)), true, false, ""},

		// Values that would be -1099511627776 in two's complement.
		equalsTestCase{uint64((1 << 64) - 1099511627776), false, false, ""},

		// Non-equal values of signed integer type.
		equalsTestCase{int64(-1099511627775), false, false, ""},

		// Non-equal values of other numeric types.
		equalsTestCase{float64(-1099511627776.1), false, false, ""},
		equalsTestCase{float64(-1099511627775.9), false, false, ""},
		equalsTestCase{complex128(-1099511627775), false, false, ""},
		equalsTestCase{complex128(-1099511627776 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1099511627776.0, true, false, ""},
		equalsTestCase{1099511627776 + 0i, true, false, ""},
		equalsTestCase{int64(1099511627776), true, false, ""},
		equalsTestCase{uint64(1099511627776), true, false, ""},
		equalsTestCase{float32(1099511627776), true, false, ""},
		equalsTestCase{float64(1099511627776), true, false, ""},
		equalsTestCase{complex64(1099511627776), true, false, ""},
		equalsTestCase{complex128(1099511627776), true, false, ""},
		equalsTestCase{interface{}(int64(1099511627776)), true, false, ""},
		equalsTestCase{interface{}(uint64(1099511627776)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1099511627775), false, false, ""},
		equalsTestCase{uint64(1099511627775), false, false, ""},
		equalsTestCase{float64(1099511627776.1), false, false, ""},
		equalsTestCase{float64(1099511627775.9), false, false, ""},
		equalsTestCase{complex128(1099511627775), false, false, ""},
		equalsTestCase{complex128(1099511627776 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{int64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{int64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), false, false, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float32(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{float64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), false, false, ""},
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
		equalsTestCase{int64(kTwoTo54 + 0), false, false, ""},
		equalsTestCase{int64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 2), false, false, ""},

		equalsTestCase{uint64(kTwoTo54 + 0), false, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{float64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 3), false, false, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), false, false, ""},
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
		equalsTestCase{17, true, false, ""},
		equalsTestCase{17.0, true, false, ""},
		equalsTestCase{17 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int8(kExpected), true, false, ""},
		equalsTestCase{int16(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint8(kExpected), true, false, ""},
		equalsTestCase{uint16(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, false, false, ""},
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int8(kExpected + 1), false, false, ""},
		equalsTestCase{int16(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint8(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{65553, true, false, ""},
		equalsTestCase{65553.0, true, false, ""},
		equalsTestCase{65553 + 0i, true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int16(17), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(17), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},
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
		equalsTestCase{int64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{int64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), false, false, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float32(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{float64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), false, false, ""},
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
		equalsTestCase{17, true, false, ""},
		equalsTestCase{17.0, true, false, ""},
		equalsTestCase{17 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int8(kExpected), true, false, ""},
		equalsTestCase{int16(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint8(kExpected), true, false, ""},
		equalsTestCase{uint16(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, false, false, ""},
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int8(kExpected + 1), false, false, ""},
		equalsTestCase{int16(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint8(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{17, true, false, ""},
		equalsTestCase{17.0, true, false, ""},
		equalsTestCase{17 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int8(kExpected), true, false, ""},
		equalsTestCase{int16(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint8(kExpected), true, false, ""},
		equalsTestCase{uint16(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, false, false, ""},
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int8(kExpected + 1), false, false, ""},
		equalsTestCase{int16(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint8(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{273, true, false, ""},
		equalsTestCase{273.0, true, false, ""},
		equalsTestCase{273 + 0i, true, false, ""},
		equalsTestCase{int16(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint16(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int8(17), false, false, ""},
		equalsTestCase{int16(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint8(17), false, false, ""},
		equalsTestCase{uint16(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},
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
		equalsTestCase{17, true, false, ""},
		equalsTestCase{17.0, true, false, ""},
		equalsTestCase{17 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int8(kExpected), true, false, ""},
		equalsTestCase{int16(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint8(kExpected), true, false, ""},
		equalsTestCase{uint16(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, false, false, ""},
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int8(kExpected + 1), false, false, ""},
		equalsTestCase{int16(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint8(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{65553, true, false, ""},
		equalsTestCase{65553.0, true, false, ""},
		equalsTestCase{65553 + 0i, true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int16(17), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(17), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},
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
		equalsTestCase{int64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{int64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), false, false, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float32(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{float64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), false, false, ""},
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
		equalsTestCase{17, true, false, ""},
		equalsTestCase{17.0, true, false, ""},
		equalsTestCase{17 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int8(kExpected), true, false, ""},
		equalsTestCase{int16(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint8(kExpected), true, false, ""},
		equalsTestCase{uint16(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{kExpected + 1, false, false, ""},
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int8(kExpected + 1), false, false, ""},
		equalsTestCase{int16(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint8(kExpected + 1), false, false, ""},
		equalsTestCase{uint16(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{4294967313.0, true, false, ""},
		equalsTestCase{4294967313 + 0i, true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric types.
		equalsTestCase{int(17), false, false, ""},
		equalsTestCase{int32(17), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(17), false, false, ""},
		equalsTestCase{uint32(17), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected + 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 1), false, false, ""},
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
		equalsTestCase{int64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{int64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{uint64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), false, false, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float32(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{float64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 2), false, false, ""},

		equalsTestCase{complex128(kTwoTo25 + 0), false, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), false, false, ""},
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
		equalsTestCase{int64(kTwoTo54 + 0), false, false, ""},
		equalsTestCase{int64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 2), false, false, ""},

		equalsTestCase{uint64(kTwoTo54 + 0), false, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{float64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 3), false, false, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), false, false, ""},
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
		equalsTestCase{ptr1, true, false, ""},
		equalsTestCase{ptr2, true, false, ""},
		equalsTestCase{uintptr(0), true, false, ""},
		equalsTestCase{uintptr(17), false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a uintptr"},
		equalsTestCase{bool(false), false, true, "which is not a uintptr"},
		equalsTestCase{int(0), false, true, "which is not a uintptr"},
		equalsTestCase{int8(0), false, true, "which is not a uintptr"},
		equalsTestCase{int16(0), false, true, "which is not a uintptr"},
		equalsTestCase{int32(0), false, true, "which is not a uintptr"},
		equalsTestCase{int64(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint8(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint16(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint32(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint64(0), false, true, "which is not a uintptr"},
		equalsTestCase{true, false, true, "which is not a uintptr"},
		equalsTestCase{[...]int{}, false, true, "which is not a uintptr"},
		equalsTestCase{make(chan int), false, true, "which is not a uintptr"},
		equalsTestCase{func() {}, false, true, "which is not a uintptr"},
		equalsTestCase{map[int]int{}, false, true, "which is not a uintptr"},
		equalsTestCase{&someInt, false, true, "which is not a uintptr"},
		equalsTestCase{[]int{}, false, true, "which is not a uintptr"},
		equalsTestCase{"taco", false, true, "which is not a uintptr"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a uintptr"},
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
		equalsTestCase{uintptr(17), true, false, ""},
		equalsTestCase{uintptr(16), false, false, ""},
		equalsTestCase{uintptr(0), false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a uintptr"},
		equalsTestCase{bool(false), false, true, "which is not a uintptr"},
		equalsTestCase{int(0), false, true, "which is not a uintptr"},
		equalsTestCase{int8(0), false, true, "which is not a uintptr"},
		equalsTestCase{int16(0), false, true, "which is not a uintptr"},
		equalsTestCase{int32(0), false, true, "which is not a uintptr"},
		equalsTestCase{int64(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint8(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint16(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint32(0), false, true, "which is not a uintptr"},
		equalsTestCase{uint64(0), false, true, "which is not a uintptr"},
		equalsTestCase{true, false, true, "which is not a uintptr"},
		equalsTestCase{[...]int{}, false, true, "which is not a uintptr"},
		equalsTestCase{make(chan int), false, true, "which is not a uintptr"},
		equalsTestCase{func() {}, false, true, "which is not a uintptr"},
		equalsTestCase{map[int]int{}, false, true, "which is not a uintptr"},
		equalsTestCase{&someInt, false, true, "which is not a uintptr"},
		equalsTestCase{[]int{}, false, true, "which is not a uintptr"},
		equalsTestCase{"taco", false, true, "which is not a uintptr"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a uintptr"},
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
		equalsTestCase{-32769.0, true, false, ""},
		equalsTestCase{-32769 + 0i, true, false, ""},
		equalsTestCase{int32(-32769), true, false, ""},
		equalsTestCase{int64(-32769), true, false, ""},
		equalsTestCase{float32(-32769), true, false, ""},
		equalsTestCase{float64(-32769), true, false, ""},
		equalsTestCase{complex64(-32769), true, false, ""},
		equalsTestCase{complex128(-32769), true, false, ""},
		equalsTestCase{interface{}(float32(-32769)), true, false, ""},
		equalsTestCase{interface{}(int64(-32769)), true, false, ""},

		// Values that would be -32769 in two's complement.
		equalsTestCase{uint64((1 << 64) - 32769), false, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(-32770), false, false, ""},
		equalsTestCase{float32(-32769.1), false, false, ""},
		equalsTestCase{float32(-32768.9), false, false, ""},
		equalsTestCase{float64(-32769.1), false, false, ""},
		equalsTestCase{float64(-32768.9), false, false, ""},
		equalsTestCase{complex128(-32768), false, false, ""},
		equalsTestCase{complex128(-32769 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{-32769.1, true, false, ""},
		equalsTestCase{-32769.1 + 0i, true, false, ""},
		equalsTestCase{float32(-32769.1), true, false, ""},
		equalsTestCase{float64(-32769.1), true, false, ""},
		equalsTestCase{complex64(-32769.1), true, false, ""},
		equalsTestCase{complex128(-32769.1), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int32(-32769), false, false, ""},
		equalsTestCase{int32(-32770), false, false, ""},
		equalsTestCase{int64(-32769), false, false, ""},
		equalsTestCase{int64(-32770), false, false, ""},
		equalsTestCase{float32(-32769.2), false, false, ""},
		equalsTestCase{float32(-32769.0), false, false, ""},
		equalsTestCase{float64(-32769.2), false, false, ""},
		equalsTestCase{complex128(-32769.1 + 2i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{0.0, true, false, ""},
		equalsTestCase{0 + 0i, true, false, ""},
		equalsTestCase{int(0), true, false, ""},
		equalsTestCase{int8(0), true, false, ""},
		equalsTestCase{int16(0), true, false, ""},
		equalsTestCase{int32(0), true, false, ""},
		equalsTestCase{int64(0), true, false, ""},
		equalsTestCase{uint(0), true, false, ""},
		equalsTestCase{uint8(0), true, false, ""},
		equalsTestCase{uint16(0), true, false, ""},
		equalsTestCase{uint32(0), true, false, ""},
		equalsTestCase{uint64(0), true, false, ""},
		equalsTestCase{float32(0), true, false, ""},
		equalsTestCase{float64(0), true, false, ""},
		equalsTestCase{complex64(0), true, false, ""},
		equalsTestCase{complex128(0), true, false, ""},
		equalsTestCase{interface{}(float32(0)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), false, false, ""},
		equalsTestCase{int64(-1), false, false, ""},
		equalsTestCase{float32(1), false, false, ""},
		equalsTestCase{float32(-1), false, false, ""},
		equalsTestCase{complex128(0 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{32769.0, true, false, ""},
		equalsTestCase{32769 + 0i, true, false, ""},
		equalsTestCase{int(32769), true, false, ""},
		equalsTestCase{int32(32769), true, false, ""},
		equalsTestCase{int64(32769), true, false, ""},
		equalsTestCase{uint(32769), true, false, ""},
		equalsTestCase{uint32(32769), true, false, ""},
		equalsTestCase{uint64(32769), true, false, ""},
		equalsTestCase{float32(32769), true, false, ""},
		equalsTestCase{float64(32769), true, false, ""},
		equalsTestCase{complex64(32769), true, false, ""},
		equalsTestCase{complex128(32769), true, false, ""},
		equalsTestCase{interface{}(float32(32769)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(32770), false, false, ""},
		equalsTestCase{uint64(32770), false, false, ""},
		equalsTestCase{float32(32769.1), false, false, ""},
		equalsTestCase{float32(32768.9), false, false, ""},
		equalsTestCase{float64(32769.1), false, false, ""},
		equalsTestCase{float64(32768.9), false, false, ""},
		equalsTestCase{complex128(32768), false, false, ""},
		equalsTestCase{complex128(32769 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{32769.1, true, false, ""},
		equalsTestCase{32769.1 + 0i, true, false, ""},
		equalsTestCase{float32(32769.1), true, false, ""},
		equalsTestCase{float64(32769.1), true, false, ""},
		equalsTestCase{complex64(32769.1), true, false, ""},
		equalsTestCase{complex128(32769.1), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int32(32769), false, false, ""},
		equalsTestCase{int32(32770), false, false, ""},
		equalsTestCase{uint64(32769), false, false, ""},
		equalsTestCase{uint64(32770), false, false, ""},
		equalsTestCase{float32(32769.2), false, false, ""},
		equalsTestCase{float32(32769.0), false, false, ""},
		equalsTestCase{float64(32769.2), false, false, ""},
		equalsTestCase{complex128(32769.1 + 2i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{uint64(0), false, false, ""},
		equalsTestCase{uint64(math.MaxUint64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{int64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{int64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{uint64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{uint64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 3), false, false, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float32(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex128(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex128(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 3), false, false, ""},
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
		equalsTestCase{-1125899906842624.0, true, false, ""},
		equalsTestCase{-1125899906842624.0 + 0i, true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},
		equalsTestCase{interface{}(float64(kExpected)), true, false, ""},

		// Values that would be kExpected in two's complement.
		equalsTestCase{uint64((1 << 64) + kExpected), false, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.5), false, false, ""},
		equalsTestCase{float64(kExpected + 0.5), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(-kTwoTo50), false, false, ""},
		equalsTestCase{int64(-kTwoTo50 - 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.25), false, false, ""},
		equalsTestCase{float64(kExpected + 0.25), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{0.0, true, false, ""},
		equalsTestCase{0 + 0i, true, false, ""},
		equalsTestCase{int(0), true, false, ""},
		equalsTestCase{int8(0), true, false, ""},
		equalsTestCase{int16(0), true, false, ""},
		equalsTestCase{int32(0), true, false, ""},
		equalsTestCase{int64(0), true, false, ""},
		equalsTestCase{uint(0), true, false, ""},
		equalsTestCase{uint8(0), true, false, ""},
		equalsTestCase{uint16(0), true, false, ""},
		equalsTestCase{uint32(0), true, false, ""},
		equalsTestCase{uint64(0), true, false, ""},
		equalsTestCase{float32(0), true, false, ""},
		equalsTestCase{float64(0), true, false, ""},
		equalsTestCase{complex64(0), true, false, ""},
		equalsTestCase{complex128(0), true, false, ""},
		equalsTestCase{interface{}(float32(0)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), false, false, ""},
		equalsTestCase{int64(-1), false, false, ""},
		equalsTestCase{float32(1), false, false, ""},
		equalsTestCase{float32(-1), false, false, ""},
		equalsTestCase{complex128(0 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1125899906842624.0, true, false, ""},
		equalsTestCase{1125899906842624.0 + 0i, true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},
		equalsTestCase{interface{}(float64(kExpected)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.5), false, false, ""},
		equalsTestCase{float64(kExpected + 0.5), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kTwoTo50), false, false, ""},
		equalsTestCase{int64(kTwoTo50 - 1), false, false, ""},
		equalsTestCase{float64(kExpected - 0.25), false, false, ""},
		equalsTestCase{float64(kExpected + 0.25), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{uint64(0), false, false, ""},
		equalsTestCase{uint64(math.MaxUint64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{int64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{int64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 3), false, false, ""},

		equalsTestCase{uint64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{uint64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{float64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 3), false, false, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), false, false, ""},
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
		equalsTestCase{-32769.0, true, false, ""},
		equalsTestCase{-32769.0 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},
		equalsTestCase{interface{}(float64(kExpected)), true, false, ""},

		// Values that would be kExpected in two's complement.
		equalsTestCase{uint32((1 << 32) + kExpected), false, false, ""},
		equalsTestCase{uint64((1 << 64) + kExpected), false, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.5), false, false, ""},
		equalsTestCase{float64(kExpected + 0.5), false, false, ""},
		equalsTestCase{complex64(kExpected - 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(-kTwoTo20), false, false, ""},
		equalsTestCase{int(-kTwoTo20 - 1), false, false, ""},
		equalsTestCase{int32(-kTwoTo20), false, false, ""},
		equalsTestCase{int32(-kTwoTo20 - 1), false, false, ""},
		equalsTestCase{int64(-kTwoTo20), false, false, ""},
		equalsTestCase{int64(-kTwoTo20 - 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.25), false, false, ""},
		equalsTestCase{float64(kExpected + 0.25), false, false, ""},
		equalsTestCase{complex64(kExpected - 0.75), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected - 0.75), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{0.0, true, false, ""},
		equalsTestCase{0 + 0i, true, false, ""},
		equalsTestCase{int(0), true, false, ""},
		equalsTestCase{int8(0), true, false, ""},
		equalsTestCase{int16(0), true, false, ""},
		equalsTestCase{int32(0), true, false, ""},
		equalsTestCase{int64(0), true, false, ""},
		equalsTestCase{uint(0), true, false, ""},
		equalsTestCase{uint8(0), true, false, ""},
		equalsTestCase{uint16(0), true, false, ""},
		equalsTestCase{uint32(0), true, false, ""},
		equalsTestCase{uint64(0), true, false, ""},
		equalsTestCase{float32(0), true, false, ""},
		equalsTestCase{float64(0), true, false, ""},
		equalsTestCase{complex64(0), true, false, ""},
		equalsTestCase{complex128(0), true, false, ""},
		equalsTestCase{interface{}(float32(0)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), false, false, ""},
		equalsTestCase{int64(-1), false, false, ""},
		equalsTestCase{float32(1), false, false, ""},
		equalsTestCase{float32(-1), false, false, ""},
		equalsTestCase{float64(1), false, false, ""},
		equalsTestCase{float64(-1), false, false, ""},
		equalsTestCase{complex64(0 + 2i), false, false, ""},
		equalsTestCase{complex128(0 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1048576.0, true, false, ""},
		equalsTestCase{1048576.0 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},
		equalsTestCase{interface{}(float64(kExpected)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.5), false, false, ""},
		equalsTestCase{float64(kExpected + 0.5), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kTwoTo20), false, false, ""},
		equalsTestCase{int64(kTwoTo20 - 1), false, false, ""},
		equalsTestCase{uint64(kTwoTo20), false, false, ""},
		equalsTestCase{uint64(kTwoTo20 - 1), false, false, ""},
		equalsTestCase{float32(kExpected - 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected - 0.25), false, false, ""},
		equalsTestCase{float64(kExpected + 0.25), false, false, ""},
		equalsTestCase{complex64(kExpected - 1), false, false, ""},
		equalsTestCase{complex64(kExpected - 1i), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected - 1i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{uint64(0), false, false, ""},
		equalsTestCase{uint64(math.MaxUint64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{int64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{int64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{int64(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{uint64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{uint64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{uint64(kTwoTo25 + 3), false, false, ""},

		// Single-precision floating point.
		equalsTestCase{float32(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float32(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float32(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex64(kTwoTo25 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{float64(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{float64(kTwoTo25 + 3), false, false, ""},

		equalsTestCase{complex128(kTwoTo25 - 2), false, false, ""},
		equalsTestCase{complex128(kTwoTo25 - 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 0), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 2), true, false, ""},
		equalsTestCase{complex128(kTwoTo25 + 3), false, false, ""},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kRealPart + kImagPart, true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kRealPart), false, false, ""},
		equalsTestCase{int8(kRealPart), false, false, ""},
		equalsTestCase{int16(kRealPart), false, false, ""},
		equalsTestCase{int32(kRealPart), false, false, ""},
		equalsTestCase{int64(kRealPart), false, false, ""},
		equalsTestCase{uint(kRealPart), false, false, ""},
		equalsTestCase{uint8(kRealPart), false, false, ""},
		equalsTestCase{uint16(kRealPart), false, false, ""},
		equalsTestCase{uint32(kRealPart), false, false, ""},
		equalsTestCase{uint64(kRealPart), false, false, ""},
		equalsTestCase{float32(kRealPart), false, false, ""},
		equalsTestCase{float64(kRealPart), false, false, ""},
		equalsTestCase{complex64(kRealPart), false, false, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5), false, false, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5i), false, false, ""},
		equalsTestCase{complex128(kRealPart), false, false, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5), false, false, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5i), false, false, ""},
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
		equalsTestCase{-32769.0, true, false, ""},
		equalsTestCase{-32769.0 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},
		equalsTestCase{interface{}(float64(kExpected)), true, false, ""},

		// Values that would be kExpected in two's complement.
		equalsTestCase{uint32((1 << 32) + kExpected), false, false, ""},
		equalsTestCase{uint64((1 << 64) + kExpected), false, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.5), false, false, ""},
		equalsTestCase{float64(kExpected + 0.5), false, false, ""},
		equalsTestCase{complex64(kExpected - 1), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(-kTwoTo20), false, false, ""},
		equalsTestCase{int(-kTwoTo20 - 1), false, false, ""},
		equalsTestCase{int32(-kTwoTo20), false, false, ""},
		equalsTestCase{int32(-kTwoTo20 - 1), false, false, ""},
		equalsTestCase{int64(-kTwoTo20), false, false, ""},
		equalsTestCase{int64(-kTwoTo20 - 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.25), false, false, ""},
		equalsTestCase{float64(kExpected + 0.25), false, false, ""},
		equalsTestCase{complex64(kExpected - 0.75), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected - 0.75), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex64(kExpected + 2i), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{0.0, true, false, ""},
		equalsTestCase{0 + 0i, true, false, ""},
		equalsTestCase{int(0), true, false, ""},
		equalsTestCase{int8(0), true, false, ""},
		equalsTestCase{int16(0), true, false, ""},
		equalsTestCase{int32(0), true, false, ""},
		equalsTestCase{int64(0), true, false, ""},
		equalsTestCase{uint(0), true, false, ""},
		equalsTestCase{uint8(0), true, false, ""},
		equalsTestCase{uint16(0), true, false, ""},
		equalsTestCase{uint32(0), true, false, ""},
		equalsTestCase{uint64(0), true, false, ""},
		equalsTestCase{float32(0), true, false, ""},
		equalsTestCase{float64(0), true, false, ""},
		equalsTestCase{complex64(0), true, false, ""},
		equalsTestCase{complex128(0), true, false, ""},
		equalsTestCase{interface{}(float32(0)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(1), false, false, ""},
		equalsTestCase{int64(-1), false, false, ""},
		equalsTestCase{float32(1), false, false, ""},
		equalsTestCase{float32(-1), false, false, ""},
		equalsTestCase{float64(1), false, false, ""},
		equalsTestCase{float64(-1), false, false, ""},
		equalsTestCase{complex64(0 + 2i), false, false, ""},
		equalsTestCase{complex128(0 + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{1048576.0, true, false, ""},
		equalsTestCase{1048576.0 + 0i, true, false, ""},
		equalsTestCase{int(kExpected), true, false, ""},
		equalsTestCase{int32(kExpected), true, false, ""},
		equalsTestCase{int64(kExpected), true, false, ""},
		equalsTestCase{uint(kExpected), true, false, ""},
		equalsTestCase{uint32(kExpected), true, false, ""},
		equalsTestCase{uint64(kExpected), true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},
		equalsTestCase{interface{}(float64(kExpected)), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kExpected + 1), false, false, ""},
		equalsTestCase{int32(kExpected + 1), false, false, ""},
		equalsTestCase{int64(kExpected + 1), false, false, ""},
		equalsTestCase{uint(kExpected + 1), false, false, ""},
		equalsTestCase{uint32(kExpected + 1), false, false, ""},
		equalsTestCase{uint64(kExpected + 1), false, false, ""},
		equalsTestCase{float32(kExpected - (1 << 30)), false, false, ""},
		equalsTestCase{float32(kExpected + (1 << 30)), false, false, ""},
		equalsTestCase{float64(kExpected - 0.5), false, false, ""},
		equalsTestCase{float64(kExpected + 0.5), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},

		// Non-numeric types.
		equalsTestCase{uintptr(0), false, true, "which is not numeric"},
		equalsTestCase{true, false, true, "which is not numeric"},
		equalsTestCase{[...]int{}, false, true, "which is not numeric"},
		equalsTestCase{make(chan int), false, true, "which is not numeric"},
		equalsTestCase{func() {}, false, true, "which is not numeric"},
		equalsTestCase{map[int]int{}, false, true, "which is not numeric"},
		equalsTestCase{&someInt, false, true, "which is not numeric"},
		equalsTestCase{[]int{}, false, true, "which is not numeric"},
		equalsTestCase{"taco", false, true, "which is not numeric"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not numeric"},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int64(kTwoTo20), false, false, ""},
		equalsTestCase{int64(kTwoTo20 - 1), false, false, ""},
		equalsTestCase{uint64(kTwoTo20), false, false, ""},
		equalsTestCase{uint64(kTwoTo20 - 1), false, false, ""},
		equalsTestCase{float32(kExpected - 1), false, false, ""},
		equalsTestCase{float32(kExpected + 1), false, false, ""},
		equalsTestCase{float64(kExpected - 0.25), false, false, ""},
		equalsTestCase{float64(kExpected + 0.25), false, false, ""},
		equalsTestCase{complex64(kExpected - 1), false, false, ""},
		equalsTestCase{complex64(kExpected - 1i), false, false, ""},
		equalsTestCase{complex128(kExpected - 1), false, false, ""},
		equalsTestCase{complex128(kExpected - 1i), false, false, ""},
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
		equalsTestCase{kExpected + 0i, true, false, ""},
		equalsTestCase{float32(kExpected), true, false, ""},
		equalsTestCase{float64(kExpected), true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{castedInt, false, false, ""},
		equalsTestCase{int64(0), false, false, ""},
		equalsTestCase{int64(math.MinInt64), false, false, ""},
		equalsTestCase{int64(math.MaxInt64), false, false, ""},
		equalsTestCase{uint64(0), false, false, ""},
		equalsTestCase{uint64(math.MaxUint64), false, false, ""},
		equalsTestCase{float32(kExpected / 2), false, false, ""},
		equalsTestCase{float64(kExpected / 2), false, false, ""},
		equalsTestCase{complex128(kExpected + 2i), false, false, ""},
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
		equalsTestCase{int64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{int64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{int64(kTwoTo54 + 3), false, false, ""},

		equalsTestCase{uint64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{uint64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{uint64(kTwoTo54 + 3), false, false, ""},

		// Double-precision floating point.
		equalsTestCase{float64(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{float64(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{float64(kTwoTo54 + 3), false, false, ""},

		equalsTestCase{complex128(kTwoTo54 - 2), false, false, ""},
		equalsTestCase{complex128(kTwoTo54 - 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 0), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 1), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 2), true, false, ""},
		equalsTestCase{complex128(kTwoTo54 + 3), false, false, ""},
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
		equalsTestCase{kExpected, true, false, ""},
		equalsTestCase{kRealPart + kImagPart, true, false, ""},
		equalsTestCase{complex64(kExpected), true, false, ""},
		equalsTestCase{complex128(kExpected), true, false, ""},

		// Non-equal values of numeric type.
		equalsTestCase{int(kRealPart), false, false, ""},
		equalsTestCase{int8(kRealPart), false, false, ""},
		equalsTestCase{int16(kRealPart), false, false, ""},
		equalsTestCase{int32(kRealPart), false, false, ""},
		equalsTestCase{int64(kRealPart), false, false, ""},
		equalsTestCase{uint(kRealPart), false, false, ""},
		equalsTestCase{uint8(kRealPart), false, false, ""},
		equalsTestCase{uint16(kRealPart), false, false, ""},
		equalsTestCase{uint32(kRealPart), false, false, ""},
		equalsTestCase{uint64(kRealPart), false, false, ""},
		equalsTestCase{float32(kRealPart), false, false, ""},
		equalsTestCase{float64(kRealPart), false, false, ""},
		equalsTestCase{complex64(kRealPart), false, false, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5), false, false, ""},
		equalsTestCase{complex64(kRealPart + kImagPart + 0.5i), false, false, ""},
		equalsTestCase{complex128(kRealPart), false, false, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5), false, false, ""},
		equalsTestCase{complex128(kRealPart + kImagPart + 0.5i), false, false, ""},
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
		equalsTestCase{nilChan1, true, false, ""},
		equalsTestCase{nilChan2, true, false, ""},
		equalsTestCase{nonNilChan1, false, false, ""},

		// uint channels
		equalsTestCase{nilChan3, false, true, "which is not a chan int"},
		equalsTestCase{nonNilChan2, false, true, "which is not a chan int"},

		// Other types.
		equalsTestCase{0, false, true, "which is not a chan int"},
		equalsTestCase{bool(false), false, true, "which is not a chan int"},
		equalsTestCase{int(0), false, true, "which is not a chan int"},
		equalsTestCase{int8(0), false, true, "which is not a chan int"},
		equalsTestCase{int16(0), false, true, "which is not a chan int"},
		equalsTestCase{int32(0), false, true, "which is not a chan int"},
		equalsTestCase{int64(0), false, true, "which is not a chan int"},
		equalsTestCase{uint(0), false, true, "which is not a chan int"},
		equalsTestCase{uint8(0), false, true, "which is not a chan int"},
		equalsTestCase{uint16(0), false, true, "which is not a chan int"},
		equalsTestCase{uint32(0), false, true, "which is not a chan int"},
		equalsTestCase{uint64(0), false, true, "which is not a chan int"},
		equalsTestCase{true, false, true, "which is not a chan int"},
		equalsTestCase{[...]int{}, false, true, "which is not a chan int"},
		equalsTestCase{func() {}, false, true, "which is not a chan int"},
		equalsTestCase{map[int]int{}, false, true, "which is not a chan int"},
		equalsTestCase{&someInt, false, true, "which is not a chan int"},
		equalsTestCase{[]int{}, false, true, "which is not a chan int"},
		equalsTestCase{"taco", false, true, "which is not a chan int"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a chan int"},
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
		equalsTestCase{nonNilChan1, true, false, ""},
		equalsTestCase{nonNilChan2, false, false, ""},
		equalsTestCase{nilChan1, false, false, ""},

		// uint channels
		equalsTestCase{nilChan2, false, true, "which is not a chan int"},
		equalsTestCase{nonNilChan3, false, true, "which is not a chan int"},

		// Other types.
		equalsTestCase{0, false, true, "which is not a chan int"},
		equalsTestCase{bool(false), false, true, "which is not a chan int"},
		equalsTestCase{int(0), false, true, "which is not a chan int"},
		equalsTestCase{int8(0), false, true, "which is not a chan int"},
		equalsTestCase{int16(0), false, true, "which is not a chan int"},
		equalsTestCase{int32(0), false, true, "which is not a chan int"},
		equalsTestCase{int64(0), false, true, "which is not a chan int"},
		equalsTestCase{uint(0), false, true, "which is not a chan int"},
		equalsTestCase{uint8(0), false, true, "which is not a chan int"},
		equalsTestCase{uint16(0), false, true, "which is not a chan int"},
		equalsTestCase{uint32(0), false, true, "which is not a chan int"},
		equalsTestCase{uint64(0), false, true, "which is not a chan int"},
		equalsTestCase{true, false, true, "which is not a chan int"},
		equalsTestCase{[...]int{}, false, true, "which is not a chan int"},
		equalsTestCase{func() {}, false, true, "which is not a chan int"},
		equalsTestCase{map[int]int{}, false, true, "which is not a chan int"},
		equalsTestCase{&someInt, false, true, "which is not a chan int"},
		equalsTestCase{[]int{}, false, true, "which is not a chan int"},
		equalsTestCase{"taco", false, true, "which is not a chan int"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a chan int"},
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
		equalsTestCase{chan1, true, false, ""},
		equalsTestCase{chan2, false, true, "which is not a chan<- int"},
		equalsTestCase{chan3, false, true, "which is not a chan<- int"},
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
		equalsTestCase{func1, true, false, ""},
		equalsTestCase{func2, false, false, ""},
		equalsTestCase{func3, false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a function"},
		equalsTestCase{bool(false), false, true, "which is not a function"},
		equalsTestCase{int(0), false, true, "which is not a function"},
		equalsTestCase{int8(0), false, true, "which is not a function"},
		equalsTestCase{int16(0), false, true, "which is not a function"},
		equalsTestCase{int32(0), false, true, "which is not a function"},
		equalsTestCase{int64(0), false, true, "which is not a function"},
		equalsTestCase{uint(0), false, true, "which is not a function"},
		equalsTestCase{uint8(0), false, true, "which is not a function"},
		equalsTestCase{uint16(0), false, true, "which is not a function"},
		equalsTestCase{uint32(0), false, true, "which is not a function"},
		equalsTestCase{uint64(0), false, true, "which is not a function"},
		equalsTestCase{true, false, true, "which is not a function"},
		equalsTestCase{[...]int{}, false, true, "which is not a function"},
		equalsTestCase{map[int]int{}, false, true, "which is not a function"},
		equalsTestCase{&someInt, false, true, "which is not a function"},
		equalsTestCase{[]int{}, false, true, "which is not a function"},
		equalsTestCase{"taco", false, true, "which is not a function"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a function"},
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
		equalsTestCase{nilMap1, true, false, ""},
		equalsTestCase{nilMap2, true, false, ""},
		equalsTestCase{nilMap3, true, false, ""},
		equalsTestCase{nonNilMap1, false, false, ""},
		equalsTestCase{nonNilMap2, false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a map"},
		equalsTestCase{bool(false), false, true, "which is not a map"},
		equalsTestCase{int(0), false, true, "which is not a map"},
		equalsTestCase{int8(0), false, true, "which is not a map"},
		equalsTestCase{int16(0), false, true, "which is not a map"},
		equalsTestCase{int32(0), false, true, "which is not a map"},
		equalsTestCase{int64(0), false, true, "which is not a map"},
		equalsTestCase{uint(0), false, true, "which is not a map"},
		equalsTestCase{uint8(0), false, true, "which is not a map"},
		equalsTestCase{uint16(0), false, true, "which is not a map"},
		equalsTestCase{uint32(0), false, true, "which is not a map"},
		equalsTestCase{uint64(0), false, true, "which is not a map"},
		equalsTestCase{true, false, true, "which is not a map"},
		equalsTestCase{[...]int{}, false, true, "which is not a map"},
		equalsTestCase{func() {}, false, true, "which is not a map"},
		equalsTestCase{&someInt, false, true, "which is not a map"},
		equalsTestCase{[]int{}, false, true, "which is not a map"},
		equalsTestCase{"taco", false, true, "which is not a map"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a map"},
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
		equalsTestCase{nonNilMap1, true, false, ""},
		equalsTestCase{nonNilMap2, false, false, ""},
		equalsTestCase{nonNilMap3, false, false, ""},
		equalsTestCase{nilMap1, false, false, ""},
		equalsTestCase{nilMap2, false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a map"},
		equalsTestCase{bool(false), false, true, "which is not a map"},
		equalsTestCase{int(0), false, true, "which is not a map"},
		equalsTestCase{int8(0), false, true, "which is not a map"},
		equalsTestCase{int16(0), false, true, "which is not a map"},
		equalsTestCase{int32(0), false, true, "which is not a map"},
		equalsTestCase{int64(0), false, true, "which is not a map"},
		equalsTestCase{uint(0), false, true, "which is not a map"},
		equalsTestCase{uint8(0), false, true, "which is not a map"},
		equalsTestCase{uint16(0), false, true, "which is not a map"},
		equalsTestCase{uint32(0), false, true, "which is not a map"},
		equalsTestCase{uint64(0), false, true, "which is not a map"},
		equalsTestCase{true, false, true, "which is not a map"},
		equalsTestCase{[...]int{}, false, true, "which is not a map"},
		equalsTestCase{func() {}, false, true, "which is not a map"},
		equalsTestCase{&someInt, false, true, "which is not a map"},
		equalsTestCase{[]int{}, false, true, "which is not a map"},
		equalsTestCase{"taco", false, true, "which is not a map"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a map"},
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
		equalsTestCase{nilInt1, true, false, ""},
		equalsTestCase{nilInt2, true, false, ""},
		equalsTestCase{nonNilInt, false, false, ""},

		// Incorrect type.
		equalsTestCase{nilUint, false, true, "which is not a *int"},
		equalsTestCase{nonNilUint, false, true, "which is not a *int"},

		// Other types.
		equalsTestCase{0, false, true, "which is not a *int"},
		equalsTestCase{bool(false), false, true, "which is not a *int"},
		equalsTestCase{int(0), false, true, "which is not a *int"},
		equalsTestCase{int8(0), false, true, "which is not a *int"},
		equalsTestCase{int16(0), false, true, "which is not a *int"},
		equalsTestCase{int32(0), false, true, "which is not a *int"},
		equalsTestCase{int64(0), false, true, "which is not a *int"},
		equalsTestCase{uint(0), false, true, "which is not a *int"},
		equalsTestCase{uint8(0), false, true, "which is not a *int"},
		equalsTestCase{uint16(0), false, true, "which is not a *int"},
		equalsTestCase{uint32(0), false, true, "which is not a *int"},
		equalsTestCase{uint64(0), false, true, "which is not a *int"},
		equalsTestCase{true, false, true, "which is not a *int"},
		equalsTestCase{[...]int{}, false, true, "which is not a *int"},
		equalsTestCase{func() {}, false, true, "which is not a *int"},
		equalsTestCase{map[int]int{}, false, true, "which is not a *int"},
		equalsTestCase{[]int{}, false, true, "which is not a *int"},
		equalsTestCase{"taco", false, true, "which is not a *int"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a *int"},
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
		equalsTestCase{nonNilInt1, true, false, ""},
		equalsTestCase{nonNilInt2, false, false, ""},
		equalsTestCase{nilInt, false, false, ""},

		// Incorrect type.
		equalsTestCase{nilUint, false, true, "which is not a *int"},
		equalsTestCase{nonNilUint, false, true, "which is not a *int"},

		// Other types.
		equalsTestCase{0, false, true, "which is not a *int"},
		equalsTestCase{bool(false), false, true, "which is not a *int"},
		equalsTestCase{int(0), false, true, "which is not a *int"},
		equalsTestCase{int8(0), false, true, "which is not a *int"},
		equalsTestCase{int16(0), false, true, "which is not a *int"},
		equalsTestCase{int32(0), false, true, "which is not a *int"},
		equalsTestCase{int64(0), false, true, "which is not a *int"},
		equalsTestCase{uint(0), false, true, "which is not a *int"},
		equalsTestCase{uint8(0), false, true, "which is not a *int"},
		equalsTestCase{uint16(0), false, true, "which is not a *int"},
		equalsTestCase{uint32(0), false, true, "which is not a *int"},
		equalsTestCase{uint64(0), false, true, "which is not a *int"},
		equalsTestCase{true, false, true, "which is not a *int"},
		equalsTestCase{[...]int{}, false, true, "which is not a *int"},
		equalsTestCase{func() {}, false, true, "which is not a *int"},
		equalsTestCase{map[int]int{}, false, true, "which is not a *int"},
		equalsTestCase{[]int{}, false, true, "which is not a *int"},
		equalsTestCase{"taco", false, true, "which is not a *int"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a *int"},
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
		equalsTestCase{nilInt1, true, false, ""},
		equalsTestCase{nilInt2, true, false, ""},
		equalsTestCase{nonNilInt, false, false, ""},

		// Incorrect type.
		equalsTestCase{nilUint, false, true, "which is not a []int"},
		equalsTestCase{nonNilUint, false, true, "which is not a []int"},

		// Other types.
		equalsTestCase{0, false, true, "which is not a []int"},
		equalsTestCase{bool(false), false, true, "which is not a []int"},
		equalsTestCase{int(0), false, true, "which is not a []int"},
		equalsTestCase{int8(0), false, true, "which is not a []int"},
		equalsTestCase{int16(0), false, true, "which is not a []int"},
		equalsTestCase{int32(0), false, true, "which is not a []int"},
		equalsTestCase{int64(0), false, true, "which is not a []int"},
		equalsTestCase{uint(0), false, true, "which is not a []int"},
		equalsTestCase{uint8(0), false, true, "which is not a []int"},
		equalsTestCase{uint16(0), false, true, "which is not a []int"},
		equalsTestCase{uint32(0), false, true, "which is not a []int"},
		equalsTestCase{uint64(0), false, true, "which is not a []int"},
		equalsTestCase{true, false, true, "which is not a []int"},
		equalsTestCase{[...]int{}, false, true, "which is not a []int"},
		equalsTestCase{func() {}, false, true, "which is not a []int"},
		equalsTestCase{map[int]int{}, false, true, "which is not a []int"},
		equalsTestCase{"taco", false, true, "which is not a []int"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a []int"},
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
		equalsTestCase{"taco1", true, false, ""},
		equalsTestCase{"taco" + "1", true, false, ""},
		equalsTestCase{expected, true, false, ""},

		equalsTestCase{"", false, false, ""},
		equalsTestCase{"taco", false, false, ""},
		equalsTestCase{"taco1\x00", false, false, ""},
		equalsTestCase{"taco2", false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a string"},
		equalsTestCase{bool(false), false, true, "which is not a string"},
		equalsTestCase{int(0), false, true, "which is not a string"},
		equalsTestCase{int8(0), false, true, "which is not a string"},
		equalsTestCase{int16(0), false, true, "which is not a string"},
		equalsTestCase{int32(0), false, true, "which is not a string"},
		equalsTestCase{int64(0), false, true, "which is not a string"},
		equalsTestCase{uint(0), false, true, "which is not a string"},
		equalsTestCase{uint8(0), false, true, "which is not a string"},
		equalsTestCase{uint16(0), false, true, "which is not a string"},
		equalsTestCase{uint32(0), false, true, "which is not a string"},
		equalsTestCase{uint64(0), false, true, "which is not a string"},
		equalsTestCase{true, false, true, "which is not a string"},
		equalsTestCase{[...]int{}, false, true, "which is not a string"},
		equalsTestCase{func() {}, false, true, "which is not a string"},
		equalsTestCase{map[int]int{}, false, true, "which is not a string"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a string"},
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
		equalsTestCase{nilPtr1, true, false, ""},
		equalsTestCase{nilPtr2, true, false, ""},
		equalsTestCase{nonNilPtr, false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{bool(false), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int8(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int16(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int32(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int64(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint8(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint16(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint32(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint64(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uintptr(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{true, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{[...]int{}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{make(chan int), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{func() {}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{map[int]int{}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{&someInt, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{[]int{}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{"taco", false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a unsafe.Pointer"},
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
		equalsTestCase{nonNilPtr1, true, false, ""},
		equalsTestCase{nonNilPtr2, false, false, ""},
		equalsTestCase{nilPtr, false, false, ""},

		// Other types.
		equalsTestCase{0, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{bool(false), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int8(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int16(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int32(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{int64(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint8(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint16(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint32(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uint64(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{uintptr(0), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{true, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{[...]int{}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{make(chan int), false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{func() {}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{map[int]int{}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{&someInt, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{[]int{}, false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{"taco", false, true, "which is not a unsafe.Pointer"},
		equalsTestCase{equalsTestCase{}, false, true, "which is not a unsafe.Pointer"},
	}

	checkTestCases(t, matcher, cases)
}
