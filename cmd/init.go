package cmd

import (
	"time"
)

const SORT_NAME = "sort_name"
const SORT_TIME = "sort_time"
const SORT_EXTENSION = "sort_extension"
const ASC_SORT = "asc"
const DESC_SORT = "desc"

var USER_ID_BASE int = 0
var FOLDER_ID_BASE int = 1000

type Base struct {
	Id          string
	Name        string
	Description string
	CreatedAt   time.Time
}

type Label struct {
	Base
	Color string
}

type UserMap map[string]*User
type FolderMap map[string]*Folder
type LabelMap map[string]*Label

type Action struct {
	Command   string
	Options   []string
	UserMap   UserMap
	FolderMap FolderMap
	LabelMap  LabelMap
}
