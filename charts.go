package cdwarehouse

type Charts interface {
	Notify(artist, title string, items int) error
	IsTop100(artist, title string) bool
	GetLowestPrice(artist string, title string) float64
}
