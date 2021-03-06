/*
Copyright 2018 The Knative Authors

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

package cmd_test

import (
	"testing"

	. "github.com/cppforlife/knctl/pkg/knctl/cmd"
)

func TestNewCreateServiceAccountCmd_Ok(t *testing.T) {
	realCmd := NewCreateServiceAccountOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewCreateServiceAccountCmd(realCmd, FlagsFactory{}))
	cmd.ExpectBasicConfig()
	cmd.Execute([]string{
		"-n", "test-namespace",
		"-a", "test-serv-account",
		"-s", "test-secret1",
		"-s", "test-secret2",
		"-p", "test-secret3",
		"-p", "test-secret4",
	})
	cmd.ExpectReachesExecution()

	DeepEqual(t, realCmd.ServiceAccountFlags,
		ServiceAccountFlags{NamespaceFlags{"test-namespace"}, "test-serv-account"})

	DeepEqual(t, realCmd.ServiceAccountCreateFlags, ServiceAccountCreateFlags{
		Secrets:          []string{"test-secret1", "test-secret2"},
		ImagePullSecrets: []string{"test-secret3", "test-secret4"},
	})
}

func TestNewCreateServiceAccountCmd_OkLongFlagNames(t *testing.T) {
	realCmd := NewCreateServiceAccountOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewCreateServiceAccountCmd(realCmd, FlagsFactory{}))
	cmd.Execute([]string{
		"--namespace", "test-namespace",
		"--service-account", "test-serv-account",
		"--secret", "test-secret1",
		"--secret", "test-secret2",
		"--image-pull-secret", "test-secret3",
		"--image-pull-secret", "test-secret4",
	})
	cmd.ExpectReachesExecution()

	DeepEqual(t, realCmd.ServiceAccountFlags,
		ServiceAccountFlags{NamespaceFlags{"test-namespace"}, "test-serv-account"})

	DeepEqual(t, realCmd.ServiceAccountCreateFlags, ServiceAccountCreateFlags{
		Secrets:          []string{"test-secret1", "test-secret2"},
		ImagePullSecrets: []string{"test-secret3", "test-secret4"},
	})
}

func TestNewCreateServiceAccountCmd_OkMinimum(t *testing.T) {
	realCmd := NewCreateServiceAccountOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewCreateServiceAccountCmd(realCmd, FlagsFactory{}))
	cmd.Execute([]string{
		"--namespace", "test-namespace",
		"--service-account", "test-serv-account",
	})
	cmd.ExpectReachesExecution()

	DeepEqual(t, realCmd.ServiceAccountFlags,
		ServiceAccountFlags{NamespaceFlags{"test-namespace"}, "test-serv-account"})

	DeepEqual(t, realCmd.ServiceAccountCreateFlags, ServiceAccountCreateFlags{})
}

func TestNewCreateServiceAccountCmd_RequiredFlags(t *testing.T) {
	realCmd := NewCreateServiceAccountOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewCreateServiceAccountCmd(realCmd, FlagsFactory{}))
	cmd.Execute([]string{})
	cmd.ExpectRequiredFlags([]string{"service-account"})
}
