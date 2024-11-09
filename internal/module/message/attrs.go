package message

type MessageAttrFunc func(u *Message)

type MessageAttrFuncs []MessageAttrFunc

func (this MessageAttrFuncs) apply(u *Message) {
	for _, f := range this {
		f(u)
	}
}

func WithMessageID(id int64) MessageAttrFunc {
	return func(u *Message) {
		u.ID = id
	}
}

func WithConversationSessionID(conversationSessionID string) MessageAttrFunc {
	return func(u *Message) {
		u.ConversationSessionID = conversationSessionID
	}
}

func WithMessageType(messageType string) MessageAttrFunc {
	return func(u *Message) {
		u.MessageType = messageType
	}
}

func WithContent(content string) MessageAttrFunc {
	return func(u *Message) {
		u.Content = content
	}
}

func WithContext(context string) MessageAttrFunc {
	return func(u *Message) {
		u.Context = context
	}
}

func WithExpand(expand string) MessageAttrFunc {
	return func(u *Message) {
		u.Expand = expand
	}
}

func WithModelID(modelID *int64) MessageAttrFunc {
	return func(u *Message) {
		u.ModelID = modelID
	}
}

func WithSequenceNumber(sequenceNumber int) MessageAttrFunc {
	return func(u *Message) {
		u.SequenceNumber = sequenceNumber
	}
}

func WithParentQuestionID(parentId int64) MessageAttrFunc {
	return func(u *Message) {
		u.ParentQuestionID = parentId
	}
}
