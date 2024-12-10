package keys

type Address string

func (addr Address) PublicKey() PublicKey {
	return MustPublicKeyFromBase58(string(addr))
}
