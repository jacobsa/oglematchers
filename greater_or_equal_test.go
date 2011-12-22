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

type LessThanTest struct {
}

func init() { RegisterTestSuite(&LessThanTest{}) }

type ltTestCase struct {
	candidate      interface{}
	expectedResult MatchResult
	expectedError  string
}

func (t *LessThanTest) checkTestCases(matcher Matcher, cases []ltTestCase) {
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

func (t *LessThanTest) IntegerCandidateBadTypes() {
	matcher := LessThan(int(-150))

	cases := []ltTestCase{
		ltTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{&ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) FloatCandidateBadTypes() {
	matcher := LessThan(float32(-150))

	cases := []ltTestCase{
		ltTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{&ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{"-151", MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) StringCandidateBadTypes() {
	matcher := LessThan("17")

	cases := []ltTestCase{
		ltTestCase{true, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int8(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int16(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int32(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{int64(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint8(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint16(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint32(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uint64(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{uintptr(17), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{float32(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{float64(0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex64(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{complex128(-151), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{[...]int{-151}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make(chan int), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{func() {}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{map[int]int{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{&ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{make([]int, 0), MATCH_UNDEFINED, "which is not comparable"},
		ltTestCase{ltTestCase{}, MATCH_UNDEFINED, "which is not comparable"},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) BadArgument() {
	panicked := false

	defer func() {
		ExpectThat(panicked, Equals(true))
	}()

	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	LessThan(complex128(0))
}

////////////////////////////////////////////////////////////
// Integer literals
////////////////////////////////////////////////////////////

func (t *LessThanTest) NegativeIntegerLiteral() {
	matcher := LessThan(-150)
	desc := matcher.Description()
	expectedDesc := "less than -150"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-(1 << 30), MATCH_TRUE, ""},
		ltTestCase{-151, MATCH_TRUE, ""},
		ltTestCase{-150, MATCH_FALSE, ""},
		ltTestCase{0, MATCH_FALSE, ""},
		ltTestCase{17, MATCH_FALSE, ""},

		ltTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int(-151), MATCH_TRUE, ""},
		ltTestCase{int(-150), MATCH_FALSE, ""},
		ltTestCase{int(0), MATCH_FALSE, ""},
		ltTestCase{int(17), MATCH_FALSE, ""},

		ltTestCase{int8(-127), MATCH_FALSE, ""},
		ltTestCase{int8(0), MATCH_FALSE, ""},
		ltTestCase{int8(17), MATCH_FALSE, ""},

		ltTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		ltTestCase{int16(-151), MATCH_TRUE, ""},
		ltTestCase{int16(-150), MATCH_FALSE, ""},
		ltTestCase{int16(0), MATCH_FALSE, ""},
		ltTestCase{int16(17), MATCH_FALSE, ""},

		ltTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int32(-151), MATCH_TRUE, ""},
		ltTestCase{int32(-150), MATCH_FALSE, ""},
		ltTestCase{int32(0), MATCH_FALSE, ""},
		ltTestCase{int32(17), MATCH_FALSE, ""},

		ltTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int64(-151), MATCH_TRUE, ""},
		ltTestCase{int64(-150), MATCH_FALSE, ""},
		ltTestCase{int64(0), MATCH_FALSE, ""},
		ltTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		ltTestCase{uint(0), MATCH_FALSE, ""},
		ltTestCase{uint(17), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_FALSE, ""},
		ltTestCase{uint8(17), MATCH_FALSE, ""},
		ltTestCase{uint8(253), MATCH_FALSE, ""},

		ltTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		ltTestCase{uint16(0), MATCH_FALSE, ""},
		ltTestCase{uint16(17), MATCH_FALSE, ""},

		ltTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		ltTestCase{uint32(0), MATCH_FALSE, ""},
		ltTestCase{uint32(17), MATCH_FALSE, ""},

		ltTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		ltTestCase{uint64(0), MATCH_FALSE, ""},
		ltTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float32(-151), MATCH_TRUE, ""},
		ltTestCase{float32(-150.1), MATCH_TRUE, ""},
		ltTestCase{float32(-150), MATCH_FALSE, ""},
		ltTestCase{float32(-149.9), MATCH_FALSE, ""},
		ltTestCase{float32(0), MATCH_FALSE, ""},
		ltTestCase{float32(17), MATCH_FALSE, ""},
		ltTestCase{float32(160), MATCH_FALSE, ""},

		ltTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float64(-151), MATCH_TRUE, ""},
		ltTestCase{float64(-150.1), MATCH_TRUE, ""},
		ltTestCase{float64(-150), MATCH_FALSE, ""},
		ltTestCase{float64(-149.9), MATCH_FALSE, ""},
		ltTestCase{float64(0), MATCH_FALSE, ""},
		ltTestCase{float64(17), MATCH_FALSE, ""},
		ltTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) ZeroIntegerLiteral() {
	matcher := LessThan(0)
	desc := matcher.Description()
	expectedDesc := "less than 0"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-(1 << 30), MATCH_TRUE, ""},
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{0, MATCH_FALSE, ""},
		ltTestCase{1, MATCH_FALSE, ""},
		ltTestCase{17, MATCH_FALSE, ""},
		ltTestCase{(1 << 30), MATCH_FALSE, ""},

		ltTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(0), MATCH_FALSE, ""},
		ltTestCase{int(1), MATCH_FALSE, ""},
		ltTestCase{int(17), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(0), MATCH_FALSE, ""},
		ltTestCase{int8(1), MATCH_FALSE, ""},

		ltTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(0), MATCH_FALSE, ""},
		ltTestCase{int16(1), MATCH_FALSE, ""},
		ltTestCase{int16(17), MATCH_FALSE, ""},

		ltTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(0), MATCH_FALSE, ""},
		ltTestCase{int32(1), MATCH_FALSE, ""},
		ltTestCase{int32(17), MATCH_FALSE, ""},

		ltTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(0), MATCH_FALSE, ""},
		ltTestCase{int64(1), MATCH_FALSE, ""},
		ltTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint((1 << 32) - 1), MATCH_FALSE, ""},
		ltTestCase{uint(0), MATCH_FALSE, ""},
		ltTestCase{uint(17), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_FALSE, ""},
		ltTestCase{uint8(17), MATCH_FALSE, ""},
		ltTestCase{uint8(253), MATCH_FALSE, ""},

		ltTestCase{uint16((1 << 16) - 1), MATCH_FALSE, ""},
		ltTestCase{uint16(0), MATCH_FALSE, ""},
		ltTestCase{uint16(17), MATCH_FALSE, ""},

		ltTestCase{uint32((1 << 32) - 1), MATCH_FALSE, ""},
		ltTestCase{uint32(0), MATCH_FALSE, ""},
		ltTestCase{uint32(17), MATCH_FALSE, ""},

		ltTestCase{uint64((1 << 64) - 1), MATCH_FALSE, ""},
		ltTestCase{uint64(0), MATCH_FALSE, ""},
		ltTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(-0.1), MATCH_TRUE, ""},
		ltTestCase{float32(-0.0), MATCH_FALSE, ""},
		ltTestCase{float32(0), MATCH_FALSE, ""},
		ltTestCase{float32(0.1), MATCH_FALSE, ""},
		ltTestCase{float32(17), MATCH_FALSE, ""},
		ltTestCase{float32(160), MATCH_FALSE, ""},

		ltTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(-0.1), MATCH_TRUE, ""},
		ltTestCase{float64(-0), MATCH_FALSE, ""},
		ltTestCase{float64(0), MATCH_FALSE, ""},
		ltTestCase{float64(17), MATCH_FALSE, ""},
		ltTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) PositiveIntegerLiteral() {
	matcher := LessThan(150)
	desc := matcher.Description()
	expectedDesc := "less than 150"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{149, MATCH_TRUE, ""},
		ltTestCase{150, MATCH_FALSE, ""},
		ltTestCase{151, MATCH_FALSE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(149), MATCH_TRUE, ""},
		ltTestCase{int(150), MATCH_FALSE, ""},
		ltTestCase{int(151), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(0), MATCH_TRUE, ""},
		ltTestCase{int8(17), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(149), MATCH_TRUE, ""},
		ltTestCase{int16(150), MATCH_FALSE, ""},
		ltTestCase{int16(151), MATCH_FALSE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(149), MATCH_TRUE, ""},
		ltTestCase{int32(150), MATCH_FALSE, ""},
		ltTestCase{int32(151), MATCH_FALSE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(149), MATCH_TRUE, ""},
		ltTestCase{int64(150), MATCH_FALSE, ""},
		ltTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(149), MATCH_TRUE, ""},
		ltTestCase{uint(150), MATCH_FALSE, ""},
		ltTestCase{uint(151), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(127), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(149), MATCH_TRUE, ""},
		ltTestCase{uint16(150), MATCH_FALSE, ""},
		ltTestCase{uint16(151), MATCH_FALSE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(149), MATCH_TRUE, ""},
		ltTestCase{uint32(150), MATCH_FALSE, ""},
		ltTestCase{uint32(151), MATCH_FALSE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(149), MATCH_TRUE, ""},
		ltTestCase{uint64(150), MATCH_FALSE, ""},
		ltTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(149), MATCH_TRUE, ""},
		ltTestCase{float32(149.9), MATCH_TRUE, ""},
		ltTestCase{float32(150), MATCH_FALSE, ""},
		ltTestCase{float32(150.1), MATCH_FALSE, ""},
		ltTestCase{float32(151), MATCH_FALSE, ""},

		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(149), MATCH_TRUE, ""},
		ltTestCase{float64(149.9), MATCH_TRUE, ""},
		ltTestCase{float64(150), MATCH_FALSE, ""},
		ltTestCase{float64(150.1), MATCH_FALSE, ""},
		ltTestCase{float64(151), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// Float literals
////////////////////////////////////////////////////////////

func (t *LessThanTest) NegativeFloatLiteral() {
	matcher := LessThan(-150.1)
	desc := matcher.Description()
	expectedDesc := "less than -150.1"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-(1 << 30), MATCH_TRUE, ""},
		ltTestCase{-151, MATCH_TRUE, ""},
		ltTestCase{-150, MATCH_FALSE, ""},
		ltTestCase{0, MATCH_FALSE, ""},
		ltTestCase{17, MATCH_FALSE, ""},

		ltTestCase{int(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int(-151), MATCH_TRUE, ""},
		ltTestCase{int(-150), MATCH_FALSE, ""},
		ltTestCase{int(0), MATCH_FALSE, ""},
		ltTestCase{int(17), MATCH_FALSE, ""},

		ltTestCase{int8(-127), MATCH_FALSE, ""},
		ltTestCase{int8(0), MATCH_FALSE, ""},
		ltTestCase{int8(17), MATCH_FALSE, ""},

		ltTestCase{int16(-(1 << 14)), MATCH_TRUE, ""},
		ltTestCase{int16(-151), MATCH_TRUE, ""},
		ltTestCase{int16(-150), MATCH_FALSE, ""},
		ltTestCase{int16(0), MATCH_FALSE, ""},
		ltTestCase{int16(17), MATCH_FALSE, ""},

		ltTestCase{int32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int32(-151), MATCH_TRUE, ""},
		ltTestCase{int32(-150), MATCH_FALSE, ""},
		ltTestCase{int32(0), MATCH_FALSE, ""},
		ltTestCase{int32(17), MATCH_FALSE, ""},

		ltTestCase{int64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{int64(-151), MATCH_TRUE, ""},
		ltTestCase{int64(-150), MATCH_FALSE, ""},
		ltTestCase{int64(0), MATCH_FALSE, ""},
		ltTestCase{int64(17), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint((1 << 32) - 151), MATCH_FALSE, ""},
		ltTestCase{uint(0), MATCH_FALSE, ""},
		ltTestCase{uint(17), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_FALSE, ""},
		ltTestCase{uint8(17), MATCH_FALSE, ""},
		ltTestCase{uint8(253), MATCH_FALSE, ""},

		ltTestCase{uint16((1 << 16) - 151), MATCH_FALSE, ""},
		ltTestCase{uint16(0), MATCH_FALSE, ""},
		ltTestCase{uint16(17), MATCH_FALSE, ""},

		ltTestCase{uint32((1 << 32) - 151), MATCH_FALSE, ""},
		ltTestCase{uint32(0), MATCH_FALSE, ""},
		ltTestCase{uint32(17), MATCH_FALSE, ""},

		ltTestCase{uint64((1 << 64) - 151), MATCH_FALSE, ""},
		ltTestCase{uint64(0), MATCH_FALSE, ""},
		ltTestCase{uint64(17), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float32(-151), MATCH_TRUE, ""},
		ltTestCase{float32(-150.2), MATCH_TRUE, ""},
		ltTestCase{float32(-150.1), MATCH_FALSE, ""},
		ltTestCase{float32(-150), MATCH_FALSE, ""},
		ltTestCase{float32(0), MATCH_FALSE, ""},
		ltTestCase{float32(17), MATCH_FALSE, ""},
		ltTestCase{float32(160), MATCH_FALSE, ""},

		ltTestCase{float64(-(1 << 30)), MATCH_TRUE, ""},
		ltTestCase{float64(-151), MATCH_TRUE, ""},
		ltTestCase{float64(-150.2), MATCH_TRUE, ""},
		ltTestCase{float64(-150.1), MATCH_FALSE, ""},
		ltTestCase{float64(-150), MATCH_FALSE, ""},
		ltTestCase{float64(0), MATCH_FALSE, ""},
		ltTestCase{float64(17), MATCH_FALSE, ""},
		ltTestCase{float64(160), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) PositiveFloatLiteral() {
	matcher := LessThan(149.9)
	desc := matcher.Description()
	expectedDesc := "less than 149.9"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{149, MATCH_TRUE, ""},
		ltTestCase{150, MATCH_FALSE, ""},
		ltTestCase{151, MATCH_FALSE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(149), MATCH_TRUE, ""},
		ltTestCase{int(150), MATCH_FALSE, ""},
		ltTestCase{int(151), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(0), MATCH_TRUE, ""},
		ltTestCase{int8(17), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(149), MATCH_TRUE, ""},
		ltTestCase{int16(150), MATCH_FALSE, ""},
		ltTestCase{int16(151), MATCH_FALSE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(149), MATCH_TRUE, ""},
		ltTestCase{int32(150), MATCH_FALSE, ""},
		ltTestCase{int32(151), MATCH_FALSE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(149), MATCH_TRUE, ""},
		ltTestCase{int64(150), MATCH_FALSE, ""},
		ltTestCase{int64(151), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(149), MATCH_TRUE, ""},
		ltTestCase{uint(150), MATCH_FALSE, ""},
		ltTestCase{uint(151), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(127), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(149), MATCH_TRUE, ""},
		ltTestCase{uint16(150), MATCH_FALSE, ""},
		ltTestCase{uint16(151), MATCH_FALSE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(149), MATCH_TRUE, ""},
		ltTestCase{uint32(150), MATCH_FALSE, ""},
		ltTestCase{uint32(151), MATCH_FALSE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(149), MATCH_TRUE, ""},
		ltTestCase{uint64(150), MATCH_FALSE, ""},
		ltTestCase{uint64(151), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(149), MATCH_TRUE, ""},
		ltTestCase{float32(149.8), MATCH_TRUE, ""},
		ltTestCase{float32(149.9), MATCH_FALSE, ""},
		ltTestCase{float32(150), MATCH_FALSE, ""},
		ltTestCase{float32(151), MATCH_FALSE, ""},

		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(149), MATCH_TRUE, ""},
		ltTestCase{float64(149.8), MATCH_TRUE, ""},
		ltTestCase{float64(149.9), MATCH_FALSE, ""},
		ltTestCase{float64(150), MATCH_FALSE, ""},
		ltTestCase{float64(151), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// Subtle cases
////////////////////////////////////////////////////////////

func (t *LessThanTest) Int64NotExactlyRepresentableBySinglePrecision() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := LessThan(int64(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "less than 33554433"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{kTwoTo25 + 0, MATCH_TRUE, ""},
		ltTestCase{kTwoTo25 + 1, MATCH_FALSE, ""},
		ltTestCase{kTwoTo25 + 2, MATCH_FALSE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{int(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(0), MATCH_TRUE, ""},
		ltTestCase{int16(32767), MATCH_TRUE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{int32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int32(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{uint(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(255), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(65535), MATCH_TRUE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{uint32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint32(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{float32(kTwoTo25 - 1), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 0), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) Int64NotExactlyRepresentableByDoublePrecision() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := LessThan(int64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "less than 18014398509481985"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{1 << 30, MATCH_TRUE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(math.MaxInt32), MATCH_TRUE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(0), MATCH_TRUE, ""},
		ltTestCase{int16(32767), MATCH_TRUE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(math.MaxInt32), MATCH_TRUE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(math.MaxUint32), MATCH_TRUE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(255), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(65535), MATCH_TRUE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(math.MaxUint32), MATCH_TRUE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo54 - 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 0), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 2), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) Uint64NotExactlyRepresentableBySinglePrecision() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := LessThan(uint64(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "less than 33554433"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{kTwoTo25 + 0, MATCH_TRUE, ""},
		ltTestCase{kTwoTo25 + 1, MATCH_FALSE, ""},
		ltTestCase{kTwoTo25 + 2, MATCH_FALSE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{int(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(0), MATCH_TRUE, ""},
		ltTestCase{int16(32767), MATCH_TRUE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{int32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int32(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{uint(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(255), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(65535), MATCH_TRUE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{uint32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint32(kTwoTo25 + 2), MATCH_FALSE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{float32(kTwoTo25 - 1), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 0), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 - 1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 + 0), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) Uint64NotExactlyRepresentableByDoublePrecision() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := LessThan(uint64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "less than 18014398509481985"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{-1, MATCH_TRUE, ""},
		ltTestCase{1 << 30, MATCH_TRUE, ""},

		ltTestCase{int(-1), MATCH_TRUE, ""},
		ltTestCase{int(math.MaxInt32), MATCH_TRUE, ""},

		ltTestCase{int8(-1), MATCH_TRUE, ""},
		ltTestCase{int8(127), MATCH_TRUE, ""},

		ltTestCase{int16(-1), MATCH_TRUE, ""},
		ltTestCase{int16(0), MATCH_TRUE, ""},
		ltTestCase{int16(32767), MATCH_TRUE, ""},

		ltTestCase{int32(-1), MATCH_TRUE, ""},
		ltTestCase{int32(math.MaxInt32), MATCH_TRUE, ""},

		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 - 1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 + 0), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint(0), MATCH_TRUE, ""},
		ltTestCase{uint(math.MaxUint32), MATCH_TRUE, ""},

		ltTestCase{uint8(0), MATCH_TRUE, ""},
		ltTestCase{uint8(255), MATCH_TRUE, ""},

		ltTestCase{uint16(0), MATCH_TRUE, ""},
		ltTestCase{uint16(65535), MATCH_TRUE, ""},

		ltTestCase{uint32(0), MATCH_TRUE, ""},
		ltTestCase{uint32(math.MaxUint32), MATCH_TRUE, ""},

		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 - 1), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 + 0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo54 - 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 0), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 2), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) Float32AboveExactIntegerRange() {
	// Single-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^25-1, 2^25+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo25 = 1 << 25
	matcher := LessThan(float32(kTwoTo25 + 1))

	desc := matcher.Description()
	expectedDesc := "less than 3.3554432e+07"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo25 - 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo25 + 0), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo25 - 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo25 + 0), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo25 + 3), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float32(-1), MATCH_TRUE, ""},
		ltTestCase{float32(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{float32(kTwoTo25 - 1), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 0), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{float32(kTwoTo25 + 3), MATCH_FALSE, ""},

		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 - 2), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo25 - 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 0), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 2), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo25 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) Float64AboveExactIntegerRange() {
	// Double-precision floats don't have enough bits to represent the integers
	// near this one distinctly, so [2^54-1, 2^54+2] all receive the same value
	// and should be treated as equivalent when floats are in the mix.
	const kTwoTo54 = 1 << 54
	matcher := LessThan(float64(kTwoTo54 + 1))

	desc := matcher.Description()
	expectedDesc := "less than 1.8014398509481984e+16"

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		// Signed integers.
		ltTestCase{int64(-1), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 - 2), MATCH_TRUE, ""},
		ltTestCase{int64(kTwoTo54 - 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo54 + 0), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo54 + 2), MATCH_FALSE, ""},
		ltTestCase{int64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Unsigned integers.
		ltTestCase{uint64(0), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 - 2), MATCH_TRUE, ""},
		ltTestCase{uint64(kTwoTo54 - 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo54 + 0), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo54 + 2), MATCH_FALSE, ""},
		ltTestCase{uint64(kTwoTo54 + 3), MATCH_FALSE, ""},

		// Floating point.
		ltTestCase{float64(-1), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo54 - 2), MATCH_TRUE, ""},
		ltTestCase{float64(kTwoTo54 - 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 0), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 1), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 2), MATCH_FALSE, ""},
		ltTestCase{float64(kTwoTo54 + 3), MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

////////////////////////////////////////////////////////////
// String literals
////////////////////////////////////////////////////////////

func (t *LessThanTest) EmptyString() {
	matcher := LessThan("")
	desc := matcher.Description()
	expectedDesc := "less than \"\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		ltTestCase{"", MATCH_FALSE, ""},
		ltTestCase{"\x00", MATCH_FALSE, ""},
		ltTestCase{"a", MATCH_FALSE, ""},
		ltTestCase{"foo", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) SingleNullByte() {
	matcher := LessThan("\x00")
	desc := matcher.Description()
	expectedDesc := "less than \"\x00\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		ltTestCase{"", MATCH_TRUE, ""},
		ltTestCase{"\x00", MATCH_FALSE, ""},
		ltTestCase{"a", MATCH_FALSE, ""},
		ltTestCase{"foo", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}

func (t *LessThanTest) LongerString() {
	matcher := LessThan("foo\x00")
	desc := matcher.Description()
	expectedDesc := "less than \"foo\x00\""

	ExpectThat(desc, Equals(expectedDesc))

	cases := []ltTestCase{
		ltTestCase{"", MATCH_TRUE, ""},
		ltTestCase{"\x00", MATCH_TRUE, ""},
		ltTestCase{"bar", MATCH_TRUE, ""},
		ltTestCase{"foo", MATCH_TRUE, ""},
		ltTestCase{"foo\x00", MATCH_FALSE, ""},
		ltTestCase{"fooa", MATCH_FALSE, ""},
		ltTestCase{"qux", MATCH_FALSE, ""},
	}

	t.checkTestCases(matcher, cases)
}
