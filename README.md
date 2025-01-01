# Calibre Metadata Organizer

A tool to organize metadata from a [Calibre](https://calibre-ebook.com/) library. Maybe a little bit overstating the solution it provides, but the intention is to keep improving it until achieves the desired outcome.

## Getting Started

...

## Building and Testing

You'll need a `.env` file with the `GOOGLE_BOOKS_API_KEY` and `LIBRARY_PATH` environment variables. There's also a `DEBUG` environment variable that can be set to `true` to enable debug logs.

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
