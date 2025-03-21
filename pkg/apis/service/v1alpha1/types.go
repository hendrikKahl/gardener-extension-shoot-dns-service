// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSConfig configuration resource
type DNSConfig struct {
	metav1.TypeMeta `json:",inline"`

	// DNSProviderReplication contains enablement for replication of DNSProviders from shoot cluster to control plane
	// +optional
	DNSProviderReplication *DNSProviderReplication `json:"dnsProviderReplication,omitempty"`

	// Providers is a list of additional DNS providers that shall be enabled for this shoot cluster.
	// The primary ("external") provider at `spec.dns.provider` is added automatically
	// +optional
	Providers []DNSProvider `json:"providers,omitempty"`

	// SyncProvidersFromShootSpecDNS is an optional flag for migrating and synchronising the providers given in the
	// shoot manifest at section `spec.dns.providers`. If true, any direct changes on the `providers` section
	// are overwritten with the content of section `spec.dns.providers`.
	// +optional
	SyncProvidersFromShootSpecDNS *bool `json:"syncProvidersFromShootSpecDNS,omitempty"`
}

// DNSProviderReplication contains enablement for replication of DNSProviders from shoot cluster to control plane
type DNSProviderReplication struct {
	// Enabled if true, the replication of DNSProviders from shoot cluster to the control plane is enabled
	Enabled bool `json:"enabled"`
}

// DNSProvider contains information about a DNS provider.
type DNSProvider struct {
	// Domains contains information about which domains shall be included/excluded for this provider.
	// +optional
	Domains *DNSIncludeExclude `json:"domains,omitempty"`
	// SecretName is a name of a secret containing credentials for the stated domain and the
	// provider.
	// +optional
	SecretName *string `json:"secretName,omitempty"`
	// Type is the DNS provider type.
	// +optional
	Type *string `json:"type,omitempty"`
	// Zones contains information about which hosted zones shall be included/excluded for this provider.
	// +optional
	Zones *DNSIncludeExclude `json:"zones,omitempty"`
}

// DNSIncludeExclude contains information about which domains shall be included/excluded.
type DNSIncludeExclude struct {
	// Include is a list of domains that shall be included.
	// +optional
	Include []string `json:"include,omitempty"`
	// Exclude is a list of domains that shall be excluded.
	// +optional
	Exclude []string `json:"exclude,omitempty"`
}
