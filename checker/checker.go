// Package checker re-exports the internal type checker API for external
// consumers of this fork. See package github.com/microsoft/typescript-go/shim
// for the governing policy.
package checker

import (
	"github.com/microsoft/typescript-go/internal/checker"
)

type (
	Checker = checker.Checker
	Type    = checker.Type

	ElementFlags = checker.ElementFlags
	TypeFlags    = checker.TypeFlags
)

// Type flags that the psgen and superschematic TypeScript readers test on a
// checker Type. The *Like flags are the compiler's own unions of several
// flags. Each one aliases the compiler's constant, so a reader never
// hardcodes a flag bit that a compiler release can move.
const (
	TypeFlagsAny            = checker.TypeFlagsAny
	TypeFlagsBoolean        = checker.TypeFlagsBoolean
	TypeFlagsBooleanLike    = checker.TypeFlagsBooleanLike
	TypeFlagsBooleanLiteral = checker.TypeFlagsBooleanLiteral
	TypeFlagsEnumLiteral    = checker.TypeFlagsEnumLiteral
	TypeFlagsIntersection   = checker.TypeFlagsIntersection
	TypeFlagsNever          = checker.TypeFlagsNever
	TypeFlagsNull           = checker.TypeFlagsNull
	TypeFlagsNumber         = checker.TypeFlagsNumber
	TypeFlagsNumberLike     = checker.TypeFlagsNumberLike
	TypeFlagsNumberLiteral  = checker.TypeFlagsNumberLiteral
	TypeFlagsObject         = checker.TypeFlagsObject
	TypeFlagsString         = checker.TypeFlagsString
	TypeFlagsStringLike     = checker.TypeFlagsStringLike
	TypeFlagsStringLiteral  = checker.TypeFlagsStringLiteral
	TypeFlagsUndefined      = checker.TypeFlagsUndefined
	TypeFlagsUnion          = checker.TypeFlagsUnion
	TypeFlagsUnknown        = checker.TypeFlagsUnknown
)

// Tuple element flags. A tuple element that is not optional, rest or
// variadic has exactly ElementFlagsRequired.
const ElementFlagsRequired = checker.ElementFlagsRequired
