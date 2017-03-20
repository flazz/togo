# togo – convert any file to .go source

Generates a Go source file with a []byte var containing the given file's
contents.

Allows files to to be compiled (and linked) into a pkg. Quite handy for use
with [go generate](https://blog.golang.org/generate).

## Example:

Given some file `foo.txt`
```{.sh}
% echo "hello world" >foo.txt
```

Convert it via *togo*
```{.sh}
% togo -pkg bar -name foo -input ./foo.txt
```

A new file is created: `foo.txt.go`
```{.sh}
% ls foo.txt*
foo.txt
foo.txt.go
```

With a var inside
```{.sh}
% cat foo.txt.go
package bar

var foo = []byte{
	// 12 bytes from ./foo.txt
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0a,
}
```
