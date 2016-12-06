# vscale-metadata

Create Metadata service for Vscale

# Install from sources and Build

First you need to setup the project. It will download the golang tools and and vendor the dependencies project

```
make init
```

To build the server binary.

```
make build
```

Now you can run the unit-tests.

```
make test
```

Test coverage information

```
make cover
```

# Other commands

```
make help

  Usage:
    make <target>
  Targets:
    all                    Install tools and run the following targets: vendor, build, validate, test, docker
    build-docker           Build rest server docker image
    build                  Build rest server binary
    clean                  Clean project
    cover-extra            Run test coverage and generate report in standard output
    cover-xml              Run test coverage and generate report in _reports folder
    cover                  Run all test coverage
    docker                 Build docker images
    get-tools              Download all tools dependencies in hack sub-directory
    help                   Display list of targets
    init                   Init project
    run                    Run currencyconverter server process
    test-xml               Run all tests and generate reports in _reports folder
    test                   Run unit-tests
    validate               Run go style validation (golint)
    vendor                 Install all vendor dependencies
```

# Test with web browser

start the server with the following command

```
make run
```

```
open http://0.0.0.0:8080/metadata/v1/
```
