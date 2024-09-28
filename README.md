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
should get directory called `scratches/<some-silly-name>` created in the
directory where you currently are. This location will be changing to a
more consistent location later on, in a new version, likely under
`$XDG_DATA_HOME`. Usage information will be updated at that time.

> Note: Play-Go requires that you have Go installed and available on your `$PATH`.

## License

This project is licensed under the terms of the MIT license.
