package conversation

import "github.com/google/wire"

var ConversationSet = wire.NewSet(
	ProvideConversationRepo,
	ProvideConversationService,
	ProvideConversationController,
)
