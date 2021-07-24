package hash

import "strconv"

type Hash struct {
	Hex   string
	Int64 int64
}

type IHash interface {
	Set(value Hash)
	SetFromHex(value string)

	ToHex() string

	Equal(compare Hash) bool
	EqualFromHex(compare string) bool
	Empty() bool
}

func (h *Hash) Set(value Hash) {
	h.Hex = value.Hex
	h.Int64 = value.Int64
}

func (h *Hash) SetFromHex(value string) {
	h.Hex = value

	v, err := strconv.ParseInt(value, 16, 64)
	if err != nil {
		h.Int64 = -1
	}

	h.Int64 = v
}

func (h *Hash) ToHex() string {
	return h.Hex
}

func (h *Hash) Equal(compare Hash) bool {
	return h.Hex == compare.Hex
}

func (h *Hash) EqualFromHex(compare string) bool {
	return h.Hex == compare
}

func (h *Hash) Empty() bool {
	return len(h.Hex) == 0
}
