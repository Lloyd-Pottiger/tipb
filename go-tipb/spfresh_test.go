package tipb

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestSPFreshFilterExprColumnSourceWireValues(t *testing.T) {
	if got := int32(SPFreshFilterExprColumnSource_SPFreshFilterExprColumnSourceUnspecified); got != 0 {
		t.Fatalf("unspecified source wire value = %d, want 0", got)
	}
	if got := int32(SPFreshFilterExprColumnSource_Stored); got != 1 {
		t.Fatalf("stored source wire value = %d, want 1", got)
	}
	if got := int32(SPFreshFilterExprColumnSource_Handle); got != 2 {
		t.Fatalf("handle source wire value = %d, want 2", got)
	}
}

func TestSPFreshFilterExprColumnMissingSourceIsUnspecified(t *testing.T) {
	// Encodes only column_id = 42, leaving source absent from the wire payload.
	var column SPFreshFilterExprColumn
	if err := proto.Unmarshal([]byte{0x08, 0x2a}, &column); err != nil {
		t.Fatalf("unmarshal column: %v", err)
	}
	if got := column.GetSource(); got != SPFreshFilterExprColumnSource_SPFreshFilterExprColumnSourceUnspecified {
		t.Fatalf("missing source = %v, want unspecified", got)
	}
}
