/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/kubetrail/qq/pkg/run"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize shared secret",
	Long: `Initialize a Diffie-Hellman key exchange protocol.
It is assumed that you are able to communicate with your peer.
The init sequence will require you to first send a key to your peer
and then input the key received from your peer.
Once the key exchange handshake is done, you both will have access
to a common shared secret key, which can be used later to
encrypt and decrypt mutually communicated messages.`,
	RunE: run.Init,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
