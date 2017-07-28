# Gosal (sal-client)

*Gosal is an alpha project, and should not be considered for production use at this time.  There is no support, pull requests accepted :)*

## Overview

Gosal is intended to be a multi platform client for sal.

## Getting Started

Currently gosal uses a `LoadConfig` funtion that accepts a path, and expects `json`.  For development, its fine to use a path local to the gosal binary/exe.

# Building

To build the project after cloning:

```
make deps
make build
```

New macOS and windows binaries will be added to the `build/` directory.

## Dependencies

Gosal uses [dep](https://github.com/golang/dep#current-status) to manage external dependencies. Run `make deps` to install/update the required dependencies. 
After adding a new dependency, run `dep ensure -update`, which will update the Gopkg.lock file. See [Adding a dependency](https://github.com/golang/dep#adding-a-dependency) for more.
