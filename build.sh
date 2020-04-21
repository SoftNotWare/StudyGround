# !/bin/sh

go build -ldflags "-s -w"
upx --brute  study-ground