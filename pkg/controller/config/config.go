/*
 * Copyright 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package config

import "k8s.io/apimachinery/pkg/types"

// DNSService contains configuration for the lifecycle controller of the dns service.
var DNSService DNSServiceConfig

// DNSServiceConfig contains configuration for the dns service.
type DNSServiceConfig struct {
	SeedID                    string
	DNSClass                  string
	RemoteDefaultDomainSecret *types.NamespacedName
	ManageDNSProviders        bool
	ReplicateDNSProviders     bool
	OwnerDNSActivation        bool
}
