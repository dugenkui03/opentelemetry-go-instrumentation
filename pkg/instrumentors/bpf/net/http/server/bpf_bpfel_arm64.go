// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package server

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfHttpRequestT struct {
	StartTime uint64
	EndTime   uint64
	Method    [7]int8
	Path      [100]int8
	Sc        bpfSpanContext
	Psc       bpfSpanContext
	_         [5]byte
}

type bpfSpanContext struct {
	TraceID [16]uint8
	SpanID  [8]uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeServerMuxServeHTTP         *ebpf.ProgramSpec `ebpf:"uprobe_ServerMux_ServeHTTP"`
	UprobeServerMuxServeHTTP_Returns *ebpf.ProgramSpec `ebpf:"uprobe_ServerMux_ServeHTTP_Returns"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	AllocMap                    *ebpf.MapSpec `ebpf:"alloc_map"`
	Events                      *ebpf.MapSpec `ebpf:"events"`
	GolangMapbucketStorageMap   *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	HttpEvents                  *ebpf.MapSpec `ebpf:"http_events"`
	ParentSpanContextStorageMap *ebpf.MapSpec `ebpf:"parent_span_context_storage_map"`
	TrackedSpans                *ebpf.MapSpec `ebpf:"tracked_spans"`
	TrackedSpansBySc            *ebpf.MapSpec `ebpf:"tracked_spans_by_sc"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	AllocMap                    *ebpf.Map `ebpf:"alloc_map"`
	Events                      *ebpf.Map `ebpf:"events"`
	GolangMapbucketStorageMap   *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	HttpEvents                  *ebpf.Map `ebpf:"http_events"`
	ParentSpanContextStorageMap *ebpf.Map `ebpf:"parent_span_context_storage_map"`
	TrackedSpans                *ebpf.Map `ebpf:"tracked_spans"`
	TrackedSpansBySc            *ebpf.Map `ebpf:"tracked_spans_by_sc"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.AllocMap,
		m.Events,
		m.GolangMapbucketStorageMap,
		m.HttpEvents,
		m.ParentSpanContextStorageMap,
		m.TrackedSpans,
		m.TrackedSpansBySc,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeServerMuxServeHTTP         *ebpf.Program `ebpf:"uprobe_ServerMux_ServeHTTP"`
	UprobeServerMuxServeHTTP_Returns *ebpf.Program `ebpf:"uprobe_ServerMux_ServeHTTP_Returns"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeServerMuxServeHTTP,
		p.UprobeServerMuxServeHTTP_Returns,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_arm64.o
var _BpfBytes []byte
