// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetupdateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetUpdateOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AssetUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"asset_to_update":`)

	{

		obj, err = j.AssetToUpdate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"issuer":`)

	{

		obj, err = j.Issuer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"fee":`)
	err = buf.Encode(&j.Fee)
	if err != nil {
		return err
	}
	buf.WriteString(`,"extensions":`)
	if j.Extensions != nil {
		buf.WriteString(`[`)
		for i, v := range j.Extensions {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=interface {} kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	if j.NewIssuer != nil {
		buf.WriteString(`,"new_issuer":`)

		{

			obj, err = j.NewIssuer.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
	} else {
		buf.WriteString(`,"new_issuer":null`)
	}
	/* Struct fall back. type=types.AssetOptions kind=struct */
	buf.WriteString(`,"new_options":`)
	err = buf.Encode(&j.NewOptions)
	if err != nil {
		return err
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAssetUpdateOperationbase = iota
	ffjtAssetUpdateOperationnosuchkey

	ffjtAssetUpdateOperationAssetToUpdate

	ffjtAssetUpdateOperationIssuer

	ffjtAssetUpdateOperationFee

	ffjtAssetUpdateOperationExtensions

	ffjtAssetUpdateOperationNewIssuer

	ffjtAssetUpdateOperationNewOptions
)

var ffjKeyAssetUpdateOperationAssetToUpdate = []byte("asset_to_update")

var ffjKeyAssetUpdateOperationIssuer = []byte("issuer")

var ffjKeyAssetUpdateOperationFee = []byte("fee")

var ffjKeyAssetUpdateOperationExtensions = []byte("extensions")

var ffjKeyAssetUpdateOperationNewIssuer = []byte("new_issuer")

var ffjKeyAssetUpdateOperationNewOptions = []byte("new_options")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetUpdateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAssetUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAssetUpdateOperationAssetToUpdate, kn) {
						currentKey = ffjtAssetUpdateOperationAssetToUpdate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAssetUpdateOperationExtensions, kn) {
						currentKey = ffjtAssetUpdateOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetUpdateOperationFee, kn) {
						currentKey = ffjtAssetUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAssetUpdateOperationIssuer, kn) {
						currentKey = ffjtAssetUpdateOperationIssuer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAssetUpdateOperationNewIssuer, kn) {
						currentKey = ffjtAssetUpdateOperationNewIssuer
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetUpdateOperationNewOptions, kn) {
						currentKey = ffjtAssetUpdateOperationNewOptions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateOperationNewOptions, kn) {
					currentKey = ffjtAssetUpdateOperationNewOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateOperationNewIssuer, kn) {
					currentKey = ffjtAssetUpdateOperationNewIssuer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateOperationExtensions, kn) {
					currentKey = ffjtAssetUpdateOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetUpdateOperationFee, kn) {
					currentKey = ffjtAssetUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateOperationIssuer, kn) {
					currentKey = ffjtAssetUpdateOperationIssuer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateOperationAssetToUpdate, kn) {
					currentKey = ffjtAssetUpdateOperationAssetToUpdate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAssetUpdateOperationAssetToUpdate:
					goto handle_AssetToUpdate

				case ffjtAssetUpdateOperationIssuer:
					goto handle_Issuer

				case ffjtAssetUpdateOperationFee:
					goto handle_Fee

				case ffjtAssetUpdateOperationExtensions:
					goto handle_Extensions

				case ffjtAssetUpdateOperationNewIssuer:
					goto handle_NewIssuer

				case ffjtAssetUpdateOperationNewOptions:
					goto handle_NewOptions

				case ffjtAssetUpdateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_AssetToUpdate:

	/* handler: j.AssetToUpdate type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.AssetToUpdate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Issuer:

	/* handler: j.Issuer type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Issuer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Extensions", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Extensions = nil
		} else {

			j.Extensions = []interface{}{}

			wantVal := true

			for {

				var tmpJExtensions interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJExtensions type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJExtensions)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Extensions = append(j.Extensions, tmpJExtensions)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NewIssuer:

	/* handler: j.NewIssuer type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.NewIssuer = nil

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			if j.NewIssuer == nil {
				j.NewIssuer = new(types.GrapheneID)
			}

			err = j.NewIssuer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NewOptions:

	/* handler: j.NewOptions type=types.AssetOptions kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetOptions kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.NewOptions)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
