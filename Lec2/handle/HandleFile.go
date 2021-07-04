package handle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFromFile(fileName string) string{
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(file)
	return string(b)
}

func ConvertStringtoArray(context string) []int{
	context = strings.ReplaceAll(context, " ", "")
	text := strings.Split(context, ",")
	var array []int
	for _, value := range text{
		x, _ := strconv.Atoi(value)
		array = append(array, x)
	}
	return array
}

func ConvertArrayToString(array []int) string {
	x := ""
	for _, value := range array{
		x += strconv.Itoa(value) + " "
	}
	x = "[ " + x + "]"
	return x
}

func WriteToText(fileName string, array []int){
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString("Read data from txt: "+ ConvertArrayToString(array) + "\n")
	_, err = f.WriteString("Min of array: "+ strconv.Itoa(GetMinOfArray(array))+ "\n")
	_, err = f.WriteString("Max of array: "+ strconv.Itoa(GetMaxOfArray(array))+ "\n")
	_, err = f.WriteString("AVG of array: "+ strconv.FormatFloat(AvgOfArray(array),'f', 2, 32)+ "\n")
	_, err = f.WriteString("List of prime in array: "+ ConvertArrayToString(checkPrimeINArray(array))+ "\n")
	_, err = f.WriteString("Array after sort: "+ ConvertArrayToString(QuickSort(array)) + "\n")
}

