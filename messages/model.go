package messages

import (
	"time"
)

type Message struct {
	Email   string `json:"email"`
	MsgBody string `json:"body"`
	PubDate string `json:"pub_date"`
}

type Messages []Message

var pubDate time.Time

func formatTime(t time.Time) string {
	cst, _ := time.LoadLocation("Asia/Shanghai")
	return t.In(cst).Format("Mon 2006-1-2 3:04PM")
}
