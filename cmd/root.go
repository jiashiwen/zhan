/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"zhan/asset"
	"zhan/commons"
	"zhan/qigua"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zhan",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		gua, zhigua, guamun := qigua.GetGuaAndZhigua()

		//result := qigua.YiGua(gua).String() + "\n" + qigua.YiGua(zhigua).String() + "\n"

		fmt.Println(guamun)
		fmt.Println("正卦: ", qigua.YiGua(gua).Rune(), " ", qigua.YiGua(gua).Name(), " ", qigua.YiGua(gua).FullName())
		fmt.Println("支卦: ", qigua.YiGua(zhigua).Rune(), " ", qigua.YiGua(zhigua).Name(), " ", qigua.YiGua(zhigua).FullName())

		guacontent, err := asset.Asset("resources/" + qigua.YiGua(gua).FileName())
		if err != nil {
			fmt.Println(err)
		}
		zhiguacontent, err := asset.Asset("resources/" + qigua.YiGua(zhigua).FileName())
		if err != nil {
			fmt.Println(err)
		}
		guacol := commons.Line2Column(string(guacontent), 25)
		zhiguacol := commons.Line2Column(string(zhiguacontent), 25)

		for _, v := range guacol {
			for i := len(v) - 1; i >= 0; i-- {
				fmt.Print(string(v[i]))
			}
			fmt.Println("")
		}

		fmt.Println("\n\n\n")
		for _, v := range zhiguacol {
			for i := len(v) - 1; i >= 0; i-- {
				fmt.Print(string(v[i]))
			}
			fmt.Println("")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zhan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".zhan" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".zhan")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
