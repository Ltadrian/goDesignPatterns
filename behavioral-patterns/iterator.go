package main

import "fmt"

/**
Iterator is a behavioral design pattern
that lets you traverse elements of a collection
without exposing its underlying representation
Ex: collection iterator
*/

// Collection
type collection interface {
	createIterator() iterator
}

// Concrete collection
type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

// Iterator
type iterator interface {
	hasNext() bool
	getNext() *user
}

// Concrete iterator
type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// Client code
type user struct {
	name string
	age  int
}

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
