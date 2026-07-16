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

func TestSPFreshEvalContextPreservesSQLMode(t *testing.T) {
	const sqlMode = uint64(1<<6 | 1<<21 | 1<<26)

	encoded, err := proto.Marshal(&SPFreshEvalContext{SqlMode: sqlMode})
	if err != nil {
		t.Fatalf("proto.Marshal(SPFreshEvalContext) failed: %v", err)
	}
	var decoded SPFreshEvalContext
	if err := proto.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("proto.Unmarshal(SPFreshEvalContext) failed: %v", err)
	}
	if got := decoded.GetSqlMode(); got != sqlMode {
		t.Fatalf("SPFreshEvalContext.GetSqlMode() = %#x, want %#x", got, sqlMode)
	}
}

func TestSPFreshSearchResponseResultIsExclusive(t *testing.T) {
	resp := &SPFreshSearchResponse{
		Result: &SPFreshSearchResponse_Success{Success: &SPFreshSearchResult{
			WarningCount: 1 << 63,
		}},
	}
	if resp.GetSuccess() == nil || resp.GetError() != nil {
		t.Fatalf("success response result = %T, want success only", resp.GetResult())
	}
	if got := resp.GetSuccess().GetWarningCount(); got != uint64(1)<<63 {
		t.Fatalf("warning_count = %d, want %d", got, uint64(1)<<63)
	}

	resp.Result = &SPFreshSearchResponse_Error{Error: &Error{Code: int32(SPFreshErrorCode_SPFreshIndexCorruption)}}
	if resp.GetSuccess() != nil || resp.GetError() == nil {
		t.Fatalf("error response result = %T, want error only", resp.GetResult())
	}
}

func TestSPFreshIndexCorruptionCode(t *testing.T) {
	if got := int32(SPFreshErrorCode_SPFreshIndexCorruption); got != 9015 {
		t.Fatalf("SPFreshIndexCorruption code = %d, want 9015", got)
	}
}
