// Copyright © 2017 G. Hussain Chinoy
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ghchinoy/rwctl/apis"
	"github.com/ghchinoy/rwctl/control"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists apis",
	Long: `lists apis available on the API platform`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfgmap map[string]interface{}
		var config control.Configuration

		if viper.IsSet(profile) {
			cfgmap = viper.GetStringMap(profile)
		}

		config, err := control.ViperToConfiguration(cfgmap, debug)
		if err != nil {
			fmt.Println("Error translating config", err.Error())
		} else {
			apis.APIList(config, debug)
		}
	},
}

func init() {
	apisCmd.AddCommand(listCmd)

}
