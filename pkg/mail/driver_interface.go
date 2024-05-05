package mail

type Driver interface {
	//核查验证码
	Send(email Email, config map[string]string) bool
}
