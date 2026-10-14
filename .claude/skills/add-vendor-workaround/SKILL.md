---
name: add-vendor-workaround
description: Make gofish accept a Redfish service's non-conformant JSON, such as a number sent as a string, an object where the schema has an array, a misspelled property, or a link in the wrong place. Use when a user reports a decode error ("cannot unmarshal string into Go struct field"), a field that comes back empty for one vendor, or asks to add a workaround or quirk handling for a specific BMC, vendor or firmware.
---

# Add a vendor workaround

The rules are in `AGENTS.md` (custom JSON decoding, required tests). This is
the procedure.

## 1. Get the payload

You need the JSON the service actually returns for the failing resource.
If the user has not provided it, ask for it, along with the vendor, model and
firmware version. Trim it to `@odata.id`, `@odata.type`, `Id`, and the
properties involved; keep their values exactly as sent.

## 2. Find the type and field

- `@odata.type` `#Processor.v1_20_0.Processor` maps to `schemas/processor.go`.
  For a nested object, find its struct with `grep -n "^type <Name> struct" schemas/*.go`.
- Note the Go type name and Go field name (`MaxSpeedMHz`, `MemberID`); config
  keys use Go names, not JSON names.
- Reproduce the failure first: decode the payload into the type in a test and
  confirm the error or wrong value.

## 3. Pick the mechanism

| Symptom | Mechanism |
|---|---|
| One field has the wrong JSON type | `LenientFields` entry with a converter |
| Value is present but under another name or location (misspelling, `Links` vs top level) | `postUnmarshal` hook |
| Need a value the generated code ignores (`@Redfish.AllowableValues`, `@odata.count`) | `ExtraFields` entry plus a `postUnmarshal` hook |
| The whole value has the wrong shape (`true` instead of an object) and the type has no links | `UnmarshalJSON` replacement |

Check the type's `schemas/<schema>_custom.go` first; it may already have a
hook or converter to extend.

### LenientFields

Add the entry to `LenientFields` in `tools/generator/internal/config/config.go`,
keeping the map sorted:

```go
"Processor.MaxSpeedMHz": "toInt",
```

Existing converters in `schemas/utils.go`: `toInt` (`*int`), `toFloat32`
(`*float32`), `toFloat64` (`*float64`), `parseString` (`string`). The converter
must return the field's exact type. For anything else, write a
`func(any) T` in the type's `_custom.go` file with a comment naming the vendor
behavior it handles; `powerControls` in `schemas/power_custom.go` is an
example. A converter cannot return an error, so return the zero value or nil
for input it cannot convert.

### postUnmarshal hook

In `schemas/<schema>_custom.go`. If the type already has a hook, add to it.

```go
// postUnmarshal accepts SupportedCollectionFuntions, a misspelling some
// services send instead of SupportedCollectionFunctions.
func (t *TelemetryService) postUnmarshal(b []byte) error {
	var v struct {
		SupportedCollectionFuntions []CollectionFunction
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if len(t.SupportedCollectionFunctions) == 0 && len(v.SupportedCollectionFuntions) > 0 {
		t.SupportedCollectionFunctions = v.SupportedCollectionFuntions
	}
	return nil
}
```

Decode only what the hook needs into a local struct. Do not overwrite a value
the standard decode already set unless the workaround requires it.

### UnmarshalJSON replacement

Only for a link-free type whose payload is the wrong shape. Follow
`IPMIHostInterface.UnmarshalJSON` in `schemas/computersystem_custom.go`: try the
standard decode through `type temp T` first, and fall back only when it fails.

## 4. Regenerate

Config changes and new hooks or replacements take effect through the generator:

```bash
make generate
git status --short schemas serviceroot.go   # expect only the type you targeted
```

## 5. Write the test

In `schemas/<schema>_test.go`. The test must decode the vendor payload, assert
the corrected values, and decode a conformant payload for the same fields:

```go
// <Vendor> <model> BMC <firmware> sends MaxSpeedMHz as a string.
func TestProcessorMaxSpeedMHzString(t *testing.T) {
	for body, want := range map[string]int{
		`{"@odata.id": "/redfish/v1/Systems/1/Processors/1", "Id": "1", "MaxSpeedMHz": "3000"}`: 3000,
		`{"@odata.id": "/redfish/v1/Systems/1/Processors/1", "Id": "1", "MaxSpeedMHz": 3000}`:   3000,
	} {
		var result Processor
		if err := json.Unmarshal([]byte(body), &result); err != nil {
			t.Fatalf("Error decoding %s: %s", body, err)
		}
		if result.MaxSpeedMHz == nil || *result.MaxSpeedMHz != want {
			t.Errorf("MaxSpeedMHz from %s = %v, want %d", body, result.MaxSpeedMHz, want)
		}
	}
}
```

If the type is writable (it has an `Update()` method), also check that an
update sends the field in its schema-defined form; `TestMemoryUpdate` in
`schemas/memory_test.go` shows the `TestClient` and `CapturedCalls()` pattern.

## 6. Verify

```bash
make test
make lint
make check-generate   # after committing, or compare git status before and after
```

In the change description, name the vendor, model and firmware, and quote the
offending JSON.
