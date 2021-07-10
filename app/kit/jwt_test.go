package kit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	token, err := JwtCreate(JwtDate{
		AccountID: 30,
		CompanyID: 40,
		EmployID:  50,
	})

	assert.Nil(t, err)
	time.Sleep(time.Millisecond * 1000)
	data, err := JwtParse(token)
	assert.Equal(t, data, JwtDate{
		AccountID: 30,
		CompanyID: 40,
		EmployID:  50,
	})
	assert.Nil(t, err)
}
