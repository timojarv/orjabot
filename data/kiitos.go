package data

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// AddKiitos adds a kiitos
func AddKiitos() (int64, error) {
	var kiitokset interface{}
	ref := fs.Collection("kiitos").Doc("kiitos")
	err := fs.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref)
		if err != nil {
			log.Println(err)
			return err
		}
		kiitokset, err = doc.DataAt("kiitokset")
		if err != nil {
			return err
		}
		return tx.Set(ref, map[string]interface{}{
			"kiitokset": kiitokset.(int64) + 1,
		})
	})
	if err != nil {
		return 0, err
	}

	return kiitokset.(int64) + 1, nil
}
