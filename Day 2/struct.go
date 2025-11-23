package main

type User struct {
	ID       int64
	Name     string
	isActive bool
}

func (u *User) InitializeUser(name string) {
	u.ID = 1001
	u.Name = name
	u.isActive = true
}

func (u User) InitializeUser1(name string) {
	u.ID = 1001
	u.Name = name
	u.isActive = true
}

func (u *User) Deactivate() {
	u.isActive = false
}

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value += 1
}