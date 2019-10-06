package data

import (
	"log"
)

// GetNostot returns a list of hatunnosto for the past 7 days
func GetNostot(group int64) map[string]int {
	query := `
	SELECT COUNT(name), name
	FROM nostot
	WHERE chat = ?
	AND created > DATETIME('NOW', '-7 day')
	GROUP BY name
	`
	nostot := map[string]int{}

	rows, err := db.Query(query, group)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		name string
		n int
	)
	for rows.Next() {
		if err := rows.Scan(&n, &name); err != nil {
			log.Fatal(err)
		}
		nostot[name] = n
	}

	return nostot
}

// AddNosto adds a hatunnosto
func AddNosto(group int64, name string) bool {
	if !IsMember(group, name) {
		return false
	}
	query := "INSERT INTO nostot (chat, name) VALUES (?, ?)"

	_, err := db.Exec(query, group, name)
	if err != nil {
		log.Fatal(err)
	}

	return true
}
