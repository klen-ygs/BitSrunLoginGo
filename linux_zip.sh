#!bin/bash

cd cmd/usc-conn || return
go build
zip usc-conn_linux.zip usc-conn userinfo.yaml
rm usc-conn
