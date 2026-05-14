package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"pomeloServe/common/config"
	"pomeloServe/common/metrics"
	"pomeloServe/framework/game"
	"pomeloServe/hall/app"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hall",
	Short: "hall 大厅相关的处理",
	Long:  `hall 大厅相关的处理`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	},
}

var (
	configFile    string
	gameConfigDir string
	serverId      string
)

//var configFile = flag.String("config", "application.yml", "config file")

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", "application.yml", "app config yml file")
	rootCmd.Flags().StringVar(&gameConfigDir, "gameDir", "../config", "game config dir")
	rootCmd.Flags().StringVar(&serverId, "serverId", "", "app server id， required")
	_ = rootCmd.MarkFlagRequired("serverId")
}

func main() {
	//1.加载配置
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	config.InitConfig(configFile)
	game.InitConfig(gameConfigDir)
	//2.启动监控
	go func() {
		err := metrics.Serve(fmt.Sprintf("0.0.0.0:%d", config.Conf.MetricPort))
		if err != nil {
			panic(err)
		}
	}()
	//3.连接nats服务 并进行订阅
	err := app.Run(context.Background(), serverId)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
