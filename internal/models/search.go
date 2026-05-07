package models

import (
	"strings"
	"unicode"
)

func tokenizeSearchQuery(query string) []string {
	normalized := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, query)

	rawTokens := strings.Fields(normalized)
	if len(rawTokens) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(rawTokens))
	tokens := make([]string, 0, len(rawTokens))
	for _, token := range rawTokens {
		if _, exists := seen[token]; exists {
			continue
		}
		seen[token] = struct{}{}
		tokens = append(tokens, token)
	}

	return tokens
}

func buildFlexibleSearchClause(columns []string, query string) (string, []interface{}) {
	tokens := tokenizeSearchQuery(query)
	if len(tokens) == 0 || len(columns) == 0 {
		return "", nil
	}

	var clauseParts []string
	args := make([]interface{}, 0, len(tokens)*len(columns))

	for _, token := range tokens {
		var tokenParts []string
		for _, column := range columns {
			tokenParts = append(tokenParts, column+" LIKE ?")
			args = append(args, "%"+token+"%")
		}
		clauseParts = append(clauseParts, "("+strings.Join(tokenParts, " OR ")+")")
	}

	return " AND " + strings.Join(clauseParts, " AND "), args
}
