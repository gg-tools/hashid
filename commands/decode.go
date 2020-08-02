package commands

import (
	"errors"
	"github.com/urfave/cli"
)

var Decode = cli.Command{
	Name:      "decode",
	ShortName: "d",
	Usage:     "decode hashed ID",
	UsageText: `
hashid decode 25yodre1
hashid decode 25yodre1 re7og60y
`,
	ArgsUsage: `decode [hashedIDs...]`,
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none hash id present")
		}

		for _, v := range c.Args() {
			id, err := decode(v)
			if err != nil {
				return err
			}

			print(id)
		}
		return nil
	},
}
