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

func TestNewDeleteBuildCmd_Ok(t *testing.T) {
	realCmd := NewDeleteBuildOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewDeleteBuildCmd(realCmd, FlagsFactory{}))
	cmd.ExpectBasicConfig()
	cmd.Execute([]string{
		"-n", "test-namespace",
		"-b", "test-build",
	})
	cmd.ExpectReachesExecution()

	DeepEqual(t, realCmd.BuildFlags,
		BuildFlags{NamespaceFlags{"test-namespace"}, "test-build"})
}

func TestNewDeleteBuildCmd_OkLongFlagNames(t *testing.T) {
	realCmd := NewDeleteBuildOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewDeleteBuildCmd(realCmd, FlagsFactory{}))
	cmd.Execute([]string{
		"--namespace", "test-namespace",
		"--build", "test-build",
	})
	cmd.ExpectReachesExecution()

	DeepEqual(t, realCmd.BuildFlags,
		BuildFlags{NamespaceFlags{"test-namespace"}, "test-build"})
}

func TestNewDeleteBuildCmd_RequiredFlags(t *testing.T) {
	realCmd := NewDeleteBuildOptions(nil, newDepsFactory())
	cmd := NewTestCmd(t, NewDeleteBuildCmd(realCmd, FlagsFactory{}))
	cmd.Execute([]string{})
	cmd.ExpectRequiredFlags([]string{"build"})
}
