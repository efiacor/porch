// Copyright 2022 Google LLC
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

package main

import (
	"context"
	"os"

	"github.com/nephio-project/porch/third_party/GoogleContainerTools/kpt-functions-sdk/go/kfn/commands"
	"github.com/spf13/cobra"
)

func main() {
	var err error
	cmd := &cobra.Command{
		Use:   "kfn",
		Short: "a CLI tool to edit your own KRM functions with ease",
	}
	ctx := context.Background()
	cmd.AddCommand(commands.NewInitRunner(ctx).Command)
	cmd.AddCommand(commands.NewBuildRunner(ctx).Command)
	err = cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
