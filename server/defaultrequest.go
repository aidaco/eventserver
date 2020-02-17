package server

type DefaultRequest struct {
	r *http.Request
	body []byte
}

func (r *DefaultRequest) Body() []byte {
	if r.body != nil {
		return r.body
	}

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		body := nil

	r.body = body

	return body
}

func (r *DefaultRequest) URL() *url.URL {
	return r.r.URL
}

func (r *DefaultRequest) Path() string {
	return r.r.URL.Path
}

func (r *DefaultRequest) Header() http.Header {
	return r.r.Header
}

func (r *DefaultRequest) Addr() string {
	return r.r.RemoteAddr
}

func (r *DefaultRequest) Method() string {
	return r.r.Method
}

func (r *DefaultRequest) Request() *http.Request {
	return r.r
}

