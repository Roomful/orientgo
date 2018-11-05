package obinary

import (
	"github.com/Roomful/orientgo/obinary/rw"
)

func ReadErrorResponse(r *rw.Reader) (serverException error) {
	return readErrorResponse(r, CurrentProtoVersion)
}
