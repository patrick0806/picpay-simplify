package enums

type UserType int

const (
	Shopkeeper UserType = iota + 1
	Common
)

func (ut UserType) String() string {
	return [...]string{"shopkeeper", "common"}[ut-1]
}

func (ut UserType) EnumIndex() int {
	return int(ut)
}
