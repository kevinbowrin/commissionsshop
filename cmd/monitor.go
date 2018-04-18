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

	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor Twitter's Steaming API for new posts",
	Long: `Using the Twitter Streaming API, find posts or retweets from
whitelisted accounts which mention 'commissions open'.`,
	Run: func(cmd *cobra.Command, args []string) {

		required := [...]string{"accesstoken", "accesstokensecret", "consumerkey", "consumersecret"}
		for _, v := range required {
			if viper.GetString(v) == "" {
				log.Fatalln("A required flag or config option is missing:", v)
			}
		}
		api := anaconda.NewTwitterApiWithCredentials(viper.GetString("accesstoken"),
			viper.GetString("accesstokensecret"),
			viper.GetString("consumerkey"),
			viper.GetString("consumersecret"))
		ok, err := api.VerifyCredentials()
		log.Println(ok, err)
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().String("accesstoken", "", "Twitter API Access Token.")
	monitorCmd.Flags().String("accesstokensecret", "", "Twitter API Access Token Secret.")
	monitorCmd.Flags().String("consumerkey", "", "Twitter API Consumer Key.")
	monitorCmd.Flags().String("consumersecret", "", "Twitter API Consumer Secret.")
	viper.BindPFlags(monitorCmd.Flags())
}
