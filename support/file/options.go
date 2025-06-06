package file

import "os"

// Option represents an option for FilePutContents
type Option func(*option)

type option struct {
	mode   os.FileMode
	append bool
}

// WithAppend sets the append mode for FilePutContents
func WithAppend(append bool) Option {
	return func(opts *option) {
		opts.append = append
	}
}

// WithMode sets the file mode for FilePutContents
func WithMode(mode os.FileMode) Option {
	return func(opts *option) {
		opts.mode = mode
	}
}
