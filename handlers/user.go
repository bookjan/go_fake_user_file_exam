package handlers

import (
	"fmt"
	"time"

	"go_fake_user_file_exam/config"
	"go_fake_user_file_exam/util"
)

func Register(args *config.Arguments) {
	userName := args.Options[0]
	_, ok := args.UserMap[userName]
	if !ok {
		config.USER_ID_BASE += 1
		args.UserMap[userName] = &config.User{
			Id:           fmt.Sprint(config.USER_ID_BASE),
			Name:         userName,
			FolderIdMap:  map[string]bool{},
			LabelNameMap: map[string]bool{},
		}

		util.PrintOrLog("Success", util.Trace)
	} else {
		util.PrintOrLog("user already existing", util.Error)
	}
}

func AddLabel(args *config.Arguments) {
	if len(args.Options) < 3 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, labelName, color := options[0], options[1], options[2]
	user, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	_, ok = args.LabelMap[labelName]
	if ok {
		util.PrintOrLog("the label name existing", util.Error)
		return
	}

	user.LabelNameMap[labelName] = true
	args.LabelMap[labelName] = &config.Label{
		Name:      labelName,
		Color:     color,
		CreatedAt: time.Now(),
	}

	util.PrintOrLog("Success", util.Trace)
}

func GetLabel(args *config.Arguments) {
	if len(args.Options) < 1 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	userName := args.Options[0]
	user, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	for k := range user.LabelNameMap {
		v := args.LabelMap[k]
		fmt.Printf("%v|%v|%v|%v", v.Name, v.Color, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}
}

func DeleteLabel(args *config.Arguments) {
	if len(args.Options) < 2 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, labelName := options[0], options[1]
	user, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	_, ok = args.LabelMap[labelName]
	if !ok {
		util.PrintOrLog("the label name not exists", util.Error)
		return
	}

	_, ok = user.LabelNameMap[labelName]
	if !ok {
		util.PrintOrLog("owner mismatch", util.Error)
		return
	}

	delete(user.LabelNameMap, labelName)

	util.PrintOrLog("Success", util.Trace)
}
