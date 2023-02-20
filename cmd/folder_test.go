package cmd

import (
	"go_fake_user_file_exam/util"
	"strings"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	userName := "user1"
	folderName := "work"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	var action2 Action = Action{
		Options:   []string{userName, folderName, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	msg, _ := action2.CreateFolder()
	if msg != "1001" {
		t.Error("wrong result")
	}
}

func TestDeleteFolder(t *testing.T) {
	userName := "user1"
	folderName := "work"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	var action2 Action = Action{
		Options:   []string{userName, folderName, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.CreateFolder()

	folderId := "1001"
	var action3 Action = Action{
		Options:   []string{userName, folderId},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	msg, _ := action3.DeleteFolder()
	if msg != "Success" {
		t.Error("wrong result")
	}
}

func TestGetFolders(t *testing.T) {
	userName := "user1"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	folderName1 := "work1"
	var action2 Action = Action{
		Options:   []string{userName, folderName1, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.CreateFolder()

	folderName2 := "work2"
	var action3 Action = Action{
		Options:   []string{userName, folderName2, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action3.CreateFolder()

	var action4 Action = Action{
		Options:   []string{userName, SORT_NAME, ASC_SORT},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	output := util.CaptureOutput(func() {
		msg, _ := action4.GetFolders()
		if msg != "" {
			t.Error("wrong result")
		}
	})

	str := strings.Split(output, "|")
	// expect the second item is "work1"
	if str[1] != folderName1 {
		t.Error("wrong result")
	}
}

func TestGetFoldersWithLabel(t *testing.T) {
	userName := "user1"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	folderName1 := "work1"
	var action2 Action = Action{
		Options:   []string{userName, folderName1, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.CreateFolder()

	folderName2 := "work2"
	var action3 Action = Action{
		Options:   []string{userName, folderName2, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action3.CreateFolder()

	labelName1 := "label1"
	labelColor1 := "red"
	var action4 Action = Action{
		Options:   []string{userName, labelName1, labelColor1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action4.AddLabel()

	folderId := "1001"
	var action5 Action = Action{
		Options:   []string{userName, folderId, labelName1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action5.AddFolderLabel()

	var action6 Action = Action{
		Options:   []string{userName, labelName1, SORT_NAME, ASC_SORT},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	output := util.CaptureOutput(func() {
		msg, _ := action6.GetFolders()
		if msg != "" {
			t.Error("wrong result")
		}
	})

	str := strings.Split(output, "|")
	// expect the second item is "work1"
	if str[1] != folderName1 {
		t.Error("wrong result")
	}
}

func TestRenameFolder(t *testing.T) {
	userName := "user1"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	folderName1 := "work1"
	var action2 Action = Action{
		Options:   []string{userName, folderName1, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.CreateFolder()

	folderId := "1001"
	folderName2 := "work2"
	var action3 Action = Action{
		Options:   []string{userName, folderId, folderName2},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	msg, _ := action3.RenameFolder()
	if msg != "Success" {
		t.Error("wrong result")
	}
}

func TestAddFolderLabel(t *testing.T) {
	userName := "user1"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	folderName1 := "work1"
	var action2 Action = Action{
		Options:   []string{userName, folderName1, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.CreateFolder()

	labelName1 := "label1"
	labelColor1 := "red"
	var action3 Action = Action{
		Options:   []string{userName, labelName1, labelColor1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action3.AddLabel()

	folderId := "1001"
	var action4 Action = Action{
		Options:   []string{userName, folderId, labelName1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	msg, _ := action4.AddFolderLabel()
	if msg != "Success" {
		t.Error("wrong result")
	}
}

func TestDeleteFolderLabel(t *testing.T) {
	userName := "user1"
	folderDescription := "it's for work."
	userMap := make(UserMap)
	folderMap := make(FolderMap)
	labelMap := make(LabelMap)
	var action1 Action = Action{
		Options:   []string{userName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action1.Register()

	folderName1 := "work1"
	var action2 Action = Action{
		Options:   []string{userName, folderName1, folderDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action2.CreateFolder()

	labelName1 := "label1"
	labelColor1 := "red"
	var action3 Action = Action{
		Options:   []string{userName, labelName1, labelColor1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action3.AddLabel()

	folderId := "1001"
	var action4 Action = Action{
		Options:   []string{userName, folderId, labelName1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action4.AddFolderLabel()

	var action5 Action = Action{
		Options:   []string{userName, folderId, labelName1},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	msg, _ := action5.DeleteFolderLabel()
	if msg != "Success" {
		t.Error("wrong result")
	}
}
