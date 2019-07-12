// Code generated by solo-kit. DO NOT EDIT.

package v2alpha1

import (
	"fmt"

	testing_solo_io_v1 "github.com/solo-io/solo-kit/test/mocks/v1"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type TestingSnapshot struct {
	Mocks MockResourceList
	Fakes testing_solo_io_v1.FakeResourceList
}

func (s TestingSnapshot) Clone() TestingSnapshot {
	return TestingSnapshot{
		Mocks: s.Mocks.Clone(),
		Fakes: s.Fakes.Clone(),
	}
}

func (s TestingSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMocks(),
		s.hashFakes(),
	)
}

func (s TestingSnapshot) hashMocks() uint64 {
	return hashutils.HashAll(s.Mocks.AsInterfaces()...)
}

func (s TestingSnapshot) hashFakes() uint64 {
	return hashutils.HashAll(s.Fakes.AsInterfaces()...)
}

func (s TestingSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("mocks", s.hashMocks()))
	fields = append(fields, zap.Uint64("fakes", s.hashFakes()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type TestingSnapshotStringer struct {
	Version uint64
	Mocks   []string
	Fakes   []string
}

func (ss TestingSnapshotStringer) String() string {
	s := fmt.Sprintf("TestingSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Mocks %v\n", len(ss.Mocks))
	for _, name := range ss.Mocks {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Fakes %v\n", len(ss.Fakes))
	for _, name := range ss.Fakes {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s TestingSnapshot) Stringer() TestingSnapshotStringer {
	return TestingSnapshotStringer{
		Version: s.Hash(),
		Mocks:   s.Mocks.NamespacesDotNames(),
		Fakes:   s.Fakes.NamespacesDotNames(),
	}
}
