package models
type Ebooks struct {
	TotalPages  int     `json:"total_pages"`
	TotalEbooks int     `json:"total_ebooks"`
	List        []Ebook `json:"ebooks"`
}