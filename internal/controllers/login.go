package controllers

import (
	"github.com/Mmx233/BitSrunLoginGo/pkg/srun"
)

// Login 登录逻辑
func Login() error {
	// 登录配置初始化

	srunClient := srun.New()

	online, ip, err := srunClient.LoginStatus()
	if err != nil {
		return err
	}

	if online {
		return nil
	}

	srunClient.DoLogin(ip)

	return nil
}
