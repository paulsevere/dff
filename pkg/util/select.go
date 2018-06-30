package util

import survey "gopkg.in/AlecAivazis/survey.v1"

type ListOption interface {
	Display() string
}

func intoOptionStrings(opts []ListOption) (options []string) {
	for _, item := range opts {
		options = append(options, item.Display())
	}
	return
}

func fromDisplayMatch(opts []ListOption, selection string) ListOption {
	for _, opt := range opts {
		if opt.Display() == selection {
			return opt
		}
	}
	return nil
}

func SelectFromList(message string, options []ListOption) ListOption {
	selection := ""
	prompt := &survey.Select{
		Message: message,
		Options: intoOptionStrings(options),
	}
	if err := survey.AskOne(prompt, &selection, nil); err == nil {
		return fromDisplayMatch(options, selection)
	}
	return nil
}
