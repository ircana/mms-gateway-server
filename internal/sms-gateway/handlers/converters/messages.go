package converters

import (
	"github.com/android-sms-gateway/client-go/smsgateway"
	"github.com/android-sms-gateway/server/internal/sms-gateway/modules/messages"
)

func MessageToMobileDTO(m messages.Message) smsgateway.MobileMessage {
	var message string
	var textMessage *smsgateway.TextMessage
	var dataMessage *smsgateway.DataMessage
	var mmsMessage *smsgateway.MmsMessage

	if m.TextContent != nil {
		message = m.TextContent.Text
		textMessage = &smsgateway.TextMessage{
			Text: m.TextContent.Text,
		}
	} else if m.DataContent != nil {
		dataMessage = &smsgateway.DataMessage{
			Data: m.DataContent.Data,
			Port: m.DataContent.Port,
		}
	} else if m.MmsContent != nil {
		// The phone reads the caption from mmsMessage.text; `message` stays
		// empty so an older client does not send the caption as a bare SMS.
		mmsMessage = &smsgateway.MmsMessage{
			Subject:     m.MmsContent.Subject,
			Text:        m.MmsContent.Text,
			Attachments: m.MmsContent.Attachments,
		}
	}

	return smsgateway.MobileMessage{
		Message: smsgateway.Message{
			ID:       m.ID,
			DeviceID: "",

			Message:     message,
			TextMessage: textMessage,
			DataMessage: dataMessage,
			MmsMessage:  mmsMessage,

			SimNumber:          m.SimNumber,
			WithDeliveryReport: m.WithDeliveryReport,
			IsEncrypted:        m.IsEncrypted,
			PhoneNumbers:       m.PhoneNumbers,
			TTL:                m.TTL,
			ValidUntil:         m.ValidUntil,
			ScheduleAt:         m.ScheduleAt,
			Priority:           m.Priority,
		},
		State:     smsgateway.ProcessingState(m.State),
		CreatedAt: m.CreatedAt,
	}
}

func MessageStateToDTO(state messages.MessageState) smsgateway.MessageState {
	return smsgateway.MessageState{
		ID:          state.ID,
		DeviceID:    state.DeviceID,
		State:       smsgateway.ProcessingState(state.State),
		IsHashed:    state.IsHashed,
		IsEncrypted: state.IsEncrypted,
		Recipients:  state.Recipients,
		States:      state.States,

		TextMessage:   state.TextContent,
		DataMessage:   state.DataContent,
		MmsMessage:    state.MmsContent,
		HashedMessage: state.HashedContent,
	}
}
