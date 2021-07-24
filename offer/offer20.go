package offer

type  State int
type CharType int

const (
	SPACE_STATE State = iota
	INITIAL_STATE
	INTEGER_STATE
	ONLY_SIGN_STATE
	ONLY_POINT_STATE
	POINT_STATE
	EXP_STATE
	ONLY_SIGN_EXP_STATE
	SIGN_EXP_STATE
	FINNAL_STATE
)

const (
	SPACE CharType = iota
	SIGN
	POINT
	NUM
	E
	UNEXPECTED
)

func isNumber(s string) bool {
	toType := func(b byte) CharType {
		switch b {
		case ' ':
			return SPACE
		case '+','-':
			return SIGN
		case '.':
			return POINT
		case '0', '1', '2', '3', '4', '5', '6', '7', '8','9':
			return NUM
		case 'e', 'E':
			return E
		default:
			return UNEXPECTED
		}
	}

	transfer := map[State]map[CharType]State {
		INITIAL_STATE: {
			SPACE: SPACE_STATE,
			SIGN: ONLY_SIGN_STATE,
			NUM: INTEGER_STATE,
			POINT: ONLY_POINT_STATE,
		},
		SPACE_STATE: {
			SPACE: SPACE_STATE,
			SIGN: ONLY_SIGN_STATE,
			NUM: INTEGER_STATE,
			POINT: ONLY_POINT_STATE,
		},
		ONLY_SIGN_STATE: {
			NUM: INTEGER_STATE,
			POINT: ONLY_POINT_STATE,
		},
		ONLY_POINT_STATE: {
			NUM: POINT_STATE,
		},
		ONLY_SIGN_EXP_STATE: {
			NUM: SIGN_EXP_STATE,
		},
		INTEGER_STATE: {
			SPACE: FINNAL_STATE,
			NUM: INTEGER_STATE,
			E: EXP_STATE,
			POINT: POINT_STATE,
		},
		POINT_STATE: {
			SPACE: FINNAL_STATE,
			NUM: POINT_STATE,
			E: EXP_STATE,
		},
		EXP_STATE: {
			SIGN: ONLY_SIGN_EXP_STATE,
			NUM: SIGN_EXP_STATE,
		},
		SIGN_EXP_STATE: {
			SPACE: FINNAL_STATE,
			NUM: SIGN_EXP_STATE,
		},
		FINNAL_STATE: {
			SPACE: FINNAL_STATE,
		},
	}

	currentState := INITIAL_STATE
	for _, b := range []byte(s) {
		charType := toType(b)
		state, ok := transfer[currentState][charType]
		if !ok {
			return false
		}
		currentState = state
	}
	if currentState == SPACE_STATE || currentState == INITIAL_STATE || currentState == ONLY_POINT_STATE || currentState == ONLY_SIGN_STATE || currentState == ONLY_SIGN_EXP_STATE || currentState == EXP_STATE {
		return false
	}
	return true
}