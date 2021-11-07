// Copyright (c) 2021 Jing-Ying Chen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bdd

import "testing"

func TestParams(t *testing.T) {
	params := Params([]string{"a: 123", "b: abc"})
	if params.Get("a") != "123" || params.Get("b") != "abc" {
		t.Fatal("error")
	}

	if params.Get("c", "xxx") != "xxx" {
		t.Fatal("bad default value")
	}
}
