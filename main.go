package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// 建立 LINE Bot 的實體
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 設置 LINE Bot 的事件處理程序
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			// 判斷是否為加入群組事件
			if event.Type == linebot.EventTypeJoin {
				// 回覆歡迎訊息給用戶
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎加入群組！")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})

	// 啟動 LINE Bot 伺服器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
