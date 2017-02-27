// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/koolay/econfig/config"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "econfig",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(preInit)
	RootCmd.PersistentFlags().StringVarP(&config.GlobalFlags.CfgFile, "config", "c", "", "config file (default is $HOME/.econfig.toml)")
	RootCmd.PersistentFlags().StringVarP(&config.GlobalFlags.App, "app", "p", "", "process special app")
	RootCmd.PersistentFlags().StringVarP(&config.GlobalFlags.Store, "store", "s", "mysql", "data type, supports: mysql, redis,vault")
	RootCmd.PersistentFlags().BoolVarP(&config.GlobalFlags.Verbose, "verbose", "v", false, "log verbose mode")

	// apps := viper.Get("apps").(map[string]interface{})
	// for appName, appProps := range apps {
	// fmt.Println("app:", appName, appProps)
	// }
}

// initConfig reads in config file and ENV variables if set.
func preInit() {

}
