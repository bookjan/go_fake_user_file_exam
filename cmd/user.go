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

func (action *Action) Register() (msg string, logLevel int) {
	userName := action.Options[0]
	_, ok := action.UserMap[userName]
	if !ok {
		USER_ID_BASE += 1
		action.UserMap[userName] = &User{
			Base: Base{
				Id:   fmt.Sprint(USER_ID_BASE),
				Name: userName,
			},
			FolderIdMap:  map[string]bool{},
			LabelNameMap: map[string]bool{},
		}

		return "Success", util.Trace
	} else {
		return "user already existing", util.Error
	}
}

func (action *Action) AddLabel() (msg string, logLevel int) {
	if len(action.Options) < 3 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, labelName, color := options[0], options[1], options[2]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	_, ok = action.LabelMap[labelName]
	if ok {
		return "the label name existing", util.Error
	}

	user.LabelNameMap[labelName] = true
	action.LabelMap[labelName] = &Label{
		Base: Base{
			Name:      labelName,
			CreatedAt: time.Now(),
		},
		Color: color,
	}

	return "Success", util.Trace
}

func (action *Action) GetLabels() (msg string, logLevel int) {
	if len(action.Options) < 1 {
		return "invalid arguments", util.Error
	}

	userName := action.Options[0]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	for k := range user.LabelNameMap {
		v := action.LabelMap[k]
		fmt.Printf("%v|%v|%v|%v", v.Name, v.Color, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}

	return "", 0
}

func (action *Action) DeleteLabel() (msg string, logLevel int) {
	if len(action.Options) < 2 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, labelName := options[0], options[1]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	_, ok = action.LabelMap[labelName]
	if !ok {
		return "the label name not exists", util.Error
	}

	_, ok = user.LabelNameMap[labelName]
	if !ok {
		return "owner mismatch", util.Error
	}

	delete(user.LabelNameMap, labelName)

	return "Success", util.Trace
}
