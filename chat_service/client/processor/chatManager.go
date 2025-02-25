package processor

import (
	"fmt"
	"go-learn/chat_service/client/model"
)


func openRoom(userId string) {
	// TODO: show chat history
	self.ChatUserId = userId
	self.Status = model.BUSY
	// TODO: notify server change status
	fmt.Printf("開始與%s對話，輸入QUIT退出對話\n", userId)
	content := ""
	for {
		fmt.Scanln(&content)
		if content == "QUIT" {
			self.ChatUserId = ""
			self.Status = model.ONLINE
			break
		}
		if err := SP.SendSms(userId, content); err != nil {
			fmt.Println("訊息發送錯誤")
		}
	}
}
