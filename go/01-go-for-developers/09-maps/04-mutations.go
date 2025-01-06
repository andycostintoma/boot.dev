package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !user.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {

	users := map[string]user{
		"John": {
			name:                 "John",
			number:               123,
			scheduledForDeletion: true,
		},
		"Jane": {
			name:                 "Jane",
			number:               456,
			scheduledForDeletion: false,
		},
	}

	johnDeleted, err := deleteIfNecessary(users, "John")
	if err != nil {
		return
	}
	println("John deleted: ", johnDeleted)

	janeDeleted, err := deleteIfNecessary(users, "Jane")
	if err != nil {
		return
	}
	println("Jane deleted: ", janeDeleted)

}
