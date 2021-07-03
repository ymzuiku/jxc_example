package main

import (
	"gewu_jxc/app/apis/account"
	"gewu_jxc/app/kit"
)

func main() {
	kit.InitTest()
	account.Seeds()
}
