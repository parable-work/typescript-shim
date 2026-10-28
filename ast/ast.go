// Package ast re-exports the internal AST API for external consumers of this
// fork. See package github.com/microsoft/typescript-go/shim for the governing
// policy.
package ast

import (
	"github.com/microsoft/typescript-go/internal/ast"
)

type (
	Node       = ast.Node
	NodeList   = ast.NodeList
	SourceFile = ast.SourceFile
	Symbol     = ast.Symbol
	Diagnostic = ast.Diagnostic
	Kind       = ast.Kind

	ModifierFlags = ast.ModifierFlags
	NodeFlags     = ast.NodeFlags
	SymbolFlags   = ast.SymbolFlags
)

// Node kinds the schema walker matches on. Expand as walker code demands.
const (
	KindNumericLiteral           = ast.KindNumericLiteral
	KindStringLiteral            = ast.KindStringLiteral
	KindIdentifier               = ast.KindIdentifier
	KindDecorator                = ast.KindDecorator
	KindPropertyDeclaration      = ast.KindPropertyDeclaration
	KindMethodDeclaration        = ast.KindMethodDeclaration
	KindTypeReference            = ast.KindTypeReference
	KindArrayType                = ast.KindArrayType
	KindUnionType                = ast.KindUnionType
	KindObjectLiteralExpression  = ast.KindObjectLiteralExpression
	KindPropertyAccessExpression = ast.KindPropertyAccessExpression
	KindCallExpression           = ast.KindCallExpression
	KindClassDeclaration         = ast.KindClassDeclaration
	KindInterfaceDeclaration     = ast.KindInterfaceDeclaration
	KindTypeAliasDeclaration     = ast.KindTypeAliasDeclaration
	KindEnumDeclaration          = ast.KindEnumDeclaration
	KindHeritageClause           = ast.KindHeritageClause
	KindSourceFile               = ast.KindSourceFile
)

// More node kinds that the psgen and superschematic TypeScript readers match
// on. Each one aliases the compiler's constant, so a reader never hardcodes a
// Kind number that a compiler release can move.
const (
	KindAnyKeyword                    = ast.KindAnyKeyword
	KindArrayLiteralExpression        = ast.KindArrayLiteralExpression
	KindAsExpression                  = ast.KindAsExpression
	KindBinaryExpression              = ast.KindBinaryExpression
	KindBooleanKeyword                = ast.KindBooleanKeyword
	KindComputedPropertyName          = ast.KindComputedPropertyName
	KindEnumMember                    = ast.KindEnumMember
	KindExportAssignment              = ast.KindExportAssignment
	KindExpressionWithTypeArguments   = ast.KindExpressionWithTypeArguments
	KindExtendsKeyword                = ast.KindExtendsKeyword
	KindFalseKeyword                  = ast.KindFalseKeyword
	KindImplementsKeyword             = ast.KindImplementsKeyword
	KindImportDeclaration             = ast.KindImportDeclaration
	KindImportSpecifier               = ast.KindImportSpecifier
	KindIndexedAccessType             = ast.KindIndexedAccessType
	KindLiteralType                   = ast.KindLiteralType
	KindMinusToken                    = ast.KindMinusToken
	KindNamedImports                  = ast.KindNamedImports
	KindNamespaceImport               = ast.KindNamespaceImport
	KindNoSubstitutionTemplateLiteral = ast.KindNoSubstitutionTemplateLiteral
	KindNullKeyword                   = ast.KindNullKeyword
	KindNumberKeyword                 = ast.KindNumberKeyword
	KindObjectKeyword                 = ast.KindObjectKeyword
	KindParameter                     = ast.KindParameter
	KindParenthesizedType             = ast.KindParenthesizedType
	KindPrefixUnaryExpression         = ast.KindPrefixUnaryExpression
	KindPropertyAssignment            = ast.KindPropertyAssignment
	KindPropertySignature             = ast.KindPropertySignature
	KindQualifiedName                 = ast.KindQualifiedName
	KindQuestionToken                 = ast.KindQuestionToken
	KindStringKeyword                 = ast.KindStringKeyword
	KindTemplateExpression            = ast.KindTemplateExpression
	KindTrueKeyword                   = ast.KindTrueKeyword
	KindTupleType                     = ast.KindTupleType
	KindTypeLiteral                   = ast.KindTypeLiteral
	KindTypeOperator                  = ast.KindTypeOperator
	KindUnknownKeyword                = ast.KindUnknownKeyword
	KindVariableDeclaration           = ast.KindVariableDeclaration
	KindVariableDeclarationList       = ast.KindVariableDeclarationList
	KindVariableStatement             = ast.KindVariableStatement
)

// Modifier flags that the readers test on a declaration.
const (
	ModifierFlagsAbstract = ast.ModifierFlagsAbstract
	ModifierFlagsExport   = ast.ModifierFlagsExport
)

// Node flags that tell let and const apart on a variable declaration list.
const (
	NodeFlagsConst = ast.NodeFlagsConst
	NodeFlagsLet   = ast.NodeFlagsLet
)

// Symbol flags that the readers test on a resolved symbol. SymbolFlagsEnum and
// SymbolFlagsModule are the compiler's own unions of two flags each.
const (
	SymbolFlagsAlias      = ast.SymbolFlagsAlias
	SymbolFlagsClass      = ast.SymbolFlagsClass
	SymbolFlagsEnum       = ast.SymbolFlagsEnum
	SymbolFlagsEnumMember = ast.SymbolFlagsEnumMember
	SymbolFlagsFunction   = ast.SymbolFlagsFunction
	SymbolFlagsInterface  = ast.SymbolFlagsInterface
	SymbolFlagsModule     = ast.SymbolFlagsModule
	SymbolFlagsOptional   = ast.SymbolFlagsOptional
	SymbolFlagsTypeAlias  = ast.SymbolFlagsTypeAlias
)
