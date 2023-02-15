package orc_demo

import (
	"fmt"
	"github.com/scritchley/orc"
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	r, err := orc.Open("student.orc")

	if err != nil {
		log.Fatal(err)
	}

	selected := r.Schema().Columns()
	c := r.Select(selected...)
	defer c.Close()

	vals := make([]interface{}, len(selected))
	ptrVals := make([]interface{}, len(selected))
	strVals := make([]string, len(selected))
	for i := range vals {
		ptrVals[i] = &vals[i]
	}

	for c.Stripes() {
		for c.Next() {
			err := c.Scan(ptrVals...)
			if err != nil {
				log.Fatal(err)
			}
			for i := range ptrVals {
				strVals[i] = fmt.Sprint(ptrVals[i])
				log.Print(strVals[i])
			}
			log.Println("====")
		}
	}

	if err := c.Err(); err != nil {
		log.Fatal(err)
	}
}
