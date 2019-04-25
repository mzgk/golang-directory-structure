package main

import (
	"zoo/animals" // モジュール名（go mod init モジュール名）/パッケージ名で各
)

func main() {
	println(AppName())

	println(animals.ElephantFeed())
	println(animals.MonkeyFeed())
	println(animals.RabbitFeed())
}
