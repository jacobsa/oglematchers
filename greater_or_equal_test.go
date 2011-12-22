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
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"math"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type GreaterOrEqualTest struct {
}

func init() { RegisterTestSuite(&GreaterOrEqualTest{}) }

type geTestCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func (t *GreaterOrEqualTest) checkTestCases(matcher Matcher, cases []geTestCase) {
	for i, c := range cases {
		result, err := matcher.Matches(c.candidate)

		ExpectThat(
			result,
			Equals(c.expectedResult),
			"Case %d (candidate %v)",
			i,
			c.candidate)

		errorMatcher := Error(Equals(c.expectedError))
		if c.expectedError == "" {
			errorMatcher = Equals(nil)
		}

		ExpectThat(
			err,
			errorMatcher,
			"Case %d (candidate %v)",
			i,
			c.candidate)
	}
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func (t *GreaterOrEqualTest) IntegerCandidateBadTypes() {
	matcher := GreaterOrEqual(int(-150))

	cases := []geTestCase{
		geTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{&geTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{geTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) FloatCandidateBadTypes() {
	matcher := GreaterOrEqual(float32(-150))

	cases := []geTestCase{
		geTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{&geTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{geTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) StringCandidateBadTypes() {
	matcher := GreaterOrEqual("17")

	cases := []geTestCase{
		geTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{int(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{int8(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{int16(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{int32(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{int64(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uint(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uint8(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uint16(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uint32(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uint64(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{float32(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{float64(0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{&geTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		geTestCase{geTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) BadArgument() {
	panicked := false

	defer func() {
		ExpectThat(panicked, Equals(true))
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	GreaterOrEqual(complex128(0))
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func (t *GreaterOrEqualTest) NegativeIntegerLiteral() {
	matcher := GreaterOrEqual(-150)
	desc := matcher.Description()
	expectedDesc := "greater than or equal to -150"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-(1 << 30), MATCH_TRUE, ""},
		geTestCase{-151, MATCH_TRUE, ""},
		geTestCase{-150, MATCH_FALSE, ""},
		geTestCase{0, MATCH_FALSE, ""},
		geTestCase{17, MATCH_FALSE, ""},

		geTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int(-151), MATCH_TRUE, ""},
		geTestCase{int(-150), MATCH_FALSE, ""},
		geTestCase{int(0), MATCH_FALSE, ""},
		geTestCase{int(17), MATCH_FALSE, ""},

		geTestCase{int8(-127), MATCH_FALSE, ""},
		geTestCase{int8(0), MATCH_FALSE, ""},
		geTestCase{int8(17), MATCH_FALSE, ""},

		geTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		geTestCase{int16(-151), MATCH_TRUE, ""},
		geTestCase{int16(-150), MATCH_FALSE, ""},
		geTestCase{int16(0), MATCH_FALSE, ""},
		geTestCase{int16(17), MATCH_FALSE, ""},

		geTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int32(-151), MATCH_TRUE, ""},
		geTestCase{int32(-150), MATCH_FALSE, ""},
		geTestCase{int32(0), MATCH_FALSE, ""},
		geTestCase{int32(17), MATCH_FALSE, ""},

		geTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int64(-151), MATCH_TRUE, ""},
		geTestCase{int64(-150), MATCH_FALSE, ""},
		geTestCase{int64(0), MATCH_FALSE, ""},
		geTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		geTestCase{uint(0), MATCH_FALSE, ""},
		geTestCase{uint(17), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_FALSE, ""},
		geTestCase{uint8(17), MATCH_FALSE, ""},
		geTestCase{uint8(253), MATCH_FALSE, ""},

		geTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		geTestCase{uint16(0), MATCH_FALSE, ""},
		geTestCase{uint16(17), MATCH_FALSE, ""},

		geTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		geTestCase{uint32(0), MATCH_FALSE, ""},
		geTestCase{uint32(17), MATCH_FALSE, ""},

		geTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		geTestCase{uint64(0), MATCH_FALSE, ""},
		geTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{float32(-151), MATCH_TRUE, ""},
		geTestCase{float32(-150.1), MATCH_TRUE, ""},
		geTestCase{float32(-150), MATCH_FALSE, ""},
		geTestCase{float32(-149.9), MATCH_FALSE, ""},
		geTestCase{float32(0), MATCH_FALSE, ""},
		geTestCase{float32(17), MATCH_FALSE, ""},
		geTestCase{float32(160), MATCH_FALSE, ""},

		geTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{float64(-151), MATCH_TRUE, ""},
		geTestCase{float64(-150.1), MATCH_TRUE, ""},
		geTestCase{float64(-150), MATCH_FALSE, ""},
		geTestCase{float64(-149.9), MATCH_FALSE, ""},
		geTestCase{float64(0), MATCH_FALSE, ""},
		geTestCase{float64(17), MATCH_FALSE, ""},
		geTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) ZeroIntegerLiteral() {
	matcher := GreaterOrEqual(0)
	desc := matcher.Description()
	expectedDesc := "greater than or equal to 0"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-(1 << 30), MATCH_TRUE, ""},
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{0, MATCH_FALSE, ""},
		geTestCase{1, MATCH_FALSE, ""},
		geTestCase{17, MATCH_FALSE, ""},
		geTestCase{(1 << 30), MATCH_FALSE, ""},

		geTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(0), MATCH_FALSE, ""},
		geTestCase{int(1), MATCH_FALSE, ""},
		geTestCase{int(17), MATCH_FALSE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(0), MATCH_FALSE, ""},
		geTestCase{int8(1), MATCH_FALSE, ""},

		geTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(0), MATCH_FALSE, ""},
		geTestCase{int16(1), MATCH_FALSE, ""},
		geTestCase{int16(17), MATCH_FALSE, ""},

		geTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(0), MATCH_FALSE, ""},
		geTestCase{int32(1), MATCH_FALSE, ""},
		geTestCase{int32(17), MATCH_FALSE, ""},

		geTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(0), MATCH_FALSE, ""},
		geTestCase{int64(1), MATCH_FALSE, ""},
		geTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint((1 << 32) - 1), MATCH_FALSE, ""},
		geTestCase{uint(0), MATCH_FALSE, ""},
		geTestCase{uint(17), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_FALSE, ""},
		geTestCase{uint8(17), MATCH_FALSE, ""},
		geTestCase{uint8(253), MATCH_FALSE, ""},

		geTestCase{uint16((1 << 16) - 1), MATCH_FALSE, ""},
		geTestCase{uint16(0), MATCH_FALSE, ""},
		geTestCase{uint16(17), MATCH_FALSE, ""},

		geTestCase{uint32((1 << 32) - 1), MATCH_FALSE, ""},
		geTestCase{uint32(0), MATCH_FALSE, ""},
		geTestCase{uint32(17), MATCH_FALSE, ""},

		geTestCase{uint64((1 << 64) - 1), MATCH_FALSE, ""},
		geTestCase{uint64(0), MATCH_FALSE, ""},
		geTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{float32(-1), MATCH_TRUE, ""},
		geTestCase{float32(-0.1), MATCH_TRUE, ""},
		geTestCase{float32(-0.0), MATCH_FALSE, ""},
		geTestCase{float32(0), MATCH_FALSE, ""},
		geTestCase{float32(0.1), MATCH_FALSE, ""},
		geTestCase{float32(17), MATCH_FALSE, ""},
		geTestCase{float32(160), MATCH_FALSE, ""},

		geTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(-0.1), MATCH_TRUE, ""},
		geTestCase{float64(-0), MATCH_FALSE, ""},
		geTestCase{float64(0), MATCH_FALSE, ""},
		geTestCase{float64(17), MATCH_FALSE, ""},
		geTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) PositiveIntegerLiteral() {
	matcher := GreaterOrEqual(150)
	desc := matcher.Description()
	expectedDesc := "greater than or equal to 150"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{149, MATCH_TRUE, ""},
		geTestCase{150, MATCH_FALSE, ""},
		geTestCase{151, MATCH_FALSE, ""},

		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(149), MATCH_TRUE, ""},
		geTestCase{int(150), MATCH_FALSE, ""},
		geTestCase{int(151), MATCH_FALSE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(0), MATCH_TRUE, ""},
		geTestCase{int8(17), MATCH_TRUE, ""},
		geTestCase{int8(127), MATCH_TRUE, ""},

		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(149), MATCH_TRUE, ""},
		geTestCase{int16(150), MATCH_FALSE, ""},
		geTestCase{int16(151), MATCH_FALSE, ""},

		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(149), MATCH_TRUE, ""},
		geTestCase{int32(150), MATCH_FALSE, ""},
		geTestCase{int32(151), MATCH_FALSE, ""},

		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(149), MATCH_TRUE, ""},
		geTestCase{int64(150), MATCH_FALSE, ""},
		geTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint(0), MATCH_TRUE, ""},
		geTestCase{uint(149), MATCH_TRUE, ""},
		geTestCase{uint(150), MATCH_FALSE, ""},
		geTestCase{uint(151), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_TRUE, ""},
		geTestCase{uint8(127), MATCH_TRUE, ""},

		geTestCase{uint16(0), MATCH_TRUE, ""},
		geTestCase{uint16(149), MATCH_TRUE, ""},
		geTestCase{uint16(150), MATCH_FALSE, ""},
		geTestCase{uint16(151), MATCH_FALSE, ""},

		geTestCase{uint32(0), MATCH_TRUE, ""},
		geTestCase{uint32(149), MATCH_TRUE, ""},
		geTestCase{uint32(150), MATCH_FALSE, ""},
		geTestCase{uint32(151), MATCH_FALSE, ""},

		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(149), MATCH_TRUE, ""},
		geTestCase{uint64(150), MATCH_FALSE, ""},
		geTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-1), MATCH_TRUE, ""},
		geTestCase{float32(149), MATCH_TRUE, ""},
		geTestCase{float32(149.9), MATCH_TRUE, ""},
		geTestCase{float32(150), MATCH_FALSE, ""},
		geTestCase{float32(150.1), MATCH_FALSE, ""},
		geTestCase{float32(151), MATCH_FALSE, ""},

		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(149), MATCH_TRUE, ""},
		geTestCase{float64(149.9), MATCH_TRUE, ""},
		geTestCase{float64(150), MATCH_FALSE, ""},
		geTestCase{float64(150.1), MATCH_FALSE, ""},
		geTestCase{float64(151), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// Float literals
////////////////////////////////////////////////////////////

func (t *GreaterOrEqualTest) NegativeFloatLiteral() {
	matcher := GreaterOrEqual(-150.1)
	desc := matcher.Description()
	expectedDesc := "greater than or equal to -150.1"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-(1 << 30), MATCH_TRUE, ""},
		geTestCase{-151, MATCH_TRUE, ""},
		geTestCase{-150, MATCH_FALSE, ""},
		geTestCase{0, MATCH_FALSE, ""},
		geTestCase{17, MATCH_FALSE, ""},

		geTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int(-151), MATCH_TRUE, ""},
		geTestCase{int(-150), MATCH_FALSE, ""},
		geTestCase{int(0), MATCH_FALSE, ""},
		geTestCase{int(17), MATCH_FALSE, ""},

		geTestCase{int8(-127), MATCH_FALSE, ""},
		geTestCase{int8(0), MATCH_FALSE, ""},
		geTestCase{int8(17), MATCH_FALSE, ""},

		geTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		geTestCase{int16(-151), MATCH_TRUE, ""},
		geTestCase{int16(-150), MATCH_FALSE, ""},
		geTestCase{int16(0), MATCH_FALSE, ""},
		geTestCase{int16(17), MATCH_FALSE, ""},

		geTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int32(-151), MATCH_TRUE, ""},
		geTestCase{int32(-150), MATCH_FALSE, ""},
		geTestCase{int32(0), MATCH_FALSE, ""},
		geTestCase{int32(17), MATCH_FALSE, ""},

		geTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{int64(-151), MATCH_TRUE, ""},
		geTestCase{int64(-150), MATCH_FALSE, ""},
		geTestCase{int64(0), MATCH_FALSE, ""},
		geTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		geTestCase{uint(0), MATCH_FALSE, ""},
		geTestCase{uint(17), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_FALSE, ""},
		geTestCase{uint8(17), MATCH_FALSE, ""},
		geTestCase{uint8(253), MATCH_FALSE, ""},

		geTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		geTestCase{uint16(0), MATCH_FALSE, ""},
		geTestCase{uint16(17), MATCH_FALSE, ""},

		geTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		geTestCase{uint32(0), MATCH_FALSE, ""},
		geTestCase{uint32(17), MATCH_FALSE, ""},

		geTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		geTestCase{uint64(0), MATCH_FALSE, ""},
		geTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{float32(-151), MATCH_TRUE, ""},
		geTestCase{float32(-150.2), MATCH_TRUE, ""},
		geTestCase{float32(-150.1), MATCH_FALSE, ""},
		geTestCase{float32(-150), MATCH_FALSE, ""},
		geTestCase{float32(0), MATCH_FALSE, ""},
		geTestCase{float32(17), MATCH_FALSE, ""},
		geTestCase{float32(160), MATCH_FALSE, ""},

		geTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		geTestCase{float64(-151), MATCH_TRUE, ""},
		geTestCase{float64(-150.2), MATCH_TRUE, ""},
		geTestCase{float64(-150.1), MATCH_FALSE, ""},
		geTestCase{float64(-150), MATCH_FALSE, ""},
		geTestCase{float64(0), MATCH_FALSE, ""},
		geTestCase{float64(17), MATCH_FALSE, ""},
		geTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) PositiveFloatLiteral() {
	matcher := GreaterOrEqual(149.9)
	desc := matcher.Description()
	expectedDesc := "greater than or equal to 149.9"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{149, MATCH_TRUE, ""},
		geTestCase{150, MATCH_FALSE, ""},
		geTestCase{151, MATCH_FALSE, ""},

		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(149), MATCH_TRUE, ""},
		geTestCase{int(150), MATCH_FALSE, ""},
		geTestCase{int(151), MATCH_FALSE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(0), MATCH_TRUE, ""},
		geTestCase{int8(17), MATCH_TRUE, ""},
		geTestCase{int8(127), MATCH_TRUE, ""},

		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(149), MATCH_TRUE, ""},
		geTestCase{int16(150), MATCH_FALSE, ""},
		geTestCase{int16(151), MATCH_FALSE, ""},

		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(149), MATCH_TRUE, ""},
		geTestCase{int32(150), MATCH_FALSE, ""},
		geTestCase{int32(151), MATCH_FALSE, ""},

		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(149), MATCH_TRUE, ""},
		geTestCase{int64(150), MATCH_FALSE, ""},
		geTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint(0), MATCH_TRUE, ""},
		geTestCase{uint(149), MATCH_TRUE, ""},
		geTestCase{uint(150), MATCH_FALSE, ""},
		geTestCase{uint(151), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_TRUE, ""},
		geTestCase{uint8(127), MATCH_TRUE, ""},

		geTestCase{uint16(0), MATCH_TRUE, ""},
		geTestCase{uint16(149), MATCH_TRUE, ""},
		geTestCase{uint16(150), MATCH_FALSE, ""},
		geTestCase{uint16(151), MATCH_FALSE, ""},

		geTestCase{uint32(0), MATCH_TRUE, ""},
		geTestCase{uint32(149), MATCH_TRUE, ""},
		geTestCase{uint32(150), MATCH_FALSE, ""},
		geTestCase{uint32(151), MATCH_FALSE, ""},

		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(149), MATCH_TRUE, ""},
		geTestCase{uint64(150), MATCH_FALSE, ""},
		geTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-1), MATCH_TRUE, ""},
		geTestCase{float32(149), MATCH_TRUE, ""},
		geTestCase{float32(149.8), MATCH_TRUE, ""},
		geTestCase{float32(149.9), MATCH_FALSE, ""},
		geTestCase{float32(150), MATCH_FALSE, ""},
		geTestCase{float32(151), MATCH_FALSE, ""},

		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(149), MATCH_TRUE, ""},
		geTestCase{float64(149.8), MATCH_TRUE, ""},
		geTestCase{float64(149.9), MATCH_FALSE, ""},
		geTestCase{float64(150), MATCH_FALSE, ""},
		geTestCase{float64(151), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// Subtle cases
////////////////////////////////////////////////////////////

func (t *GreaterOrEqualTest) Int64NotExactlyRepresentableBySinglePrecision() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := GreaterOrEqual(int64(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than or equal to 33554433"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{kTwoTo25 + 0, MATCH_TRUE, ""},
		geTestCase{kTwoTo25 + 1, MATCH_FALSE, ""},
		geTestCase{kTwoTo25 + 2, MATCH_FALSE, ""},

		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{int(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(127), MATCH_TRUE, ""},

		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(0), MATCH_TRUE, ""},
		geTestCase{int16(32767), MATCH_TRUE, ""},

		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{int32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int32(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint(0), MATCH_TRUE, ""},
		geTestCase{uint(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{uint(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_TRUE, ""},
		geTestCase{uint8(255), MATCH_TRUE, ""},

		geTestCase{uint16(0), MATCH_TRUE, ""},
		geTestCase{uint16(65535), MATCH_TRUE, ""},

		geTestCase{uint32(0), MATCH_TRUE, ""},
		geTestCase{uint32(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{uint32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint32(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-1), MATCH_TRUE, ""},
		geTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{float32(kTwoTo25 - 1), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 0), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) Int64NotExactlyRepresentableByDoublePrecision() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := GreaterOrEqual(int64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than or equal to 18014398509481985"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{1 << 30, MATCH_TRUE, ""},

		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(math.MaxInt32), MATCH_TRUE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(127), MATCH_TRUE, ""},

		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(0), MATCH_TRUE, ""},
		geTestCase{int16(32767), MATCH_TRUE, ""},

		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(math.MaxInt32), MATCH_TRUE, ""},

		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint(0), MATCH_TRUE, ""},
		geTestCase{uint(math.MaxUint32), MATCH_TRUE, ""},

		geTestCase{uint8(0), MATCH_TRUE, ""},
		geTestCase{uint8(255), MATCH_TRUE, ""},

		geTestCase{uint16(0), MATCH_TRUE, ""},
		geTestCase{uint16(65535), MATCH_TRUE, ""},

		geTestCase{uint32(0), MATCH_TRUE, ""},
		geTestCase{uint32(math.MaxUint32), MATCH_TRUE, ""},

		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo54 - 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 0), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 2), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) Uint64NotExactlyRepresentableBySinglePrecision() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := GreaterOrEqual(uint64(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than or equal to 33554433"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{kTwoTo25 + 0, MATCH_TRUE, ""},
		geTestCase{kTwoTo25 + 1, MATCH_FALSE, ""},
		geTestCase{kTwoTo25 + 2, MATCH_FALSE, ""},

		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{int(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(127), MATCH_TRUE, ""},

		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(0), MATCH_TRUE, ""},
		geTestCase{int16(32767), MATCH_TRUE, ""},

		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{int32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int32(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint(0), MATCH_TRUE, ""},
		geTestCase{uint(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{uint(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{uint8(0), MATCH_TRUE, ""},
		geTestCase{uint8(255), MATCH_TRUE, ""},

		geTestCase{uint16(0), MATCH_TRUE, ""},
		geTestCase{uint16(65535), MATCH_TRUE, ""},

		geTestCase{uint32(0), MATCH_TRUE, ""},
		geTestCase{uint32(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{uint32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint32(kTwoTo25 + 2), MATCH_FALSE, ""},

		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-1), MATCH_TRUE, ""},
		geTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{float32(kTwoTo25 - 1), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 0), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) Uint64NotExactlyRepresentableByDoublePrecision() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := GreaterOrEqual(uint64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than or equal to 18014398509481985"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{-1, MATCH_TRUE, ""},
		geTestCase{1 << 30, MATCH_TRUE, ""},

		geTestCase{int(-1), MATCH_TRUE, ""},
		geTestCase{int(math.MaxInt32), MATCH_TRUE, ""},

		geTestCase{int8(-1), MATCH_TRUE, ""},
		geTestCase{int8(127), MATCH_TRUE, ""},

		geTestCase{int16(-1), MATCH_TRUE, ""},
		geTestCase{int16(0), MATCH_TRUE, ""},
		geTestCase{int16(32767), MATCH_TRUE, ""},

		geTestCase{int32(-1), MATCH_TRUE, ""},
		geTestCase{int32(math.MaxInt32), MATCH_TRUE, ""},

		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint(0), MATCH_TRUE, ""},
		geTestCase{uint(math.MaxUint32), MATCH_TRUE, ""},

		geTestCase{uint8(0), MATCH_TRUE, ""},
		geTestCase{uint8(255), MATCH_TRUE, ""},

		geTestCase{uint16(0), MATCH_TRUE, ""},
		geTestCase{uint16(65535), MATCH_TRUE, ""},

		geTestCase{uint32(0), MATCH_TRUE, ""},
		geTestCase{uint32(math.MaxUint32), MATCH_TRUE, ""},

		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo54 - 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 0), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 2), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) Float32AboveExactIntegerRange() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := GreaterOrEqual(float32(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than or equal to 3.3554432e+07"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo25 - 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo25 - 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float32(-1), MATCH_TRUE, ""},
		geTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{float32(kTwoTo25 - 1), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 0), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo25 - 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) Float64AboveExactIntegerRange() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := GreaterOrEqual(float64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "greater than or equal to 1.8014398509481984e+16"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		// Signed integers.
		geTestCase{int64(-1), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 - 2), MATCH_TRUE, ""},
		geTestCase{int64(kTwoTo54 - 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo54 + 0), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},
		geTestCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Unsigned integers.
		geTestCase{uint64(0), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 - 2), MATCH_TRUE, ""},
		geTestCase{uint64(kTwoTo54 - 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo54 + 0), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},
		geTestCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Floating point.
		geTestCase{float64(-1), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		geTestCase{float64(kTwoTo54 - 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 0), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 1), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 2), MATCH_FALSE, ""},
		geTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// String literals
////////////////////////////////////////////////////////////

func (t *GreaterOrEqualTest) EmptyString() {
	matcher := GreaterOrEqual("")
	desc := matcher.Description()
	expectedDesc := "greater than or equal to \"\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		geTestCase{"", MATCH_FALSE, ""},
		geTestCase{"\x00", MATCH_FALSE, ""},
		geTestCase{"a", MATCH_FALSE, ""},
		geTestCase{"foo", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) SingleNullByte() {
	matcher := GreaterOrEqual("\x00")
	desc := matcher.Description()
	expectedDesc := "greater than or equal to \"\x00\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		geTestCase{"", MATCH_TRUE, ""},
		geTestCase{"\x00", MATCH_FALSE, ""},
		geTestCase{"a", MATCH_FALSE, ""},
		geTestCase{"foo", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *GreaterOrEqualTest) LongerString() {
	matcher := GreaterOrEqual("foo\x00")
	desc := matcher.Description()
	expectedDesc := "greater than or equal to \"foo\x00\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []geTestCase{
		geTestCase{"", MATCH_TRUE, ""},
		geTestCase{"\x00", MATCH_TRUE, ""},
		geTestCase{"bar", MATCH_TRUE, ""},
		geTestCase{"foo", MATCH_TRUE, ""},
		geTestCase{"foo\x00", MATCH_FALSE, ""},
		geTestCase{"fooa", MATCH_FALSE, ""},
		geTestCase{"qux", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}
