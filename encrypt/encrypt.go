package encrypt

func EncryptComunication(keyInput, messageInput *string) string {

	if messageInput == nil || *messageInput == "" {
		return ""
	}

	key := "DCJ"
	if keyInput != nil && *keyInput != "" {
		key = *keyInput
	}

	messageEncrypt := make([]rune, 0, len(*messageInput)*2)
	vowels := "aeiouAEIOU"

	for _, char := range *messageInput {
		if isVowel(char, vowels) {
			messageEncrypt = append(messageEncrypt, []rune(key)...)
		}
		messageEncrypt = append(messageEncrypt, char)
	}

	return string(messageEncrypt)
}

func isVowel(c rune, vowels string) bool {
	for _, v := range vowels {
		if c == v {
			return true
		}
	}
	return false
}
