// Copyright © 2018 Kevin Bowrin <kjbowrin@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"testing"

	"github.com/spf13/viper"
)

func TestCheckRequired(t *testing.T) {
	//None set
	viper.Reset()
	if err := checkRequired(); err == nil {
		t.Error("checkRequired() returns nil even when no required config options are set.")
	}
	//Some set
	viper.Reset()
	viper.Set("accesstoken", "a")
	viper.Set("consumersecret", "a")
	if err := checkRequired(); err == nil {
		t.Error("checkRequired() returns nil even when only some of the required config options are set.")
	}
	//All set
	viper.Set("accesstokensecret", "a")
	viper.Set("consumerkey", "a")
	if err := checkRequired(); err != nil {
		t.Error("checkRequired() returns error even when all required config options are set:", err)
	}
}
