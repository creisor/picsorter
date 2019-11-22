package imagefile

type MetaGetter interface {
	Get(string, string) (string, error)
}

type DateTimeGetter struct{}

func NewDateTimeGetter() *DateTimeGetter {
	return &DateTimeGetter{}
}

func (g *DateTimeGetter) Get(field string, filename string) (string, error) {
	return "", nil
}
