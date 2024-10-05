package whoami

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWhoami(t *testing.T) {
	expected := "danil\n"

	old := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	fmt.Println(Whoami())

	w.Close()
	os.Stdout = old

	got, err := io.ReadAll(r)
	if err != nil {
		t.Errorf("error reading from r")
	}

	userName := string(got)

	if len(userName) != len(expected) {
		t.Fatalf("len(userName) is of not expected length. expected=%d, got=%d",
			len(expected), len(userName))
	}

	if userName != expected {
		t.Fatalf("incorrect whoami. got=%v", string(got))
	}

}
