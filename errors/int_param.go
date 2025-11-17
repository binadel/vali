package errors

import "github.com/mailru/easyjson/jwriter"

type IntParamError struct {
	BasicError
	ParamName  string
	ParamValue int64
}

func (e *IntParamError) Error() string {
	return e.Message
}

func (e *IntParamError) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(keyCode)
	w.String(string(e.Code))
	w.RawString(keyMessage)
	w.String(e.Message)
	w.RawString(",\"")
	w.RawString(e.ParamName)
	w.RawString("\":")
	w.Int64(e.ParamValue)
	w.RawByte('}')
}
