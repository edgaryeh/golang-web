package skyhdd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"encoding/base64"
)

func putCookie(res http.ResponseWriter, req *http.Request, fname string) ([]string, error) {
	var xs []string
	cookie, _ := req.Cookie("file-names")
	if cookie != nil {
		bs, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler base64.URLEncoding.DecodeString: %s", err)
		}
		err = json.Unmarshal(bs, &xs)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler json.Unmarshal: %s", err)
		}
	}

	xs = append(xs, fname)
	bs, err := json.Marshal(xs)
	if err != nil {
		return xs, fmt.Errorf("ERROR putCookie json.Marshal: ", err)
	}
	b64 := base64.URLEncoding.EncodeToString(bs)

	http.SetCookie(res, &http.Cookie{
		Name:  "file-names",
		Value: b64,
	})
	return xs, nil
}