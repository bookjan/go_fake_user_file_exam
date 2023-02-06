package handlers

import (
	"fmt"
	"time"

	"go_fake_user_file_exam/config"
)

func Register(args []string, userMap map[string]config.User) {
	userName := args[0]
	_, ok := userMap[userName]
	if !ok {
		config.USER_ID_BASE += 1
		userMap[userName] = config.User{
			Id:           fmt.Sprint(config.USER_ID_BASE),
			Name:         userName,
			FolderIdMap:  map[string]bool{},
			LabelNameMap: map[string]bool{},
		}

		fmt.Printf("Success")
	} else {
		fmt.Printf("Error - user already existing")
	}
}

func AddLabel(args []string, userMap map[string]config.User, labelMap map[string]*config.Label) {
	if len(args) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, labelName, color := args[0], args[1], args[2]
	user, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = labelMap[labelName]
	if ok {
		fmt.Println("Error - the label name existing")
		return
	}

	user.LabelNameMap[labelName] = true
	labelMap[labelName] = &config.Label{
		Name:      labelName,
		Color:     color,
		CreatedAt: time.Now(),
	}

	fmt.Println("Success")
}

func GetLabel(args []string, userMap map[string]config.User, labelMap map[string]*config.Label) {
	if len(args) < 1 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName := args[0]
	user, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	for k := range user.LabelNameMap {
		v := labelMap[k]
		fmt.Printf("%v|%v|%v|%v", v.Name, v.Color, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}
}

func DeleteLabel(args []string, userMap map[string]config.User, labelMap map[string]*config.Label) {
	if len(args) < 2 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, labelName := args[0], args[1]
	user, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = labelMap[labelName]
	if !ok {
		fmt.Println("Error - the label name not exists")
		return
	}

	_, ok = user.LabelNameMap[labelName]
	if !ok {
		fmt.Println("Error - owner mismatch")
		return
	}

	delete(user.LabelNameMap, labelName)

	fmt.Println("Success")
}
