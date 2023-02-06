package handlers

import (
	"fmt"

	"go_fake_user_file_exam/config"
)

func Register(args []string, userMap map[string]config.User) {
	userName := args[0]
	_, ok := userMap[userName]
	if !ok {
		config.USER_ID_BASE += 1
		userMap[userName] = config.User{
			Id:          fmt.Sprint(config.USER_ID_BASE),
			Name:        userName,
			FolderIdMap: map[string]bool{},
		}

		fmt.Printf("Success")
	} else {
		fmt.Printf("Error - user already existing")
	}
}
