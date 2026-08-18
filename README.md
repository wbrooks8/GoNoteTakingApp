# Structs-Practice

A small Go example that demonstrates using structs and a simple `note` package to read/write JSON-backed notes.

## Repository structure

- `main.go` - application entrypoint.
- `go.mod` - Go module file.
- `firstnote.json` - example JSON note used by the app.
- `note/` - package implementing note-related types and helpers ([note/note.go](note/note.go)).

## Prerequisites

- Go 1.18 or newer installed. Verify with:

```bash
go version
```

## Build & Run

To run the app directly:

```bash
go run .
```

To build a binary:

```bash
go build -o structspractice
./structspractice
```

## Usage

The app ships with an example note file, `firstnote.json` ([firstnote.json](firstnote.json)). Edit the JSON to try different inputs. The `note` package ([note/note.go](note/note.go)) contains the data structures and helpers used by `main.go`.

## Contributing

Feel free to open issues or submit pull requests. Keep changes small and focused.

## License

This project is unlicensed. Add a license if you plan to publish it.
