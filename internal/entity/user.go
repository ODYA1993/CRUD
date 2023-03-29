package entity

import "sync"

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
	sync.Mutex
}

func (u *User) AddFriend(user *User) bool {
	u.Lock()
	u.Friends = append(u.Friends, user)
	u.Unlock()
	return true
}

func (u *User) DeleteFriend(id int) bool {
	u.Lock()
	u.Friends = append(u.Friends[:id], u.Friends[id+1:]...)
	u.Unlock()
	return true
}
