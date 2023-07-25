package books

const (
	CategoryNovel        string = "Novel"
	CategoryShortStory   string = "ShortStory"
	CategoryThreeHundred string = "ThreeHundred"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) Category() string {
	if b.Pages > 300 {
		return CategoryNovel
	} else if b.Pages < 300 {
		return CategoryShortStory
	}
	return CategoryThreeHundred
}
