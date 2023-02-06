#!/usr/bin/env bash

./run.out - << EOF
register user1
create_folder user1 ‘Work’ ‘The working files and necessary files are here’
get_folders user1
rename_folder user1 1001 ‘Temp’
delete_folder user1 1001
create_folder user1 ‘Testing’ ‘The testing folders’
upload_file user1 1002 ‘1.tc’ ‘first test case for a company’
upload_file user1 1002 ‘1.png’ ‘the picture for first test case’
get_files user1 1002 sort_extension asc
delete_file user1 1002 1.png
delete_file user1 1001 1.tc
exit
EOF