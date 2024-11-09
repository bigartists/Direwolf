package conversation

type CreateConversationRequest struct {
	UserID   int64   `json:"user_id" binding:"required"`
	Title    string  `json:"title" binding:"required"`
	ModelIDs []int64 `json:"model_ids" binding:"required,min=1"`
}