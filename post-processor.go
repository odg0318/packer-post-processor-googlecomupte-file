package main

import (
	"io/ioutil"

	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

type Config struct {
	OutputPath string `mapstructure:"output"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	if err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...); err != nil {
		return err
	}

	return nil
}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	imageName := artifact.State("ImageName").(string)

	if err := ioutil.WriteFile(p.config.OutputPath, []byte(imageName), 0644); err != nil {
		return artifact, false, err
	}

	return artifact, true, nil
}
