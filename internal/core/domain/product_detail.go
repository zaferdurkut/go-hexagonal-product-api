package domain

type ProductDetail struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
	FetchTime string `json:"fetch_time"`
}
