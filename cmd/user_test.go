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

	msg, _ := action.Register()
	if msg != "Success" {
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

	msg, _ := action.Register()

	if msg != "user already existing" {
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
	msg, _ := action2.AddLabel()
	if msg != "Success" {
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

	var action3 Action = Action{
		Command:   Register,
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	msg, _ := action3.GetLabels()
	if msg != "" {
		t.Error("wrong result")
	}
}
