// Copyright © 2018 Raul Sampedro
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

package cli

import (
	"strings"

	"github.com/rsrdesarrollo/SaSSHimi/server"
	"github.com/rsrdesarrollo/SaSSHimi/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var idFile string
var osArch string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server <user@host:port|host_id>",
	Short: "Run local server to create tunnels",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokens := strings.Split(args[0], "@")

		user, remoteHost := strings.Join(tokens[:len(tokens)-1], "@"), tokens[len(tokens)-1]

		subv := viper.Sub(remoteHost)

		if subv == nil {
			subv = viper.GetViper()
		}

		utils.Logger.Debug("Parsed User:", user)
		utils.Logger.Debug("Parsed Remote Host:", remoteHost)
		utils.Logger.Debug("Parsed OSARCH:", osArch)

		if user != "" {
			subv.Set("User", user)
		}

		subv.SetDefault("RemoteHost", remoteHost)
		subv.SetDefault("PrivateKey", idFile)
		subv.SetDefault("OSArch", osArch)

		server.Run(subv, bindAddress, verboseLevel)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVar(&bindAddress, "bind", "127.0.0.1:1080", "Set local bind address and port")
	serverCmd.Flags().StringVarP(&idFile, "identity_file", "i", "", "Path to private key")
	serverCmd.Flags().StringVarP(&osArch, "os_arch", "o", "", "Server OS_ARCH if different from client (linux_amd64, macos_arm)")
}
