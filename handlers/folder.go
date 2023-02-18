package handlers

import (
	"fmt"
	"sort"
	"time"

	"go_fake_user_file_exam/config"
	"go_fake_user_file_exam/util"
)

func CreateFolder(args *config.Arguments) {
	if len(args.Options) < 3 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderName, description := options[0], options[1], options[2]
	user, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	config.FOLDER_ID_BASE += 1
	folderId := fmt.Sprint(config.FOLDER_ID_BASE)
	args.FolderMap[folderId] = &config.Folder{
		Id:           folderId,
		Name:         folderName,
		Description:  description,
		CreatedAt:    time.Now(),
		FileMap:      make(map[string]*config.File),
		LabelNameMap: make(map[string]bool),
	}

	user.FolderIdMap[folderId] = true

	util.PrintOrLog(folderId, util.Trace)
}

func DeleteFolder(args *config.Arguments) {
	if len(args.Options) < 2 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderId := options[0], options[1]
	user, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	_, ok = args.FolderMap[folderId]
	if !ok {
		util.PrintOrLog("folder doesn’t exist", util.Error)
		return
	}

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		util.PrintOrLog("folder owner not match", util.Error)
		return
	}

	delete(args.FolderMap, folderId)
	delete(user.FolderIdMap, folderId)

	util.PrintOrLog("Success", util.Trace)
}

func GetFolders(args *config.Arguments) {
	if len(args.Options) < 1 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName := options[0]
	user, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	folders := []*config.Folder{}
	folderIds := []string{}
	for k := range user.FolderIdMap {
		folders = append(folders, args.FolderMap[k])
		folderIds = append(folderIds, k)
	}

	if len(folderIds) == 0 {
		util.PrintOrLog("empty folders", util.Warn)
		return
	}

	labelName := ""
	index := 0
	if len(options) == 4 { // with label_name
		index++
		labelName = options[1]
	}

	_, ok = args.LabelMap[labelName]
	if labelName != "" && !ok {
		util.PrintOrLog("the label is not exists", util.Error)
		return
	}

	orderField := config.SORT_NAME
	if len(options) > 1+index {
		orderField = options[1+index]
	}
	order := config.ASC_SORT
	if len(options) > 2+index {
		order = options[2+index]
	}

	if orderField == config.SORT_NAME && config.ASC_SORT == order {
		sort.Sort(config.SortFolderByName(folders))
	}
	if orderField == config.SORT_NAME && config.DESC_SORT == order {
		sort.Sort(sort.Reverse(config.SortFolderByName(folders)))
	}

	if orderField == config.SORT_TIME && config.ASC_SORT == order {
		sort.Sort(config.SortFolderByTime(folders))
	}
	if orderField == config.SORT_TIME && config.DESC_SORT == order {
		sort.Sort(sort.Reverse(config.SortFolderByTime(folders)))
	}

	for _, v := range folders {
		if labelName != "" {
			fmt.Printf("%v|%v|%v|%v|%v|%v\n", v.Id, labelName, v.Name, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
		} else {
			fmt.Printf("%v|%v|%v|%v|%v\n", v.Id, v.Name, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
		}
	}
}

func RenameFolder(args *config.Arguments) {
	if len(args.Options) < 3 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderId, newFolderName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	_, ok = args.FolderMap[folderId]
	if !ok {
		util.PrintOrLog("folder_id not found", util.Error)
		return
	}

	args.FolderMap[folderId].Name = newFolderName

	util.PrintOrLog("Success", util.Trace)
}

func AddFolderLabel(args *config.Arguments) {
	if len(args.Options) < 3 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderId, labelName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	_, ok = args.LabelMap[labelName]
	if !ok {
		util.PrintOrLog("the label name not exists", util.Error)
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		util.PrintOrLog("folder not exists", util.Error)
		return
	}

	folder.LabelNameMap[labelName] = true

	util.PrintOrLog("Success", util.Trace)
}

func DeleteFolderLabel(args *config.Arguments) {
	if len(args.Options) < 3 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderId, labelName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		util.PrintOrLog("folder not exists", util.Error)
		return
	}

	delete(folder.LabelNameMap, labelName)

	util.PrintOrLog("Success", util.Trace)
}
