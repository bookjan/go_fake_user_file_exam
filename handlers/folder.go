package handlers

import (
	"fmt"
	"sort"
	"time"

	"go_fake_user_file_exam/config"
)

func CreateFolder(args []string, userMap map[string]config.User, folderMap map[string]*config.Folder) {
	if len(args) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, folderName, description := args[0], args[1], args[2]
	user, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	config.FOLDER_ID_BASE += 1
	folderId := fmt.Sprint(config.FOLDER_ID_BASE)
	folderMap[folderId] = &config.Folder{
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

func DeleteFolder(args []string, userMap map[string]config.User, folderMap map[string]*config.Folder) {
	if len(args) < 2 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, folderId := args[0], args[1]
	user, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = folderMap[folderId]
	if !ok {
		fmt.Println("Error - folder doesn’t exist")
		return
	}

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		fmt.Println("Error - folder owner not match")
		return
	}

	delete(folderMap, folderId)
	delete(user.FolderIdMap, folderId)

	fmt.Println("Success")
}

func GetFolders(args []string, userMap map[string]config.User, folderMap map[string]*config.Folder, labelMap map[string]*config.Label) {
	if len(args) < 1 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName := args[0]
	user, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	folders := []config.Folder{}
	folderIds := []string{}
	for k := range user.FolderIdMap {
		folders = append(folders, *folderMap[k])
		folderIds = append(folderIds, k)
	}

	if len(folderIds) == 0 {
		fmt.Println("Warning - empty folders")
		return
	}

	labelName := ""
	index := 0
	if len(args) == 4 { // with label_name
		index++
		labelName = args[1]
	}

	_, ok = labelMap[labelName]
	if labelName != "" && !ok {
		fmt.Println("Error - the label is not exists")
		return
	}

	orderField := config.SORT_NAME
	if len(args) > 1+index {
		orderField = args[1+index]
	}
	order := config.ASC_SORT
	if len(args) > 2+index {
		order = args[2+index]
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

func RenameFolder(args []string, userMap map[string]config.User, folderMap map[string]*config.Folder) {
	if len(args) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, folderId, newFolderName := args[0], args[1], args[2]
	_, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = folderMap[folderId]
	if !ok {
		fmt.Println("Error - folder_id not found")
		return
	}

	folderMap[folderId].Name = newFolderName

	fmt.Println("Success")
}

func AddFolderLabel(args []string, userMap map[string]config.User, folderMap map[string]*config.Folder, labelMap map[string]*config.Label) {
	if len(args) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, folderId, labelName := args[0], args[1], args[2]
	_, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	_, ok = labelMap[labelName]
	if !ok {
		fmt.Println("Error - the label name not exists")
		return
	}

	folder, ok := folderMap[folderId]
	if !ok {
		fmt.Println("Error - folder not exists")
		return
	}

	folder.LabelNameMap[labelName] = true

	fmt.Println("Success")
}

func DeleteFolderLabel(args []string, userMap map[string]config.User, folderMap map[string]*config.Folder, labelMap map[string]*config.Label) {
	if len(args) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	userName, folderId, labelName := args[0], args[1], args[2]
	_, ok := userMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	folder, ok := folderMap[folderId]
	if !ok {
		fmt.Println("Error - folder not exists")
		return
	}

	delete(folder.LabelNameMap, labelName)

	fmt.Println("Success")
}
