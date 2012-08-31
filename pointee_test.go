// Copyright 2012 Aaron Jacobs. All Rights Reserved.
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
	"testing"
)

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type PointeeTest struct {}
func init()                     { RegisterTestSuite(&PointeeTest{}) }

func TestPointee(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *PointeeTest) Description() {
	wrapped := &fakeMatcher{nil, "taco"}
	matcher := Pointee(wrapped)

	ExpectEq("pointee(taco)", matcher.Description())
}

func (t *PointeeTest) CandidateIsNotAPointer() {
	matcher := Pointee(HasSubstr(""))
	err := matcher.Matches([]byte{})

	ExpectThat(err, Error(Equals("which is not a pointer")))
	ExpectTrue(isFatal(err))
}

func (t *PointeeTest) CandidateIsANilLiteral() {
	matcher := Pointee(HasSubstr(""))
	err := matcher.Matches(nil)

	ExpectThat(err, Error(Equals("which is not a pointer")))
	ExpectTrue(isFatal(err))
}

func (t *PointeeTest) CandidateIsANilPointer() {
	matcher := Pointee(HasSubstr(""))
	err := matcher.Matches((*int)(nil))

	ExpectThat(err, Error(Equals("")))
	ExpectTrue(isFatal(err))
}

func (t *PointeeTest) CallsWrapped() {
	var suppliedCandidate interface{}
	matchFunc := func(c interface{}) error {
		suppliedCandidate = c
		return nil
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Pointee(wrapped)

	someSlice := []byte{}
	matcher.Matches(&someSlice)
	ExpectThat(suppliedCandidate, IdenticalTo(someSlice))
}

func (t *PointeeTest) WrappedReturnsOkay() {
	matchFunc := func(c interface{}) error {
		return nil
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Pointee(wrapped)

	err := matcher.Matches(new(int))
	ExpectEq(nil, err)
}

func (t *PointeeTest) WrappedReturnsNonFatalError() {
	matchFunc := func(c interface{}) error {
		return errors.New("taco")
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Pointee(wrapped)

	err := matcher.Matches(new(int))
	ExpectThat(err, Error(Equals("taco")))
	ExpectFalse(isFatal(err))
}

func (t *PointeeTest) WrappedReturnsFatalError() {
	matchFunc := func(c interface{}) error {
		return NewFatalError("taco")
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Pointee(wrapped)

	err := matcher.Matches(new(int))
	ExpectThat(err, Error(Equals("taco")))
	ExpectTrue(isFatal(err))
}
