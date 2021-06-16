package response

type Response struct {
	StatusCode int         `json:"status_code,omitempty"`
	Size       int         `json:"size,omitempty"`
	Status     bool        `json:"status,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Total      int64       `json:"total,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *Response) Success(data interface{}, totalData int64, size int) *Response {
	r.Status = true
	r.StatusCode = 200
	r.Data = data
	r.Message = "success"
	r.Total = totalData
	return r
}

func (r *Response) Failed(message string) *Response {
	r.StatusCode = 400
	r.Status = false
	r.Data = nil
	r.Message = message
	r.Total = 0
	return r
}
