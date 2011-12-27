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
	"errors"
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type PanicsTest struct {
	matcherCalled bool
	suppliedCandidate interface{}
	wrappedResult bool
	wrappedError error

	matcher Matcher
}

func init() { RegisterTestSuite(&PanicsTest{}) }

func (t *PanicsTest) SetUp() {
	wrapped := &fakeMatcher{
		func(c interface{}) (bool, error) {
			t.matcherCalled = true
			t.suppliedCandidate = c
			return t.wrappedResult, t.wrappedError
		},
		"foo",
	}

	t.matcher = Panics(wrapped)
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *PanicsTest) Description() {
	ExpectThat(t.matcher.Description(), Equals("panics with: foo"))
}

func (t *PanicsTest) CandidateIsNil() {
	res, err := t.matcher.Matches(nil)

	ExpectThat(res, Equals(false))
	ExpectThat(err, Error(Equals("which is not a zero-arg function")))
	ExpectTrue(isFatal(err))
}

func (t *PanicsTest) CandidateIsString() {
	res, err := t.matcher.Matches("taco")

	ExpectThat(res, Equals(false))
	ExpectThat(err, Error(Equals("which is not a zero-arg function")))
	ExpectTrue(isFatal(err))
}

func (t *PanicsTest) CandidateTakesArgs() {
	res, err := t.matcher.Matches(func(i int) string { return "" })

	ExpectThat(res, Equals(false))
	ExpectThat(err, Error(Equals("which is not a zero-arg function")))
	ExpectTrue(isFatal(err))
}

func (t *PanicsTest) CallsFunction() {
	callCount := 0
	t.matcher.Matches(func() string {
		callCount++
		return ""
	})

	ExpectThat(callCount, Equals(1))
}

func (t *PanicsTest) FunctionDoesntPanic() {
	res, err := t.matcher.Matches(func() {})

	ExpectThat(res, Equals(false))
	ExpectThat(err, Error(Equals("which didn't panic")))
	ExpectFalse(isFatal(err))
}

func (t *PanicsTest) CallsWrappedMatcher() {
	expectedErr := 17
	t.matcher.Matches(func() { panic(expectedErr) })

	ExpectThat(t.suppliedCandidate, Equals(expectedErr))
}

func (t *PanicsTest) WrappedReturnsTrue() {
	t.wrappedResult = true
	res, err := t.matcher.Matches(func() { panic("") })

	ExpectTrue(res)
	ExpectEq(nil, err)
}

func (t *PanicsTest) WrappedReturnsFatalErrorWithoutText() {
	t.wrappedResult = false
	t.wrappedError = NewFatalError("")
	res, err := t.matcher.Matches(func() { panic(17) })

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which panicked with: 17")))
	ExpectFalse(isFatal(err))
}

func (t *PanicsTest) WrappedReturnsFatalErrorWithText() {
	t.wrappedResult = false
	t.wrappedError = NewFatalError("which blah")
	res, err := t.matcher.Matches(func() { panic(17) })

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which panicked with: 17, which blah")))
	ExpectFalse(isFatal(err))
}

func (t *PanicsTest) WrappedReturnsNonFatalErrorWithoutText() {
	t.wrappedResult = false
	t.wrappedError = errors.New("")
	res, err := t.matcher.Matches(func() { panic(17) })

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which panicked with: 17")))
	ExpectFalse(isFatal(err))
}

func (t *PanicsTest) WrappedReturnsNonFatalErrorWithText() {
	t.wrappedResult = false
	t.wrappedError = errors.New("which blah")
	res, err := t.matcher.Matches(func() { panic(17) })

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which panicked with: 17, which blah")))
	ExpectFalse(isFatal(err))
}
