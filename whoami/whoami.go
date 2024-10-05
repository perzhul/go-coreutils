package whoami

import (
	"fmt"
	"os/user"
)

func whoami() {
	user, err := user.Current()
	if err != nil {
		fmt.Errorf("error retrieving current user")
	}

	fmt.Println(user.Username)

}
