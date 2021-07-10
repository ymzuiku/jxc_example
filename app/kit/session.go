package kit

import (
	"context"
	"fmt"
	"time"
)

const SESSION_KEY = "session:"

func SessionCreate(key string, val string) (string, error) {
	session := Sha256(key + Env.Session + time.Now().String())
	if err := Redis.SetEX(context.Background(), SESSION_KEY+session, val, loginTimeOut).Err(); err != nil {
		return "", err
	}
	return session, nil
}
func SessionRemove(session string) error {
	res := Redis.Del(context.Background(), SESSION_KEY+session)
	err := res.Err()
	if err != nil {
		return err
	}
	if res.Val() != 1 {
		return fmt.Errorf("session val is error")
	}
	return nil
}

func SessionGet(session string) (string, error) {
	res := Redis.Get(context.Background(), SESSION_KEY+session)
	err := res.Err()
	if err != nil {
		return "", err
	}
	val := res.Val()
	if val == "" {
		return "", fmt.Errorf("session val is error")
	}
	return val, nil
}

func SessionHas(session string) error {
	res := Redis.Get(context.Background(), SESSION_KEY+session)
	err := res.Err()
	if err != nil {
		return err
	}
	val := res.Val()
	if val == "" {
		return fmt.Errorf("session val is error")
	}
	return nil
}
