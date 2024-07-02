package models

import (
	"sync"
)

type InMemoryDatabase struct {
	Accounts     map[string]*Account
	Transactions map[string]*Transaction
	Mu           sync.Mutex
}
