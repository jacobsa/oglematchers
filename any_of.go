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

package ogletest

import (
)

// AnyOf accepts a set of values S and returns a matcher that follows the
// algorithm below:
//
//  1. If there exists a value m in S such that m implements the Matcher
//     interface and m matches c, return MATCH_TRUE.
//
//  2. Otherwise, if there exists a value v in S such that v does not implement
//     the Matcher interface and the matcher Equals(v) matches c, return
//     MATCH_TRUE.
//
//  3. Otherwise, if there is a value m in S such that m implements the Matcher
//     interface and m returns MATCH_UNDEFINED for c, return MATCH_UNDEFINED
//     with that matcher's error message.
//
//  4. Otherwise, return  MATCH_FALSE.
//
// This is akin to a logical OR operation for matchers.
func AnyOf(vals ...interface{}) Matcher {
	return Equals(0)
}
