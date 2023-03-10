package cmd

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"go_fake_user_file_exam/util"
)

type File struct {
	Base
	Extension string
}

type SortFileByName []*File

func (x SortFileByName) Len() int {
	return len(x)
}
func (x SortFileByName) Less(i, j int) bool {
	return x[i].Name < x[j].Name
}
func (x SortFileByName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type SortFileByTime []*File

func (x SortFileByTime) Len() int {
	return len(x)
}
func (x SortFileByTime) Less(i, j int) bool {
	return x[i].CreatedAt.Before(x[j].CreatedAt)
}
func (x SortFileByTime) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type SortFileByExtension []*File

func (x SortFileByExtension) Len() int {
	return len(x)
}
func (x SortFileByExtension) Less(i, j int) bool {
	return x[i].Extension < x[j].Extension
}
func (x SortFileByExtension) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (action *Action) UploadFile() (msg string, logLevel int) {
	if len(action.Options) < 4 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId, fileName, description := options[0], options[1], options[2], options[3]
	_, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	folder, ok := action.FolderMap[folderId]
	if !ok {
		return "folder_id not found", util.Error
	}

	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	match := re.FindStringSubmatch(fileName)

	Name := match[2]
	Extension := strings.ReplaceAll(match[3], ".", "")
	folder.FileMap[fileName] = &File{
		Base: Base{
			Name:        Name,
			Description: description,
			CreatedAt:   time.Now(),
		},
		Extension: Extension,
	}

	return "Success", util.Trace
}

func (action *Action) DeleteFile() (msg string, logLevel int) {
	if len(action.Options) < 3 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId, fileName := options[0], options[1], options[2]
	_, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	folder, ok := action.FolderMap[folderId]
	if !ok {
		return "folder_id not found", util.Error
	}

	_, ok = folder.FileMap[fileName]
	if !ok {
		return "file_name not found", util.Error
	}

	delete(folder.FileMap, fileName)

	return "Success", util.Trace
}

func (action *Action) GetFiles() (msg string, logLevel int) {
	if len(action.Options) < 2 {
		return "invalid arguments", util.Error
	}

	options := action.Options
	userName, folderId := options[0], options[1]
	user, ok := action.UserMap[userName]
	if !ok {
		return "unknown user", util.Error
	}

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		return "folder_name not found", util.Error
	}

	files := []*File{}
	folder := action.FolderMap[folderId]
	for k := range folder.FileMap {
		files = append(files, folder.FileMap[k])
	}

	if len(files) == 0 {
		return "empty files", util.Warn
	}

	orderField := SORT_NAME
	if len(options) > 2 {
		orderField = options[2]
	}
	order := ASC_SORT
	if len(options) > 3 {
		order = options[3]
	}

	if orderField == SORT_NAME && ASC_SORT == order {
		sort.Sort(SortFileByTime(files))
	}
	if orderField == SORT_NAME && DESC_SORT == order {
		sort.Sort(sort.Reverse(SortFileByName(files)))
	}

	if orderField == SORT_TIME && ASC_SORT == order {
		sort.Sort(SortFileByTime(files))
	}
	if orderField == SORT_TIME && DESC_SORT == order {
		sort.Sort(sort.Reverse(SortFileByTime(files)))
	}

	if orderField == SORT_EXTENSION && ASC_SORT == order {
		sort.Sort(SortFileByExtension(files))
	}
	if orderField == SORT_EXTENSION && DESC_SORT == order {
		sort.Sort(sort.Reverse(SortFileByExtension(files)))
	}

	for _, v := range files {
		fullFileName := v.Name + "." + v.Extension
		fmt.Printf("%v|%v|%v|%v|%v\n", fullFileName, v.Extension, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}

	return "", 0
}
