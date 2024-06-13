package repos

import (
	//"errors"
	//"reflect"
	"sync"
	//"time"
)

func NewStore(mu *sync.Mutex, mp map[int]Event) *Store {
	return &Store{mu: mu, mp: mp}
}
