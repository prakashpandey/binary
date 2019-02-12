# Binary

Convert text to binary (011111....)

## Why?

I was having some free time and was in fun mood. I and my friend used it to chat on whatsapp group.

### Download binaries

- Download the binaries for your os using the [link](https://github.com/prakashpandey/binary/tree/master/target).

### How to use it?

- **String to binary conversion**: 
    
    **command**: `./binary -s="mad on the mars"`
    
    **output**: `01101101 01100001 01100100 00100000 01101111 01101110 00100000 01110100 01101000 01100101 00100000 01101101 01100001 01110010 01110011`

- **Binary string to string conversion**: 
    
    **command**: `./binary -b="01101101 01100001 01100100 00100000 01101111 01101110 00100000 01110100 01101000 01100101 00100000 01101101 01100001 01110010 01110011"`
    
    **output**: `mad on the mars` 

[OPTIONAL] export it to your `PATH` for a convenient use

- `export PATH=$PATH:$(pwd)`

## For hackers

### Get the code

- `go get -u github.com/prakashpandey/binary`

### How to install?

- `go install github.com/prakashpandey/binary`

### Build it natively

- `cd $GOPATH/src/github.com/prakashpandey/binary`

- `go build`

## LICENSE

`The Unlicense`, please visit [LICENSE](LICENSE) for more information.
