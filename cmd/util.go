package cmd

import (
	"encoding/json"
	"fmt"
	"log"
)

func printJSON(o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(b))
}
