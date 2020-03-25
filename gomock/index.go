//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package gomock

type Foo interface {
	Bar(x int) int
}
