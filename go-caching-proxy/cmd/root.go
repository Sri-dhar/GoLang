package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sri-dhar/GoLang/go-caching-proxy/internal/proxy-server/proxy"
	"github.com/sri-dhar/GoLang/go-caching-proxy/internal/app"
)

var (
  Port       string
  Origin     string
  ClearCache bool
  Proxy      *proxy.Proxy
)

var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Short: "Caching server that caches responses from other server",
	Long: `Caching proxy server that forwards requests to the actual server 
and caches the responses. If the same request is made again, 
it will return the cached response instead of forwarding the request to the server.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "1.0.0",


	Run: func(cmd *cobra.Command, args []string) {
		var port int
		var err error

		if Port != "" {
			port, err = strconv.Atoi(Port)
			if err != nil {
				fmt.Printf("Invalid port number: %v\n", err)
				return
			}
		}

		Proxy = proxy.NewProxy(Origin, ClearCache)

		server := &app.Server{
			Port:       port,
			Origin:     Origin,
			ClearCache: ClearCache,
			Proxy:      Proxy,
		}

		server.StartServer()
	},
}

func Execute(){
  err := rootCmd.Execute()
  if err != nil{
    os.Exit(1)
  }
}

func init(){
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.caching-proxy.yaml)")
  root.Cmd.PersistentFlags().BoolVarP(&ClearCache, "clear-cache", "c", false, "Clear the cache of the proxy server")
  root.Cmd.PersistentFlags().StringVarP(&Port, "port","p", "", "Port number to start the cache server")
  root.Cmd.PersistentFlags().StringVarP(&Origin, "origin", "o", "", "Origin url of the server to cache the response")
  root.Cmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
}
