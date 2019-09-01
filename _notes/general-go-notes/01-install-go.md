# B"H


---

### Installing Go (Linux example):
- https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04


1. Get latest stable version number and SHA256 Checksum: 
    - https://golang.org/dl/
    - As of this writing: 
        - `go1.12.9`
        - `ac2a6efcc1f5ec8bdc0db0a988bb1d301d64b6d61b7e8d9e42f662fbb75a2b9b`


2. Run the following one line at a time:

```sh

cd ~

curl -O https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz

sha256sum go1.12.9.linux-amd64.tar.gz

tar xvf go1.12.9.linux-amd64.tar.gz
```


3. You should now have a directory called `go` in your home directory. 
    - Recursively change `go`'s owner and group to `root`, and move it to `/usr/local`:

```sh
sudo chown -R root:root ./go
sudo mv go /usr/local
```


---
---

## DO NOT DO FOLLOWING STEPS
- THIS WAS BEFORE I READ https://golang.org/doc/code.html


4. Add following to `~/.zshrc`:

```sh
export GOPATH=$HOME/go-sandbox
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

5. Test Install

```sh
mkdir $HOME/go-sandbox

cd $HOME/go-sandbox

mkdir -p src/test-repo/hello

cd src/test-repo/hello

touch hello.go
```

- Add following lines to `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

- Compile:

```sh
go install test-repo/hello
```

- Run:

```sh
hello

which hello
```

- ...and some more:

```sh
go version

go cd ~/go-sandbox

tree .
```

---

### Notes:

#### `go build` vs `go install`:
- https://stackoverflow.com/questions/30612611/what-does-go-build-build


