# Week 2

### Making an executable program :sunglasses:
1. Code the program in a `main.go` file
1. Test the code by running `go run main.go`
1. When happy, run `go build main.go` to compile the binary
1. `./main` to run the file!
1. Every time the `.go` file is updated, re-compile

### Pointers
- [Intro reading](https://www.golang-book.com/books/intro/8)
- Pointers are shortcuts to variables (like a shortcut on desktop)
- [Example code](notes/pointers.go)

### Go packages
- Split code into reusable chunks
- [Example packages](animals) and [packages in use](animals.go)

### Web server
Url example: `https://golang.org/`

- `https://` - protocol
- `golang.org` - unique address, like a PostCode
- `/` - root, like a 'drawer'
- [Example server](main.go)
