package tools

import "github.com/go-toast/toast"

func Notify(msg string) {
	notify := toast.Notification{
		AppID:   "usc-conn",
		Message: msg,
	}
	notify.Push()
}
