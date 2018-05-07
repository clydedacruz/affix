package affix

type topic struct {
	ID           string
	Metadata     []byte
	Title        string
	sheetPreview []byte
	sheets       [][]byte
}
