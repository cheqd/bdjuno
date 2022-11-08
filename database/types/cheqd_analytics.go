package types

type AnalyticsItem struct {
	Height uint64 `json:"height,omitempty" db:"height"`
	Total  uint64 `json:"total" db:"total"`
}
