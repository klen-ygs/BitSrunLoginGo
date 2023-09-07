package main

import (
	"flag"
	"fmt"
	"github.com/Mmx233/BitSrunLoginGo/internal/config"
	"github.com/Mmx233/BitSrunLoginGo/internal/controllers"
	"github.com/Mmx233/BitSrunLoginGo/tools"
	"time"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "f", "", "配置文件")
	flag.Parse()

	config.LoadConfig(configPath)
	fmt.Println(config.UserInfo)
	for i := 0; i < 50; i++ {
		if tools.TestInUSCButLogout() {
			if err := controllers.Login(); err != nil {
				continue
			}
			return
		}
		time.Sleep(time.Second * 3)
	}

}
