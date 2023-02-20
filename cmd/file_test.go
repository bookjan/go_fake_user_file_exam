package cmd

import (
	"go_fake_user_file_exam/util"
	"strings"
	"testing"
)

func TestUploadFile(t *testing.T) {
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
	fileName := "1.png"
	fileDescription := "photo for working space"
	var action3 Action = Action{
		Options:   []string{userName, folderId, fileName, fileDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}

	msg, _ := action3.UploadFile()
	if msg != "Success" {
		t.Error("wrong result")
	}
}

func TestDeleteFile(t *testing.T) {
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
	fileName := "1.png"
	fileDescription := "photo for working space"
	var action3 Action = Action{
		Options:   []string{userName, folderId, fileName, fileDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action3.UploadFile()

	var action4 Action = Action{
		Options:   []string{userName, folderId, fileName},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	msg, _ := action4.DeleteFile()

	if msg != "Success" {
		t.Error("wrong result")
	}
}

func TestGetFiles(t *testing.T) {
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
	fileName1 := "1.png"
	fileDescription := "photo for working space"
	var action3 Action = Action{
		Options:   []string{userName, folderId, fileName1, fileDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action3.UploadFile()

	fileName2 := "2.png"
	fileDescription = "photo for working space"
	var action4 Action = Action{
		Options:   []string{userName, folderId, fileName2, fileDescription},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	action4.UploadFile()

	var action5 Action = Action{
		Options:   []string{userName, folderId, SORT_NAME, ASC_SORT},
		UserMap:   userMap,
		FolderMap: folderMap,
		LabelMap:  labelMap,
	}
	output := util.CaptureOutput(func() {
		msg, _ := action5.GetFiles()
		if msg != "" {
			t.Error("wrong result")
		}
	})

	str := strings.Split(output, "|")
	// expect the first item is "1.png"
	if str[0] != fileName1 {
		t.Error("wrong result")
	}
}
