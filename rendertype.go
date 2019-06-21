package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/oberlademeister/replacer"
)

// RenderType renders the type
func (t *T) RenderType(out io.Writer) {
	fmt.Fprintf(out, "package %s\n\n", t.PackageName)
	fmt.Fprint(out, "import \"strings\"\n\n")
	fmt.Fprintf(out, "// %s %s\n", t.Name, t.TypeDescription)
	fmt.Fprintf(out, "type %s int\n", t.Name)
	fmt.Fprintf(out, "// Num%s is the number of %s variants\nconst Num%s = %d\n\n", t.Name, t.Name, t.Name, len(t.Values))
	fmt.Fprint(out, "// enum constants generated by gstringer\n")
	fmt.Fprint(out, "const (\n")
	first := true
	for _, v := range t.Values {
		if first {
			fmt.Fprintf(out, "    %s %s = %d\n", v.Symbol, t.Name, v.Index)
			first = false
			continue
		}
		fmt.Fprintf(out, "    %s = %d\n", v.Symbol, v.Index)
	}
	fmt.Fprint(out, ")\n")
}

// RenderString renders the String() function
func (t *T) RenderString(out io.Writer) {
	fmt.Fprint(out, "// String satisfies Stringer interface\n")
	fmt.Fprintf(out, "func (t %s) String() string {\n", t.Name)
	fmt.Fprint(out, " switch t {")
	for _, v := range t.Values {
		fmt.Fprintf(out, "  case %s:\n", v.Symbol)
		fmt.Fprintf(out, "  return %q\n", v.Symbol)
	}
	fmt.Fprint(out, " }\n")
	fmt.Fprintf(out, " return %q\n", t.UnknownDescription)
	fmt.Fprint(out, "}\n")
}

// RenderVerboseString renders the String() function
func (t *T) RenderVerboseString(out io.Writer) {
	fmt.Fprint(out, "// VerboseString like String, but more verbose\n")
	fmt.Fprintf(out, "func (t %s) VerboseString() string {\n", t.Name)
	fmt.Fprint(out, " switch t {")
	for _, v := range t.Values {
		fmt.Fprintf(out, "  case %s:\n", v.Symbol)
		fmt.Fprintf(out, "  return %q\n", v.VerboseDescription)
	}
	fmt.Fprint(out, " }\n")
	fmt.Fprintf(out, " return %q\n", t.UnknownDescription)
	fmt.Fprint(out, "}\n")
}

// RenderExTypeMaker creates and initializes the ExtypeMaker
func (t *T) RenderExTypeMaker(out io.Writer) {
	var symbol, symbolLC, descr, descrLC string
	{
		var buf bytes.Buffer
		for _, v := range t.Values {
			fmt.Fprintf(&buf, "  %q: %s,\n", v.Symbol, v.Symbol)
		}
		symbol = buf.String()
	}
	{
		var buf bytes.Buffer
		for _, v := range t.Values {
			fmt.Fprintf(&buf, "  %q: %s,\n", strings.ToLower(v.Symbol), v.Symbol)
		}
		symbolLC = buf.String()
	}
	{
		var buf bytes.Buffer
		for _, v := range t.Values {
			fmt.Fprintf(&buf, "  %q: %s,\n", v.VerboseDescription, v.Symbol)
		}
		descr = buf.String()
	}
	{
		var buf bytes.Buffer
		for _, v := range t.Values {
			fmt.Fprintf(&buf, "  %q: %s,\n", strings.ToLower(v.VerboseDescription), v.Symbol)
		}
		descrLC = buf.String()
	}
	fmt.Fprint(out, replacer.Positional(`
// {0}Maker is a factory type to make ExTypes from string
type {0}Maker struct {
	fromSymbol               map[string]{0}
	fromSymbolLowerCase      map[string]{0}
	fromDescription     	 map[string]{0}
	fromDescriptionLowerCase map[string]{0}
}

// NewExTypeMaker initialises an ExTypeFactory
func NewExTypeMaker() *{0}Maker {
	return &{0}Maker{
		fromSymbol: map[string]{0}{
{1}
		},
		fromSymbolLowerCase: map[string]{0}{
{2}
		},
		fromDescription: map[string]{0}{
{3}
		},
		fromDescriptionLowerCase: map[string]{0}{
{4}
		},
	}
}
`, false, t.Name, symbol, symbolLC, descr, descrLC))
}

// RenderFromString renders the FromString Function
func (t *T) RenderFromString(out io.Writer) {
	fmt.Fprint(out, replacer.Positional(`
// FromString creates a {0} from string
// This could be done better using a trie
func  (t *{0}Maker) FromString(s string, verbose, cases bool) {0} {
var key string
var m map[string]{0}

if cases {
	key = s
	if verbose {
		m = t.fromDescription
	} else {
		m = t.fromSymbol
	}
} else {
 	key = strings.ToLower(s) 
	if verbose {
		m = t.fromDescriptionLowerCase
	} else {
		m = t.fromSymbolLowerCase
	}
}
if val, ok := m[key]; ok {
	return val
}

return -1

}

`, false, t.Name))
}
