package main

import (
	"io/ioutil"

	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

const cExampleDir = "./example"
const cGstringerYaml = "gstringer.yaml"

func exampleOut(filename string) error {
	box := packr.NewBox(cExampleDir)

	b, err := box.Find(cGstringerYaml)
	if err != nil {
		return errors.Wrap(err, "failed to find "+cGstringerYaml)
	}
	err = ioutil.WriteFile(filename, b, 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create file "+filename)
	}
	return nil
}
