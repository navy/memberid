# idconv

`idconv` is a user ID converter.

## Installation

You can download a binary from the [releases](https://github.com/navy/idconv/releases) and put in your `$PATH`.

## Usage

```
Usage: idconv -c <CONFIG> [OPTION] <ID>

  -c="": config file path (required)
  -from="": from type
  -to="": to type
```

### Config file format

```json
{
    "<base_id>": {"<service1>": "<another_id1>", "<service2>": "<another_id2>"}
}
```

### Example

`sample.json`:

```json
{
    "jojo": {"github": "jotarok", "facebook": "jotaro.kujo"}
}
```

You can convert to the github ID with the `-to` option:

```bash
% idconv -c sample.json -to github jojo
jotarok
```

You can also convert from the github ID with the `-from` option:

```bash
% idconv -c sample.json -from github -to facebook jotarok
jotaro.kujo
```

If you specify undefined from/to value, use the base ID instead:

```bash
% idconv -c sample.json -from github -to google jotarok
jojo
```
