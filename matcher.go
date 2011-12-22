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

package oglematchers

// A MatchResult is an integer equal to one of the MATCH_* constants below.
// Matchers use a tri-state logic in order to make the semantics of matchers
// that wrap other matchers make more sense. The constants below represent the
// three values that a matcher may return.
type MatchResult int

const (
	// MATCH_FALSE indicates that the supplied value didn't match. For example,
	// IsNil would return this when presented with any non-nil value, and
	// GreaterThan(17) would return this when presented with 16.
	MATCH_FALSE MatchResult = 0

	// MATCH_TRUE indicates that the supplied value did match. For example, IsNil
	// would return this when presented with nil, and GreaterThan(17) would
	// return this when presented with 19.
	MATCH_TRUE MatchResult = 1

	// MATCH_UNDEFINED indicates that the matcher doesn't process values of the
	// supplied type, or otherwise doesn't know how to handle the value. This is
	// akin to returning MATCH_FALSE, except that wrapper matchers should
	// propagagate undefined values.
	//
	// For example, if GreaterThan(17) returned MATCH_FALSE for the value "taco",
	// then Not(GreaterThan(17)) would return MATCH_TRUE. This is technically
	// correct, but is surprising and may mask failures where the wrong sort of
	// matcher is accidentally used. Instead, GreaterThan(17) can return
	// MATCH_UNDEFINED, which will be propagated by Not().
	MATCH_UNDEFINED MatchResult = -1
)

// A Matcher is some predicate implicitly defining a set of values that it
// matches. For example, GreaterThan(17) matches all numeric values greater
// than 17, and HasSubstr("taco") matches all strings with the substring
// "taco".
type Matcher interface {
	// Matches returns a MatchResult indicating whether the supplied value
	// belongs to the set defined by the matcher.
	//
	// If the result is MATCH_FALSE or MATCH_UNDEFINED, it may additionally
	// return an error describing why the value doesn't match. The error text is
	// a relative clause that is suitable for being placed after the value. For
	// example, a predicate that matches strings with a particular substring may,
	// when presented with a numerical value, return the following error text:
	//
	//     "which is not a string"
	//
	// Then the failure message may look like:
	//
	//     Expected: has substring "taco"
	//     Actual:   17, which is not a string
	//
	Matches(candidate interface{}) (MatchResult, error)

	// Description returns a string describing the property that values matching
	// this matcher have, as a verb phrase where the subject is the value. For
	// example, "is greather than 17" or "has substring "taco"".
	Description() string
}
