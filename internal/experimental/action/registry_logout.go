/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package action

import (
	"io"

	"helm.sh/helm/v3/pkg/action"
)

// RegistryLogout performs a registry login operation.
type RegistryLogout struct {
	cfg *action.Configuration
}

// NewRegistryLogout creates a new RegistryLogout object with the given configuration.
func NewRegistryLogout(cfg *action.Configuration) *RegistryLogout {
	return &RegistryLogout{
		cfg: cfg,
	}
}

// Run executes the registry logout operation
func (a *RegistryLogout) Run(out io.Writer, hostname string) error {
	return a.cfg.RegistryClient.Logout(hostname)
}
