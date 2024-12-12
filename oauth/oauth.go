package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "encoding/json"

    "npkg.dev/oauth.v2"
    "npkg.dev/oauth/google.v2"
)

var config *oauth.Config

func init() {
    config = &oauth.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        RedirectURL:  "http://localhost:8080/callback",
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }
}

func main() {
    // Handle main page
    http.HandleFunc("/", handleHome)
    // Handle login
    http.HandleFunc("/login", handleLogin)
    // Handle callback from OAuth provider
    http.HandleFunc("/callback", handleCallback)

    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    var html = `<html><body><a href="/login">Google Login</a></body></html>`
    fmt.Fprint(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    // Create OAuth state cookie
    oauthState := generateStateToken()
    cookie := &http.Cookie{
        Name:     "oauthstate",
        Value:    oauthState,
        Path:     "/",
        MaxAge:   3600,
        HttpOnly: true,
    }
    http.SetCookie(w, cookie)

    // Redirect to Google's consent page
    url := config.AuthCodeURL(oauthState)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
    // Read oauthState from Cookie
    oauthState, _ := r.Cookie("oauthstate")

    if r.FormValue("state") != oauthState.Value {
        fmt.Println("Invalid oauth state")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    token, err := config.Exchange(r.Context(), r.FormValue("code"))
    if err != nil {
        fmt.Printf("config.Exchange() failed with %s\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    // Get user info
    client := config.Client(r.Context(), token)
    userInfo, err := getUserInfo(client)
    if err != nil {
        fmt.Printf("Failed to get user info: %s\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    // Respond with user info
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userInfo)
}

func getUserInfo(client *http.Client) (map[string]interface{}, error) {
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var userInfo map[string]interface{}
    if err = json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        return nil, err
    }

    return userInfo, nil
}

// generateStateToken creates a random state token
func generateStateToken() string {
    // In a production environment, use a proper random token generator
    return "random-state-token"
}
