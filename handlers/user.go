package handlers

import (
	"fmt"
	"time"

	"go_fake_user_file_exam/config"
)

func Register(args *config.Arguments) {
	userName := args.Options[0]
	_, ok := args.UserMap[userName]
	if !ok {
		config.USER_ID_BASE += 1
		args.UserMap[userName] = config.User{
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

func AddLabel(args *config.Arguments) {
	if len(args.Options) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, labelName, color := options[0], options[1], options[2]
	user, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = args.LabelMap[labelName]
	if ok {
		fmt.Println("Error - the label name existing")
		return
	}

	user.LabelNameMap[labelName] = true
	args.LabelMap[labelName] = &config.Label{
		Name:      labelName,
		Color:     color,
		CreatedAt: time.Now(),
	}

	fmt.Println("Success")
}

func GetLabel(args *config.Arguments) {
	if len(args.Options) < 1 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName := args.Options[0]
	user, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	for k := range user.LabelNameMap {
		v := args.LabelMap[k]
		fmt.Printf("%v|%v|%v|%v", v.Name, v.Color, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}
}

func DeleteLabel(args *config.Arguments) {
	if len(args.Options) < 2 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, labelName := options[0], options[1]
	user, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = args.LabelMap[labelName]
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
