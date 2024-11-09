package internal

import (
	"github.com/bigartists/Direwolf/internal/module/conversation"
	"github.com/bigartists/Direwolf/internal/module/maas"
	"github.com/bigartists/Direwolf/internal/module/message"
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/google/wire"
)

var WholeSets = wire.NewSet(
	user.Userset,
	maas.MaasSet,
	conversation.ConversationSet,
	message.MessageSet,
)
