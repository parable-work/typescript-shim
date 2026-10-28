// Package shim and its subpackages re-export the subset of the TypeScript 7
// compiler's internal API that Parable's TypeScript readers need (psgen and
// superschematic). Each subpackage is a pure veneer over the corresponding
// internal package: type aliases for types and var/const bindings for
// functions and constants. No wrapper logic lives here.
//
// The surface is intentionally minimal and is expanded only as consumer code
// demands. See the repository README for the policy governing these packages
// and the compiler upgrade procedure.
package shim
