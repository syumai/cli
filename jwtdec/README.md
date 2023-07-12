# jwtdec

* jwtdec decodes JWT's header and payload parts and prints them.

## Usage

* example JWT (copied from https://jwt.io)

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

```
$ cat jwt.txt | jwtdec # jwt.txt contains JWT above
{"alg":"HS256","typ":"JWT"}
{"sub":"1234567890","name":"John Doe","iat":1516239022}
```

## Installation

```
go install github.com/syumai/cli/jwtdec@latest
```

## License

MIT

## Author

syumai
