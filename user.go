package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID      int
	Balance int
}

func (u *User) bet() {
	amount := rand.Intn(100) - 50 // Random amount between -50 and 50
	if amount > 0 {
		u.deposit(amount)
	} else {
		u.withdraw(amount)
	}
}

func (u *User) deposit(amount int) {
	u.Balance += amount
}

func (u *User) withdraw(amount int) {
	u.Balance -= amount
}

// UserService manages users
type UserService struct {
	users map[int]*User // In-memory storage (replace with database for persistence)
}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	return &UserService{users: make(map[int]*User)}
}

// AddUser adds a new user with the provided ID and balance
func (s *UserService) AddUser(id int, balance int) error {
	if _, ok := s.users[id]; ok {
		return fmt.Errorf("user with ID %d already exists", id)
	}
	s.users[id] = &User{ID: id, Balance: balance}
	return nil
}

// GetUsers returns a list of all users
func (s *UserService) GetUsers() []*User {
	userList := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		userList = append(userList, user)
	}
	return userList
}

func (s *UserService) GetUser(userID int) *User {
	if _, ok := userService.users[userID]; !ok {
		userService.users[userID] = &User{ID: userID, Balance: 0}
	}
	return userService.users[userID]
}
