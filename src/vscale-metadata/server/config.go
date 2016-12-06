package server

import (
	"fmt"
	"net"
  "encoding/json"
  "github.com/spf13/viper"
)

type Config struct {
	Port          int
	Host          string
  ApiKey        string
}

func getPort(port int) string {
	return fmt.Sprintf("%d", port)
}

func (c Config) getAddr() string {
	return net.JoinHostPort(c.Host, getPort(c.Port))
}

func (c Config) getUserData() string{
  viper.SetConfigName("metadata")
  viper.AddConfigPath("/etc/metadata/")
  viper.AddConfigPath("$HOME/.metadata")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig() 
  if err != nil {
      return ""
  }
  user_data, _ := json.Marshal(viper.AllSettings())
  return string(user_data)
}
