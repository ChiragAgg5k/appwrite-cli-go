package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/client"
)

// `login --switch` was reported as "unknown flag". It was not alone: the port
// implemented only the cloud browser flow, so six of the TypeScript's seven
// options did not exist (generic.ts:121).
func TestLoginOffersEveryOption(t *testing.T) {
	command := newLoginCommand()

	for _, name := range []string{
		"endpoint", "email", "password", "mfa", "code", "switch", "new",
	} {
		if command.Flags().Lookup(name) == nil {
			t.Errorf("--%s is missing", name)
		}
	}
}

// Two ways to spell "not the current account" that mean different things:
// --switch moves to one already signed in, --new signs in again. Asking for
// both is a contradiction, not a preference.
func TestLoginRejectsSwitchWithNew(t *testing.T) {
	err := runLogin(newLoginCommand(), loginOptions{Switch: true, New: true})

	if err == nil || !strings.Contains(err.Error(), "not both") {
		t.Errorf("err = %v, want a complaint about using both", err)
	}
}

// Cloud sign-in happens in a browser, so an email and password given for it
// would be silently ignored. Saying so beats appearing to accept them.
func TestLoginRejectsPasswordOptionsAgainstCloud(t *testing.T) {
	err := runLogin(newLoginCommand(), loginOptions{
		Endpoint: "https://cloud.appwrite.io/v1",
		Email:    "someone@example.com",
		Password: "hunter2",
	})

	if err == nil || !strings.Contains(err.Error(), "browser") {
		t.Errorf("err = %v, want it to explain that Cloud signs in via the browser", err)
	}
}

// The endpoint is checked BEFORE anything is prompted for, so a typo fails
// immediately rather than after an email and a password have been typed.
func TestVerifyEndpointRejectsAServerThatIsNotAppwrite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	if err := verifyEndpoint(server.URL); err == nil {
		t.Error("a server with no version was accepted as an Appwrite instance")
	}
}

func TestVerifyEndpointAcceptsAnAppwriteServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health/version" {
			t.Errorf("asked for %s, want /health/version", r.URL.Path)
		}
		w.Write([]byte(`{"version":"1.8.0"}`))
	}))
	defer server.Close()

	if err := verifyEndpoint(server.URL); err != nil {
		t.Errorf("a healthy server was rejected: %v", err)
	}
}

func TestVerifyEndpointRejectsAMalformedURL(t *testing.T) {
	for _, endpoint := range []string{"not a url", "ftp://example.com/v1", ""} {
		if err := verifyEndpoint(endpoint); err == nil {
			t.Errorf("%q was accepted as an endpoint", endpoint)
		}
	}
}

// The email-and-password flow has no token: the session cookie IS the
// credential, and it arrives only on the Set-Cookie header of the sign-in
// response. Dropping it leaves a session that cannot authenticate anything.
func TestClientCapturesTheConsoleSessionCookie(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The session cookie comes FIRST and an unrelated one after it, so a
		// capture that does not filter ends up holding the wrong one. With the
		// order reversed, last-write-wins would pass either way.
		w.Header().Add("Set-Cookie", "a_session_console=secret; Path=/; HttpOnly")
		w.Header().Add("Set-Cookie", "unrelated=value; Path=/")
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	api := client.New(server.URL, "test")
	if err := api.Call("POST", "/account/sessions/email", nil, nil); err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(api.SessionCookie, "a_session_console=") {
		t.Errorf("SessionCookie = %q, want the console session cookie", api.SessionCookie)
	}
	if strings.Contains(api.SessionCookie, "unrelated") {
		t.Error("an unrelated cookie was stored as the session")
	}
}

// A response that sets no cookie must not clear one already held, or a later
// request in the same flow would go out unauthenticated.
func TestClientKeepsTheCookieWhenAResponseSetsNone(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	api := client.New(server.URL, "test").SetCookie("a_session_console=kept")
	if err := api.Call("GET", "/account", nil, nil); err != nil {
		t.Fatal(err)
	}

	if api.SessionCookie != "" {
		t.Errorf("SessionCookie = %q, want it untouched by a response that set none",
			api.SessionCookie)
	}
}
