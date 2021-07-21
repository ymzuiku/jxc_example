package account

import (
	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
)

func SignInWithSession(body SignInWithSessionBody) (AccountRes, error) {
	sessionCache := rds.New(SESSION_CACHE)
	if !sessionCache.Is(body.AccountID, body.Session) {
		return AccountRes{}, errox.Errorf("您的登入状态已过期，请重新登入")
	}

	return LoadAccount(body.AccountID)
}
