// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jobs/jobspb/schedule.proto

package jobspb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// WaitBehavior describes how to handle previously  started
// jobs that have not completed yet.
type ScheduleDetails_WaitBehavior int32

const (
	// Wait for the previous run to complete
	// before starting the next one.
	ScheduleDetails_WAIT ScheduleDetails_WaitBehavior = 0
	// Do not wait for the previous run to complete.
	ScheduleDetails_NO_WAIT ScheduleDetails_WaitBehavior = 1
	// If the previous run is still running, skip this run
	// and advance schedule to the next recurrence.
	ScheduleDetails_SKIP ScheduleDetails_WaitBehavior = 2
)

var ScheduleDetails_WaitBehavior_name = map[int32]string{
	0: "WAIT",
	1: "NO_WAIT",
	2: "SKIP",
}
var ScheduleDetails_WaitBehavior_value = map[string]int32{
	"WAIT":    0,
	"NO_WAIT": 1,
	"SKIP":    2,
}

func (x ScheduleDetails_WaitBehavior) String() string {
	return proto.EnumName(ScheduleDetails_WaitBehavior_name, int32(x))
}
func (ScheduleDetails_WaitBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_2c4135f09ebfe46a, []int{0, 0}
}

// ErrorHandlingBehavior describes how to handle failed job runs.
type ScheduleDetails_ErrorHandlingBehavior int32

const (
	// By default, failed jobs will run again, based on their schedule.
	ScheduleDetails_RETRY_SCHED ScheduleDetails_ErrorHandlingBehavior = 0
	// Retry failed jobs soon.
	ScheduleDetails_RETRY_SOON ScheduleDetails_ErrorHandlingBehavior = 1
	// Stop running this schedule
	ScheduleDetails_PAUSE_SCHED ScheduleDetails_ErrorHandlingBehavior = 2
)

var ScheduleDetails_ErrorHandlingBehavior_name = map[int32]string{
	0: "RETRY_SCHED",
	1: "RETRY_SOON",
	2: "PAUSE_SCHED",
}
var ScheduleDetails_ErrorHandlingBehavior_value = map[string]int32{
	"RETRY_SCHED": 0,
	"RETRY_SOON":  1,
	"PAUSE_SCHED": 2,
}

func (x ScheduleDetails_ErrorHandlingBehavior) String() string {
	return proto.EnumName(ScheduleDetails_ErrorHandlingBehavior_name, int32(x))
}
func (ScheduleDetails_ErrorHandlingBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_2c4135f09ebfe46a, []int{0, 1}
}

// ScheduleDetails describes how to schedule and execute the job.
type ScheduleDetails struct {
	// How to handle running jobs.
	Wait ScheduleDetails_WaitBehavior `protobuf:"varint,1,opt,name=wait,proto3,enum=cockroach.jobs.jobspb.ScheduleDetails_WaitBehavior" json:"wait,omitempty"`
	// How to handle failed jobs.
	OnError ScheduleDetails_ErrorHandlingBehavior `protobuf:"varint,2,opt,name=on_error,json=onError,proto3,enum=cockroach.jobs.jobspb.ScheduleDetails_ErrorHandlingBehavior" json:"on_error,omitempty"`
}

func (m *ScheduleDetails) Reset()         { *m = ScheduleDetails{} }
func (m *ScheduleDetails) String() string { return proto.CompactTextString(m) }
func (*ScheduleDetails) ProtoMessage()    {}
func (*ScheduleDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_2c4135f09ebfe46a, []int{0}
}
func (m *ScheduleDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ScheduleDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleDetails.Merge(dst, src)
}
func (m *ScheduleDetails) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleDetails proto.InternalMessageInfo

// ExecutionArguments describes data needed to execute scheduled jobs.
type ExecutionArguments struct {
	Args *types.Any `protobuf:"bytes,1,opt,name=args,proto3" json:"args,omitempty"`
}

func (m *ExecutionArguments) Reset()         { *m = ExecutionArguments{} }
func (m *ExecutionArguments) String() string { return proto.CompactTextString(m) }
func (*ExecutionArguments) ProtoMessage()    {}
func (*ExecutionArguments) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_2c4135f09ebfe46a, []int{1}
}
func (m *ExecutionArguments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecutionArguments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ExecutionArguments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionArguments.Merge(dst, src)
}
func (m *ExecutionArguments) XXX_Size() int {
	return m.Size()
}
func (m *ExecutionArguments) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionArguments.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionArguments proto.InternalMessageInfo

// Message representing sql statement to execute.
type SqlStatementExecutionArg struct {
	Statement string `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
}

func (m *SqlStatementExecutionArg) Reset()         { *m = SqlStatementExecutionArg{} }
func (m *SqlStatementExecutionArg) String() string { return proto.CompactTextString(m) }
func (*SqlStatementExecutionArg) ProtoMessage()    {}
func (*SqlStatementExecutionArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_2c4135f09ebfe46a, []int{2}
}
func (m *SqlStatementExecutionArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SqlStatementExecutionArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SqlStatementExecutionArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlStatementExecutionArg.Merge(dst, src)
}
func (m *SqlStatementExecutionArg) XXX_Size() int {
	return m.Size()
}
func (m *SqlStatementExecutionArg) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlStatementExecutionArg.DiscardUnknown(m)
}

var xxx_messageInfo_SqlStatementExecutionArg proto.InternalMessageInfo

// ScheduleState represents mutable schedule state.
// The members of this proto may be mutated during each schedule execution.
type ScheduleState struct {
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *ScheduleState) Reset()         { *m = ScheduleState{} }
func (m *ScheduleState) String() string { return proto.CompactTextString(m) }
func (*ScheduleState) ProtoMessage()    {}
func (*ScheduleState) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_2c4135f09ebfe46a, []int{3}
}
func (m *ScheduleState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ScheduleState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleState.Merge(dst, src)
}
func (m *ScheduleState) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleState) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleState.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScheduleDetails)(nil), "cockroach.jobs.jobspb.ScheduleDetails")
	proto.RegisterType((*ExecutionArguments)(nil), "cockroach.jobs.jobspb.ExecutionArguments")
	proto.RegisterType((*SqlStatementExecutionArg)(nil), "cockroach.jobs.jobspb.SqlStatementExecutionArg")
	proto.RegisterType((*ScheduleState)(nil), "cockroach.jobs.jobspb.ScheduleState")
	proto.RegisterEnum("cockroach.jobs.jobspb.ScheduleDetails_WaitBehavior", ScheduleDetails_WaitBehavior_name, ScheduleDetails_WaitBehavior_value)
	proto.RegisterEnum("cockroach.jobs.jobspb.ScheduleDetails_ErrorHandlingBehavior", ScheduleDetails_ErrorHandlingBehavior_name, ScheduleDetails_ErrorHandlingBehavior_value)
}
func (m *ScheduleDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Wait != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Wait))
	}
	if m.OnError != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.OnError))
	}
	return i, nil
}

func (m *ExecutionArguments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecutionArguments) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Args != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Args.Size()))
		n1, err := m.Args.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SqlStatementExecutionArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SqlStatementExecutionArg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Statement) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Statement)))
		i += copy(dAtA[i:], m.Statement)
	}
	return i, nil
}

func (m *ScheduleState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	return i, nil
}

func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ScheduleDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Wait != 0 {
		n += 1 + sovSchedule(uint64(m.Wait))
	}
	if m.OnError != 0 {
		n += 1 + sovSchedule(uint64(m.OnError))
	}
	return n
}

func (m *ExecutionArguments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Args != nil {
		l = m.Args.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *SqlStatementExecutionArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Statement)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *ScheduleState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduleDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wait", wireType)
			}
			m.Wait = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wait |= (ScheduleDetails_WaitBehavior(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnError", wireType)
			}
			m.OnError = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnError |= (ScheduleDetails_ErrorHandlingBehavior(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecutionArguments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecutionArguments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutionArguments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Args == nil {
				m.Args = &types.Any{}
			}
			if err := m.Args.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SqlStatementExecutionArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SqlStatementExecutionArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SqlStatementExecutionArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduleState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSchedule
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSchedule(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSchedule = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("jobs/jobspb/schedule.proto", fileDescriptor_schedule_2c4135f09ebfe46a)
}

var fileDescriptor_schedule_2c4135f09ebfe46a = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0xa8, 0xea, 0xb6, 0x57, 0xd8, 0x22, 0x6b, 0x43, 0xa5, 0x42, 0x16, 0xca, 0x85,
	0x9e, 0x1c, 0x69, 0xbb, 0x70, 0x40, 0x48, 0x19, 0x8b, 0x58, 0x85, 0xb4, 0x4e, 0xc9, 0x50, 0x05,
	0x97, 0xca, 0xc9, 0x4c, 0x1a, 0x08, 0xf6, 0x70, 0x1c, 0x60, 0xdf, 0x82, 0x6f, 0xc2, 0xd7, 0xd8,
	0x71, 0xc7, 0x1d, 0x21, 0xfd, 0x22, 0xc8, 0x4e, 0x0a, 0x15, 0xea, 0x61, 0x97, 0x28, 0xef, 0xbd,
	0xdf, 0xfb, 0x3d, 0xeb, 0x0f, 0xa3, 0x8f, 0x32, 0xad, 0x02, 0xf3, 0xb9, 0x4a, 0x83, 0x2a, 0x5b,
	0xf0, 0xcb, 0xba, 0xe4, 0xf4, 0x4a, 0x49, 0x2d, 0xf1, 0x41, 0x26, 0xb3, 0x4f, 0x4a, 0xb2, 0x6c,
	0x41, 0x0d, 0x40, 0x5b, 0x6a, 0xb4, 0x9f, 0xcb, 0x5c, 0x5a, 0x22, 0x30, 0x7f, 0x2d, 0x3c, 0x7a,
	0x9c, 0x4b, 0x99, 0x97, 0x3c, 0xb0, 0x55, 0x5a, 0x7f, 0x08, 0x98, 0xb8, 0x6e, 0x47, 0xfe, 0x4f,
	0x17, 0xf6, 0x92, 0x4e, 0x7d, 0xc2, 0x35, 0x2b, 0xca, 0x0a, 0xbf, 0x86, 0xde, 0x37, 0x56, 0xe8,
	0x21, 0x7a, 0x8a, 0xc6, 0xbb, 0x87, 0x47, 0x74, 0xe3, 0x29, 0xfa, 0xdf, 0x16, 0x9d, 0xb1, 0x42,
	0x1f, 0xf3, 0x05, 0xfb, 0x5a, 0x48, 0x15, 0x5b, 0x01, 0x9e, 0xc1, 0xb6, 0x14, 0x73, 0xae, 0x94,
	0x54, 0x43, 0xd7, 0xca, 0x5e, 0xdc, 0x53, 0x16, 0x99, 0x9d, 0x53, 0x26, 0x2e, 0xcb, 0x42, 0xe4,
	0x7f, 0xad, 0x5b, 0x52, 0xd8, 0x81, 0x1f, 0xc0, 0x83, 0xf5, 0x73, 0x78, 0x1b, 0x7a, 0xb3, 0x70,
	0x72, 0xe1, 0x39, 0x78, 0x00, 0x5b, 0x67, 0xd3, 0xb9, 0x2d, 0x90, 0x69, 0x27, 0x6f, 0x26, 0xe7,
	0x9e, 0xeb, 0x4f, 0xe0, 0x60, 0xa3, 0x12, 0xef, 0xc1, 0x20, 0x8e, 0x2e, 0xe2, 0x77, 0xf3, 0xe4,
	0xd5, 0x69, 0x74, 0xe2, 0x39, 0x78, 0x17, 0xa0, 0x6b, 0x4c, 0xa7, 0x67, 0x1e, 0x32, 0xc0, 0x79,
	0xf8, 0x36, 0x89, 0x3a, 0xc0, 0xf5, 0x5f, 0x02, 0x8e, 0xbe, 0xf3, 0xac, 0xd6, 0x85, 0x14, 0xa1,
	0xca, 0xeb, 0xcf, 0x5c, 0xe8, 0x0a, 0x8f, 0xa1, 0xc7, 0x54, 0x5e, 0xd9, 0xcc, 0x06, 0x87, 0xfb,
	0xb4, 0x4d, 0x9c, 0xae, 0x12, 0xa7, 0xa1, 0xb8, 0x8e, 0x2d, 0xe1, 0x3f, 0x87, 0x61, 0xf2, 0xa5,
	0x4c, 0x34, 0xd3, 0xdc, 0xac, 0xae, 0xbb, 0xf0, 0x13, 0xd8, 0xa9, 0x56, 0x03, 0xab, 0xda, 0x89,
	0xff, 0x35, 0xfc, 0x67, 0xf0, 0x70, 0x95, 0x93, 0x5d, 0xc7, 0x8f, 0xa0, 0x6f, 0xa6, 0x75, 0xd5,
	0xb1, 0x5d, 0x75, 0x3c, 0xbe, 0xf9, 0x4d, 0x9c, 0x9b, 0x86, 0xa0, 0xdb, 0x86, 0xa0, 0xbb, 0x86,
	0xa0, 0x5f, 0x0d, 0x41, 0x3f, 0x96, 0xc4, 0xb9, 0x5d, 0x12, 0xe7, 0x6e, 0x49, 0x9c, 0xf7, 0xfd,
	0x36, 0xf7, 0xb4, 0x6f, 0x1f, 0x78, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xd6, 0x1a, 0xd9,
	0x6b, 0x02, 0x00, 0x00,
}
