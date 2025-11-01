package models

type Link struct {
	ID    int64
	URL   string
	Short string
}

type UserAgentStats struct {
	UserAgent string
	Count     int
}

type LinkStatistics []UserAgentStats
