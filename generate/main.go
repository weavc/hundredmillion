package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const number int = 100000000

func main() {
	n := number
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	}

	g := newGenerator()
	uid, _ := uuid.NewV7()
	id := uid.String()
	f := openFile(id)
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < n; i++ {
		_, err := w.WriteString(g.generateLine())
		if err != nil {
			fmt.Print(err)
		}
	}
	err := w.Flush()
	if err != nil {
		fmt.Print(err)
	}
}

func openFile(id string) *os.File {
	f, err := os.OpenFile("./out/"+id, os.O_CREATE|os.O_WRONLY, os.FileMode(0666))
	if err != nil {
		panic(err)
	}
	f.Seek(0, 0)
	return f
}

type generator struct {
	inc     int
	dateInc time.Time
	start   int
	end     int
}

func newGenerator() *generator {
	return &generator{
		inc:     0,
		dateInc: time.Unix(1704067200, 0),
		start:   1000,
		end:     4000,
	}
}

func (g *generator) generateLine() string {
	uid, _ := uuid.NewV7()
	g.inc++
	var days int = g.inc / ((g.end - g.start) * 200)
	var stores int = 100 + (g.inc / (g.end - g.start) % 200)

	return fmt.Sprintf(
		"%d,%s,%d,%d,%d,%d,%s\n",
		g.start+(g.inc%(g.end-g.start)),
		g.dateInc.Add(time.Hour*time.Duration(24*days)).String(),
		rand.IntN(1000),
		rand.IntN(10000),
		rand.Int(),
		100+stores,
		uid.String())
}
