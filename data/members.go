package data

import (
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
)

// GetMembers returns a list of chat members
func GetMembers(group int64) []string {
	dsnap, err := fs.Collection("groups").Doc(chatID(group)).Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	members := []string{}
	for member := range dsnap.Data() {
		members = append(members, member)
	}

	return members
}

// AddMember adds a member
func AddMember(group int64, name string) {
	newName := make(map[string]interface{})
	newName[name] = true
	_, err := fs.Collection("groups").Doc(chatID(group)).Set(ctx, newName, firestore.MergeAll)
	if err != nil {
		log.Fatal(err)
	}
}

// IsMember checks if name is member of chat
func IsMember(group int64, name string) bool {
	members := GetMembers(group)

	for _, member := range members {
		if member == name {
			return true
		}
	}

	return false
}

func chatID(i int64) string {
	return strconv.Itoa(int(i))
}
