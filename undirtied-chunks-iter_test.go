package torrent

import (
	"testing"

	typedRoaring "github.com/anacrolix/torrent/typed-roaring"
)

func BenchmarkIterUndirtiedRequestIndexesInPiece(b *testing.B) {
	var bitmap typedRoaring.Bitmap[RequestIndex]
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// This is the worst case, when Torrent.iterUndirtiedRequestIndexesInPiece can't find a
		// usable cached iterator. This should be the only allocation.
		it := bitmap.Iterator()
		iterBitmapUnsetInRange(&it, 69, 420, func(RequestIndex) {})
	}
}
