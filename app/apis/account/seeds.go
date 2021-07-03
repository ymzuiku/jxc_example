//+build !test

package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"sync"
	"sync/atomic"
)

var n int64 = 13840000000

// 用于创建 account seeds, 其实并不是测试
func Seeds() {
	var count int64
	kit.ORM.Table("account").Count(&count)
	if count < 100_0000 {
		for i := 0; i < 100_0000; i += 100 {
			runHundred()
		}
	}
}

func seedOne() error {
	atomic.AddInt64(&n, 1)
	phone := fmt.Sprintf("%v", n)
	err := signUpSendCode(&sendCodeBody{Phone: phone})
	if err != nil {
		return err
	}
	_, err = signUp(&signUpBody{Phone: phone, Name: "seed名称" + phone, Company: "seed企业" + phone, People: "10~20", Password: "123456", Code: "999999"})

	if err != nil {
		return err
	}
	return nil
}

func runHundred() {
	var wg sync.WaitGroup
	for u := 0; u < 100; u++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := seedOne()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}
