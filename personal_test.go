package monobank

import (
	"testing"
	"time"
)

func TestGetUserInfo(t *testing.T) {
	token := "token"
	c := NewClient(token)
	user, err := c.GetUserInfo()
	// TODO: mock requests
	// TODO: add real validationj
	if err != nil {
		t.Error(err)
	} else {
		t.Error(user)
	}
}

func TestGetTransactions(t *testing.T) {
	token := "token"
	c := NewClient(token)
	user, err := c.GetUserInfo()
	data, err := c.GetTransactions(user.Accounts[2].AccountID, time.Now().Add(-48*time.Hour), time.Now())
	// TODO: mock requests
	// TODO: add real validationj
	if err != nil {
		t.Error(err)
	} else {
		t.Error(data)
	}
}
