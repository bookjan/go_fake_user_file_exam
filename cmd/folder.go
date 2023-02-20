package cmd

import (
	"fmt"
	"sort"
	"time"

	"go_fake_user_file_exam/util"
)

type Folder struct {
	Base
	FileMap      map[string]*File
	LabelNameMap map[string]bool
}

type SortFolderByName []*Folder

func (x SortFolderByName) Len() int {
	return len(x)
}
func (x SortFolderByName) Less(i, j int) bool {
	return x[i].Name < x[j].Name
}
func (x SortFolderByName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type SortFolderByTime []*Folder

func (x SortFolderByTime) Len() int {
	return len(x)
}
func (x SortFolderByTime) Less(i, j int) bool {
	return x[i].CreatedAt.Before(x[j].CreatedAt)
}
func (x SortFolderByTime) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (action *Action) CreateFolder() (msg string, logLevel int) {
	if len(action.Options) < 3 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderName, description := options[0], options[1], options[2]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	FOLDER_ID_BASE += 1
	folderId := fmt.Sprint(FOLDER_ID_BASE)
	action.FolderMap[folderId] = &Folder{
		Base: Base{
			Id:          folderId,
			Name:        folderName,
			Description: description,
			CreatedAt:   time.Now(),
		},
		FileMap:      make(map[string]*File),
		LabelNameMap: make(map[string]bool),
	}

	user.FolderIdMap[folderId] = true

	return folderId, util.Trace
}

func (action *Action) DeleteFolder() (msg string, logLevel int) {
	if len(action.Options) < 2 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId := options[0], options[1]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	_, ok = action.FolderMap[folderId]
	if !ok {
		return "folder doesn’t exist", util.Error
	}

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		return "folder owner not match", util.Error
	}

	delete(action.FolderMap, folderId)
	delete(user.FolderIdMap, folderId)

	return "Success", util.Trace
}

func (action *Action) GetFolders() (msg string, logLevel int) {
	if len(action.Options) < 1 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName := options[0]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	folders := []*Folder{}
	folderIds := []string{}
	for k := range user.FolderIdMap {
		folders = append(folders, action.FolderMap[k])
		folderIds = append(folderIds, k)
	}

	if len(folderIds) == 0 {
		return "empty folders", util.Warn
	}

	labelName := ""
	index := 0
	if len(options) == 4 { // with label_name
		index++
		labelName = options[1]
	}

	_, ok = action.LabelMap[labelName]
	if labelName != "" && !ok {
		return "the label is not exists", util.Error
	}

	orderField := SORT_NAME
	if len(options) > 1+index {
		orderField = options[1+index]
	}
	order := ASC_SORT
	if len(options) > 2+index {
		order = options[2+index]
	}

	if orderField == SORT_NAME && ASC_SORT == order {
		sort.Sort(SortFolderByName(folders))
	}
	if orderField == SORT_NAME && DESC_SORT == order {
		sort.Sort(sort.Reverse(SortFolderByName(folders)))
	}

	if orderField == SORT_TIME && ASC_SORT == order {
		sort.Sort(SortFolderByTime(folders))
	}
	if orderField == SORT_TIME && DESC_SORT == order {
		sort.Sort(sort.Reverse(SortFolderByTime(folders)))
	}

	for _, v := range folders {
		if labelName != "" {
			hasLabel := v.LabelNameMap[labelName]
			folderLabelName := ""
			if hasLabel {
				folderLabelName = labelName
			}
			fmt.Printf("%v|%v|%v|%v|%v|%v\n", v.Id, folderLabelName, v.Name, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
		} else {
			fmt.Printf("%v|%v|%v|%v|%v\n", v.Id, v.Name, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
		}
	}

	return "", 0
}

func (action *Action) RenameFolder() (msg string, logLevel int) {
	if len(action.Options) < 3 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId, newFolderName := options[0], options[1], options[2]
	_, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	_, ok = action.FolderMap[folderId]
	if !ok {
		return "folder_id not found", util.Error
	}

	action.FolderMap[folderId].Name = newFolderName

	return "Success", util.Trace
}

func (action *Action) AddFolderLabel() (msg string, logLevel int) {
	if len(action.Options) < 3 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId, labelName := options[0], options[1], options[2]
	_, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	_, ok = action.LabelMap[labelName]
	if !ok {
		return "the label name not exists", util.Error
	}

	folder, ok := action.FolderMap[folderId]
	if !ok {
		return "folder not exists", util.Error
	}

	folder.LabelNameMap[labelName] = true

	return "Success", util.Trace
}

func (action *Action) DeleteFolderLabel() (msg string, logLevel int) {
	if len(action.Options) < 3 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId, labelName := options[0], options[1], options[2]
	_, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	folder, ok := action.FolderMap[folderId]
	if !ok {
		return "folder not exists", util.Error
	}

	delete(folder.LabelNameMap, labelName)

	return "Success", util.Trace
}
