package handle

import (
	"github.com/NamDuongkiwi/GoLab/Lec4/models"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func (ebooks * models.Ebooks) GetTotalPages(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	lastPageLink, _ := doc.Find("ul.pagination li:last-child a").Attr("href") // Đọc dữ liệu từ thẻ a của ul.pagination
	if lastPageLink == "javascript:void();" { // Trường hợp chỉ có 1 page thì sẽ không có url
		ebooks.TotalPages = 1
		return nil
	}
	split := strings.Split(lastPageLink, "?page=")
	totalPages, _ := strconv.Atoi(split[1])
	ebooks.TotalPages = totalPages
	return nil
}