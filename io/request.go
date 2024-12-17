package io

import (
	"encoding/json"
	"github.com/hootuu/gelato/crtpto/rand"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/idx"
	"github.com/hootuu/gelato/io/serializer"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
	"time"
)

type RequestData interface {
	Validate() *errors.Error
}

type Request[T RequestData] struct {
	ID         string                                `json:"id" bson:"id"`
	Data       *T                                    `json:"data" bson:"data"`
	Accounts   AccountSet                            `json:"accounts" bson:"accounts"`
	Timestamp  int64                                 `json:"timestamp" bson:"timestamp"`
	Nonce      int64                                 `json:"nonce" bson:"nonce"`
	Signatures map[keys.Address]keys.SignatureBase58 `json:"signatures" bson:"signatures"`
}

func NewRequest[T RequestData](data *T) *Request[T] {
	randInt64, _ := rand.Int64()
	return &Request[T]{
		ID:         idx.New(),
		Data:       data,
		Accounts:   NewAccountSet(),
		Timestamp:  time.Now().UnixMilli(),
		Nonce:      randInt64,
		Signatures: make(map[keys.Address]keys.SignatureBase58),
	}
}

func (r *Request[T]) AddPayer(addr keys.Address) *Request[T] {
	if r.Accounts == nil {
		r.Accounts = NewAccountSet()
	}
	_ = r.Accounts.AddAccount(Account{
		Address: addr,
		Payer:   true,
		Signer:  false,
	})
	return r
}

func (r *Request[T]) AddSigner(addr keys.Address) *Request[T] {
	if r.Accounts == nil {
		r.Accounts = NewAccountSet()
	}
	_ = r.Accounts.AddAccount(Account{
		Address: addr,
		Payer:   false,
		Signer:  true,
	})
	return r
}

func (r *Request[T]) HasPayer(addr keys.Address) bool {
	if r.Accounts != nil {
		return r.Accounts.HasPayer(addr)
	}
	return false
}
func (r *Request[T]) HasSigner(addr keys.Address) bool {
	if r.Accounts != nil {
		return r.Accounts.HasSigner(addr)
	}
	return false
}

func (r *Request[T]) HasSignerAtLeast() bool {
	if r.Accounts != nil {
		return r.Accounts.ExistsSignerAtLeast()
	}
	return false
}

func (r *Request[T]) Sign(keyArr ...*keys.Key) *errors.Error {
	priKeys := make(map[keys.Address]keys.PrivateKey)
	for _, item := range keyArr {
		priKeys[item.Public.Address()] = item.Private
	}
	if err := r.Validate(); err != nil {
		return err
	}
	serializeData, err := r.Serialize()
	if err != nil {
		return err
	}
	for _, account := range r.Accounts {
		if !(account.Payer || account.Signer) {
			continue
		}
		curPriKey, ok := priKeys[account.Address]
		if !ok {
			return errors.E("require_private_key", "require private key for %s", account.Address)
		}
		cusSignature, err := curPriKey.Sign(serializeData)
		if err != nil {
			return err
		}
		r.Signatures[account.Address] = cusSignature.HexEncode()
	}
	return nil
}

func (r *Request[T]) Verify() *errors.Error {
	if err := r.Validate(); err != nil {
		return err
	}
	if len(r.Signatures) == 0 {
		return errors.E("require_signatures", "no any signature")
	}
	serializeData, err := r.Serialize()
	if err != nil {
		return err
	}
	for _, account := range r.Accounts {
		if !(account.Payer || account.Signer) {
			continue
		}
		curSignatureHex, ok := r.Signatures[account.Address]
		if !ok {
			return errors.E("require_signature", "require signature for %s", account.Address)
		}
		curSignature, err := curSignatureHex.Decode()
		if err != nil {
			return err
		}
		curPubKey, err := account.Address.PublicKey()
		if err != nil {
			return err
		}
		bVerify := curPubKey.Verify(serializeData, *curSignature)
		if !bVerify {
			return errors.E("invalid_signature", "invalid signature for %s", account.Address)
		}
	}
	return nil
}

func (r *Request[T]) Serialize() ([]byte, *errors.Error) {
	serializeStr, err := serializer.Serialize(map[string]interface{}{
		"id":        r.ID,
		"data":      r.Data,
		"accounts":  r.Accounts,
		"timestamp": r.Timestamp,
		"nonce":     r.Nonce,
	})
	if err != nil {
		return nil, err
	}
	return []byte(serializeStr), nil
}

func (r *Request[T]) Marshal() ([]byte, *errors.Error) {
	byteArr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.System("json_marshal_failed", err)
	}
	return byteArr, nil
}

func UnmarshalRequest[T RequestData](data []byte) (*Request[T], *errors.Error) {
	var req Request[T]
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, errors.E("json_unmarshal_failed", "unmarshal request failed: %w", err)
	}
	return &req, nil
}

func (r *Request[T]) Validate() *errors.Error {
	if r.Data == nil {
		return errors.Verify("require data")
	}
	if err := (*r.Data).Validate(); err != nil {
		return err
	}
	if len(r.Accounts) == 0 {
		return errors.Verify("require accounts")
	}
	payerCount := 0
	for _, account := range r.Accounts {
		if account.Payer {
			payerCount++
		}
	}
	if payerCount != 1 {
		return errors.Verify("there can only be one payer")
	}
	return nil
}

func (r *Request[T]) MustGetSignature() (domains.Signature, *errors.Error) {
	if r.Signatures == nil {
		return "", errors.Verify("require signatures")
	}
	if r.Accounts == nil {
		return "", errors.Verify("require accounts")
	}
	payer, err := r.Accounts.MustGetPayer()
	if err != nil {
		return "", err
	}
	for address, hex := range r.Signatures {
		if address == payer.Address {
			return domains.Signature(hex), nil
		}
	}
	return "", errors.Verify("require payer signature")
}
