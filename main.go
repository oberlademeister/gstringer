package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func run() error {
	ctx := getContext()
	if ctx.Example != "" {
		err := exampleOut(ctx.Example)
		if err != nil {
			return errors.Wrap(err, "exampleOut failed")
		}
		ctx.Logger.Infof("wrote file %s", ctx.Example)
		return nil
	}
	if ctx.In != "" {
		t := &T{}
		b, err := ioutil.ReadFile(ctx.In)
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(b, t)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal")
		}
		ctx.Logger.Debugf(pretty.Sprint(t))

		out, err := os.Create(t.FileName)
		if err != nil {
			return errors.Wrap(err, "error creating "+t.FileName)
		}
		defer out.Close()
		t.RenderType(out)
		t.RenderString(out)
		t.RenderVerboseString(out)
		t.RenderExTypeMaker(out)
		t.RenderFromString(out)
		{
			cmd := exec.Command("gofmt", "-w", t.FileName)
			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrap(err, "failed to run gofmt "+string(out))
			}
			ctx.Logger.Debug("ran gofmt -w " + t.FileName)
		}
		return nil
	}
	return errors.New("end of run reached. this should never happen")
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
