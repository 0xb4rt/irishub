# iriscli distribution withdraw-rewards

## Description

Withdraw rewards

## Usage

```
iriscli distribution withdraw-rewards <flags>
```

Print help messages:

```
iriscli distribution withdraw-rewards --help
```

## Unique Flags

| Name, shorthand       | type   | Required | Default  | Description                                                         |
| --------------------- | -----  | -------- | -------- | ------------------------------------------------------------------- |
| --only-from-validator | string |          |          | Only withdraw from this validator address (in bech) |
| --is-validator        | bool   |          | false    | Also withdraw validator's commission |

Keep in mind, don't specify the above options both.

## Examples

1. Only withdraw a delegation rewards from a given validator
    ```
    iriscli distribution withdraw-rewards --only-from-validator=<validator_address> --from=<key_name> --fee=0.3iris --chain-id=<chain-id>
    ```
2. Withdraw all delegation rewards of a delegator
    ```
    iriscli distribution withdraw-rewards --from=<key_name> --fee=0.3iris --chain-id=<chain-id>
    ```
3. If the delegator is a owner of a validator, withdraw all delegation rewards and validator commission rewards:
    ```
    iriscli distribution withdraw-rewards --is-validator=true --from=<key_name> --fee=0.3iris --chain-id=<chain-id>
    ```