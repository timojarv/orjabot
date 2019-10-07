package data

import (
	"fmt"
	"time"
	"strings"

	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"
)

type Keitto struct {
	Name        string
	Date        int64
	KeittoIndex int
}

func AddKeitto(group int64, keitto Keitto) error {
	_, _, err := fs.Collection("keitot").Add(ctx, map[string]interface{}{
		"name":  keitto.Name,
		"group": chatID(group),
		"index": keitto.KeittoIndex,
		"date":  keitto.Date,
	})
	return err
}

func GetKeittoList(group int64) (*[]Keitto, error) {
	iter := fs.Collection("keitot").
		Where("date", ">", time.Now().Unix()).
		Where("group", "==", chatID(group)).
		OrderBy("date", firestore.Asc).
		Documents(ctx)

	keitot := []Keitto{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		keitto := doc.Data()

		keitot = append(keitot, Keitto{
			Name:        keitto["name"].(string),
			Date:        keitto["date"].(int64),
			KeittoIndex: int(keitto["index"].(int64)),
		})
	}

	return &keitot, nil
}

func (k Keitto) String() string {
	return fmt.Sprintf("%s (%s) %s", k.Name, time.Unix(k.Date, 0).Format("2.1.2006"), strings.Repeat("⭑", k.KeittoIndex))
}
