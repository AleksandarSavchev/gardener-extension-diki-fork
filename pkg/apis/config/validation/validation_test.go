// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/gardener/gardener-extension-diki/pkg/apis/config"
	. "github.com/gardener/gardener-extension-diki/pkg/apis/config/validation"
)

var _ = Describe("ValidateConfiguration", func() {
	var cfg *config.Configuration

	BeforeEach(func() {
		cfg = &config.Configuration{}
	})

	It("should allow an empty configuration", func() {
		errs := ValidateConfiguration(cfg)
		Expect(errs).To(BeEmpty())
	})

	It("should allow a valid baseDikiConfig", func() {
		data := "providers:\n- id: managedk8s\n"
		cfg.BaseDikiConfig = &data
		errs := ValidateConfiguration(cfg)
		Expect(errs).To(BeEmpty())
	})

	It("should reject baseDikiConfig with empty data", func() {
		data := ""
		cfg.BaseDikiConfig = &data
		errs := ValidateConfiguration(cfg)
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Type).To(Equal(field.ErrorTypeRequired))
		Expect(errs[0].Field).To(Equal("baseDikiConfig"))
	})
})
