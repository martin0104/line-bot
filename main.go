package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/line/line-bot-sdk-go/linebot"
)

var sm sync.Mutex
var groupMap = make(map[string]int)

// func getGroupCount(groupID string, b *linebot.Client) bool {
// 	value, isExist := groupMap[groupID]
// 	if isExist {
// 		realTimeCount := b.GetGroupMemberCount(groupID)
// 		if realTimeCount > value {

// 		}

// 	}
// 	b.GetGroupMemberCount(groupID)

// }

func goGetGroupCount(groupID string) {
	url := "https://api.line.me/v2/bot/group/" + groupID + "/members/count"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer {Hvy37MFNxht9BmRRDr++k1HqAP4VH44sgYUjtxbpBM9YlRWX+cEBRsYOKvlkzbrvxOZ376VfsuOSGiQV6KFsk27kOl+jHSdMokY8L4zN/tS7R5Onumwm43n9BX2X6uqlJmu9aLaWtfKa2CSuEo8KjAdB04t89/1O/w1cDnyilFU=}")
	client := &http.Client{}
	resp, _ := client.Do(req)
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("group member count", r)
	resp.Body.Close()
}

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
		log.Println("trigget callback")
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
			fmt.Println("trigger events")
			groupID := event.Source.GroupID
			fmt.Println("got group id", groupID)
			goGetGroupCount(groupID)

			// 判斷是否為加入群組事件
			if event.Type == linebot.EventTypeJoin {
				fmt.Println("trigger event response")
				// 回覆歡迎訊息給用戶
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎加入群組！")).Do(); err != nil {
					fmt.Println("trigger  response")
					log.Print(err)
				}
			} else {
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這不是加入群組訊息！")).Do(); err != nil {
					fmt.Println(event.Type)
					log.Print(err)
				}
			}
		}
	})

	// 啟動 LINE Bot 伺服器
	port := os.Getenv("PORT")
	if port == "" {
		port = "443"
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
