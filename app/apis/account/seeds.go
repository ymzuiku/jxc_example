//+build !test

package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"sync"
)

// 用于创建 account seeds, 其实并不是测试
func Seeds() {
	for i := 0; i < 100_0000; i += 100 {
		runHundred()
	}
}

func runHundred() {
	var wg sync.WaitGroup
	for u := 0; u < 100; u++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := seedOne(); err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

func seedOne() error {
	// atomic.AddInt64(&n, 1)
	phone := fmt.Sprintf("9%v", kit.RandomCodeBase(30))
	if err := registerSendCode(sendCodeBody{Phone: phone}); err != nil {
		return err
	}

	if _, err := registerCompany(registerCompanyBody{Phone: phone, Name: "seed名称" + phone, Company: "seed企业" + phone, People: 50, Password: "123456", Code: "999999"}); err != nil {
		return err
	}

	return nil
}
