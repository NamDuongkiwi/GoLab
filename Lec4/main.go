package main

import (
	"github.com/NamDuongkiwi/GoLab/Lec4/handle"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	link := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	films := handle.GetFilmInfo(link)
	handle.Insertdb(films)
}