package data

import (
	"log"
)

// GetMembers returns a list of chat members
func GetMembers(group int64) []string {
	query := "SELECT name FROM chat_members WHERE chat = ?"
	members := []string{}

	rows, err := db.Query(query, group)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		members = append(members, name)
	}

	return members
}

// AddMember adds a member
func AddMember(group int64, name string) {
	query := "REPLACE INTO chat_members (chat, name) VALUES (?, ?)"
	_, err := db.Exec(query, group, name)
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
