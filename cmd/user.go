package cmd

import (
	"fmt"
	"time"

	"go_fake_user_file_exam/util"
)

type User struct {
	Base
	FolderIdMap  map[string]bool
	LabelNameMap map[string]bool
}

func Register(args *Arguments) {
	userName := args.Options[0]
	_, ok := args.UserMap[userName]
	if !ok {
		USER_ID_BASE += 1
		args.UserMap[userName] = &User{
			Base: Base{
				Id:   fmt.Sprint(USER_ID_BASE),
				Name: userName,
			},
			FolderIdMap:  map[string]bool{},
			LabelNameMap: map[string]bool{},
		}

		util.PrintOrLog("Success", util.Trace)
	} else {
		util.PrintOrLog("user already existing", util.Error)
	}
}

func AddLabel(args *Arguments) {
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
	args.LabelMap[labelName] = &Label{
		Base: Base{
			Name:      labelName,
			CreatedAt: time.Now(),
		},
		Color: color,
	}

	util.PrintOrLog("Success", util.Trace)
}

func GetLabel(args *Arguments) {
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

func DeleteLabel(args *Arguments) {
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
