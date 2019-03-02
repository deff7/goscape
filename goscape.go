package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

var usage = `goscape is simple tool to encoding and decoding web-specific values

Usage:
	goscape <command> <entity>

Commands:
	encode, e - Encode specified entity
	decode, d - Decode specified entity

Entities:
	html   - Escape/unescape HTML offending characters
	url    - Encode/decode URL string
	base64 - Encode/decode Base64
`

type entityType int

const (
	typeHTML entityType = iota
	typeURL
	typeBase64
)

func encode(src string, t entityType) (string, error) {
	switch t {
	case typeHTML:
		return html.EscapeString(src), nil
	case typeURL:
		log.Printf("%q", src)
		return url.QueryEscape(src), nil
	case typeBase64:
		return base64.StdEncoding.EncodeToString([]byte(src)), nil
	default:
		return "", fmt.Errorf("unknown type %d", t)
	}
	return "", nil
}

func decode(src string, t entityType) (string, error) {
	switch t {
	case typeHTML:
		return html.UnescapeString(src), nil
	case typeURL:
		return url.QueryUnescape(src)
	case typeBase64:
		out, err := base64.StdEncoding.DecodeString(src)
		return string(out), err
	default:
		return "", fmt.Errorf("unknown type %d", t)
	}
	return "", nil
}

type commandType int

const (
	commandEncode commandType = iota
	commandDecode
)

func getCommand(cmd string) (commandType, error) {
	var ok bool
	for _, c := range []string{
		"e", "encode",
		"d", "decode",
	} {
		if cmd == c {
			ok = true
			break
		}
	}
	if !ok {
		return -1, errors.New("unknown command " + cmd)
	}

	if cmd == "e" || cmd == "encode" {
		return commandEncode, nil
	}
	return commandDecode, nil
}

func getEntity(entity string) (entityType, error) {
	var typeMapping = map[string]entityType{
		"html":   typeHTML,
		"url":    typeURL,
		"base64": typeBase64,
	}

	t, ok := typeMapping[entity]
	if !ok {
		return -1, errors.New("unknown entity type " + entity)
	}
	return t, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(usage)
		os.Exit(0)
	}

	cmd, err := getCommand(args[0])
	checkError(err)

	t, err := getEntity(args[1])
	checkError(err)

	in, err := ioutil.ReadAll(os.Stdin)
	checkError(err)

	var out string
	if cmd == commandEncode {
		out, err = encode(string(in), t)
	}
	if cmd == commandDecode {
		out, err = decode(string(in), t)
	}

	checkError(err)

	fmt.Fprint(os.Stdout, out)
}
