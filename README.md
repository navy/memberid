# idconv

idconv is a user ID converter.

## Usage

```
Usage: iconv -c <CONFIG> [OPTION] <ID>

  -c="": config file path (required)
  -from="": from type
  -to="": to type
```

### Config file format

```json
{
    "<default_id>": {"<service1>": "<another_id1>", "<service2>": "<another_id2>"}
}
```

### Example

`sample.json`:

```json
{
    "jojo": {"github": "jotarok", "facebook": "jotaro.kujo"}
}
```

You can convert from base ID to the github ID with the `-to` option:

```bash
% idconv -c sample.json -to github jojo
jotarok
```

You can also convert from the github ID with the `-from` option:

```bash
% idconv -c sample.json -from github -to facebook jotarok
jotaro.kujo
```

If you specify undefined from/to value, use base ID instead:

```bash
% idconv -c sample.json -from github -to google jotarok
jojo
```
