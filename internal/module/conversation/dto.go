package conversation

type CreateConversationRequest struct {
	//UserID   int64   `json:"user_id" binding:"required"`
	Title   string  `json:"title" binding:"required"`
	MaasIds []int64 `json:"maas_ids" binding:"required,min=1"`
}
