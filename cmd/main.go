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

	var args Arguments = Arguments{
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
			args.Command = fields[0]
			args.Options = fields[1:]
			mapCommandFunction(args)
		}

	}

	if scanner.Err() != nil {
		util.PrintOrLog(scanner.Err().Error(), util.Error)
	}
}

func mapCommandFunction(args Arguments) {
	switch args.Command {
	case "register":
		Register(&args)
	case "add_label":
		AddLabel(&args)
	case "get_labels":
		GetLabel(&args)
	case "delete_label":
		DeleteLabel(&args)

	case "create_folder":
		CreateFolder(&args)
	case "delete_folder":
		DeleteFolder(&args)
	case "get_folders":
		GetFolders(&args)
	case "rename_folder":
		RenameFolder(&args)
	case "add_folder_label":
		AddFolderLabel(&args)
	case "delete_folder_label":
		DeleteFolderLabel(&args)

	case "upload_file":
		UploadFile(&args)
	case "delete_file":
		DeleteFile(&args)
	case "get_files":
		GetFiles(&args)
	default:
		fmt.Printf("No command called: %s", args.Command)
	}
}
