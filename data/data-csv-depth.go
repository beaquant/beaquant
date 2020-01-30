package data

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beaquant/utils"
	goex "github.com/nntaoli-project/GoEx"
	"log"
	"os"
	"strings"
	"time"

	"github.com/beaquant/beaquant/backtest"
)

// TickerEventFromCSVeData loads the market data from a SQLite database.
// It expands the underlying data struct.
type DepthEventFromCSVeData struct {
	backtest.Data
	fileDir string
}

func NewDepthEventFromCSVeData(fileDir string) *DepthEventFromCSVeData {
	return &DepthEventFromCSVeData{
		fileDir: fileDir,
	}
}

// Load single data events into a stream ordered by date (latest first).
func (d *DepthEventFromCSVeData) Load(fileName string) error {
	// check file location
	if len(d.fileDir) == 0 {
		return errors.New("no directory for data provided: ")
	}
	if fileName == "" {
		return errors.New("no file name for data provided: ")
	}
	if !strings.Contains(fileName, ".csv") {
		return errors.New("no csv file for data provided: " + fileName)
	}

	// read file for each fileName
	log.Printf("Loading %s file.\n", fileName)

	// open file for corresponding symbol
	lines, err := readDepthFromCSVFile(d.fileDir + fileName)
	if err != nil {
		return err
	}
	log.Printf("%v data lines found.\n", len(lines))

	// for each found record create an event
	for _, line := range lines {
		event, err := createDepthEventFromLine(line, fileName)
		if err != nil {
			log.Println(err)
		}
		d.Data.SetStream(append(d.Data.Stream(), event))
	}

	// sort data stream
	d.Data.SortStream()

	return nil
}

// readOrderbookFromCSVFile opens and reads a csv file line by line
// and returns a slice with a key/value map for each line.
func readDepthFromCSVFile(path string) (lines []map[string]string, err error) {
	log.Printf("Loading from %s.\n", path)
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create scanner on top of file
	reader := csv.NewReader(file)
	// set delimeter
	reader.Comma = ','
	// read first line for keys and fill in array
	keys, err := reader.Read()

	// read each line and create a map of values combined to the keys
	for line, err := reader.Read(); err == nil; line, err = reader.Read() {
		l := make(map[string]string)
		for i, v := range line {
			l[keys[i]] = v
		}
		// put found line as map into stream holder item
		lines = append(lines, l)
	}

	return lines, nil
}

// createTickerEventFromLine takes a key/value map and a string and builds a bar struct.
func createDepthEventFromLine(line map[string]string, symbol string) (dep backtest.DepthFrame, err error) {
	// parse each string in line to corresponding record value
	timestamp := utils.ToInt64(line["t"])
	date := time.Unix(0, timestamp*1000000)
	a := line["a"]
	b := line["b"]

	asks := make([][]float64, 0)
	err = json.Unmarshal([]byte(a), &asks)
	if err != nil {
		fmt.Println(err)
	}
	bids := make([][]float64, 0)
	err = json.Unmarshal([]byte(b), &bids)
	if err != nil {
		fmt.Println(err)
	}
	asklist := make(goex.DepthRecords, 0)
	for _, v := range asks {
		asklist = append(asklist, goex.DepthRecord{
			Price:  v[0],
			Amount: v[1],
		})
	}
	bidlist := make(goex.DepthRecords, 0)
	for _, v := range bids {
		bidlist = append(bidlist, goex.DepthRecord{
			Price:  v[0],
			Amount: v[1],
		})
	}
	// create and populate new event
	info := strings.Split(symbol, "_")

	mark := backtest.Mark{}
	mark.SetTime(date)
	mark.SetExchange(strings.ToUpper(info[1]))
	mark.SetSymbol(strings.ToUpper(info[2] + "_" + info[3]))

	dep = &backtest.Depth{
		Mark: mark,
		Depth: goex.Depth{
			AskList: asklist,
			BidList: bidlist,
		},
	}

	return dep, nil
}
