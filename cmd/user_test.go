package cmd

import "testing"

func TestRegisterUser1(t *testing.T) {
	userName := "user1"
	var action Action = Action{
		Command:   Register,
		Options:   []string{"user1"},
		UserMap:   make(UserMap),
		FolderMap: make(FolderMap),
		LabelMap:  make(LabelMap),
	}

	action.Register()

	if userName != action.UserMap[userName].Base.Name {
		t.Error("wrong result")
	}
}
