package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ayatk/memo/content"
	"github.com/urfave/cli"
)

var (
	tag         string
	description string
)

func NewCmd() cli.Command {
	return cli.Command{
		Name:   "new",
		Usage:  "Create new memo.",
		Flags:  flags,
		Action: action,
	}
}

var flags = []cli.Flag{
	cli.StringFlag{
		Name:        "tag, t",
		Usage:       "Add tags in memo.",
		Destination: &tag,
	},
	cli.StringFlag{
		Name:        "description, d",
		Usage:       "Add description in memo.",
		Destination: &description,
	},
}

func action(c *cli.Context) error {
	arg := c.Args().First()
	if arg == "" {
		return cli.NewExitError("Please input file name.", 2)
	}

	os.MkdirAll(time.Now().Format("2006/01/02"), 0777)
	fpath := fmt.Sprintf("%s/%s.md", time.Now().Format("2006/01/02"), arg)
	if _, err := os.Stat(fpath); err == nil {
		message := fmt.Sprintf("Already added %s.md.", arg)
		return cli.NewExitError(message, 1)
	}
	file, err := os.Create(fpath)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	defer file.Close()

	var data []string
	if c.NumFlags() > 0 {
		meta := content.MetaData{}

		if tag != "" {
			meta.Tag = tag
		}

		if description != "" {
			meta.Description = description
		}

		data = append(data, meta.GenerateHeader())
	}
	data = append(data, fmt.Sprintf("# %s\n", arg))

	file.WriteString(strings.Join(data, "\n"))
	return nil
}
