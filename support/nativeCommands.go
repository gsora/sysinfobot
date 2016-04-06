package support

import "fmt"

// list of internal, recognized commands
var command = map[string]string{
	"uname":    "uname -a",
	"hostname": "uname -n",
	"ip":       "",
}

func execCommand(com string) error {
	if command == nil {
		fmt.Println("empty")
	}

	fmt.Println(command[com])
	return nil
}
