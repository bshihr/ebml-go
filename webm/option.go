package webm

import (
	"errors"

	"github.com/at-wat/ebml-go"
)

var (
	// DefaultBlockInterceptor is the default BlockInterceptor used by BlockWriter.
	DefaultBlockInterceptor = NewMultiTrackBlockSorter(16, BlockSorterDropOutdated)
)

var (
	errInvalidTrackNumber = errors.New("invalid track number")
)

// BlockWriterOption configures a BlockWriterOptions.
type BlockWriterOption func(*BlockWriterOptions) error

// BlockWriterOptions stores options for BlockWriter.
type BlockWriterOptions struct {
	ebmlHeader          interface{}
	segmentInfo         interface{}
	seekHead            interface{}
	marshalOpts         []ebml.MarshalOption
	onError             func(error)
	onFatal             func(error)
	interceptor         BlockInterceptor
	mainTrackNumber     uint64
	maxKeyframeInterval int64
}

// WithEBMLHeader sets EBML header of WebM.
func WithEBMLHeader(h interface{}) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.ebmlHeader = h
		return nil
	}
}

// WithSegmentInfo sets Segment.Info of WebM.
func WithSegmentInfo(i interface{}) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.segmentInfo = i
		return nil
	}
}

// WithSeekHead sets Segment.SeekHead of WebM.
func WithSeekHead(s interface{}) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.seekHead = s
		return nil
	}
}

// WithMarshalOptions passes ebml.MarshalOption to ebml.Marshal.
func WithMarshalOptions(opts ...ebml.MarshalOption) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.marshalOpts = opts
		return nil
	}
}

// WithOnErrorHandler registers marshal error handler.
func WithOnErrorHandler(handler func(error)) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.onError = handler
		return nil
	}
}

// WithOnFatalHandler registers marshal error handler.
func WithOnFatalHandler(handler func(error)) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.onFatal = handler
		return nil
	}
}

// WithBlockInterceptor registers BlockInterceptor.
func WithBlockInterceptor(interceptor BlockInterceptor) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		o.interceptor = interceptor
		return nil
	}
}

// WithMaxKeyframeInterval sets maximum keyframe interval of the main (video) track.
// Using this option starts the cluster with a key frame if possible.
// interval must be given in the scale of timecode.
func WithMaxKeyframeInterval(mainTrackNumber uint64, interval int64) BlockWriterOption {
	return func(o *BlockWriterOptions) error {
		if mainTrackNumber <= 0 {
			return errInvalidTrackNumber
		}
		o.mainTrackNumber = mainTrackNumber
		o.maxKeyframeInterval = interval
		return nil
	}
}