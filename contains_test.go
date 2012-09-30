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
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type ContainsTest struct {}
func init() { RegisterTestSuite(&ContainsTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *ContainsTest) WrongTypeCandidates() {
	m := Contains("")

	var err error

	// Nil candidate
	err = m.Matches(nil)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("array")))
	ExpectThat(err, Error(HasSubstr("slice")))

	// String candidate
	err = m.Matches("")
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("array")))
	ExpectThat(err, Error(HasSubstr("slice")))

	// Map candidate
	err = m.Matches(make(map[string]string))
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("array")))
	ExpectThat(err, Error(HasSubstr("slice")))
}

func (t *ContainsTest) NilArgument() {
	m := Contains(nil)

	var c interface{}
	var err error

	// Empty array of pointers
	c = [...]*int{}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Empty slice of pointers
	c = []*int{}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Non-empty array of integers
	c = [...]int{17, 0, 19}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Non-empty slice of integers
	c = []int{17, 0, 19}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Non-matching array of pointers
	c = [...]*int{new(int), new(int)}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Non-matching slice of pointers
	c = []*int{new(int), new(int)}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Matching array of pointers
	c = [...]*int{new(int), nil, new(int)}
	err = m.Matches(c)
	ExpectEq(nil, err)

	// Matching slice of pointers
	c = []*int{new(int), nil, new(int)}
	err = m.Matches(c)
	ExpectEq(nil, err)
}

func (t *ContainsTest) IntegerArgument() {
	ExpectFalse(true, "TODO")
}

func (t *ContainsTest) StringArgument() {
	ExpectFalse(true, "TODO")
}

func (t *ContainsTest) MatcherArgument() {
	ExpectFalse(true, "TODO")
}
