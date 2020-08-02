package commands

import (
	"errors"
	"github.com/urfave/cli"
	"strconv"
)

var Encode = cli.Command{
	Name:      "encode",
	ShortName: "e",
	Usage:     "encode ID",
	UsageText: `
hashid encode 1988
hashid encode 1988 1990
`,
	ArgsUsage: "encode [ids...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none id present")
		}

		for _, v := range c.Args() {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}

			hashed, err := encode(id)
			print(hashed)
		}

		return nil
	},
}
