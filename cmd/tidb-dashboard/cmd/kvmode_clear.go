// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pingcap-incubator/tidb-dashboard/pkg/pd"
	globalUtil "github.com/pingcap-incubator/tidb-dashboard/pkg/utils"
)

var kvModeAuthClearCmd = &cobra.Command{
	Use:   "auth-clear",
	Short: "clear tikv mode auth secret key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pd.NewEtcdClientNoLC(cfg.CoreConfig)
		if err != nil {
			fmt.Println("Failed to create etcdClient")
			os.Exit(1)
		}

		if globalUtil.ClearKvModeAuthKey(client) != nil {
			fmt.Println("Failed to clear kv mode auth secret key")
			os.Exit(1)
		}
	},
}

func init() {
	kvModeCmd.AddCommand(kvModeAuthClearCmd)
}