package keys

type Signer struct {
	PrivateKey PrivateKey `bson:"private_key" json:"private_key"`
}
