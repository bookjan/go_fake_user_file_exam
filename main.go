package main

import (
	"bufio"
	"fmt"
	"os"
)

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
  `)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("# ") // Prompt

		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		}

		if text == "exit" {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}
