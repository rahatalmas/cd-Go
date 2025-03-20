package utils

type role struct {
	power      string
	key        string
	permission int
}

var admin = role{power: "admin", key: "_admin_key_", permission: 3}
var editor = role{power: "editor", key: "_editor_key_", permission: 2}
var user = role{power: "user", key: "_user_key_", permission: 1}

var ROLES = [3]role{admin, editor, user}
