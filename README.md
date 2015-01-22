# memberid

`memberid` is a tool for helping team members management.

## Installation

You can download a binary from the [releases](https://github.com/navy/memberid/releases) and put in your `$PATH`.

## Usage

```bash
Usage: memberid <COMMAND> [-c <CONFIG>] ...

random: [-g <GROUP>] [-to <TO>]
  -c="memberid.json": Path to memberid.json file
  -g="": group name
  -to="": id-type to

resolve: [-from <FROM>] [-to <TO>] <ID>
  -c="memberid.json": Path to memberid.json file
  -from="": id-type from
  -to="": id-type to

list: [-g <GROUP>] [-to <TO>]
  -c="memberid.json": Path to memberid.json file
  -g="": group name
  -shuffle=false: Shuffle ids
  -to="": id-type to
```

## Config file

By default, this tool looks for a `memberid.json` file on CWD. If you specify your own named config file, you can use the `-c` option.

```json
{
    "members": {
        "<base_id>": {"<service1>": "<another_id1>", "<service2>": "<another_id2>"}
    },
    "group": {
        "<group>": ["id1", "id2"]
    }
}
```

### Sample

`sample.json`:

```json
{
    "members": {
        "joseph": {"facebook": "joseph.joestar"},
        "jotaro": {"github": "jotarok", "facebook": "jotaro.kujo"},
        "dio": {}
    },
    "group": {
        "part1": ["dio"],
        "part2": ["joseph"],
        "part3": ["joseph", "jotaro", "dio"]
    }
}
```

## Commands

### resolve command

You can convert to the github ID with the `-to` option:

```bash
% memberid resolve -c sample.json -to github jotaro
jotarok
```

You can also convert from the github ID with the `-from` option:

```bash
% memberid resolve -c sample.json -from github -to facebook jotarok
jotaro.kujo
```

If you specify undefined from/to value, use the base ID instead:

```bash
% memberid resolve -c sample.json -from github -to google jotarok
jotaro
```

### random command

You can get an ID randomly:

```bash
% memberid random -c sample.json
jotaro
```

You can use the `-to` option to convert ID:

```bash
% memberid random -c sample.json -to facebook
jotaro.kujo
```

You can specify the group with the `-g` option:

```bash
% memberid random -c sample.json -g=part1
dio
```

You can also specify multi groups using `,`:

```bash
% memberid random -c sample.json -g=part1,part2
joseph
```

### list command

You can get the space-separated IDs list:

```bash
% memberid list -c sample.json
joseph jotaro dio
```

You can shuffle the list with `-shuffle` option:

```bash
% memberid list -c sample.json -shuffle
jotaro dio joseph
```

You can specify the group with the `-g` option:

```bash
% memberid list -c sample.json -g part1,part2
dio joseph
```

