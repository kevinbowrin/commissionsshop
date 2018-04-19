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
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/ChimeraCoder/anaconda"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
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
		var wg sync.WaitGroup
		wg.Add(1)
		go runMonitor(&wg)
		wg.Wait()
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

func checkRequired() error {
	var missing []string
	for _, v := range [...]string{"accesstoken", "accesstokensecret", "consumerkey", "consumersecret"} {
		if viper.GetString(v) == "" {
			missing = append(missing, v)
		}
	}
	if len(missing) == 1 {
		return fmt.Errorf("A required flag or config option is missing: %v", missing[0])
	} else if len(missing) > 1 {
		return fmt.Errorf("Some required flags or config options are missing: %v", strings.Join(missing, ", "))
	} else {
		return nil
	}
}

func runMonitor(wg *sync.WaitGroup) {
	defer wg.Done()

	if err := checkRequired(); err != nil {
		log.Fatal(err)
	}
	api := anaconda.NewTwitterApiWithCredentials(viper.GetString("accesstoken"),
		viper.GetString("accesstokensecret"),
		viper.GetString("consumerkey"),
		viper.GetString("consumersecret"))
	_, err := api.VerifyCredentials()
	if err != nil {
		log.Fatal("Unable to connect to Twitter API:", err)
	}
	log.Println("Connected to Twitter API.")

	api.SetLogger(anaconda.BasicLogger)
	v := url.Values{}
	v.Set("follow", "575930104")
	s := api.PublicStreamFilter(v)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		<-stop
		log.Println("Recieved Stop Signal... please wait...")
		s.Stop()
		<-stop
		log.Fatalln("OK OK I GET IT")
	}()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		wg.Add(1)
		s.Stop()
		go runMonitor(wg)
	})
	viper.WatchConfig()

	for t := range s.C {
		v, ok := t.(anaconda.Tweet)
		if ok {
			spew.Dump(v)
		}
	}

	log.Println("Done!")
}
