package api

import (
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

const RequestIdKey = "request-id"

var blankSessionCookie = &http.Cookie{
	Name: sessionContextKey,
	Value: "",
	MaxAge: -1, //TODO: investigate
	Path: "/",
}

func (s *Server) googleLoginHandler(ctx *gin.Context) {
	session, err := s.sessions.New(ctx.Request, sessionContextKey)
	if err != nil {
		s.logger.Errorw("Couldn't get session", err)

		ctx.SetCookie(sessionContextKey, "", -1, "/", s.serverBaseUrl, true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, s.serverBaseUrl + "/api/auth/login")
		return
	}

	state := uniuri.NewLen(64)
	session.Values["state"] = state
	s.sessions.Save(ctx.Request, ctx.Writer, session)

	url := s.oauthConfig.AuthCodeURL(state)

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (s *Server) googleCallBackHandler(ctx *gin.Context) {
	auth_code := ctx.Request.FormValue("code")
	session, err := s.sessions.Get(ctx.Request, sessionContextKey)
	if (err != nil) {
		s.logger.Errorw("Failed to get session", err)
		ctx.SetCookie(sessionContextKey, "", -1, "/", s.serverBaseUrl, true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, s.serverBaseUrl + "/api/auth/login")
		return
	}

	auth_token, err := s.oauthConfig.Exchange(ctx, auth_code)

	if (err != nil){
		s.logger.Errorw("Couldn't get auth token", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	service, err := goauth2.NewService(ctx, option.WithTokenSource(s.oauthConfig.TokenSource(ctx, auth_token)))	

	if (err != nil){
		s.logger.Errorw("Couldn't make new auth service", err)
		return
	}

	info, err := service.Userinfo.V2.Me.Get().Do()

	if (err != nil){
		s.logger.Errorw("Couldn't get info from new auth service", err)
		return
	}
	
	session.Values["email"] = info.Email
	session.Values["first_name"] = info.GivenName

	s.sessions.Save(ctx.Request, ctx.Writer, session)

	s.logger.Infof("Successfully authenticated user: ", info.Email)
	ctx.Redirect(http.StatusOK, s.clientBaseUrl)
}