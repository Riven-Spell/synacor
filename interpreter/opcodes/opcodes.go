package opcodes

const (
	HALT uint16 = iota
	SET
	PUSH
	POP
	EQ
	GT
	JMP
	JT
	JF
	ADD
	MULT
	MOD
	AND
	OR
	NOT
	RMEM
	WMEM
	CALL
	RET
	OUT
	IN
	NOOP
)

var OpcodeLength = map[uint16]uint16{
	HALT: 1,
	SET:  3,
	PUSH: 2,
	POP:  2,
	EQ:   4,
	GT:   4,
	JMP:  6,
	JT:   3,
	JF:   3,
	ADD:  4,
	MULT: 4,
	MOD:  4,
	AND:  4,
	OR:   4,
	NOT:  3,
	RMEM: 3,
	WMEM: 3,
	CALL: 2,
	RET:  1,
	OUT:  2,
	IN:   2,
	NOOP: 1,
}
