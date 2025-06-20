package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const seperator byte = byte(',')


func main() {
	
	if len(os.Args) < 2 {
		fmt.Print("Requires a file name")
		return
	}
	f1Name := os.Args[1]
	processFile(f1Name)
}

func processFile(name string) {
	skuSales := map[string]map[string]int64{}
	days := map[string]int{}

	f := openFile(name)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		y := strings.Split(scanner.Text(), ",")
		plu := y[0]
		store := y[5]
		sales, _ := strconv.ParseInt(y[2], 10, 64)
		date := y[1]

		days[date]++
		getSkuSalesElement(skuSales, store)[plu]+=sales
	}

	type pluSales struct {
		plu string
		sales int64
	}

	totalDays := len(days)
	storeSales := map[string]int64{}
	storeTopSeller := map[string]pluSales{}
	for k, v := range skuSales {
		fmt.Printf("=====\n")
		fmt.Printf(k+"\n")
		storeTopSeller[k]=pluSales{"", 0}
		for k1, v1 := range v {
			fmt.Printf("-%s=%d\n",k1,v1)
			storeSales[k]+=v1
			if v1 > storeTopSeller[k].sales {
				storeTopSeller[k]=pluSales{k1, v1}
			}
		}
		fmt.Printf("===\n")
		fmt.Printf("Store: %s\n", k)
		fmt.Printf("Total Sales: %d\n", storeSales[k])
		fmt.Printf("Daily Average: %d\n", storeSales[k]/int64(totalDays))
		fmt.Printf("Top Seller: %s %d\n", storeTopSeller[k].plu, storeTopSeller[k].sales)
	}
}

func getSkuSalesElement(m map[string]map[string]int64, store string) map[string]int64 {
	if m[store] == nil {
		m[store] = map[string]int64{}
	}
	return m[store]
}

func openFile(id string) *os.File {
	f, err := os.OpenFile(id, os.O_RDONLY, os.FileMode(0666))
	if err != nil {
		panic(err)
	}
	f.Seek(0, 0)
	return f
}


