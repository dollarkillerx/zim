package enums

type discoverKey string

func (d discoverKey) String() string {
	return string(d)
}

const (
	DiscoverManager discoverKey = "Manager"
)
