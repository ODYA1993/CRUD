package entity

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

func (u *User) AddFriend(user *User) bool {
	u.Friends = append(u.Friends, user)
	return true
}

func (u *User) DeleteFriend(id int) bool {
	u.Friends = append(u.Friends[:id], u.Friends[id+1:]...)
	return true
}
