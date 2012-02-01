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

// IdenticalTo(x) returns a matcher that matches values v such that all of the
// following hold:
//
//  *  v and x have identical types.
//
//  *  If v and x are of a reference type (slice, map, function, channel), then
//     they are either both nil or are references to the same object.
//
//  *  If v and x are not of a reference type, then it is legal to compare them
//     using the == operator, and v == x.
//
// It is illegal for x to be of struct type, or of a container type that uses
// structs as keys or elements.
func IdenticalTo(x interface{}) Matcher {
	// TODO
	return &hasSubstrMatcher{"asd"}
}
