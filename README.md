# ccat

A simple command-line tool written in Go for processing text input and files, inspired by Unix utilities. It handles standard input, file concatenation, and line numbering.

## Features
- **Standard Input**:
  - Print input as-is (`-`).
  - Number non-empty lines (`-b`).
  - Number all lines (`-n`).
- **File Handling**:
  - Read and print single file contents.
  - Concatenate and print multiple files.

## Prerequisites
- [Go](https://golang.org/doc/install) installed on your system.

## Usage

### Clone the project
```bash
    git clone https://github.com/TG199/ccat.git
```

### Build the Program
```bash
make build
```
### Testing
```bash
make test
```

### Example commands
1. Read from standard in
```bash
./bin/ccat
```
2. Process files
    
    Single file
    ```bash
    ./bin/ccat test.txt
    ```
    Multiple files
    ```bash
    ./bin/ccat test.txt test2.txt ...
    ```
3. Number lines

    Number lines as they are printed
    ```bash
    head -n3 test.txt | ./bin/ccat -n
    ```
    Number empty lines
    ```bash
    sed G test.txt | ./bin/ccat -n | head -n4
    ```
    Exclude empty lines
    ```bash
    sed G test.txt | ./bin/ccat -b | head -n5
    ```

### Clean build artifacts
```bash
make clean
```

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue.

## License
This project is open-source and available under the MIT License.








