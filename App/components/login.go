package components

import (
	"stack/frameGo"
)

type Login struct {
}

func (l *Login) Render() string {
	component := `
<form method="POST">
	<div class="input-group">
		<label for="username">Enter username or email: </label>
		<input type="text" name="login" id="loginUsername">
	</div>

	<div class="input-group">
		<label for="password">Enter Password: </label>
		<input type="password" name="password" id="loginPassword">
	</div>

	<button name="submit" id="submitLogin">Submit</button>
</form>
`
	return frameGo.ParseComponent(component, l)
}

func NewLogin() Login {
	return Login{}
}
