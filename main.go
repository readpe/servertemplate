// Server template layout inspired by "Mat Ryer - How I build HTTP services after eight years" talk
// https://github.com/matryer/2019-talks/tree/master/Mat%20Ryer%20-%20How%20I%20build%20HTTP%20services%20after%20eight%20years

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// minimize code in main, return error to main using this layout
func run() error {
	srv := newServer()
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		return err
	}
	return nil
}
