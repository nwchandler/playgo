# Play-Go

This is a small utility for creating scratchpad Go modules. When you need
a quick Go module to test something out, but you would rather use your own
local development environment instead of the Go Playground, Play-Go can help!

## Installation

Because of the nature of this tool, we'll assume that you have Go installed 
and available on your `$PATH`. As such, the best installation method is:

`go install github.com/nwchandler/playgo/cmd/playgo@latest`

## Usage

Using Play-Go is (hopefully!) simple. If you run the command `playgo`, you
should get directory called `scratches/<some-silly-name>` created in one of 
the following locations (in order):

1. If you have `$XDG_DATA_HOME` set, `$XDG_DATA_HOME/scratches` will be used.
2. If you have a home directory set (for example, `$HOME` on Unix/Linux/Darwin
systems or `$USERPROFILE` or `%userprofile%` on Windows), then
`<home>/.local/share/scratches` will be used.
3. If neither of those are set, then the `scratches` directory will be put
into whatever your current working directory is.

> Note: Play-Go requires that you have Go installed and available on your `$PATH`.

## License

This project is licensed under the terms of the MIT license.
