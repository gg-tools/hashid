package commands

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strconv"
)

var Encode = &cli.Command{
	Name:    "encode",
	Aliases: []string{"e"},
	Usage:   "encode ID",
	UsageText: `hashid e [numbers...]

hashid e 1988
hashid e 1988 1990
`,
	ArgsUsage: "encode [ids...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none id present")
		}

		for _, v := range c.Args().Slice() {
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
