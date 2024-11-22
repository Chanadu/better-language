// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package tokentype

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LeftParentheses-0]
	_ = x[RightParentheses-1]
	_ = x[LeftCurlyBrace-2]
	_ = x[RightCurlyBrace-3]
	_ = x[Comma-4]
	_ = x[Dot-5]
	_ = x[Semicolon-6]
	_ = x[Minus-7]
	_ = x[Plus-8]
	_ = x[Star-9]
	_ = x[Percent-10]
	_ = x[Slash-11]
	_ = x[Not-12]
	_ = x[NotEqual-13]
	_ = x[Equal-14]
	_ = x[EqualEqual-15]
	_ = x[Greater-16]
	_ = x[GreaterEqual-17]
	_ = x[Less-18]
	_ = x[LessEqual-19]
	_ = x[Identifier-20]
	_ = x[String-21]
	_ = x[Integer-22]
	_ = x[Double-23]
	_ = x[Or-24]
	_ = x[And-25]
	_ = x[False-26]
	_ = x[True-27]
	_ = x[If-28]
	_ = x[Else-29]
	_ = x[For-30]
	_ = x[While-31]
	_ = x[Break-32]
	_ = x[Return-33]
	_ = x[Function-34]
	_ = x[Print-35]
	_ = x[Var-36]
	_ = x[Null-37]
	_ = x[EndOfFile-38]
	_ = x[Class-39]
	_ = x[This-40]
	_ = x[Super-41]
}

const _TokenType_name = "LeftParenthesesRightParenthesesLeftCurlyBraceRightCurlyBraceCommaDotSemicolonMinusPlusStarPercentSlashNotNotEqualEqualEqualEqualGreaterGreaterEqualLessLessEqualIdentifierStringIntegerDoubleOrAndFalseTrueIfElseForWhileBreakReturnFunctionPrintVarNullEndOfFileClassThisSuper"

var _TokenType_index = [...]uint16{0, 15, 31, 45, 60, 65, 68, 77, 82, 86, 90, 97, 102, 105, 113, 118, 128, 135, 147, 151, 160, 170, 176, 183, 189, 191, 194, 199, 203, 205, 209, 212, 217, 222, 228, 236, 241, 244, 248, 257, 262, 266, 271}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}