package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/ymzuiku/gewu_jxc/internal/services/account"

	"github.com/google/uuid"
	"github.com/ymzuiku/errox"
)

var n int64 = 0

const _SEED_ONCE = 20

// 用于创建 account seeds, 其实并不是测试
func Seeds() {
	for i := 0; i < 100_0000; i += _SEED_ONCE {
		runHundred()
	}
}

func runHundred() {
	var wg sync.WaitGroup
	for u := 0; u < _SEED_ONCE; u++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := seedOne(); err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("seeded: %v\n", atomic.AddInt64(&n, _SEED_ONCE))
}

func seedOne() error {
	// atomic.AddInt64(&n, 1)
	u1 := uuid.New().String()

	phone := fmt.Sprintf("8%v", u1)
	if err := account.RegisterSendCode(account.SendCodeBody{Phone: phone}); err != nil {
		return errox.Wrap(err)
	}

	if _, err := account.RegisterCompany(account.RegisterCompanyBody{Phone: phone, Name: "seed名称" + phone, Company: "seed企业" + phone, People: 50, Password: "123456", Code: "999999"}); err != nil {
		return errox.Wrap(err)
	}

	return nil
}
