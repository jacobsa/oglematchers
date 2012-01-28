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

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type ElementsAreTest struct {
}

func init()                     { RegisterTestSuite(&ElementsAreTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *ElementsAreTest) EmptySet() {
	m := ElementsAre()
	ExpectEq("elements are: []", m.Description())

	var c []interface{}
	var res bool
	var err error

	// No candidates.
	c = []interface{}{}
	res, err = m.Matches(c)
	ExpectTrue(res)
	ExpectEq(nil, err)

	// One candidate.
	c = []interface{}{17}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 1"))
}

func (t *ElementsAreTest) OneMatcher() {
	m := ElementsAre(LessThan(17))
	ExpectEq("elements are: [less than 17]", m.Description())

	var c []interface{}
	var res bool
	var err error

	// No candidates.
	c = []interface{}{}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 0"))

	// Matching candidate.
	c = []interface{}{16}
	res, err = m.Matches(c)
	ExpectTrue(res)
	ExpectEq(nil, err)

	// Non-matching candidate.
	c = []interface{}{19}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectNe(nil, err)

	// Two candidates.
	c = []interface{}{17, 19}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 2"))
}

func (t *ElementsAreTest) OneValue() {
	m := ElementsAre(17)
	ExpectEq("elements are: [17]", m.Description())

	var c []interface{}
	var res bool
	var err error

	// No candidates.
	c = []interface{}{}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 0"))

	// Matching int.
	c = []interface{}{int(17)}
	res, err = m.Matches(c)
	ExpectTrue(res)
	ExpectEq(nil, err)

	// Matching float.
	c = []interface{}{float32(17)}
	res, err = m.Matches(c)
	ExpectTrue(res)
	ExpectEq(nil, err)

	// Non-matching candidate.
	c = []interface{}{19}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectNe(nil, err)

	// Two candidates.
	c = []interface{}{17, 19}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 2"))
}

func (t *ElementsAreTest) MultipleElements() {
	m := ElementsAre("taco", LessThan(17))
	ExpectEq("elements are: [taco, less than 17]", m.Description())

	var c []interface{}
	var res bool
	var err error

	// One candidate.
	c = []interface{}{17}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 1"))

	// Both matching.
	c = []interface{}{"taco", 16}
	res, err = m.Matches(c)
	ExpectTrue(res)
	ExpectEq(nil, err)

	// First non-matching.
	c = []interface{}{"burrito", 16}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, Error(Equals("whose element 0 doesn't match")))

	// Second non-matching.
	c = []interface{}{"taco", 17}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, Error(Equals("whose element 1 doesn't match")))

	// Three candidates.
	c = []interface{}{"taco", 17, 19}
	res, err = m.Matches(c)
	ExpectFalse(res)
	ExpectThat(err, HasSubstr("length 3"))
}

func (t *ElementsAreTest) NonFatalError() {
}

func (t *ElementsAreTest) FatalError() {
}
