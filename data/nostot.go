package data

import (
	"log"
	"time"

	"google.golang.org/api/iterator"
)

// GetNostot returns a list of hatunnosto for the past 7 days
func GetNostot(group int64) map[string]int {
	nostot := map[string]int{}
	// 7 days
	interval := int64(7 * 24 * 60 * 60)

	iter := fs.Collection("nostot").
		Where("group", "==", chatID(group)).
		Where("timestamp", ">", time.Now().Unix()-interval).
		Documents(ctx)

	var name string
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		name = doc.Data()["name"].(string)
		nostot[name]++
	}

	return nostot
}

// AddNosto adds a hatunnosto
func AddNosto(group int64, name string) bool {
	if !IsMember(group, name) {
		return false
	}
	_, _, err := fs.Collection("nostot").Add(ctx, map[string]interface{}{
		"name":      name,
		"timestamp": time.Now().Unix(),
		"group":     chatID(group),
	})
	if err != nil {
		log.Fatal(err)
	}

	return true
}
