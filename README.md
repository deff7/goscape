# goscape

goscape is the simple tool for encoding and decoding web-specific entities such as HTML and URL. This utility follows UNIX philosophy and can be easily composed with another tools

## Getting started

### Prerequisites

You need to have Go distribution installed on your machine

### Installing

```
go get github.com/deff7/goscape
```

### Usage

```
goscape <command> <entity>
```

Commands:

- `encode`, `e` - encode (or escape) input text in terms of specified entity
- `decode`, `d` - decode (or unescape) input text

Supported entities:

- `html` - escape and unescape HTML offending characters
- `url` - escape and unescape URL query
- `base64` - encode and decode Base64 strings
- `json` - escape and unescape JSON stored in outher JSON's strings

### Examples

```
> printf "abc def/абв" | goscape encode url
abc+def%2F%D0%B0%D0%B1%D0%B2
```

```
> goscape decode url
> abc+def%2F%D0%B0%D0%B1%D0%B2
> ^D
abc def/абв
```
