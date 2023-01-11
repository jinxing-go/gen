package util

import (
	"bufio"
	"bytes"
	"go/format"
	"html/template"
	"os"
)

func WriteFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
	return nil
}

func Parse(temp string, data interface{}) (string, error) {
	tmpl, err := template.New("model").Parse(temp)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	var body []byte
	if body, err = format.Source(buf.Bytes()); err != nil {
		return "", err
	}

	return string(body), nil
}
