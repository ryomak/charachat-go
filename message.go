package charachat

type MessageUserKind string

const (
	MessageUserKindUser MessageUserKind = "user"
	MessageUserKindBot  MessageUserKind = "bot"
)

type Message struct {
	Kind MessageUserKind
	Text string
}
