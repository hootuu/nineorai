package keys

import (
	"encoding/hex"
	"fmt"
	"github.com/hootuu/gelato/errors"
)

type Signature [64]byte

type SignatureHex string

func (s *Signature) HexEncode() SignatureHex {
	return SignatureHex(hex.EncodeToString(s[:]))
}

func (s *SignatureHex) Decode() (*Signature, *errors.Error) {
	str := string(*s)
	if len(str) != 128 {
		return nil, errors.Verify(fmt.Sprintf("invalid signature length: got %d want 128", len(str)))
	}

	decoded, err := hex.DecodeString(str)
	if err != nil {
		return nil, errors.E("invalid_signature", err.Error())
	}
	signature := new(Signature)
	copy(signature[:], decoded)
	return signature, nil
}

//
//func (s *Signature) MarshalJSON() ([]byte, error) {
//	return json.Marshal(hex.EncodeToString(s[:]))
//}
//
//func (s *Signature) UnmarshalJSON(data []byte) error {
//	var hexStr string
//	if err := json.Unmarshal(data, &hexStr); err != nil {
//		return fmt.Errorf("invalid signature format: %w", err)
//	}
//
//	decoded, err := hex.DecodeString(hexStr)
//	if err != nil {
//		return fmt.Errorf("invalid hex string: %w", err)
//	}
//
//	if len(decoded) != 64 {
//		return fmt.Errorf("invalid signature length: got %d, want 64", len(decoded))
//	}
//
//	copy(s[:], decoded)
//	return nil
//}
