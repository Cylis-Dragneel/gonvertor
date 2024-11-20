# Gonvertor (Name needs work ik)

This is a currency convertor app I made in Go to learn how to work with APIs.

## Features

- Convert from any currency to any other
- As fast as your internet

## Environment Variables

To run this project, you will need to add the following environment variables to your `.envrc`:

`API_KEY`

I used the API key from [Open Exchange Rates](openexchangerates.org) as it is free.

## Building

Build the project using `go build` you can later move it to a directory in `PATH`.

```bash
  go build -o gonvert main.go
```

## Usage/Examples

```bash
  ./gonvert 40 USD JPY
  ./gonvert 13560 INR EUR
```

## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `CODE OF CONDUCT.md`.

## TODO

- ~Be able to provide lowercase currency abbreviations as input~
- Turn it into a TUI application rather than a CLI

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
