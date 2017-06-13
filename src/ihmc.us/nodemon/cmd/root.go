// Copyright © 2017 Enrico Casini <ecasini@ihmc.us>
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
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"ihmc.us/nodemon/config"
	"ihmc.us/nodemon/broker"
	"ihmc.us/nodemon/sensors/netsensor"
	"ihmc.us/nodemon/sensors/mockets"
	"ihmc.us/nodemon/sensors/disservice"
	"ihmc.us/nodemon/http"
)

const (
	Authors    = "Enrico Casini <ecasini@ihmc.us>";
	Name       = "nodemon"
	DescrShort = "A sensor data collector and fusion service for IoT"
	DescrLong  = "NodeMon is a sensor fusion service that collects, reshapes and distributes data from local sensors.\n" +
		"NodeMon uses NATS (http://nats.io/) to distribute the aggregated data to potential clients along the network."
	DescrSignature = Name + " - " + DescrShort
	TAG            = "NodeMon"
)

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Use:   Name,
	Short: DescrShort,
	Long:  DescrLong,
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

//config manager
var cfg config.Manager

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	cfg = config.Manager{}

	RootCmd.PersistentFlags().StringVar(&cfg.File, "config", "",
		"config file (default is $HOME/."+Name+".yaml)")
	RootCmd.PersistentFlags().StringVar(&cfg.NATSHost, "nats-address", broker.DefaultHost, "NATS server address")
	RootCmd.PersistentFlags().Uint16Var(&cfg.NATSPort, "nats-port", broker.DefaultPort, "NATS server port")
	RootCmd.PersistentFlags().Uint16Var(&cfg.NetSensorPort, "netsensor-port", netsensor.DefaultPort, "NetSensor statistics port")
	RootCmd.PersistentFlags().Uint16Var(&cfg.MocketsSensorPort, "mockets-port", mockets.DefaultPort, "Mockets statistics port")
	RootCmd.PersistentFlags().Uint16Var(&cfg.DisServiceSensorPort, "disservice-port", disservice.DefaultPort, "DisService statistics port")
	RootCmd.PersistentFlags().Uint16Var(&cfg.HTTPServerPort, "http-port", http.DefaultPort, "HTTP server port")
	RootCmd.PersistentFlags().StringVar(&cfg.LogLevel, "log-level", log.DebugLevel.String(), "Log level. Values: debug, info, warn, error, fatal, panic")
	RootCmd.PersistentFlags().BoolVar(&cfg.NetSensorLogDebug, "netsensor-debug", false, "Enable detailed debug log level for NetSensor")
	RootCmd.PersistentFlags().BoolVar(&cfg.MocketsLogDebug, "mockets-debug", false, "Enable detailed debug log level for MocketsSensor")
	RootCmd.PersistentFlags().BoolVar(&cfg.DisServiceLogDebug, "disservice-debug", false, "Enable detailed debug log level for DisServiceSensor")
	RootCmd.PersistentFlags().BoolVar(&cfg.HTTPServerLogDebug, "http-debug", false, "Enable detailed debug log level for HTTP Server")
	RootCmd.PersistentFlags().BoolVar(&cfg.NATSLogDebug, "nats-debug", true, "Enable detailed debug log level for NATS Server")
	RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", Authors)
	viper.SetDefault("license", "apache")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfg.File != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfg.File)
	}

	viper.SetConfigName("." + Name) // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}
