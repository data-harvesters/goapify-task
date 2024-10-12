# Goapify-Task
a goapify Task module allows you to create multi threaded tasks on actors

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

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		s, err := newScraper(input, a)
		if err != nil {
			fmt.Printf("failed to create scraper: %v\\n", err)
			continue
		}

		go func() {
			defer wg.Done()
			goapifyscraper.Run(s)
		}()
	}
	wg.Wait()
}

```