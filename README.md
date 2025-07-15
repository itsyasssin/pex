# PEX
`pex` is a simple parameter extractor written in Go that automatically extracts potential parameters from various sources:

- **JavaScript variables** (var, let, const declarations)
- **JSON object keys** (from API responses and configuration)
- **HTML form inputs** (name and id attributes)
- **URL query parameters** (from GET requests and links)
- **Function parameters** (JavaScript function calls)

### Installation
```console
go install -v github.com/itsyasssin/pex@latest
```

### Usage
```console
$ pex -h
USAGE: pex [-strings]
```

Simple usage:
```console
$ curl -s https://domain.tld/file.js | pex # or `pex -strings` to extract all javascript strings
```
<img width="913" height="1054" alt="Screenshot From 2025-07-16 00-30-04" src="https://github.com/user-attachments/assets/b8ae4263-7d0f-47d0-bc60-10e408d38906" />


