package gosnow

import (
	"encoding/base64"
	"encoding/json"
)

type BatchReponse struct {
	Id            string              `json:"id"`
	RawBody       string              `json:"body"`
	StatusCode    int                 `json:"status_code"`
	StatusText    string              `json:"status_text"`
	Headers       []map[string]string `json:"headers"`
	ExecutionTime int                 `json:"execution_time"`
	Body          ResponseEntry
}

type BatchedResponse struct {
	Id                string         `json:"batch_request_id"`
	ServicedRequests  []BatchReponse `json:"serviced_requests"`
	UnservicedRequets []BatchReponse `json:"unserviced_requests"`
	response          Response
}

// FROMJSON converts JSON to BatchResponse
func (B *BatchedResponse) FromJSON(JSON string) error {
	err := json.Unmarshal([]byte(JSON), B)

	for i, r := range B.ServicedRequests {
		rawBody := []byte(r.RawBody)
		b64 := make([]byte, base64.StdEncoding.DecodedLen(len(rawBody)))
		n, err := base64.StdEncoding.Decode(b64, rawBody)

		if err != nil {
			return err
		}

		if err := json.Unmarshal(b64[:n], &r.Body.Entry); err != nil {
			panic(err)
		}

		r.RawBody = ""
		B.ServicedRequests[i] = r
	}

	if err != nil {
		return err
	}
	return nil
}
