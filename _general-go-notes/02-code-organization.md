# B"H


### This is super super important link:
- Read over in-depth: https://golang.org/doc/code.html
- [Writing, building, installing, and testing Go code (5 min)](https://www.youtube.com/watch?v=XCsL89YtqCs)

---

To give you an idea of how a workspace looks in practice, here's an example:

```sh

bin/
    hello                          # command executable
    outyet                         # command executable
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

