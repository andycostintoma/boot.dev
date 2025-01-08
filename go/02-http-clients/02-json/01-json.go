package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const issueList = `
[
  {
    "id": 0,
    "name": "Fix the thing",
    "estimate": 0.5,
    "completed": false
  },
  {
    "id": 1,
    "name": "Unstick the widget",
    "estimate": 30,
    "completed": false
  }
]
`

const userObject = `
{
  "name": "Wayne Lagner",
  "role": "Developer",
  "remote": true
}
`

func isValidJSON(input string) bool {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	return err == nil
}

func main() {
	fmt.Println(isValidJSON(issueList))
	fmt.Println(isValidJSON(userObject))
}
