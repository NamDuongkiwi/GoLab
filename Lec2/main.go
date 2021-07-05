package main

import (
	"fmt"
	"github.com/NamDuongkiwi/GoLab/Lec2/handle"
)

const InputFile = "IOFile/input.txt"
const OutputFile = "IOFile/output.txt"

func main(){
	//cmt to up git
	context := handle.ReadFromFile(InputFile)
	array := handle.ConvertStringtoArray(context)

	//TODO: Bubble, Merge, Quick sort
	fmt.Println("Array after bubble sort: ", handle.BubbleSort(handle.ConvertStringtoArray(context)))
	fmt.Println("Array after merge sort:", handle.MergeSort(handle.ConvertStringtoArray(context)))
	fmt.Println("Array after quick sort:", handle.QuickSort(handle.ConvertStringtoArray(context)))

	//TODO: Write avg, min, max, array after sort to file
	handle.WriteToText(OutputFile, array)

	//TODO: Create api add, sub, mul, div (localhost:10000/add?a=&b=)
	handle.HandleRequests()
}
