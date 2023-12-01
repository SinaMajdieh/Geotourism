package commands

type Flag struct {
	Name string
	Long string
}
type Command struct {
	Flags []Flag
	Run   func(args []string)
	Help  func()string
}

func (c *Command) AddFlag(flag *Flag) {
	c.Flags = append(c.Flags, *flag)
}
