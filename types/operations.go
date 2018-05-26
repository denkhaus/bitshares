package types

import (
	"encoding/json"
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	OperationMap = make(map[OperationType]Operation)
)

type Operation interface {
	util.TypeMarshaller
	ApplyFee(fee AssetAmount)
	Type() OperationType
}

type OperationResult interface {
}

type OperationEnvelopeHolder struct {
	Op OperationEnvelope `json:"op"`
}

type OperationEnvelopeHolders []OperationEnvelopeHolder

func (p OperationEnvelopeHolders) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, op := range p {
		if err := enc.Encode(op.Op); err != nil {
			return errors.Annotate(err, "encode Op")
		}
	}

	return nil
}

type OperationEnvelope struct {
	Type      OperationType
	Operation Operation
}

func (p OperationEnvelope) Marshal(enc *util.TypeEncoder) error {
	// type is marshaled by operation
	if err := enc.Encode(p.Operation); err != nil {
		return errors.Annotate(err, "encode Operation")
	}

	return nil
}

func (p OperationEnvelope) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal([]interface{}{
		p.Type,
		p.Operation,
	})
}

func (p *OperationEnvelope) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	if len(raw) != 2 {
		return ErrInvalidInputLength
	}

	if err := ffjson.Unmarshal(raw[0], &p.Type); err != nil {
		return errors.Annotate(err, "unmarshal OperationType")
	}

	descr := fmt.Sprintf("Operation %s", p.Type)

	if op, ok := OperationMap[p.Type]; ok {
		p.Operation = op
		if err := ffjson.Unmarshal(raw[1], p.Operation); err != nil {
			util.DumpUnmarshaled(descr, raw[1])
			return errors.Annotatef(err, "unmarshal Operation %s", p.Type)
		}
	} else {
		fmt.Printf("Operation type %s not yet supported\n", p.Type)
		util.DumpUnmarshaled(descr, raw[1])
	}

	return nil
}

type Operations []Operation

func (p Operations) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode Operations length")
	}

	for _, op := range p {
		if err := enc.Encode(op); err != nil {
			return errors.Annotate(err, "encode Operation")
		}
	}

	return nil
}

func (p Operations) MarshalJSON() ([]byte, error) {
	env := make([]OperationEnvelope, len(p))
	for idx, op := range p {
		env[idx] = OperationEnvelope{
			Type:      op.Type(),
			Operation: op,
		}
	}

	return ffjson.Marshal(env)
}

func (p *Operations) UnmarshalJSON(data []byte) error {
	var envs []OperationEnvelope
	if err := ffjson.Unmarshal(data, &envs); err != nil {
		return err
	}

	ops := make(Operations, len(envs))
	for idx, env := range envs {
		if env.Operation != nil {
			ops[idx] = env.Operation.(Operation)
		}
	}

	*p = ops
	return nil
}

func (p Operations) ApplyFees(fees AssetAmounts) error {
	if len(p) != len(fees) {
		return errors.New("count of fees must match count of operations")
	}

	for idx, op := range p {
		op.ApplyFee(fees[idx])
	}

	return nil
}

func (p Operations) Types() [][]OperationType {
	ret := make([][]OperationType, len(p))
	for idx, op := range p {
		ret[idx] = []OperationType{op.Type()}
	}

	return ret
}
