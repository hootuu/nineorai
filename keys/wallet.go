package keys

type Wallet struct {
	PrivateKey PrivateKey `bson:"private_key" json:"private_key"`
}
