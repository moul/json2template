package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/moul/advanced-ssh-config/pkg/templates"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "json2template"
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/json2template"
	app.Action = json2template
	app.Run(os.Args)
}

func json2template(c *cli.Context) error {
	var tmplSource, input []byte

	switch len(os.Args) {
	case 2:
		var err error

		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		tmplSource, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		break
	case 3:
		var err error
		input, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		tmplSource, err = ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		break
	default:
		return fmt.Errorf("Invalid usage")
	}

	tmpl, err := templates.New(string(tmplSource))
	if err != nil {
		panic(err)
	}

	var data interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}

	return nil
}
