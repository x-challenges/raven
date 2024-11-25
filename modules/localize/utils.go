package localize

func MustParseLanguageCodes(codes ...string) []LanguageCode {
	var result = make([]LanguageCode, 0, len(codes))

	result = append(result, codes...)

	return result
}
