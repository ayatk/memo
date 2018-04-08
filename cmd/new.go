package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func NewCmd() cli.Command {
	return cli.Command{
		Name:   "new",
		Usage:  "Create new memo.",
		Action: action,
	}
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
	file.WriteString(fmt.Sprintf("# %s", arg))

	return nil
}
