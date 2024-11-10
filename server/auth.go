package server

import (
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/42LoCo42/einauth/config"
	"github.com/42LoCo42/einauth/utils"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {
	headers := c.Request().Header
	target := url.URL{
		Scheme: headers.Get("x-forwarded-proto"),
		Host:   headers.Get("x-forwarded-host"),
		Path:   headers.Get("x-forwarded-uri"),
	}

	accept := func() error {
		return c.NoContent(http.StatusOK)
	}

	reject := func() error {
		c.SetCookie(utils.MakeCookie("redir", target.String()))
		return c.Redirect(http.StatusSeeOther, config.CONFIG.URL)
	}

	match := func(pat, val string) bool {
		ok, err := regexp.MatchString("^"+pat+"$", val)
		return ok && err == nil
	}

	rules := utils.Filter(config.CONFIG.Rules, func(rule config.Rule) bool {
		return utils.And([]bool{
			// domain must always match exactly
			rule.Domain == target.Host,

			// if there are paths, at least one must match
			utils.EmptyOrAny(rule.Paths, func(path string) bool {
				return match(path, target.Path)
			}),
		})
	})

	log.Print(rules)

	// rules with no policy just accept w/o loading the cookie
	if utils.Any(rules, func(rule config.Rule) bool {
		return rule.Policy == nil
	}) {
		return accept()
	}

	cookie, err := c.Cookie("einauth-token")
	if err != nil {
		return reject()
	}

	user, err := utils.VerifyCookie[CookieUser](cookie)
	if err != nil {
		return reject()
	}

	rules = utils.Filter(rules, func(rule config.Rule) bool {
		// at least one subject must match
		return utils.Any(rule.Policy.Subjects, func(subject string) bool {
			// a subject can match via...
			return utils.Or(
				append(
					// ...any group of the current user
					utils.Map(user.Groups, func(group string) bool {
						return match(subject, "group:"+group)
					}),

					// ...or the user itself
					match(subject, "user:"+user.Name),
				),
			)
		})
	})

	if len(rules) == 0 {
		return reject()
	}

	// TODO validate two factor auth
	log.Print(rules)
	return accept()
}
