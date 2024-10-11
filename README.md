# Goapify-Scraper
a goapify Scraper module allows you to create scraper that do specific things

## Usage
```go

type scraper struct {
	goapifyscraper.Base
}

func newScraper(actor *goapify.Actor) (*scraper, error) {
	return &scraper{
		Base:   *goapifyscraper.New(actor),
	}, nil
}

const (
	ScrapeProducts goapifyscraper.State = iota
)

func (p *scraper) Next(state goapifyscraper.State) (goapifyscraper.State, error) {
	switch state {
	case goapifyscraper.Initialize:
		return ScrapeProducts, nil
	case ScrapeProducts:
		return ScrapeProducts, nil
	}

	p.Stop() // should never get here, theoretically, unless wanted
	return 0, nil
}

func main() {
    // INITIALIZE ACTOR
	a := goapify.NewActor()

	s, err := newScraper(i, a)
	if err != nil {
		fmt.Printf("failed to create scraper: %v\\n", err)
		panic(err)
	}

	goapifyscraper.Run(s)
}

```