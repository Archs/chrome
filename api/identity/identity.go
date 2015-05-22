package identity

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	identity = chrome.Get("identity")
)

/*
* Types
 */

type AccountInfo struct {
	*js.Object
	Id string `js:"id"`
}

/*
* Methods
 */

// GetAccounts retrieves a list of AccountInfo js.Ms describing the accounts present on the profile.
func GetAccounts(callback func(accounts []AccountInfo)) {
	identity.Call("getAccounts", callback)
}

// GetAuthToken gets an OAuth2 access token using the client ID and scopes specified in the oauth2 section of manifest.json.
/* The Identity API caches access tokens in memory, so it's ok to call getAuthToken non-interactively any time a token is required. The token cache automatically handles expiration.
For a good user experience it is important interactive token requests are initiated by UI in your app explaining what the authorization is for. Failing to do this will cause your users to get authorization requests, or Chrome sign in screens if they are not signed in, with with no context. In particular, do not use getAuthToken interactively when your app is first launched. */
func GetAuthToken(details js.M, callback func(token string)) {
	identity.Call("getAuthToken", details, callback)
}

// GetProfileUserInfo retrieves email address and obfuscated gaia id of the user signed into a profile.
// This API is different from identity.getAccounts in two ways. The information returned is available offline,
// and it only applies to the primary account for the profile.
func GetProfileUserInfo(callback func(userInfo js.M)) {
	identity.Call("getProfileUserInfo", callback)
}

// RemoveCacheAuthToken removes an OAuth2 access token from the Identity API's token cache.
// If an access token is discovered to be invalid, it should be passed to removeCachedAuthToken
// to remove it from the cache. The app may then retrieve a fresh token with getAuthToken.
func RemoveCacheAuthToken(details js.M, callback func()) {
	identity.Call("removeCacheAuthToken", details, callback)
}

// LaunchWebAuthFrom starts an auth flow at the specified URL.
/* This method enables auth flows with non-Google identity providers by launching a web view and navigating it to the first URL in the provider's auth flow. When the provider redirects to a URL matching the pattern https://<app-id>.chromiumapp.org/*, the window will close, and the final redirect URL will be passed to the callback function.
For a good user experience it is important interactive auth flows are initiated by UI in your app explaining what the authorization is for. Failing to do this will cause your users to get authorization requests with no context. In particular, do not launch an interactive auth flow when your app is first launched. */
func LaunchWebAuthFrom(details js.M, callback func(response string)) {
	identity.Call("launchWebAuthFrom", details, callback)
}

// GetRedirectURL generates a redirect URL to be used in |launchWebAuthFlow|.
// The generated URLs match the pattern https://<app-id>.chromiumapp.org/*.
func GetRedirectURL(path string) {
	identity.Call("getRedirectURL", path)
}

/*
* Events
 */

// OnSignInChanged fired when signin state changes for an account on the user's profile.
func OnSignInChanged(callback func(account AccountInfo, signedIn bool)) {
	identity.Get("onSignInChanged").Call("addListener", callback)
}
