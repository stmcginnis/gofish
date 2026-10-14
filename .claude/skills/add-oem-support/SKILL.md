---
name: add-oem-support
description: Add or extend vendor-specific (Oem) support in gofish, such as exposing Oem.<Vendor> properties on a standard resource, an OEM action under Actions.Oem, a vendor-only resource reached through an Oem link, or a new vendor package under oem/. Use when asked to support Dell, HPE, Supermicro, AMI, ZT or another vendor's Redfish extensions, iDRAC/iLO/BMC-specific features, or properties that appear only in a payload's "Oem" section.
---

# Add OEM support

The rules are in `AGENTS.md` (OEM support, required tests). OEM code never
goes in `schemas/`; it lives in `oem/<vendor>/`, package name `<vendor>`
(`ami`, `dell`, `hpe`, `smc`, `zt`). For a new vendor, create the directory
with the same file header as the existing packages.

## 1. Get the payload

You need the real JSON for the resource, including its `Oem` section and any
`Actions.Oem` entries. Ask the user for it if they have not provided it, and
note the vendor, model and firmware version. Trim unrelated properties but
keep the `Oem` structure and `@odata.type` values exactly as sent.

## 2. Pick the shape

| What the vendor adds | Pattern | Example |
|---|---|---|
| Properties under `Oem.<Vendor>` on a standard resource | wrapper type plus `From<Type>` | `oem/smc/drive.go` |
| An action under `Actions.Oem` | unexported target on the wrapper, method that posts to it | `Drive.Indicate` in `oem/smc/drive.go` |
| A vendor-only resource linked from `Oem.<Vendor>` | its own struct embedding `schemas.Entity`, a `Get<Type>` function, a getter on the wrapper | `NodeManager` in `oem/smc/nodemanager.go`, `ComputerSystem.NodeManager` in `oem/smc/computersystem.go` |

Check the vendor package for an existing wrapper of the same type and extend
it rather than adding a second one.

## 3. Write the wrapper

```go
// Drive extends a schemas.Drive with Supermicro OEM fields.
type Drive struct {
	schemas.Drive

	Temperature int

	indicateTarget string
}

// FromDrive returns an OEM-extended schemas.Drive.
func FromDrive(drive *schemas.Drive) (Drive, error) {
	smcDrive := Drive{Drive: *drive}

	var t struct {
		Oem struct {
			Supermicro struct {
				Temperature int
			} `json:"Supermicro"`
		} `json:"Oem"`
		Actions struct {
			Oem struct {
				DriveIndicate schemas.ActionTarget `json:"#Drive.Indicate"`
			} `json:"Oem"`
		} `json:"Actions"`
	}
	if err := json.Unmarshal(drive.RawData, &t); err != nil {
		return smcDrive, err
	}

	smcDrive.Temperature = t.Oem.Supermicro.Temperature
	smcDrive.indicateTarget = t.Actions.Oem.DriveIndicate.Target
	return smcDrive, nil
}
```

- Read from `RawData` when the standard type has it (writable resources do).
  Otherwise unmarshal the standard type's `OEM` field, which holds the raw
  `Oem` object; see `FromComputerSystem` in `oem/smc/computersystem.go`.
- Copy the client to the wrapper (`SetClient(x.GetClient())`) if it calls the
  service, unless it embeds the original value, which carries the client.
- When an OEM action's target is empty, return an error the caller can check
  and fall back to the standard path, as `Drive.Indicate` does with
  `ErrActionNotSupported` in `oem/smc`. Reuse the package's error if it has one.
- Comment any firmware-specific handling with the vendor behavior it covers,
  as `FromDrive` does for the `#SmcDrive.Indicate` variant.

For a vendor-only resource, embed `schemas.Entity`, add the fields, and fetch
it with `schemas.GetObject[T](client, uri)` in a `Get<Type>` function.

## 4. Test

In `oem/<vendor>/<file>_test.go`, with the payload inline as a raw string and
a comment naming the vendor, model and firmware:

- Decode into the standard type, call `From<Type>`, and assert each OEM field
  and each action target. Tests in the same package can read unexported
  fields such as `indicateTarget`.
- For each action, attach a `schemas.TestClient`, call the method, and check
  `CapturedCalls()` for the method, URL and payload. Test that an empty target
  returns the not-supported error.
- If the vendor sends a non-conformant value in the standard part of the
  payload, fix that in `schemas/` with the `add-vendor-workaround` skill, not
  in the OEM package.

`oem/smc/drive_test.go` and `oem/dell/manager_test.go` show both kinds of test.

## 5. Verify

```bash
make test
make lint
```
