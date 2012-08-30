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

	// Int alias candidate.
	type intAlias int
	err = m.Matches(intAlias(x))
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("intAlias")))
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

func (t *DeepEqualsTest) WrongTypeCandidateWithNilLiteralValue() {
	m := DeepEquals(nil)

	var err error

	// String candidate.
	err = m.Matches("taco")
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("string")))
	ExpectThat(err, Error(HasSubstr("TODO")))

	// Nil slice candidate.
	err = m.Matches([]byte(nil))
	AssertNe(nil, err)
	ExpectTrue(isFatal(err))
	ExpectThat(err, Error(HasSubstr("type")))
	ExpectThat(err, Error(HasSubstr("[]byte")))
	ExpectThat(err, Error(HasSubstr("TODO")))
}

func (t *DeepEqualsTest) NilLiteralValue() {
	m := DeepEquals(nil)
	ExpectEq("deep equals: nil", m.Description())

	var c interface{}
	var err error

	// Nil literal candidate.
	c = nil
	err = m.Matches(c)
	ExpectEq(nil, err)
}

func (t *DeepEqualsTest) IntValue() {
	m := DeepEquals(int(17))
	ExpectEq("deep equals: 17", m.Description())

	var c interface{}
	var err error

	// Matching int.
	c = int(17)
	err = m.Matches(c)
	ExpectEq(nil, err)

	// Non-matching int.
	c = int(18)
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))
}

func (t *DeepEqualsTest) SliceValue() {
	x := []byte{17, 19}
	m := DeepEquals(x)
	ExpectEq("deep equals: [17 19]", m.Description())

	var c []byte
	var err error

	// Matching.
	c = make([]byte, len(x))
	AssertEq(len(x), copy(c, x))

	err = m.Matches(c)
	ExpectEq(nil, err)

	// Prefix.
	AssertGt(len(x), 1)
	c = make([]byte, len(x)-1)
	AssertEq(len(x)-1, copy(c, x))

	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))

	// Suffix.
	c = make([]byte, len(x)+1)
	AssertEq(len(x), copy(c, x))

	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))
}

func (t *DeepEqualsTest) NilSliceValue() {
	var x []byte
	m := DeepEquals(x)
	ExpectEq("deep equals: TODO", m.Description())

	var c []byte
	var err error

	// Nil slice.
	c = []byte(nil)
	err = m.Matches(c)
	ExpectEq(nil, err)

	// Non-nil slice.
	c = []byte{}
	err = m.Matches(c)
	ExpectThat(err, Error(Equals("")))
}
