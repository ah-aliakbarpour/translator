package export

import "translator/dictionary"

type Exporter interface {
	Export(data []dictionary.Result) error
}
