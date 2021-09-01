package account

type Setting struct {
	Proxy string
	BAuth *BAuth
}

type BAuth struct {
	Cookie    string
	Csrf      string
}
