package handlers

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"go_fake_user_file_exam/config"
	"go_fake_user_file_exam/util"
)

func UploadFile(args *config.Arguments) {
	if len(args.Options) < 4 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderId, fileName, description := options[0], options[1], options[2], options[3]
	_, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		util.PrintOrLog("folder_id not found", util.Error)
		return
	}

	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	match := re.FindStringSubmatch(fileName)

	Name := match[2]
	Extension := strings.ReplaceAll(match[3], ".", "")
	folder.FileMap[fileName] = &config.File{
		Name:        Name,
		Extension:   Extension,
		Description: description,
		CreatedAt:   time.Now(),
	}

	util.PrintOrLog("Success", util.Trace)
}

func DeleteFile(args *config.Arguments) {
	if len(args.Options) < 3 {
		util.PrintOrLog("invalid arguments", util.Error)
		return
	}

	options := args.Options
	userName, folderId, fileName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		util.PrintOrLog("unknown user", util.Error)
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		util.PrintOrLog("folder_id not found", util.Error)
		return
	}

	_, ok = folder.FileMap[fileName]
	if !ok {
		util.PrintOrLog("file_name not found", util.Error)
		return
	}

	delete(folder.FileMap, fileName)

	util.PrintOrLog("Success", util.Trace)
}

func GetFiles(args *config.Arguments) {
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

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		util.PrintOrLog("folder_name not found", util.Error)
	}

	files := []*config.File{}
	folder := args.FolderMap[folderId]
	for k := range folder.FileMap {
		files = append(files, folder.FileMap[k])
	}

	if len(files) == 0 {
		util.PrintOrLog("empty files", util.Warn)
		return
	}

	orderField := config.SORT_NAME
	if len(options) > 2 {
		orderField = options[2]
	}
	order := config.ASC_SORT
	if len(options) > 3 {
		order = options[3]
	}

	if orderField == config.SORT_NAME && config.ASC_SORT == order {
		sort.Sort(config.SortFileByTime(files))
	}
	if orderField == config.SORT_NAME && config.DESC_SORT == order {
		sort.Sort(sort.Reverse(config.SortFileByName(files)))
	}

	if orderField == config.SORT_TIME && config.ASC_SORT == order {
		sort.Sort(config.SortFileByTime(files))
	}
	if orderField == config.SORT_TIME && config.DESC_SORT == order {
		sort.Sort(sort.Reverse(config.SortFileByTime(files)))
	}

	if orderField == config.SORT_EXTENSION && config.ASC_SORT == order {
		sort.Sort(config.SortFileByExtension(files))
	}
	if orderField == config.SORT_EXTENSION && config.DESC_SORT == order {
		sort.Sort(sort.Reverse(config.SortFileByExtension(files)))
	}

	for _, v := range files {
		fullFileName := v.Name + "." + v.Extension
		fmt.Printf("%v|%v|%v|%v|%v\n", fullFileName, v.Extension, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}
}
