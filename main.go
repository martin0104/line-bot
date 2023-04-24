package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

// var groupMap = make(map[string]int)

// type GroupCount struct {
// 	C int `json:"count" `
// }

// func getGroupCount(groupID string) bool {
// 	realTimeCount := goGetGroupCount(groupID)
// 	value, isExist := groupMap[groupID]
// 	// fmt.Println("realTimeCount", realTimeCount)
// 	// fmt.Println("value1", value)
// 	if isExist {
// 		if realTimeCount > value {
// 			groupMap[groupID] = realTimeCount
// 			fmt.Println("value2", groupMap[groupID])
// 			return true
// 		}
// 		groupMap[groupID] = realTimeCount
// 	} else {
// 		// fmt.Println("not in map and set groupID to map")
// 		groupMap[groupID] = realTimeCount
// 	}
// 	return false
// }

// func goGetGroupCount(groupID string) int {
// 	var gp GroupCount
// 	url := "https://api.line.me/v2/bot/group/" + groupID + "/members/count"
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Add("Authorization", "Bearer {Hvy37MFNxht9BmRRDr++k1HqAP4VH44sgYUjtxbpBM9YlRWX+cEBRsYOKvlkzbrvxOZ376VfsuOSGiQV6KFsk27kOl+jHSdMokY8L4zN/tS7R5Onumwm43n9BX2X6uqlJmu9aLaWtfKa2CSuEo8KjAdB04t89/1O/w1cDnyilFU=}")
// 	client := &http.Client{}
// 	resp, _ := client.Do(req)
// 	r, _ := ioutil.ReadAll(resp.Body)
// 	err := json.Unmarshal(r, &gp)
// 	if err != nil {
// 		fmt.Println("json err", err)
// 	}
// 	// fmt.Println("group member count", gp.C)
// 	resp.Body.Close()
// 	return gp.C
// }

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
		// log.Println("trigget callback")
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
			fmt.Println("trigger events type", event.Type)
			groupID := event.Source.GroupID
			fmt.Println("got group id", groupID)

			// checkMemberAdd := getGroupCount(groupID)
			// if checkMemberAdd {
			// 	fmt.Println("trigger join member response")
			// 	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎加入群組！")).Do(); err != nil {
			// 		log.Print(err)
			// 	}
			// }

			switch event.Type {
			case linebot.EventTypeMemberJoined:
				fmt.Println("trigger join member response 2")
				//加入人員清單取得function event.Joined.Members
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎各位老師/爸爸媽媽加入波雀絲小姐社團❤️這邊可以許願想買的教具或其他商品。(需要發票跟收據，也可以私訊)\n🌟目前記事本也有商品持續增加中，歡迎參觀選購。")).Do(); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if strings.HasPrefix(message.Text, "/") {
						switch strings.ToLower(message.Text) {
						case "/help":
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我可以幫你做很多事")).Do(); err != nil {
								log.Print(err)
							}
						case "/新增活動":
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請提供活動代碼")).Do(); err != nil {
								log.Print(err)
							}
						case "/joinGroup":
							if _, err := bot.PushMessage(C6f08f3f4dad5b138fe306b17697ce14d, linebot.NewTextMessage("歡迎各位老師/爸爸媽媽加入波雀絲小姐社團❤️這邊可以許願想買的教具或其他商品。(需要發票跟收據，也可以私訊)\n🌟目前記事本也有商品持續增加中，歡迎參觀選購。")).Do(); err != nil {
								log.Print(err)
							}
						default:
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("未知指令")).Do(); err != nil {
								log.Print(err)
							}
						}
					} else {
						fmt.Println("userID", event.Source.UserID)
						//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						//	log.Print(err)
						//}
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			default:
				fmt.Println("unknow event type ", event.Type)
				// if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我不懂你的明白！")).Do(); err != nil {
				// 	log.Print(err)
				// }
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
