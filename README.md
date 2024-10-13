# mms-go

Write a Micromouse maze solver in Go.

For use with [mackorone/mms](https://github.com/mackorone/mms), a Micromouse simulator.

## Setup

1. Clone this repository
2. [Download the Micromouse simulator](https://github.com/mackorone/mms#download)
3. Run the simulator and click the "+" button to configure a new algorithm
4. Enter the config for your algorithm (name, directory, build command, and run command)
5. Click the "Run" button

## Example

A test algorithm has been implemented as example at `cmd/main.go`. To run it, create the following config in MMS:

- Name: Go example
- Directory: `/path/to/mms-go`
- Build command: `go build /path/to/mms-go/cmd/main.go`
- Run command: `./main`

## Notes

- Communication with the simulator is done via stdin/stdout, use stderr to print output
- Descriptions of all available API methods can be found at [mackorone/mms#mouse-api](https://github.com/mackorone/mms#mouse-api)
- The example code is a simple left wall following algorithm