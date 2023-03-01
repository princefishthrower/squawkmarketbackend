package googlecustomsearch

import (
	"context"
	"fmt"
	"log"
	"os"
	"squawkmarketbackend/models"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func CustomSearch(query string) (*models.Squawk, error) {
	ctx := context.Background()

	customsearchService, err := customsearch.NewService(ctx, option.WithAPIKey(os.Getenv("GOOGLE_CUSTOM_SEARCH_API_KEY")))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Set up a new Search call with the query
	cx := "018034913134727186954:zpihqmjxd6p"
	dateRestrict := "h1" // last hour
	sort := "date:r:0"   // sort by date in descending order (most recent first)
	resp, err := customsearchService.Cse.List().Q(query).Cx(cx).DateRestrict(dateRestrict).Sort(sort).Do()
	if err != nil {
		return nil, err
	}

	// For debugging:
	// fmt.Println("Search results for", query)
	// for _, result := range resp.Items {
	// 	fmt.Println(result.Title)
	// 	fmt.Println(result.Link)
	// 	fmt.Println(result.Snippet)
	// 	fmt.Println()
	// }

	// convert items to squawks
	squawks := make([]models.Squawk, len(resp.Items))
	for i, item := range resp.Items {
		squawks[i] = models.Squawk{
			Squawk: item.Title,
			Link:   item.Link,
		}
	}

	// return only first squawk
	if len(squawks) > 0 {
		squawk := &squawks[0]
		if squawk.Squawk == "" || squawk.Squawk == "Google news" {
			return nil, fmt.Errorf("squawk is empty")
		}
		return &squawks[0], nil
	}

	// return nil and new error saying no squawks found
	return nil, fmt.Errorf("no squawks found for query: %s", query)
}
