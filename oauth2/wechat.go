package oauth2

import (
	"net/http"

	"log"

	"errors"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/oauth"
)

type Oauth2 struct {
	wx       *wechat.Wechat
	request  *http.Request
	response http.ResponseWriter
}

var wx *wechat.Wechat

func init() {
	WechatInstance(nil)
}

func NewOauth(r *http.Request, w http.ResponseWriter) *Oauth2 {
	return &Oauth2{
		wx:       WechatInstance(nil),
		request:  r,
		response: w,
	}
}

func WechatInstance(config *wechat.Config) *wechat.Wechat {
	if wx == nil {
		config := config
		if config == nil {
			config = &wechat.Config{
				AppID:          "xxxx",
				AppSecret:      "xxxx",
				Token:          "xxxx",
				EncodingAESKey: "xxxx",
				//Cache:          memCache
			}
		}
		wx = wechat.NewWechat(config)
	}
	return wx

}

func (o *Oauth2) GetOpenid() *oauth.ResAccessToken {
	code := o.request.Form.Get("code")
	oauth := o.wx.GetOauth(o.request, o.response)
	if code != "" {
		err := oauth.Redirect("url", "snsapi_userinfo", "state")
		log.Println("GetOpenid:", err)
		return nil
	}
	token, err := oauth.GetUserAccessToken(code)
	if err != nil {
		return nil
	}
	return &token
}

func (o *Oauth2) GetOauth() *oauth.Oauth {
	return o.wx.GetOauth(o.request, o.response)
}

func (o *Oauth2) GetUserInfo(token *oauth.ResAccessToken) (result oauth.UserInfo, err error) {
	if token == nil {
		return oauth.UserInfo{}, errors.New("GetUserInfo() token nil")
	}
	oauth := o.wx.GetOauth(o.request, o.response)
	return oauth.GetUserInfo(token.AccessToken, token.OpenID)
}
