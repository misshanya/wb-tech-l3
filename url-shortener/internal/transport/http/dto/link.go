package dto

type CreateShortLinkRequest struct {
	URL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ShortURL string `json:"short_url"`
}

type UserAgentStats struct {
	UserAgent string `json:"user_agent"`
	Count     int    `json:"count"`
}

type GetLinkStatisticsResponse struct {
	UserAgents []UserAgentStats `json:"user_agents"`
}
