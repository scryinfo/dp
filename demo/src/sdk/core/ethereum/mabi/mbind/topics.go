// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package mbind

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	abi "../../mabi"
)

// makeTopics converts a filter query argument list into a filter topic set.
func makeTopics(query ...[]interface{}) ([][]common.Hash, error) {
	topics := make([][]common.Hash, len(query))
	for i, filter := range query {
		for _, rule := range filter {
			var topic common.Hash

			// Try to generate the topic based on simple types
			switch rule := rule.(type) {
			case common.Hash:
				copy(topic[:], rule[:])
			case common.Address:
				copy(topic[common.HashLength-common.AddressLength:], rule[:])
			case *big.Int:
				blob := rule.Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case bool:
				if rule {
					topic[common.HashLength-1] = 1
				}
			case int8:
				blob := big.NewInt(int64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case int16:
				blob := big.NewInt(int64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case int32:
				blob := big.NewInt(int64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case int64:
				blob := big.NewInt(rule).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint8:
				blob := new(big.Int).SetUint64(uint64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint16:
				blob := new(big.Int).SetUint64(uint64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint32:
				blob := new(big.Int).SetUint64(uint64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint64:
				blob := new(big.Int).SetUint64(rule).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case string:
				hash := crypto.Keccak256Hash([]byte(rule))
				copy(topic[:], hash[:])
			case []byte:
				hash := crypto.Keccak256Hash(rule)
				copy(topic[:], hash[:])

			default:
				// Attempt to generate the topic from funky types
				val := reflect.ValueOf(rule)

				switch {
				case val.Kind() == reflect.Array && reflect.TypeOf(rule).Elem().Kind() == reflect.Uint8:
					reflect.Copy(reflect.ValueOf(topic[common.HashLength-val.Len():]), val)

				default:
					return nil, fmt.Errorf("unsupported indexed type: %T", rule)
				}
			}
			topics[i] = append(topics[i], topic)
		}
	}
	return topics, nil
}

// Big batch of reflect types for topic reconstruction.
var (
	reflectHash    = reflect.TypeOf(common.Hash{})
	reflectAddress = reflect.TypeOf(common.Address{})
	reflectBigInt  = reflect.TypeOf(new(big.Int))
)

// parseTopics converts the indexed topic fields into actual log field values.
//
// Note, dynamic types cannot be reconstructed since they get mapped to Keccak256
// hashes as the topic value!
func parseTopics(out abi.JSONObj, fields abi.Arguments, topics []common.Hash) error {
	// Sanity check that the fields and topics match up
	if len(fields) != len(topics) {
		return errors.New("topic/field count mismatch")
	}
	// Iterate over all the fields and reconstruct them from topics
	for _, arg := range fields {
		if !arg.Indexed {
			return errors.New("non-indexed field in topic reconstruction")
		}
		name := arg.Name
		// Try to parse the topic back into the fields based on primitive types
		switch arg.Type.Kind {
		case reflect.Bool:
			if topics[0][common.HashLength-1] == 1 {
				out.Set(name, true)
			} else {
				out.Set(name, false)
			}
		case reflect.Int8:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Int16:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Int32:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Int64:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Int64())

		case reflect.Uint8:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, uint8(num.Int64()))

		case reflect.Uint16:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, uint16(num.Int64()))

		case reflect.Uint32:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, uint32(num.Int64()))

		case reflect.Uint64:
			num := new(big.Int).SetBytes(topics[0][:])
			out.Set(name, num.Uint64())

		default:
			// Ran out of plain primitive types, try custom types
			switch arg.Type.Type {
			case reflectHash: // Also covers all dynamic types
				out.Set(name, topics[0].Hex())

			case reflectAddress:
				var addr common.Address
				copy(addr[:], topics[0][common.HashLength-common.AddressLength:])
				out.Set(name, addr.Hex())

			case reflectBigInt:
				num := new(big.Int).SetBytes(topics[0][:])
				out.Set(name, num)

			default:
				// Ran out of custom types, try the crazies
				switch {
				case arg.Type.T == abi.FixedBytesTy:
					out.Set(name, topics[0][common.HashLength-arg.Type.Size:])

				default:
					return fmt.Errorf("unsupported indexed type: %v", arg.Type)
				}
			}
		}
		topics = topics[1:]
	}
	return nil
}
