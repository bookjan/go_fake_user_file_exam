package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"go_fake_user_file_exam/util"
)

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

	var action Action = Action{
		UserMap:   make(UserMap),
		FolderMap: make(FolderMap),
		LabelMap:  make(LabelMap),
	}

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
			action.Command = fields[0]
			action.Options = fields[1:]
			mapCommandFunction(action)
		}

	}

	if scanner.Err() != nil {
		util.PrintOrLog(scanner.Err().Error(), util.Error)
	}
}

func mapCommandFunction(action Action) {
	switch action.Command {
	case "register":
		action.Register()
	case "add_label":
		action.AddLabel()
	case "get_labels":
		action.GetLabel()
	case "delete_label":
		action.DeleteLabel()

	case "create_folder":
		action.CreateFolder()
	case "delete_folder":
		action.DeleteFolder()
	case "get_folders":
		action.GetFolders()
	case "rename_folder":
		action.RenameFolder()
	case "add_folder_label":
		action.AddFolderLabel()
	case "delete_folder_label":
		action.DeleteFolderLabel()

	case "upload_file":
		action.UploadFile()
	case "delete_file":
		action.DeleteFile()
	case "get_files":
		action.GetFiles()
	default:
		fmt.Printf("No command called: %s", action.Command)
	}
}
