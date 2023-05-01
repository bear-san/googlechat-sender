package main

import (
	"github.com/bear-san/googlechat-sender/backend/cmd"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		cmd.Start()
	} else if args[1] == "scheduler" {
		// 予約投稿の処理用
	}
}
