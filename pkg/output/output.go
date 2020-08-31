package output

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/kuangxc/child-naming/pkg/types"
)

const (
	defaultSheet = "Sheet1"
	firstColumn  = "A"
	secondColumn = "B"
	thirdColumn  = "C"
	fourthColumn = "D"
)

//SaveExcel save result to excel file
func SaveExcel(nameChn chan *types.NameInfo, limits int) {
	log.Println("start save result to excel.")
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet(defaultSheet)
	f.SetCellValue(defaultSheet, firstColumn+"1", "名字")
	f.SetCellValue(defaultSheet, secondColumn+"1", "相关名字")
	f.SetCellValue(defaultSheet, thirdColumn+"1", "来源")
	f.SetCellValue(defaultSheet, fourthColumn+"1", "意义")
	f.SetActiveSheet(index)

	for count := 0; count < limits; count++ {
		info, ok := <-nameChn
		if !ok {
			break
		}
		f.SetCellValue(defaultSheet, firstColumn+strconv.Itoa(count+2), info.Name)
		f.SetCellValue(defaultSheet, secondColumn+strconv.Itoa(count+2), info.RelatedNames)
		f.SetCellValue(defaultSheet, thirdColumn+strconv.Itoa(count+2), info.From)
		f.SetCellValue(defaultSheet, fourthColumn+strconv.Itoa(count+2), info.Meaning)
		fmt.Println(time.Now(), "current number is:", count)

	}

	// Save xlsx file by the given path.
	if err := f.SaveAs(time.Now().Format("2006-01-02") + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	log.Println("save result to excel finish.")
	os.Exit(0)
}

//SaveText save  result to text file
func SaveText(names []*types.NameInfo, limits int) {
	log.Println("start save result to text.")
	f, err := os.Create(time.Now().Format("2006-01-02") + ".txt")
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	title := fmt.Sprintf("%s  %s  %s  %s", "名字", "相关名字", "来源", "意义")
	fmt.Fprintln(w, title)
	for n, info := range names {
		if n >= limits {
			break
		}
		content := fmt.Sprintf("%s  %v  %s  %s", info.Name, info.RelatedNames, info.From, info.Meaning)
		fmt.Fprintln(w, content)
	}
	w.Flush()
	log.Println("save result to text finish.")
}
