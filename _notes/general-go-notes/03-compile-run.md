# B"H




### Confirm `GOPATH env variable`:

```sh
echo $GOPATH
```

---

### Test the files in the `stringutil` package:

```sh
go test github.com/Ylazerson/go-shenanigans/stringutil
```

---

### Compile all the files in the `stringutil` package:
- This won't produce an output file. Instead it saves the compiled package in the local build cache.

```sh
go build github.com/Ylazerson/go-shenanigans/stringutil
```

---

### Install and run the hello program:

```sh
go install github.com/Ylazerson/go-shenanigans/hello

hello

which hello
```


---

### See resulting files:

```sh
cd ~/repos/go-workspace

tree -a bin & tree -a pkg
```