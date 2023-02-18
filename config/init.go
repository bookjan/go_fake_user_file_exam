package config

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

type SortFolderByName []*Folder
type SortFolderByTime []*Folder

func (x SortFolderByName) Len() int {
	return len(x)
}
func (x SortFolderByName) Less(i, j int) bool {
	return x[i].Name < x[j].Name
}
func (x SortFolderByName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x SortFolderByTime) Len() int {
	return len(x)
}
func (x SortFolderByTime) Less(i, j int) bool {
	return x[i].CreatedAt.Before(x[j].CreatedAt)
}
func (x SortFolderByTime) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type SortFileByName []*File
type SortFileByTime []*File
type SortFileByExtension []*File

func (x SortFileByName) Len() int {
	return len(x)
}
func (x SortFileByName) Less(i, j int) bool {
	return x[i].Name < x[j].Name
}
func (x SortFileByName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x SortFileByTime) Len() int {
	return len(x)
}
func (x SortFileByTime) Less(i, j int) bool {
	return x[i].CreatedAt.Before(x[j].CreatedAt)
}
func (x SortFileByTime) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x SortFileByExtension) Len() int {
	return len(x)
}
func (x SortFileByExtension) Less(i, j int) bool {
	return x[i].Extension < x[j].Extension
}
func (x SortFileByExtension) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type UserMap map[string]*User
type FolderMap map[string]*Folder
type LabelMap map[string]*Label

type Arguments struct {
	Command   string
	Options   []string
	UserMap   UserMap
	FolderMap FolderMap
	LabelMap  LabelMap
}
