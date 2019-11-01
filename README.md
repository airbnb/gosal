# Gosal (sal-client)

*Due to major changes in Sal4's checkin method, people who run Sal3 should checkout the Sal3 branch and build from there.*

## Overview

Gosal is intended to be a multi platform client for sal.

## Getting Started

Your configuration file should be `json` formatted as follows:

```json
{
  "key": "your gigantic machine group key",
  "url": "https://urltoyourserver.com/",
  "management": {
    "tool": "puppet",
    "path": "C:\\Program Files\\Puppet Labs\\Puppet\\bin\\puppet.bat",
    "command": "facts"
  }
}

```
# Running gosal
Gosal requires the configuration file to be passed in as an argument like so...

#### Windows Example
`gosal.exe --config "C:\path\to\config.json"`


# Building Sal3

To build the project after cloning:

```
make deps
make build
```

New macOS and windows binaries will be added to the `build/` directory.

# Building Sal4

Sal4 (master) is switching to go.mod!

## Dependencies

Gosal uses [dep](https://github.com/golang/dep#current-status) to manage external dependencies. Run `make deps` to install/update the required dependencies.
After adding a new dependency, run `dep ensure -update`, which will update the Gopkg.lock file. See [Adding a dependency](https://github.com/golang/dep#adding-a-dependency) for more.

## Formatting your code

Go has an exceptional formatter - please use it!
```
gofmt -s -w *.go
gofmt -s -w ./*/*.go
```
