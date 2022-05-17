package cdwarehouse

type Charts interface {
	Notify(artist, title string, items int) error
}
