package charachat_test

import (
	"github.com/ryomak/charachat-go"
	"reflect"
	"testing"
)

func TestLocalMemory(t *testing.T) {
	t.Run("AddMessage", func(t *testing.T) {
		mem := charachat.NewLocalMemory()
		msg := charachat.Message{Kind: charachat.MessageUserKindBot, Text: "Test message"} // Assuming UserKindA exists
		mem.AddMessage(msg)

		if !reflect.DeepEqual(mem.Messages[0], msg) {
			t.Errorf("got %v, want %v", mem.Messages[0], msg)
		}
	})

	t.Run("GetAllMessage", func(t *testing.T) {
		mem := charachat.NewLocalMemory()
		msg := charachat.Message{Kind: charachat.MessageUserKindBot, Text: "Test message"} // Assuming UserKindA exists
		mem.AddMessage(msg)

		messages := mem.GetAllMessage()

		if len(messages) != 1 || !reflect.DeepEqual(messages[0], msg) {
			t.Errorf("got %v, want %v", messages[0], msg)
		}
	})

	t.Run("DeleteAll", func(t *testing.T) {
		mem := charachat.NewLocalMemory()
		msg := charachat.Message{Kind: charachat.MessageUserKindBot, Text: "Test message"} // Assuming UserKindA exists
		mem.AddMessage(msg)
		mem.DeleteAll()

		if len(mem.Messages) != 0 {
			t.Errorf("expected memory to be empty, but got %v", mem.Messages)
		}
	})
}
