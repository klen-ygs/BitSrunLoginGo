package tools

import "testing"

func TestNotify(t *testing.T) {
	Notify("Hello")
}

func TestTestInUSCButLogout(t *testing.T) {
	t.Log(TestInUSCButLogout())
}
