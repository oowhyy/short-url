package service

type ShortUrlService interface {
	Shorten(someString string) (string, error)
	Reverse(shortLink string) (string,error)
}
