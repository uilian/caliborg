# Calibre Metadata Organizer

A tool to organize metadata from a [Calibre](https://calibre-ebook.com/) library. Maybe a little bit overstating the solution it provides, but the intention is to keep improving it until achieves the desired outcome.

## Getting Started

Calibre provides several options for categorizing and tagging books programmatically. Here's how Caliborg tries to achieve the desired result:

1. Export Metadata
    * Use Calibre's `calibredb` command-line tool to export the metadata of all books into a JSON file (*which will be referenced by the environment variable `LIBRARY_PATH`*):

        ```bash
        <CALIBRE-INSTALL-PATH>/calibredb list --fields title,authors,author_sort,tags,isbn --for-machine > library.json
        ```

2. Analyze Metadata
    * Parse the metadata file and analyze the titles, authors, and descriptions to categorize books.

3. Set Categories
    * Use keywords in titles or descriptions to assign categories (e.g., "programming," "fiction," "history").
    * Match against a predefined list of authors or keywords for each category.

4. Apply Tags
    * Use `calibredb` set_metadata or a Calibre plugin to apply the generated tags or categories back to your library.
    * The program will generate a shell script containing several instructions following this template (one for each book):

        ```bash
        calibredb set_metadata --ids <book_id> --tags <tags>
        ```

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
