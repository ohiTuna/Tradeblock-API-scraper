package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	//	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// One fee in the JSON response
type AutoGenerated struct {
	Data []Stat `json:"data"`
}

type Stat struct {
	Timestamp int         //`json:"num_int"`
	Fee    json.Number `json:"num_float"`
}

func (s *Stat) UnmarshalJSON(data []byte) error {
	var tmp []int
	if s.Timestamp < 6000 {
		json.Unmarshal(data, &tmp)
	}
		s.Timestamp = tmp[0]
	
	var blk []json.Number
	if err := json.Unmarshal(data, &blk); err != nil {
		return err
	}
	s.Fee = blk[1]

	return nil
}

func main() {
	
	fmt.Println("Enter data type: txsize_new, txfee_new, txes_new, blocks_new")
    var inType string
    fmt.Scanf("%s", &inType)
    fmt.Println("Enter interval to query: 1h, 2h, 6h, 1d")
    var inInterval string
    fmt.Scanf("%s", &inInterval)
    fmt.Printf("Enter latest date value to pull: 1487008800 = Feb 13 2017 6PM, 86400=24hours")
    var inDate string
    fmt.Scanf("%s", &inDate)
    time.Sleep(2000 * time.Millisecond)
    
    var bufUrl bytes.Buffer
    bufUrl.WriteString("https://tradeblock.com/api/blockchain/statistics/")
    bufUrl.WriteString(inType)
    bufUrl.WriteString("/")
    bufUrl.WriteString(inInterval)
    bufUrl.WriteString("/")
    bufUrl.WriteString(inDate)
    rdyUrl := bufUrl.String()
    
	hour, min := time.Now().Hour(), time.Now().Minute()
	_, month, day := time.Now().Date()
	h := strconv.Itoa(hour)
	n := strconv.Itoa(min)
	d := strconv.Itoa(day)
	m := strconv.Itoa(int(month))
	var zuffer bytes.Buffer
	/*if inType  = "txsize_new" {
		zuffer.WriteString("txsize")
	}
	if inType = "txfee_new"	{
		zuffer.WriteString("txfee")
		}
	if inType = "txes_new" {
		zuffer.WriteString("txes")
		}
	*/
	zuffer.WriteString(inType)
	zuffer.WriteString(inInterval)
	zuffer.WriteString("-")
	zuffer.WriteString(m)
	zuffer.WriteString("-")
	zuffer.WriteString(d)
	zuffer.WriteString("at")
	zuffer.WriteString(h)
	zuffer.WriteString(n)
	zuffer.WriteString(".csv")
	xx := zuffer.String()

	file, err := os.Create(xx)
	if err != nil {
		fmt.Println("Cannot create")
	}
	defer file.Close()
	csvw := csv.NewWriter(file)
	ticker := time.NewTicker(time.Second * 6)

	go func() {
		for range ticker.C {
			response, err := http.Get(rdyUrl)
			if err != nil {
				log.Fatalf("failed to get JSON data: %v", err)
			}
			defer response.Body.Close()
			dec := json.NewDecoder(response.Body)
			dec.UseNumber()
			fees := AutoGenerated{}
			err = dec.Decode(&fees)
			if err != nil {
				log.Fatalf("Failed to decode the JSON body: %v", err)
			}
			noter := string("looks good")
			fmt.Printf(noter)
			csvw.Flush()
			var counter int = 1

			for _, fee := range fees.Data {
				datecol := string("=(A1/24/60/60)+DATE(1970,1,1)")
				counter = counter + 1
				numFee := fee.Fee.String()
				record := []string{strconv.Itoa(fee.Timestamp), (numFee), (datecol), strconv.Itoa(counter)}
				
				csvw.Flush()
				csvw.Write(record)
				
				if err := csvw.Error(); err != nil {
					log.Fatalf("Could not write to CSV file:", err)
				}
			}
		}
	}()

	time.Sleep(time.Second * 11)
	ticker.Stop()
}
