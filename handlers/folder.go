package handlers

import (
	"fmt"
	"sort"
	"time"

	"go_fake_user_file_exam/config"
)

func CreateFolder(args *config.Arguments) {
	if len(args.Options) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderName, description := options[0], options[1], options[2]
	user, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
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

	fmt.Println(folderId)
}

func DeleteFolder(args *config.Arguments) {
	if len(args.Options) < 2 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderId := options[0], options[1]
	user, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = args.FolderMap[folderId]
	if !ok {
		fmt.Println("Error - folder doesn’t exist")
		return
	}

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		fmt.Println("Error - folder owner not match")
		return
	}

	delete(args.FolderMap, folderId)
	delete(user.FolderIdMap, folderId)

	fmt.Println("Success")
}

func GetFolders(args *config.Arguments) {
	if len(args.Options) < 1 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName := options[0]
	user, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	folders := []config.Folder{}
	folderIds := []string{}
	for k := range user.FolderIdMap {
		folders = append(folders, *args.FolderMap[k])
		folderIds = append(folderIds, k)
	}

	if len(folderIds) == 0 {
		fmt.Println("Warning - empty folders")
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
		fmt.Println("Error - the label is not exists")
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
		sort.Slice(folders, func(i, j int) bool {
			return folders[i].Name < folders[j].Name
		})
	}
	if orderField == config.SORT_NAME && config.DESC_SORT == order {
		sort.Slice(folders, func(i, j int) bool {
			return folders[i].Name > folders[j].Name
		})
	}

	if orderField == config.SORT_TIME && config.ASC_SORT == order {
		sort.Slice(folders, func(i, j int) bool {
			return folders[i].CreatedAt.Before(folders[j].CreatedAt)
		})
	}
	if orderField == config.SORT_TIME && config.DESC_SORT == order {
		sort.Slice(folders, func(i, j int) bool {
			return folders[i].CreatedAt.After(folders[j].CreatedAt)
		})
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
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderId, newFolderName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = args.FolderMap[folderId]
	if !ok {
		fmt.Println("Error - folder_id not found")
		return
	}

	args.FolderMap[folderId].Name = newFolderName

	fmt.Println("Success")
}

func AddFolderLabel(args *config.Arguments) {
	if len(args.Options) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderId, labelName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = args.LabelMap[labelName]
	if !ok {
		fmt.Println("Error - the label name not exists")
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		fmt.Println("Error - folder not exists")
		return
	}

	folder.LabelNameMap[labelName] = true

	fmt.Println("Success")
}

func DeleteFolderLabel(args *config.Arguments) {
	if len(args.Options) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderId, labelName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		fmt.Println("Error - folder not exists")
		return
	}

	delete(folder.LabelNameMap, labelName)

	fmt.Println("Success")
}
