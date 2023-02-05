package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const MAX_USER_NUMBER int = 99999

var userMap = map[string]int{}

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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n# ") // Prompt

		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		fields := strings.Fields(text)
		if len(fields) < 2 {
			fmt.Println(text)
		} else {
			fmt.Println(fields)
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
	if userMap[userName] == 0 {
		rand.Seed(time.Now().UnixNano())
		userId := rand.Intn(MAX_USER_NUMBER) + 1
		userMap[userName] = userId

		fmt.Printf("Success")
	} else {
		fmt.Printf("Error - user already existing")
	}
}

func create_folder(args []string) {

}

func delete_folder(args []string) {

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
