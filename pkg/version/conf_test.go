// Copyright 2016 CNI authors
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

package version_test

import (
	"fmt"
	"github.com/containernetworking/cni/pkg/version"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decoding the version of network config", func() {
	var (
		decoder     *version.ConfigDecoder
		configBytes []byte
	)

	BeforeEach(func() {
		decoder = &version.ConfigDecoder{}
		configBytes = []byte(`{ "cniVersion": "4.3.2" }`)
	})

	Context("when the version is explicit", func() {
		It("returns the version", func() {
			version, err := decoder.Decode(configBytes)
			Expect(err).NotTo(HaveOccurred())

			Expect(version).To(Equal("4.3.2"))
		})
	})

	Context("when the version is not present in the config", func() {
		BeforeEach(func() {
			configBytes = []byte(`{ "not-a-version-field": "foo" }`)
		})

		It("assumes the config is version 0.1.0", func() {
			version, err := decoder.Decode(configBytes)
			Expect(err).NotTo(HaveOccurred())

			Expect(version).To(Equal("0.1.0"))
		})
	})

	Context("when the config data is malformed", func() {
		BeforeEach(func() {
			configBytes = []byte(`{{{`)
		})

		It("returns a useful error", func() {
			_, err := decoder.Decode(configBytes)
			Expect(err).To(MatchError(HavePrefix(
				"decoding version from network config: invalid character",
			)))
		})
	})
})

func TestAa(t *testing.T) {
	a := []byte{123, 34, 99, 110, 105, 86, 101, 114, 115, 105, 111, 110, 34, 58, 34, 49, 46, 48, 46, 48, 34, 44, 34, 107, 117, 98, 101, 114, 110, 101, 116, 101, 115, 34, 58, 123, 34, 99, 110, 105, 95, 98, 105, 110, 95, 100, 105, 114, 34, 58, 34, 47, 118, 97, 114, 47, 108, 105, 98, 47, 114, 97, 110, 99, 104, 101, 114, 47, 107, 51, 115, 47, 100, 97, 116, 97, 47, 99, 117, 114, 114, 101, 110, 116, 47, 98, 105, 110, 34, 44, 34, 101, 120, 99, 108, 117, 100, 101, 95, 110, 97, 109, 101, 115, 112, 97, 99, 101, 115, 34, 58, 91, 34, 105, 115, 116, 105, 111, 45, 115, 121, 115, 116, 101, 109, 34, 44, 34, 107, 117, 98, 101, 45, 115, 121, 115, 116, 101, 109, 34, 93, 44, 34, 107, 117, 98, 101, 99, 111, 110, 102, 105, 103, 34, 58, 34, 47, 118, 97, 114, 47, 108, 105, 98, 47, 114, 97, 110, 99, 104, 101, 114, 47, 107, 51, 115, 47, 97, 103, 101, 110, 116, 47, 101, 116, 99, 47, 99, 110, 105, 47, 110, 101, 116, 46, 100, 47, 90, 90, 90, 45, 105, 115, 116, 105, 111, 45, 99, 110, 105, 45, 107, 117, 98, 101, 99, 111, 110, 102, 105, 103, 34, 125, 44, 34, 108, 111, 103, 95, 108, 101, 118, 101, 108, 34, 58, 34, 100, 101, 98, 117, 103, 34, 44, 34, 108, 111, 103, 95, 117, 100, 115, 95, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 47, 118, 97, 114, 47, 114, 117, 110, 47, 105, 115, 116, 105, 111, 45, 99, 110, 105, 47, 108, 111, 103, 46, 115, 111, 99, 107, 34, 44, 34, 110, 97, 109, 101, 34, 58, 34, 99, 98, 114, 48, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 105, 115, 116, 105, 111, 45, 99, 110, 105, 34, 125}
	//fmt.Println(string(a))
	decoder := version.ConfigDecoder{}
	m, n := decoder.Decode(a)
	fmt.Println(m)
	fmt.Println(n)
}
