## iGo, Indented Golang

This project provides indentation love to `golang`.

Most of the code comes directly from the standard library of `pkg/go` however the core part, `parser`
and `scanner` has been heavily modified, the *igo* `ast` instead is pretty much unchanged to allow
an easily swap with the original one.

**It's in alpha stage**

### How it works?

You can [TRY IT ONLINE](http://igo.herokuapp.com) or with the `cli`:

```
usage: igo [compile|parse|build] [flags] [path ...]
  -comments=true: print comments
  -dest="": destination directory
  -tabs=true: indent with tabs
  -tabwidth=8: tab width
$ igo parse # will convert any *.go file in *.igo
$ igo compile # will convert *.igo source code in *.go
```

Note that `build` currently is not yet implemented.

### Manually convert go code:

```python
import
	"bytes"
	"fmt"
	"github.com/daddye/igo/from_go"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"testing"

func ExampleFromGo()
  # Initialize the scanner.
  fset := token.NewFileSet() # positions are relative to fset

  const filename = "../ast/ast.go"

  src, err := ioutil.ReadFile(filename)
  if err != nil
    log.Fatalf("%s", err)


  file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
  if err != nil
    fmt.Println(err)
    return

  # Print the function body into buffer buf.
  # The file set is provided to the printer so that it knows
  # about the original source formatting and can add additional
  # line breaks where they were present in the source.
  var buf bytes.Buffer
  from_go.Fprint(&buf, fset, file)

  # Print the cleaned-up body text to stdout.
  fmt.Println(&buf)
```

### Manually convert iGo code:

```python
import
	"bytes"
	"fmt"
	"github.com/daddye/igo/parser"
	"github.com/daddye/igo/to_go"
	"github.com/daddye/igo/token"
	"io/ioutil"
	"log"
	"testing"

func ExampleToGo()
	# Initialize the scanner.
	fset := token.NewFileSet() # positions are relative to fset

	const filename = "../ast/ast.igo"

	src, err := ioutil.ReadFile(filename)
	if err != nil
		log.Fatalf("%s", err)

	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil
		fmt.Println(err)
		return

	# Print the function body into buffer buf.
	# The file set is provided to the printer so that it knows
	# about the original source formatting and can add additional
	# line breaks where they were present in the source.
	var buf bytes.Buffer
	to_go.Fprint(&buf, fset, file)

	# Print the cleaned-up body text to stdout.
	fmt.Println(&buf)
```

### How looks like?

Just browse `*.igo` files, but soon I'll make a `ebnf` notation.

### What will change?

Pretty much really few thins, `golang` itself is almost perfect, this parser will allow you to skip
some annoyance. Nothing more.

### What's left?

In my roadmap there is:

- Hidden builds (aka `igo build|run|test`)
- iGo format (aka `igo fmt`)
- iGo doc (aka `igo doc`)
- Expose `ast` (aka `little macros`)
- Expose `__filename__`, `__fname__`

### License MIT/BSD-style

Author Davide D'Agostino (@DAddYE) and `The Go Authors`
