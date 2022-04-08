package helpers

func GetArg(position int, args []string, default_value string) string {
	if len(args) < position+1 {
		return default_value
	}

	return args[position]
}
