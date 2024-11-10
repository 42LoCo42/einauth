package server

import (
	"github.com/42LoCo42/einauth/utils"
	"github.com/labstack/echo/v4"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func Page(c echo.Context, title string, nodes ...Node) error {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head:     []Node{},
		Body:     nodes,
	}).Render(c.Response())
}

func UI(c echo.Context) error {
	cookie, err := c.Cookie("einauth-token")
	if err != nil {
		return LoginUI(c)
	}

	user, err := utils.VerifyCookie[CookieUser](cookie)
	if err != nil {
		return LoginUI(c)
	}

	return Page(c, "einauth",
		H1(Text("Welcome to einauth")),
		H3(Textf("You're logged in as %s", user.Name)),

		Form(Method("POST"), Action("/logout"),
			Input(Type("submit"), Value("Logout")),
		),
	)
}

func LoginUI(c echo.Context) error {
	return Page(c, "einauth - login",
		H1(Text("Log in to einauth")),

		Form(Method("POST"), Action("/login"),
			Table(
				Tr(
					Td(Text("Username: ")),
					Td(Input(Type("text"), Name("username"))),
				),
				Tr(
					Td(Text("Password: ")),
					Td(Input(Type("password"), Name("password"))),
				),
			),
			Input(Type("submit"), Value("Login")),
		),
	)
}
