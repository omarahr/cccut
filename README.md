
# cccut

`cccut` is a command-line tool written in Go that allows users to extract specific fields from each line of a file or standard input. It's similar to the Unix `cut` command.

---

## Features

- Extract specific fields from each line of a file or standard input.
- Supports custom delimiters for parsing fields.
---

## Installation

To use `cccut`, you need to have [Go](https://golang.org/) installed. Once Go is set up, you can build and run the program as follows:

```bash
go build -o cccut main.go
```

---

## Usage

```bash
cccut [options] [file]
```

### Options

- `-d`, `--delimiter` : Specify the field delimiter (default is tab `\t`).
- `-f`, `--field`     : Specify the fields to extract, separated by delimiters (e.g., `-f1,2` or `-f"1 2,3"`).

### Examples

1. **Extract specific fields from a file**:
    ```bash
    cccut -f1,3 -d"," input.txt
    ```

    Extracts the 1st and 3rd fields from `input.txt`, assuming fields are separated by commas.

2. **Extract fields from standard input**:
    ```bash
    cat input.txt | cccut -f2
    ```

    Extracts the 2nd field from each line of the piped input.

3. **Custom delimiter**:
    ```bash
    cccut -f2 -d"|" input.txt
    ```

    Extracts the 2nd field using `|` as a delimiter.
---

## Implementation

The tool uses the [Cobra](https://github.com/spf13/cobra) library for command-line argument parsing. It handles fields with a custom regular expression to support flexible input formats.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

## Acknowledgments

- Inspired the solution from [Coding Challenges by John Crickett](https://codingchallenges.fyi/challenges/challenge-cut)
