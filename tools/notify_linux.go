package tools

import "os/exec"

func Notify(msg string) {
	exec.Command("notify-send", "-a", "usc-conn", msg).Run()
}
