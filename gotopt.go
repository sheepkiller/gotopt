// Package to generate/validate TOPT from 4 to 8 digits with SHA1/SHA256/SHA512 hash functions
package gotopt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"math"
	"time"
)

const window = 30

type topt struct {
	secret []byte
	digits int
	shaX   string
	msg    []byte
	ts     uint64
}

func (t *topt) getHmac() (hash.Hash, error) {
	switch t.shaX {
	case "sha1":
		return hmac.New(sha1.New, t.secret), nil
	case "sha256":
		return hmac.New(sha256.New, t.secret), nil
	case "sha512":
		return hmac.New(sha512.New, t.secret), nil
	}
	return nil, errors.New("Unsupported SHA function")
}

func newTOPT(str_secret string, digits int, shaX string) (t *topt, err error) {
	if digits > 8 || digits < 4 {
		return nil, errors.New("digits must be beetween 4 and 8")
	}
	t = new(topt)
	t.secret, err = base32.StdEncoding.DecodeString(str_secret)
	if err != nil {
		return nil, err
	}
	t.digits = digits
	t.shaX = shaX
	t.msg = make([]byte, 8)
	t.ts = uint64(time.Now().Unix())
	return t, nil
}

func (t *topt) TOPT() (str_token string, remain uint64, err error) {
	mac, _ := t.getHmac()
	binary.BigEndian.PutUint64(t.msg, uint64(t.ts/window))
	x := uint64(t.ts / window)
	remain = (x+1)*window - t.ts
	mac.Write(t.msg)
	h := mac.Sum(nil)
	offset := int(h[19]) & 0xF
	token :=
		((int(h[offset+0]) & 0x7F) << 24) |
			((int(h[offset+1]) & 0xFF) << 16) |
			((int(h[offset+2]) & 0xFF) << 8) |
			(int(h[offset+3]) & 0xFF)
	token = token % int(math.Pow10(t.digits))
	format := fmt.Sprintf("%%0%dd", t.digits)
	str_token = fmt.Sprintf(format, token)
	return str_token, remain, nil
}

// Generate a TOPT
//   str_secret: secret key (strict base32)
//   digits: number of digits (from 4 to 8)
//   shaX: hash function (SHA1, SHA256, or SHA512)
// returns:
//  str_token: string representation of token (type string)
//  remain: remaining validity time (type uint64)
//  err: error
func GetTOPT(str_secret string, digits int, shaX string) (str_token string, remain uint64, err error) {
	t, err := newTOPT(str_secret, digits, shaX)
	if err != nil {
		return "a", 0, err
	}
	return t.TOPT()
}

// Validate a TOPT
//   str_secret: secret key (strict base32)
//   digits: number of digits (from 4 to 8)
//   shaX: hash function (SHA1, SHA256, or SHA512)
//   topt: token to validate (type string)
//   interval: validate against previous interval window (30s)
// returns
//   valid_token: boolean : true if valid, false if not
//   err: error
// XXX: need test

func ValidateTOPT(str_secret string, digits int, shaX string, topt string, interval int) (valid_token bool, err error) {
	t, err := newTOPT(str_secret, digits, shaX)
	if err != nil {
		return false, err
	}
	base_ts := t.ts
	for i := 0; i <= interval; i++ {
		t.ts = base_ts - uint64(window*i)
		token, _, err := t.TOPT()
		if err != nil {
			return false, err
		}
		if topt == token {
			return true, nil
		}
	}
	return false, nil
}
