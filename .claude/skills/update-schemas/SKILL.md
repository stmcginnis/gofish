---
name: update-schemas
description: Move gofish to a newer DMTF Redfish or SNIA Swordfish schema bundle and regenerate schemas/. Use when asked to update to a new Redfish release (for example 2026.3), pick up new schema properties, resources or actions from DMTF or SNIA, bump REDFISH_TAG or SWORDFISH_TAG, or regenerate the generated code.
---

# Update the schema bundles

`tools/generator/README.md` describes the generator and its config. This is
the procedure for a bundle bump. It needs network access to GitHub.

## 1. Find the release tags

```bash
git ls-remote --tags --refs https://github.com/DMTF/Redfish-Publications.git | awk -F/ '{print $3}' | sort -V | tail -3
git ls-remote --tags --refs https://github.com/SNIA/Swordfish-Publications.git | awk -F/ '{print $3}' | grep _Release | sort -V | tail -3
```

Use Swordfish `_Release` tags only, not `_WorkingDraft`.

## 2. Save the current API and regenerate

```bash
go run golang.org/x/exp/cmd/apidiff@latest -w /tmp/gofish-schemas.api ./schemas
go run golang.org/x/exp/cmd/apidiff@latest -w /tmp/gofish-root.api .
```

Set `REDFISH_TAG` and `SWORDFISH_TAG` in the `Makefile`, then:

```bash
make generate 2>&1 | tee /tmp/generate.log
```

## 3. Resolve generator warnings

Each `WARNING:` line is a definition two schemas now define differently,
which would be merged into one Go type. For each:

- If the copies are the same concept and the owner's version is a superset,
  add a `TypeOwners` entry naming the owner; enum values are merged
  automatically.
- If they are different concepts, add a `TypeRenames` entry
  (`"Schema.Definition": "SchemaDefinition"`) for the newer or less-used one.

Regenerate until the run prints no warnings.

## 4. Build and test

```bash
go build ./... && go vet ./... && make test
```

A compile error after a bump usually means an upstream change collides with
generator config or a hand-written declaration: a renamed or removed property
that a `FieldOverrides`, `LenientFields` or `ExtraFields` entry names, a new
definition whose name matches a hand-written type, or a hook that references
a field that changed. Fix the config or the `_custom.go` file. Never edit a
generated file.

## 5. Review custom code for changed types

List the schemas that changed and the custom code attached to them:

```bash
for f in $(git diff --name-only -- 'schemas/*.go' | grep -v _custom); do
  c=${f%.go}_custom.go; [ -e "$c" ] && echo "$f -> $c"
done
grep -l "func (.*) UnmarshalJSON" schemas/*_custom.go
```

For each custom file next to a changed schema, check that its hooks and
methods still match the schema. A full `UnmarshalJSON` replacement does not
pick up new links or actions; if the type gained any, convert the replacement
to a `LenientFields` entry or a `postUnmarshal` hook (see `AGENTS.md`).

## 6. Check the exported API

```bash
go run golang.org/x/exp/cmd/apidiff@latest -w /tmp/new-schemas.api ./schemas
go run golang.org/x/exp/cmd/apidiff@latest -w /tmp/new-root.api .
go run golang.org/x/exp/cmd/apidiff@latest -incompatible /tmp/gofish-schemas.api /tmp/new-schemas.api
go run golang.org/x/exp/cmd/apidiff@latest -incompatible /tmp/gofish-root.api /tmp/new-root.api
```

Typical breaks: an action gaining a fourth parameter switches to a
`<Type><Action>Parameters` struct, and a property whose schema type changed
changes the Go field type. For each, either pin the old shape with a
`FieldOverrides` entry or accept the break. Report every accepted break to the
user, with the reason, and list it in the change description.

## 7. Finish

```bash
make lint
make check-generate
```

Commit the `Makefile` bump, config changes, custom-file changes and the
regenerated files together. The change description names the old and new tags
and lists any API breaks.
