package handlers

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"go_fake_user_file_exam/config"
)

func UploadFile(args *config.Arguments) {
	if len(args.Options) < 4 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderId, fileName, description := options[0], options[1], options[2], options[3]
	_, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		fmt.Println("Error - folder_id not found")
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

	fmt.Println("Success")
}

func DeleteFile(args *config.Arguments) {
	if len(args.Options) < 3 {
		fmt.Println("Error - invalid arguments")
		return
	}

	options := args.Options
	userName, folderId, fileName := options[0], options[1], options[2]
	_, ok := args.UserMap[userName]
	if !ok {
		fmt.Println("Error - unknown user")
		return
	}

	folder, ok := args.FolderMap[folderId]
	if !ok {
		fmt.Println("Error - folder_id not found")
		return
	}

	_, ok = folder.FileMap[fileName]
	if !ok {
		fmt.Println("Error - file_name not found")
		return
	}

	delete(folder.FileMap, fileName)

	fmt.Println("Success")
}

func GetFiles(args *config.Arguments) {
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

	_, ok = user.FolderIdMap[folderId]
	if !ok {
		fmt.Println("folder_name not found")
	}

	files := []config.File{}
	folder := args.FolderMap[folderId]
	for k := range folder.FileMap {
		files = append(files, *folder.FileMap[k])
	}

	if len(files) == 0 {
		fmt.Println("Warning - empty files")
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
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name < files[j].Name
		})
	}
	if orderField == config.SORT_NAME && config.DESC_SORT == order {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name > files[j].Name
		})
	}

	if orderField == config.SORT_TIME && config.ASC_SORT == order {
		sort.Slice(files, func(i, j int) bool {
			return files[i].CreatedAt.Before(files[j].CreatedAt)
		})
	}
	if orderField == config.SORT_TIME && config.DESC_SORT == order {
		sort.Slice(files, func(i, j int) bool {
			return files[i].CreatedAt.After(files[j].CreatedAt)
		})
	}

	if orderField == config.SORT_EXTENSION && config.ASC_SORT == order {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Extension < files[j].Extension
		})
	}
	if orderField == config.SORT_EXTENSION && config.DESC_SORT == order {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Extension < files[j].Extension
		})
	}

	for _, v := range files {
		fullFileName := v.Name + "." + v.Extension
		fmt.Printf("%v|%v|%v|%v|%v\n", fullFileName, v.Extension, v.Description, v.CreatedAt.Format("2006-01-02 15:04:05"), userName)
	}
}
