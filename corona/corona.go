package corona

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
)

// func main() {
// 	fileUrl := "https://covid19.mhlw.go.jp/public/opendata/newly_confirmed_cases_daily.csv"

// 	if err := DownloadFile("corona.csv", fileUrl); err != nil {
// 		panic(err)
// 	}

// 	file, err := os.Open("test.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	var line []string
// 	displayList := map[string]string{"ALL": "全国", "Aichi": "愛知", "Tokyo": "東京"}
// 	t := time.Date(2021, 8, 9, 4, 5, 6, 0, time.Local)
// 	var retval string

// 	for {
// 		line, err = reader.Read()
// 		if err != nil {
// 			break
// 		}
// 		isSelectedDate := t.Format("2006/1/2") == line[0]
// 		_, isSelectedRegion := displayList[line[1]]
// 		if isSelectedDate && isSelectedRegion {
// 			retval += displayList[line[1]] + "の感染者数は" + line[2] + "人です\n"
// 		}
// 	}

// 	retval += "本日も感染予防をして過ごしましょう\n"
// 	fmt.Println(retval)
// }

func DisPlayTodayCorona(filename string, date string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string
	displayList := map[string]string{"ALL": "全国", "Aichi": "愛知", "Tokyo": "東京"}
	var retval string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		isSelectedDate := date == line[0]
		_, isSelectedRegion := displayList[line[1]]
		if isSelectedDate && isSelectedRegion {
			retval += displayList[line[1]] + "の感染者数は" + line[2] + "人です\n"
		}
	}

	retval += "本日も感染予防をして過ごしましょう\n"
	return retval
}

func DownloadFile(filepath string, url string) error {
	// filepath = "../" + filepath
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
