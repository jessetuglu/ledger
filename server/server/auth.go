package auth

import (
	"net/http"

	"golang.org/x/oauth2"
)

type GoogleUser struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func generateGoogleConfigs() *oauth2.Config {
	var auth = &oauth2.Config{
		RedirectURL:  "http://localhost:"+os.Getenv("API_MAIN_PORT")+"/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/accounts.reauth openid"},
		Endpoint:     google.Endpoint,
	}
	return auth
}

func (h *AuthHandler) AuthGoogleHandler(w http.ResponseWriter, r *http.Request) {
	authGoogleConfig := generateGoogleConfigs()
	oauthState := generateStateOauthCookie(w)
	u := authGoogleConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(2 * time.Hour)

	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (h *AuthHandler) AuthGoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")
	authGoogleConfig := generateGoogleConfigs()

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "mingl://home?error=temp", http.StatusBadRequest)
		return
	}

	data, err := getUserDataFromGoogle(r, authGoogleConfig, r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "mingl://home?error=temp", http.StatusBadRequest)
		return
	}

	user, err := helpers.FindOrCreateUser(r.Context(), h.QueryEngine, &data.Email, &data.FirstName, &data.LastName)
	accToken, refToken := auth.GenerateTokenPair(user, h.QueryEngine)
	log.Println("acc: " + accToken + " ref: " + refToken)
	http.Redirect(w, r, "mingl://home?access_token=" + accToken + "&refresh_token=" + refToken, http.StatusTemporaryRedirect)
}

func getUserDataFromGoogle(r *http.Request, config *oauth2.Config, code string) (*GoogleUser, error) {
	// Use code to get token and get user info from Google.
	token, err :=	 config.Exchange(r.Context(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token="+token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	user := &GoogleUser{}
	err = json.Unmarshal(contents, user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %s", err.Error())
	}
	return user, nil
}