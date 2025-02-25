package processor

import (
	"encoding/json"
	"go-learn/chat_service/server/model"
	"go-learn/chat_service/server/utils"
)

type SmsProcessor struct {
}

var SP *SmsProcessor

func init() {
	SP = &SmsProcessor{}
}

func (sp *SmsProcessor) TransferClientMessage(msg *model.Message) (err error) {
	var smsMsg model.SmsMsg
	if err = json.Unmarshal([]byte(msg.Data), &smsMsg); err != nil {
		return err
	}
	msgJson, err := utils.MakeConnMessage(model.SMS_MSG_TYPE, smsMsg)
	if err != nil {
		return err
	}
	if smsMsg.To == "" {
		for _, up := range userMgr.GetAllOnlineUser() {
			if up.UserId != smsMsg.From {
				go utils.WriteConnMessage(up.Conn, msgJson)
			}
		}
		return nil
	}
	up, ok := userMgr.onlineUsers[smsMsg.To]
	if !ok {
		// TODO: send offline msg or not exists user check
		return nil
	}
	return utils.WriteConnMessage(up.Conn, msgJson)
}
