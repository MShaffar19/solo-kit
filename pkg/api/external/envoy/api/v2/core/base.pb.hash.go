// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-kit/api/external/envoy/api/v2/core/base.proto

package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *Locality) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.Locality")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetZone())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSubZone())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Node) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.Node")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetLocality()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetLocality(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetBuildVersion())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Metadata) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.Metadata")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetFilterMetadata() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if val, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if err := binary.Write(innerHash, binary.LittleEndian, val); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RuntimeUInt32) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.RuntimeUInt32")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDefaultValue())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RuntimeFeatureFlag) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.RuntimeFeatureFlag")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDefaultValue()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetDefaultValue(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HeaderValue) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.HeaderValue")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HeaderValueOption) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.HeaderValueOption")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAppend()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetAppend(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	switch m.HeaderOption.(type) {

	case *HeaderValueOption_Header:

		if h, ok := interface{}(m.GetHeader()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetHeader(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *HeaderValueOption_HeaderSecretRef:

		if h, ok := interface{}(m.GetHeaderSecretRef()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetHeaderSecretRef(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HeaderMap) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.HeaderMap")); err != nil {
		return 0, err
	}

	for _, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DataSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.DataSource")); err != nil {
		return 0, err
	}

	switch m.Specifier.(type) {

	case *DataSource_Filename:

		if _, err = hasher.Write([]byte(m.GetFilename())); err != nil {
			return 0, err
		}

	case *DataSource_InlineBytes:

		if _, err = hasher.Write(m.GetInlineBytes()); err != nil {
			return 0, err
		}

	case *DataSource_InlineString:

		if _, err = hasher.Write([]byte(m.GetInlineString())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RemoteDataSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.RemoteDataSource")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHttpUri()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetHttpUri(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetSha256())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AsyncDataSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.AsyncDataSource")); err != nil {
		return 0, err
	}

	switch m.Specifier.(type) {

	case *AsyncDataSource_Local:

		if h, ok := interface{}(m.GetLocal()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetLocal(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AsyncDataSource_Remote:

		if h, ok := interface{}(m.GetRemote()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetRemote(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TransportSocket) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.TransportSocket")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	switch m.ConfigType.(type) {

	case *TransportSocket_Config:

		if h, ok := interface{}(m.GetConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetConfig(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *TransportSocket_TypedConfig:

		if h, ok := interface{}(m.GetTypedConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetTypedConfig(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *SocketOption) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.SocketOption")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDescription())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetLevel())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetName())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	switch m.Value.(type) {

	case *SocketOption_IntValue:

		err = binary.Write(hasher, binary.LittleEndian, m.GetIntValue())
		if err != nil {
			return 0, err
		}

	case *SocketOption_BufValue:

		if _, err = hasher.Write(m.GetBufValue()); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RuntimeFractionalPercent) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.RuntimeFractionalPercent")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDefaultValue()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetDefaultValue(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ControlPlane) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core.ControlPlane")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetIdentifier())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
