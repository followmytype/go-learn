package processor

import (
	"encoding/json"
	"fmt"
	"go-learn/chat_service/client/model"
	"go-learn/chat_service/client/utils"
	"net"
	"os"
)

const (
	SHOW_USERS = iota + 1
	SHOW_CHATS
	BROADCAST
	QUIT
)

var SP *SmsProcessor

func StartService() {
	SP = &SmsProcessor{}
	for {
		action := ShowMenu()
		switch action {
		case SHOW_USERS:
			showOnlineUsers(true)
		case SHOW_CHATS:
			// TODO: show history chat room
		case BROADCAST:
			content := ""
			fmt.Println("請輸入廣播內容：")
			fmt.Scanln(&content)
			SP.BroadcastSms(content)
		case QUIT:
			fmt.Println("掰")
			os.Exit(0)
		default:
			fmt.Println("重新輸入")
		}
	}
}

func ShowMenu() int {
	fmt.Println("歡迎\n" +
		"1. 顯示在線用戶\n" +
		"2. 訊息列表\n" +
		"3. 廣播\n" +
		"4. 退出\n" +
		"請選擇1-4:")
	var key int
	fmt.Scanln(&key)
	return key
}

func serverMonitor(conn net.Conn) {
	for {
		msg, err := utils.ReadConnMessage(conn)
		if err != nil {
			return
		}
		switch msg.Type {
		case model.ONLINE_NOTIFY_MSG_TYPE:
			var onlineNotifyMsg model.OnlineNotifyMsg
			if err := json.Unmarshal([]byte(msg.Data), &onlineNotifyMsg); err != nil {
				fmt.Println("用戶上線通知有問題")
			} else {
				updateUserStatus(onlineNotifyMsg.UserId, onlineNotifyMsg.Status)
			}
		case model.SMS_MSG_TYPE:
			var smsMsg model.SmsMsg
			if err := json.Unmarshal([]byte(msg.Data), &smsMsg); err != nil {
				fmt.Println("訊息錯誤")
			} else {
				if smsMsg.To == "" {
					fmt.Printf("『廣播』%s: %s\n", smsMsg.From, smsMsg.Content)
				} else if smsMsg.To == self.UserId {
					if smsMsg.From == self.ChatUserId {
						fmt.Printf("%s: %s\n", smsMsg.From, smsMsg.Content)
					} else {
						fmt.Printf("%s傳送新訊息\n", smsMsg.From)
						// TODO: save sms
					}
				}
			}
		}
	}
}
