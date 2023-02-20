package cmd

import (
	"testing"
)

func TestRegisterUser1(t *testing.T) {
	userName := "user1"
	var action Action = Action{
		Command:   Register,
		Options:   []string{userName},
		UserMap:   make(UserMap),
		FolderMap: make(FolderMap),
		LabelMap:  make(LabelMap),
	}

	action.Register()

	if userName != action.UserMap[userName].Base.Name {
		t.Error("wrong result")
	}
}

func TestRegisterUser1Twice(t *testing.T) {
	var action Action = Action{
		Command:   Register,
		Options:   []string{"user1"},
		UserMap:   make(UserMap),
		FolderMap: make(FolderMap),
		LabelMap:  make(LabelMap),
	}

	action.Register()
	action.Register()

	if USER_ID_BASE == 1 {
		t.Error("wrong result")
	}
}

func TestAddLabelNameAndColor(t *testing.T) {
	userName := "user1"
	labelName := "label1"
	labelColor := "red"
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)

	var action1 Action = Action{
		Command:   Register,
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	var action2 Action = Action{
		Command:   Register,
		Options:   []string{userName, labelName, labelColor},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.AddLabel()

	if !action2.UserMap[userName].LabelNameMap[labelName] {
		t.Error("wrong result")
	}
	if action2.LabelMap[labelName].Color != labelColor {
		t.Error("wrong result")
	}
}

func TestGetLabels(t *testing.T) {
	userName := "user1"
	labelName := "label1"
	labelColor := "red"
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)

	var action1 Action = Action{
		Command:   Register,
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	var action2 Action = Action{
		Command:   Register,
		Options:   []string{userName, labelName, labelColor},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.AddLabel()

	if !action2.UserMap[userName].LabelNameMap[labelName] {
		t.Error("wrong result")
	}
	if action2.LabelMap[labelName].Color != labelColor {
		t.Error("wrong result")
	}
}
