package processor

import (
	"go-learn/chat_service/client/model"
	"go-learn/chat_service/client/utils"
)

type SmsProcessor struct {
}

func (sp *SmsProcessor) SendSms(userId string, content string) (err error) {
	sms := model.SmsMsg{
		From: self.UserId,
		To: userId,
		Content: content,
	}
	
	msgBytes, err := utils.MakeConnMessage(model.SMS_MSG_TYPE, sms)
	if err != nil {
		return err
	}

	err = utils.WriteConnMessage(self.Conn, msgBytes)
	if err != nil {
		return err
	}

	return
}

func (sp *SmsProcessor) BroadcastSms(content string) (err error) {
	sms := model.SmsMsg{
		From: self.UserId,
		Content: content,
	}
	
	msgBytes, err := utils.MakeConnMessage(model.SMS_MSG_TYPE, sms)
	if err != nil {
		return err
	}

	err = utils.WriteConnMessage(self.Conn, msgBytes)
	if err != nil {
		return err
	}

	return
}
