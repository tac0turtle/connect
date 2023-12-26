package coingecko

import (
	"strings"

	"github.com/skip-mev/slinky/x/oracle/types"
)

// getUniqueBaseAndQuoteDenoms returns a list of unique base and quote denoms
// from a list of currency pairs.
func getUniqueBaseAndQuoteDenoms(pairs []types.CurrencyPair) (string, string) {
	seenQuotes := make(map[string]struct{})
	quotes := make([]string, 0)

	seenBases := make(map[string]struct{})
	bases := make([]string, 0)

	for _, pair := range pairs {
		if _, ok := seenQuotes[pair.Quote]; !ok {
			seenQuotes[pair.Quote] = struct{}{}
			quotes = append(quotes, strings.ToLower(pair.Quote))
		}

		if _, ok := seenBases[pair.Base]; !ok {
			seenBases[pair.Base] = struct{}{}
			bases = append(bases, strings.ToLower(pair.Base))
		}
	}

	return strings.Join(bases, ","), strings.Join(quotes, ",")
}
