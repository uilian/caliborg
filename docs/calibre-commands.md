## General Commands

### Show Help

```bash
calibredb --help
```

> Displays the full list of calibredb commands and options.

### List All Books

```bash
calibredb list
```

> Displays all books in the library.

### Count the Number of Books

```bash
calibredb list | wc -l
```

> Counts the number of books in the library (subtract 1 for the header row).

### Search for Books

```bash
calibredb search "search_query"
```

Example:

```bash
calibredb search "title:Golang"
# Finds books with "Golang" in the title.
```

### Show Details of a Book

```bash
calibredb show <book_id>
```

Example:

```bash
calibredb show 123
# Displays detailed metadata for the book with ID 123.
```

## Library Management

### Add Books

```bash
calibredb add /path/to/bookfile.epub
```

> Adds a book to the library.

### Remove a Book

```bash
calibredb remove <book_id>
```

Example:

```bash
calibredb remove 123
# Removes the book with ID 123.
```

### Export Books

```bash
calibredb export --all --to-dir /path/to/export/folder
```

> Exports all books in the library to the specified folder.

### Import Metadata

```bash
calibredb import --with-library /path/to/library /path/to/metadata.json
```

> Imports metadata from a JSON file into the library.

### Backup the Library

```bash
calibredb backup_metadata --with-library /path/to/library
```

> Backs up metadata to a JSON file in the library folder.

## Tag and Metadata Management

### List Tags

```bash
calibredb list --fields tags
```

> Displays all tags used in the library.

### Add Tags to a Book

```bash
calibredb set_metadata --ids <book_id> --tags "new_tag1, new_tag2"
```

Example:

```bash
calibredb set_metadata --ids 123 --tags "Golang, programming"
```

### Search by Tag

```bash
calibredb search "tags:programming"
```

> Lists books with the "programming" tag.

### Change Title or Author

```bash
calibredb set_metadata --ids <book_id> --title "New Title" --authors "New Author"
```

Example:

```bash
calibredb set_metadata --ids 123 --title "Golang Programming" --authors "John Doe"
```

## Search and Filtering

### Filter by Author

```bash
calibredb search "authors:John Doe"
```

### Filter by Format

```bash
calibredb search "formats:EPUB"
```

### Search by ISBN

```bash
calibredb search "isbn:9780134190440"
```

### Filter by Date Added

```bash
calibredb search "date:>2025-01-01"
```

> Lists books added after January 1, 2025.

### Combine Search Queries

```bash
calibredb search "tags:Golang and authors:John Doe"
```

## Server Management

### Start Content Server

```bash
calibredb server --with-library /path/to/library --port 8080
```

> Starts the Calibre Content Server on port 8080.
> Stop Content Server Stop the server by killing the process running it (use Ctrl+C in the terminal).

## Advanced Features

### Custom Queries Use advanced search queries to combine multiple conditions. Example:

```bash
calibredb search "(title:Golang or tags:programming) and date:>2025-01-01"
```

### List Specific Fields

```bash
calibredb list --fields title,authors,tags
```

> Displays only the specified fields.

### Repair Metadata

```bash
calibredb check_library
```

> Checks and repairs the library database if corrupted (works sometimes).

### Sort Output

```bash
calibredb list --sort-by title
```

Sorts the output by title.

For a more extensive list of the Calibre command line reference, please [refer to this page](https://manpages.ubuntu.com/manpages/focal/id/man1/calibredb.1.html)
