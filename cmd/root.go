// Copyright © 2018 Kevin Bowrin <kjbowrin@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// EnvPrefix is the prefix for the environment variables.
	EnvPrefix string = "commissionsshop"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "commissionsshop",
	Short: "Generate a website from Twitter 'commissions open!' messages.",
	Long: `commissionsshop is an application which looks for 'commissions open'
messages from a list of followed users. If those users send a message which 
says 'commissions open' or if they retweet a message from another artist which 
says 'commissions open', they will be added to the website.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file location.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err == nil {
			log.Println("Using config file:", viper.ConfigFileUsed())
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				log.Println("Config file changed:", e.Name)
			})
		} else {
			log.Fatalln("Unable to parse config file:", viper.ConfigFileUsed(), err)
		}
	}
}
