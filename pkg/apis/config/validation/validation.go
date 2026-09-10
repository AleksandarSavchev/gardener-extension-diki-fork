// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/gardener/gardener-extension-diki/pkg/apis/config"
)

// ValidateConfiguration validates the passed configuration instance.
func ValidateConfiguration(cfg *config.Configuration) field.ErrorList {
	allErrs := field.ErrorList{}

	if cfg.BaseDikiConfig != nil {
		if len(*cfg.BaseDikiConfig) == 0 {
			allErrs = append(allErrs, field.Required(field.NewPath("baseDikiConfig"), "must not be empty when specified"))
		}
	}

	return allErrs
}
