

func getMessage(validationRule string, configMessage []string, defaultMessage string) string {
	var message string

	for i := 0; i < len(configMessage); i++ {
		validation := strings.Split(configMessage[i], ":")
		if validationRule == validation[0] {
			message = validation[1]
			break
		}
	}

	Log(message)
	if message == "" {
		message = defaultMessage
	}

	return message;
}
