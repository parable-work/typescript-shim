# typescript-shim

This repository holds one Go module, `github.com/microsoft/typescript-go/shim`.
The module re-exports a small part of the internal Go API of Microsoft's
TypeScript 7 compiler. Parable's TypeScript readers use it: the psgen reader in
parable-platform and the superschematic reader.

The compiler itself is not in this repository. The module imports Microsoft's
compiler without change, from `github.com/microsoft/TypeScript` at the
`v7.0.2` release tag. This repository replaces the shim in the
`parable-work/typescript-go` fork. The fork is no longer necessary.

## Packages

| Package | Re-exports |
|---|---|
| `shim/ast` | Node, SourceFile, Symbol and Diagnostic types. Node kinds, modifier flags, node flags and symbol flags. |
| `shim/bundled` | The embedded `lib.d.ts` files. |
| `shim/checker` | The type checker, Type, type flags and tuple element flags. |
| `shim/compiler` | Program and CompilerHost. |
| `shim/core` | Compiler options and the Tristate values. |
| `shim/tsoptions` | The `tsconfig.json` parser. |
| `shim/vfs` | The virtual file system interface. |
| `shim/vfs/osvfs` | The file system on disk. |

Each package is only aliases. A type is a type alias. A function is a `var`
binding. A constant is a `const` binding to the compiler's own constant. The
module contains no wrapper logic.

A reader must not hardcode a compiler number, such as a node kind or a flag
bit. A compiler release can change these numbers, and the reader then
compiles but reads schemas incorrectly. TypeScript 7.1 adds
`KindSourceKeyword` and moves many Kind numbers by one. When a reader needs a
new kind or flag, add an alias for it here.

## Why Go lets this module import internal packages

Go applies the `internal` rule to import paths, not to repositories. A package
can import `a/b/internal/c` only when its own import path starts with `a/b/`.
This module's path, `github.com/microsoft/typescript-go/shim`, starts with the
compiler's module path, `github.com/microsoft/typescript-go`. So Go lets the
shim import the compiler's `internal/` packages, although the two modules live
in different repositories.

The oxc-project [tsgolint](https://github.com/oxc-project/tsgolint) linter
ships a shim over the TypeScript 7 compiler in the same way.

At `v7.0.2` the compiler's module path is still
`github.com/microsoft/typescript-go`, but Microsoft publishes the code from
the `tsc/` directory of `github.com/microsoft/TypeScript`. The `replace` line
in `go.mod` connects the two:

```
require github.com/microsoft/typescript-go v0.0.0

replace github.com/microsoft/typescript-go => github.com/microsoft/TypeScript/tsc v0.0.0-20260708042240-1e4744d68260
```

Pin only a commit that a TypeScript release tag names. Do not pin an arbitrary
commit from the upstream `main` branch.

## How a consumer uses the module

The module path is under `github.com/microsoft`, so `go get` cannot find this
repository from the path alone. A consumer adds two `replace` lines to its own
`go.mod`. The first line points the shim path at this repository. The second
line points the compiler path at the same compiler commit that this module
pins.

```
require (
	github.com/microsoft/typescript-go v0.0.0
	github.com/microsoft/typescript-go/shim v0.0.0
)

replace github.com/microsoft/typescript-go/shim => github.com/parable-work/typescript-shim <version>

replace github.com/microsoft/typescript-go => github.com/microsoft/TypeScript/tsc v0.0.0-20260708042240-1e4744d68260
```

Go uses `replace` lines only from the main module. So each consumer must
declare both lines, and the compiler line must match this module's `go.mod`.
Set `<version>` to a tag or a commit of this repository, then run
`go mod tidy`.

The consumer imports the packages by their shim paths, for example:

```go
import tsast "github.com/microsoft/typescript-go/shim/ast"
```

## Tests

`shim_test.go` is a smoke test. It parses a `tsconfig.json`, builds a Program,
gets a type checker and walks the AST. It uses the shim packages only. If a
compiler upgrade moves a re-exported API, the test fails.

```
go test ./...
go vet ./...
```

CI runs `go test ./...` on every pull request and on every push to `main`.

## Policy and upgrades

The psgen skill page in parable-platform owns the policy for this module, the
compiler pin and the upgrade procedure. See its
[TypeScript compiler](https://github.com/parable-work/parable-platform/blob/main/docs/project/agents/skills/schema/psgen.mdx#typescript-compiler)
section. parable-platform is a private repository.

## License

Apache License 2.0, the same license as the TypeScript compiler. See
[LICENSE](LICENSE).
