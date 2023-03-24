package util

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"os"
)

func EstimateReadFile(filename string, defaultValue string) (string, error) {
	if filename == "" {
		return defaultValue, nil
	}

	body, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("read template file %s error: %w", filename, err)
	}

	return string(body), nil
}

func WriteFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
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
