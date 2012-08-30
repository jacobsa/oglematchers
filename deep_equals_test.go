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

type DeepEqualsTest struct {}
func init() { RegisterTestSuite(&DeepEqualsTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *DeepEqualsTest) WrongTypeCandidateWithScalarValue() {
	var x int = 17
	m := DeepEquals(x)

	var err error

	// Nil candidate.
	err = m.Matches(nil)
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("TODO")))
	ExpectThat(err, Error(HasSubstr("int")))

	// String candidate.
	err = m.Matches("taco")
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("string")))
	ExpectThat(err, Error(HasSubstr("int")))

	// Slice candidate.
	err = m.Matches([]byte{})
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("[]byte")))
	ExpectThat(err, Error(HasSubstr("int")))

	// Unsigned int candidate.
	err = m.Matches(uint(17))
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("uint")))
	ExpectThat(err, Error(HasSubstr("int")))
}

func (t *DeepEqualsTest) WrongTypeCandidateWithSliceValue() {
	x := []byte{}
	m := DeepEquals(x)

	var err error

	// Nil candidate.
	err = m.Matches(nil)
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("TODO")))
	ExpectThat(err, Error(HasSubstr("[]byte")))

	// String candidate.
	err = m.Matches("taco")
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("string")))
	ExpectThat(err, Error(HasSubstr("[]byte")))

	// Slice candidate with wrong value type.
	err = m.Matches([]uint8{})
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("[]uint8")))
	ExpectThat(err, Error(HasSubstr("[]byte")))
}

func (t *DeepEqualsTest) NilValue() {
	ExpectEq("TODO", "")
}

func (t *DeepEqualsTest) IntValue() {
	ExpectEq("TODO", "")
}

func (t *DeepEqualsTest) IntAliasValue() {
	ExpectEq("TODO", "")
}

func (t *DeepEqualsTest) SliceValue() {
	ExpectEq("TODO", "")
}

func (t *DeepEqualsTest) DoubleSliceValue() {
	ExpectEq("TODO", "")
}
