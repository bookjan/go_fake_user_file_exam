package config

import "time"

const SORT_NAME = "sort_name"
const SORT_TIME = "sort_time"
const SORT_EXTENSION = "sort_extension"
const ASC_SORT = "asc"
const DESC_SORT = "desc"

var USER_ID_BASE int = 0
var FOLDER_ID_BASE int = 1000

type Label struct {
	Name      string
	Color     string
	CreatedAt time.Time
}

type User struct {
	Id           string
	Name         string
	FolderIdMap  map[string]bool
	LabelNameMap map[string]bool
}

type File struct {
	Name        string
	Extension   string
	Description string
	CreatedAt   time.Time
}

type Folder struct {
	Id           string
	Name         string
	Description  string
	CreatedAt    time.Time
	FileMap      map[string]*File
	LabelNameMap map[string]bool
}
