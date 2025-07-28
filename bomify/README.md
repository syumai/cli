# bomify

* bomify is a simple command to add UTF-8 BOM (Byte Order Mark) to input.

## Installation

```
go install github.com/syumai/cli/bomify@latest
```

or

```
npm install -g bomify
# or without install
npx bomify
```

## Usage

```
cat file.csv | bomify > file_with_bom.csv
```

## License

MIT

## Author

syumai
