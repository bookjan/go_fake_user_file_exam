#!/usr/bin/env bash

./run.out - << EOF
register user1
create_folder user1 'Work' 'The working files and necessary files are here'
create_folder user1 'Fun' 'The fun files'
create_folder user1 'Apple' 'The apple files'
create_folder user1 'Book' 'The book files'
create_folder user1 'Car' 'The Car files'
get_folders user1
get_folders user1 sort_name desc
get_folders user1 sort_time asc
get_folders user1 sort_time desc
rename_folders user1 1001 Jello
get_folders user1
upload_file user1 1002 '1.tc' 'first test case for a company'
delete_file user1 1002 1.tc
exit
EOF