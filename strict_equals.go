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

package oglematchers

import (
)

// StrictEquals(x) returns a matcher that matches values v such that v == x.
// The matcher follows the same strict rules about types that the built-in ==
// operator does, and does not do any type folding like Equals does.
//
// For example:
//
//     type stringAlias string
//
//     ExpectThat("taco", Equals(stringAlias("taco")))        // Passes
//     ExpectThat("taco", StrictEquals(stringAlias("taco")))  // Fails
//     ExpectThat(stringAlias("taco"), StrictEquals("taco"))  // Fails
//
//     ExpectThat(int(17), Equals(int8(17)))                  // Passes
//     ExpectThat(int(17), StrictEquals(int8(17)))            // Fails
//
func StrictEquals(x interface{}) Matcher {
	// TODO
	return &hasSubstrMatcher{s}
}
