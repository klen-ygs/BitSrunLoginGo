package srun

import (
	"encoding/json"
	"errors"
	"github.com/Mmx233/BitSrunLoginGo/internal/config"
	"github.com/Mmx233/BitSrunLoginGo/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const acid = "5"

type Conf struct {
	//调用 API 时直接访问 https URL
	Https bool
	//登录参数，不可缺省
	LoginInfo LoginInfo
	Client    *http.Client
}

func New() *Srun {
	srun := &Srun{}
	srun.api.Init(false, "210.43.112.9", &http.Client{
		Timeout: time.Second * 5,
	})
	return srun
}

type Srun struct {
	api Api
}

func (c Srun) LoginStatus() (online bool, ip string, err error) {
	res, err := c.api.GetUserInfo()
	if err != nil {
		return false, "", err
	}

	errRes, ok := res["error"]
	if !ok {
		return false, "", ErrResultCannotFound
	}

	ipInterface, ok := res["client_ip"]
	if !ok {
		ipInterface, ok = res["online_ip"]
		if !ok {
			return false, "", ErrResultCannotFound
		}
	}

	ip = ipInterface.(string)
	online = errRes.(string) == "ok"
	return
}

func (c Srun) DoLogin(clientIP string) error {
	log.Debugln("正在获取 Token")

	res, err := c.api.GetChallenge(config.UserInfo.Username, clientIP)
	if err != nil {
		return errors.New("访问校园网服务器失败")
	}
	token, ok := res["challenge"]
	if !ok {
		return errors.New("您输入的用户信息无效")
	}
	tokenStr := token.(string)
	log.Debugln("token: ", tokenStr)

	log.Debugln("发送登录请求")

	info, err := json.Marshal(map[string]string{
		"username": config.UserInfo.Username,
		"password": config.UserInfo.Password,
		"ip":       clientIP,
		"acid":     acid,
		"enc_ver":  "srun_bx1",
	})
	if err != nil {
		return err
	}
	EncryptedInfo := "{SRBX1}" + Base64(XEncode(string(info), tokenStr))
	Md5Str := Md5(tokenStr)
	EncryptedMd5 := "{MD5}" + Md5Str
	EncryptedChkstr := Sha1(
		tokenStr + config.UserInfo.Username + tokenStr + Md5Str +
			tokenStr + acid + tokenStr + clientIP +
			tokenStr + "200" + tokenStr + "1" +
			tokenStr + EncryptedInfo,
	)

	for i := 0; i < 10; i++ {
		if i != 0 {
			time.Sleep(time.Second * 10)
		}

		login, err := c.loginAndTest(
			EncryptedMd5,
			clientIP,
			EncryptedInfo,
			EncryptedChkstr,
		)
		if err != nil {
			tools.Notify("登录失败，稍后重试")
			continue
		}
		if !login {
			tools.Notify("登录失败，稍后重试")
			continue
		}
		tools.Notify("usc 登录成功")
		return nil
	}

	tools.Notify("usc无法登录")
	return nil
}

func (c Srun) loginAndTest(md5 string, ip string, info string, Chkstr string) (bool, error) {
	online, _, err := c.LoginStatus()
	if online {
		return true, nil
	}

	res, err := c.api.Login(
		config.UserInfo.Username,
		md5,
		acid,
		ip,
		info,
		Chkstr,
		"200",
		"1",
	)
	if err != nil {
		return false, err
	}
	var result interface{}
	result, ok := res["error"]
	if !ok {
		return false, ErrResultCannotFound
	}
	LoginResult := result.(string)

	if LoginResult != "ok" {
		return false, errors.New(LoginResult)
	}

	time.Sleep(time.Second * 1)
	online, _, err = c.LoginStatus()
	return online, err
}

func (c Srun) DetectAcid() (string, error) {
	return c.api.DetectAcid()
}
