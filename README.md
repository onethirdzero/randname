# randname

Random name generator written in Go for practice. Inspired by [phobologic's version](https://github.com/phobologic/random_name/) and the [Docker implementation](https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go).

## Installation

```go
$ go get github.com/onethirdzero/randname
```

## Usage

```bash
# Assuming you have $GOPATH/bin added to your $PATH.
$ randname
> blurrycoral

# Short for --delimiter.
$ randname -d .
> pretty.sapphire

$ randname -d " "
> zippy rose
```

## Wishlist

- [x] Make this a cli app
- [ ] Pull words from a public API
- [ ] Add tests
- [ ] Make camel case an option?
- [ ] Make this a proper Go package
