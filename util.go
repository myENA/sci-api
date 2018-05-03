package api

import "net/http"

func TryExtractAccessTokenCookie(resp *http.Response) *http.Cookie {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == AccessTokenKey {
			return cookie
		}
	}
	return nil
}
