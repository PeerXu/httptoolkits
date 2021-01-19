package httptoolkits

import (
	"encoding/json"
	"net/http"

	"github.com/stretchr/objx"
)

type JSONRequest struct {
	*http.Request
}

func (r *JSONRequest) JSON() (objx.Map, error) {
	var msi map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&msi); err != nil {
		return nil, err
	}

	return objx.New(msi), nil
}

func WrapJSONRequest(r *http.Request) *JSONRequest {
	return &JSONRequest{r}
}
