package flagutil

import (
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCommaStringSlice(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	var (
		ss []string
	)

	fs.Var((*CommaStringSliceFlag)(&ss), "ss", "string slice flag")

	args := []string{
		"-ss", "a,b,c",
		"-ss", "d,e,f",
	}

	err := fs.Parse(args)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(ss, []string{"d", "e", "f"}) {
		t.Fatal("not correct output")
	}

}

func TestRepeatingStringSlice(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	var (
		ss []string
	)

	fs.Var((*RepeatingStringSliceFlag)(&ss), "ss", "string slice flag")

	args := []string{
		"-ss", "a,b,c",
		"-ss", "d,e,f",
	}

	err := fs.Parse(args)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(ss, []string{"a,b,c", "d,e,f"}) {
		t.Fatal("not correct output")
	}
	ok := false
	fs.Visit(func(f *flag.Flag) {
		s := f.Value.String()
		if !cmp.Equal(s, "a,b,c AND d,e,f") {
			t.Fatal("not expected", s)
		}
		ok = true
	})
	if !ok {
		t.Fatal("did not parse")
	}
}
