package data

import (
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Message is a chat message
type Message struct {
	Chat      string
	Message   string
	Username  string
	Author    string
	Timestamp time.Time
}

// Map converts Message into a map
func (msg *Message) Map() map[string]interface{} {
	return map[string]interface{}{
		"chat":      msg.Chat,
		"message":   msg.Message,
		"username":  msg.Username,
		"author":    msg.Author,
		"timestamp": msg.Timestamp.Unix(),
	}
}

// GetMessages returns a list of chat sessages
func GetMessages(group int64) (*[]Message, error) {
	iter := fs.Collection("messages").
		Where("chat", "==", chatID(group)).
		OrderBy("date", firestore.Asc).
		Documents(ctx)

	messages := []Message{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		message := doc.Data()

		messages = append(messages, Message{
			Chat:      message["chat"].(string),
			Message:   message["message"].(string),
			Username:  message["username"].(string),
			Author:    message["author"].(string),
			Timestamp: time.Unix(message["timestamp"].(int64), 0),
		})
	}

	return &messages, nil
}

// SaveMessage persists a message
func SaveMessage(msg Message) {
	_, _, err := fs.Collection("messages").Add(ctx, msg.Map())
	if err != nil {
		log.Fatal(err)
	}
}
