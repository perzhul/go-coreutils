package whoami

import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println(Whoami())
}

func Whoami() string {
	user, err := user.Current()
	if err != nil {
		fmt.Errorf("error retrieving current user")
	}

	return user.Username
}
