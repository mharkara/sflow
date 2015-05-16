package sflow

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeGenericInterfaceCountersRecord(t *testing.T) {
	rec := GenericInterfaceCounters{
		Index:            9,
		Type:             6,
		Speed:            100000000,
		Direction:        1,
		Status:           3,
		InOctets:         79282473,
		InUcastPkts:      329128,
		InMulticastPkts:  0,
		InBroadcastPkts:  1493,
		InDiscards:       0,
		InErrors:         0,
		InUnknownProtos:  0,
		OutOctets:        764247430,
		OutUcastPkts:     9470970,
		OutMulticastPkts: 780342,
		OutBroadcastPkts: 877721,
		OutDiscards:      0,
		OutErrors:        0,
		PromiscuousMode:  1,
	}

	b := &bytes.Buffer{}

	err := rec.encode(b)
	if err != nil {
		t.Fatal(err)
	}

	// Skip the header section. It's 8 bytes.
	var headerBytes [8]byte

	_, err = b.Read(headerBytes[:])
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := decodeGenericInterfaceCountersRecord(b, uint32(b.Len()))
	if err != nil {
		t.Fatal(err)
	}

	if decoded != rec {
		t.Errorf("expected\n%+#v\n, got\n%+#v", rec, decoded)
	}
}

func TestEncodeDecodeHostCPUCountersRecord(t *testing.T) {
	rec := HostCpuCounters{
		Load1m:       0.1,
		Load5m:       0.2,
		Load15m:      0.3,
		ProcsRunning: 4,
		ProcsTotal:   5,
		NumCPU:       6,
		SpeedCPU:     7,
		Uptime:       8,

		CpuUser:         9,
		CpuNice:         10,
		CpuSys:          11,
		CpuIdle:         12,
		CpuWio:          13,
		CpuIntr:         14,
		CpuSoftIntr:     15,
		Interrupts:      16,
		ContextSwitches: 17,

		CpuSteal:     18,
		CpuGuest:     19,
		CpuGuestNice: 20,
	}

	b := &bytes.Buffer{}

	err := rec.encode(b)
	if err != nil {
		t.Fatal(err)
	}

	// Skip the header section. It's 8 bytes.
	var headerBytes [8]byte

	_, err = b.Read(headerBytes[:])
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := decodeHostCpuCountersRecord(b, uint32(b.Len()))
	if err != nil {
		t.Fatal(err)
	}

	if decoded != rec {
		t.Errorf("expected\n%+#v\n, got\n%+#v", rec, decoded)
	}
}
