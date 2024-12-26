package config

import (
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Options struct {
	Config    any
	Paths     []string
	Filenames []string
}

func LoadEnv(opt Options) error {
	for _, p := range opt.Paths {
		fp := filepath.Join(p, ".env")
		if _, fileErr := os.Stat(fp); fileErr == nil {
			_ = cleanenv.ReadConfig(fp, opt.Config)
		}
	}

	var err error
	for _, f := range opt.Filenames {
		for _, p := range opt.Paths {
			fp := filepath.Join(p, f)
			if _, fileErr := os.Stat(fp); fileErr != nil {
				return fileErr
			}
			if err = cleanenv.ReadConfig(fp, opt.Config); err != nil {
				return err
			}
		}
	}

	return err
}
