package gsh

import (
	"encoding/csv"
	//"fmt"
	"io"
	"log"
	"os"
)

type Table struct {
	d []string
	s []int
}

func ReadCsv(u string, s ...int) *Table {
	f, err := os.Open(u)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)

	recs := make([]string, 0)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range rec {
			recs = append(recs, v)
		}
	}
	return &Table{
		d: recs,
		s: s,
	}
}

func (t *Table) PrintTable() {
	for {
		fmt.Println
	}
}

