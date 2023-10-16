package filter

/**
*不用鉴权对的url
 */
var excludePath = []string{
	"/v1/api/register",
	"/v1/api/login",
	"/email/send",
	"/test",
}
