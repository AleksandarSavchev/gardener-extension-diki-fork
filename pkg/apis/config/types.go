// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Configuration contains information about the diki extension configuration.
type Configuration struct {
	metav1.TypeMeta

	// BaseDikiConfig is the YAML content for base diki options.
	// +optional
	BaseDikiConfig *string
}
