// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.  See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Spencer Kimball (spencer.kimball@gmail.com)

package util

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

// TestErrorf verifies the pass through to fmt.Errorf as well as
// file/line prefix.
func TestErrorf(t *testing.T) {
	err := Errorf("foo: %d %f", 1, 3.14)
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("could not get runtime info for test")
	}
	expected := fmt.Sprintf("%sfoo: 1 3.140000", fmt.Sprintf(errorPrefixFormat, filepath.Base(file), line-1))
	if expected != err.Error() {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

// TestError verifies the pass through to fmt.Error as well as
// file/line prefix.
func TestError(t *testing.T) {
	err := Error("foo ", 1, 3.14)
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("could not get runtime info for test")
	}
	expected := fmt.Sprintf("%sfoo 1 3.14", fmt.Sprintf(errorPrefixFormat, filepath.Base(file), line-1))
	if expected != err.Error() {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}
