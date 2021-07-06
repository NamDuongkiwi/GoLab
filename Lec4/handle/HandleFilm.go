package handle

import (
	"fmt"
	"github.com/NamDuongkiwi/GoLab/Lec4/database"
	"github.com/NamDuongkiwi/GoLab/Lec4/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetFilmInfo(pathURL string) ([]models.Film) {
	doc, err := goquery.NewDocument(pathURL)
	if err != nil {
		log.Println(err)
	}
	infoList := make([]models.Film, 0)
	doc.Find("table tbody").Each(func (index int, tableHtml *goquery.Selection){
		var info models.Film
		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
			row := make([]string, 0)
			rowHtml.Find("td").Each(func(indexTd int, tableCell *goquery.Selection) {
				row = append(row, tableCell.Text())
				tableCell.Find("a").Each(func(indexA int, imgTag *goquery.Selection) {
					imgTag.Find("img").Each(func(indexImg int, attribute *goquery.Selection) {
						info.Poster, _ = attribute.Attr("src")
					})
				})
			})
			info.Title = strings.ReplaceAll(strings.TrimSpace(row[1]),"\n","")
			info.Rating = row[2]
			infoList = append(infoList, info)
		})
	})
	return infoList
}
func Insertdb(films []models.Film) {
	//var mutex = &sync.Mutex{}
	fmt.Println(len(films))

	db := database.Connect()
	defer db.Close()
	_, _ = db.Query("CREATE TABLE IF NOT EXISTS Movies(id int NOT NULL AUTO_INCREMENT PRIMARY KEY, Title varchar(100), Rating varchar(40),Url varchar(180));")
	for _, k := range films{
		_, err := db.Query("INSERT INTO Movies (Title, Rating, Url) VALUES (?, ?, ?)", k.Title, k.Rating, k.Poster)
		if err != nil {
			panic(err)
		}
	}
}
