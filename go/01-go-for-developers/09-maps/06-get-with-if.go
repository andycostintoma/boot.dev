package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}

func main() {

	users := map[string]int{
		"John":  0,
		"Jane":  0,
		"David": 0,
	}

	messagedUsers := []string{
		"John",
		"Jane",
		"Bob",
		"David",
	}

	getCounts(messagedUsers, users)

	for user, count := range users {
		println(user, count)
	}
}
