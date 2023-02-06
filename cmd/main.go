package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"go_fake_user_file_exam/config"
	"go_fake_user_file_exam/handlers"
)

var userMap = make(map[string]config.User)
var folderMap = make(map[string]*config.Folder)

func Execute() {
	fmt.Println(`Go's fake user and file CLI program`)
	fmt.Println(`
Commonds: 
  register       {username}
  create_folder  {username} {folder_name} {description}
  delete_folder  {username} {folder_id}
  get_folders    {username} {sort_name | sort_time} {asc|dsc}
  rename_folder {username} {folder_id} {new_folder_name}
  upload_file    {username} {folder_id} {file_name} {description}
  delete_file    {username} {folder_id} {file_name}
  get_files      {username} {folder_id} {sort_name|sort_time|sort_extension} {asc|dsc}
  exit 
  `)

	re := regexp.MustCompile(`(?i)‘(.*?)’|'(.*?)'|([\S]+)`)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n\n# ") // Prompt

		scanner.Scan()
		text := scanner.Text()

		fmt.Println(text)

		if text == "exit" {
			break
		}

		fields := []string{}

		for _, match := range re.FindAllStringSubmatch(text, -1) {
			s := match[0]
			s = strings.ReplaceAll(s, "‘", "")
			s = strings.ReplaceAll(s, "’", "")
			s = strings.ReplaceAll(s, "'", "")
			s = strings.ReplaceAll(s, "'", "")
			fields = append(fields, s)
		}

		if len(fields) < 2 {
			fmt.Println(text)
		} else {
			mapCommandFunction(fields[0], fields[1:])
		}

	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func mapCommandFunction(commnd string, args []string) {
	switch commnd {
	case "register":
		handlers.Register(args, userMap)
	case "add_label":
		handlers.AddLabel(args, userMap)
	case "create_folder":
		handlers.CreateFolder(args, userMap, folderMap)
	case "delete_folder":
		handlers.DeleteFolder(args, userMap, folderMap)
	case "get_folders":
		handlers.GetFolders(args, userMap, folderMap)
	case "rename_folder":
		handlers.RenameFolder(args, userMap, folderMap)
	case "upload_file":
		handlers.UploadFile(args, userMap, folderMap)
	case "delete_file":
		handlers.DeleteFile(args, userMap, folderMap)
	case "get_files":
		handlers.GetFiles(args, userMap, folderMap)
	default:
		fmt.Printf("No command called: %s", commnd)
	}
}
