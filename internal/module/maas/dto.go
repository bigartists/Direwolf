package maas

type MaasRequest struct {
	Name        string `json:"name" `
	Description string `json:"description"`
	Model       string `json:"maas" binding:"required"`
	ModelType   string `json:"model_type" binding:"required"`
	BaseURL     string `json:"base_url" binding:"required,url"`
	APIKey      string `json:"api_key" binding:"required"`
	Avatar      string `json:"avatar"`
}
