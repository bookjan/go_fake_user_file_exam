package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type User struct {
	id          string
	name        string
	folderIdMap map[string]bool
}

type Folder struct {
	id          string
	name        string
	description string
	createdAt   string
}

var USER_ID_BASE int = 0
var FOLDER_ID_BASE int = 1000

var userMap = make(map[string]User)
var folderMap = make(map[string]Folder)

func main() {
	fmt.Println(`Go's fake user and file CLI program`)
	fmt.Println(`
Commonds: 
  register       {username}
  create_folder  {username} {folder_name} {description}
  delete_folder  {username} {folder_id}
  get_folders    {username} {sort_name | sort_time} {asc|dsc}
  rename_folders {username} {folder_id} {new_folder_name}
  upload_file    {username} {folder_id} {file_name} {description}
  delete_file    {username} {folder_id} {file_name}
  get_files      {username} {folder_id} {sort_name|sort_time|sort_extension} {asc|dsc}
  exit 
  `)

	re := regexp.MustCompile(`(?i)‘(.*?)’|([\S]+)`)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n\n# ") // Prompt

		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		fields := []string{}

		for _, match := range re.FindAllStringSubmatch(text, -1) {
			s := match[0]
			s = strings.ReplaceAll(s, "‘", "")
			s = strings.ReplaceAll(s, "’", "")
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
		register(args)
	case "create_folder":
		create_folder(args)
	case "delete_folder":
		delete_folder(args)
	case "get_folders":
		get_folders(args)
	case "rename_folders":
		rename_folders(args)
	case "upload_file":
		upload_file(args)
	case "delete_file":
		delete_file(args)
	case "get_files":
		get_files(args)
	default:
		fmt.Printf("No command called: %s", commnd)
	}
}

func register(args []string) {
	userName := args[0]
	_, ok := userMap[userName]
	if !ok {
		USER_ID_BASE += 1
		userMap[userName] = User{
			id:          fmt.Sprint(USER_ID_BASE),
			name:        userName,
			folderIdMap: map[string]bool{},
		}

		fmt.Printf("Success")
	} else {
		fmt.Printf("Error - user already existing")
	}
}

func create_folder(args []string) {
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

	FOLDER_ID_BASE += 1
	folderId := fmt.Sprint(FOLDER_ID_BASE)
	folderMap[folderId] = Folder{
		id:          folderId,
		name:        folderName,
		description: description,
		createdAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	user.folderIdMap[folderId] = true

	fmt.Println(folderId)
}

func delete_folder(args []string) {
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

	_, ok = user.folderIdMap[folderId]
	if !ok {
		fmt.Println("Error - folder owner not match")
		return
	}

	delete(folderMap, folderId)
	delete(user.folderIdMap, folderId)

	fmt.Println(folderId)
}

func get_folders(args []string) {

}

func rename_folders(args []string) {

}

func upload_file(args []string) {

}

func delete_file(args []string) {

}

func get_files(args []string) {

}
