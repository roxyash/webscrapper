package internal

type Scrapper interface {
}

type scrapperRepo struct {
}

func NewScrapper() Scrapper {
	return &scrapperRepo{}
}
