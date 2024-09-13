package client

type TwistlockConfig struct {
	ENDPOINT   string
	ACCOUNT    string
	API_KEY    string
	API_SECRET string
}

type Spec struct {
	// plugin spec goes here
	TWISTLOCK []TwistlockConfig
}
