#!bin/bash

cd cmd/usc-conn || return
go env -w GOOS=windows
go build
zip usc-conn_win.zip usc-conn.exe userinfo.yaml
rm usc-conn.exe
go env -w GOOS=linux
