# iriscli asset

Asset module allows you to manage assets on IRIS Hub

## Available Commands

| Name                                            | Description                        |
| ----------------------------------------------- | ---------------------------------- |
| [token issue](#iriscli-asset-token-issue)       | Issue a new token                  |
| [token edit](#iriscli-asset-token-edit)         | Edit an existing token             |
| [token transfer](#iriscli-asset-token-transfer) | Transfer the ownership of a token  |
| [token mint](#iriscli-asset-token-mint)         | Mint tokens to a specified address |
| [token token](#iriscli-asset-token-token)       | Query a token by symbol |
| [token tokens](#iriscli-asset-token-tokens)     | Query tokens by owner |
| [token fee](#iriscli-asset-token-fee)           | Query the token related fees       |

## iriscli asset token issue

Issue a new token

```bash
iriscli asset token issue [flags]
```

**Flags:**

| Name, shorthand  | Type    | Required | Default       | Description                                                                                                                    |
| ---------------- | ------- | -------- | ------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| --name           | string  | Yes      |               | Name of the newly issued token, limited to 32 unicode characters, e.g. "IRIS Network"                                          |
| --symbol         | string  | Yes      |               | The symbol of the token, length between 3 and 8, alphanumeric characters beginning with alpha, case insensitive                                     |
| --initial-supply | uint64  | Yes      |               | The initial supply of this token. The amount before boosting should not exceed 100 billion.                                    |
| --max-supply     | uint64  |          | 1000000000000 | The hard cap of this token, total supply can not exceed max supply. The amount before boosting should not exceed 1000 billion. |
| --min-unit       | string  |          |               | The alias of minimum uint                                                                                                      |
| --scale          | uint8   | Yes      |               | A token can have a maximum of 18 digits of decimal                                                                             |
| --mintable       | boolean |          | false         | Whether this token could be minted(increased) after the initial issuing                                                        |

### Issue a token

```bash
iriscli asset token issue --symbol="kitty" --name="Kitty Token" --initial-supply=100000000000 --max-supply=1000000000000 --scale=0 --mintable=true --fee=1iris --chain-id=irishub --from=<key-name> --commit
```

### Send tokens

You can send any tokens you have just like [sending iris](./bank.md#iriscli-bank-send)

#### Send tokens

```bash
iriscli bank send --from=<key-name> --to=<address> --amount=10kitty --fee=0.3iris --chain-id=irishub --commit
```

## iriscli asset token edit

Edit an existing token

```bash
iriscli asset token edit [symbol] [flags]
```

**Flags:**

| Name         | Type   | Required | Default | Description                                    |
| ------------ | ------ | -------- | ------- | ---------------------------------------------- |
| --name       | string |          |         | The token name, e.g. IRIS Network              |
| --max-supply | uint64   |          | 0       | The max supply of the token                    |
| --mintable   | bool   |          | false   | Whether the token can be minted, default to false |

`max-supply` should not be less than the current total supply

### Edit Token

```bash
iriscli asset token edit kitty --name="Cat Token" --max-supply=100000000000 --mintable=true --from=<key-name> --chain-id=irishub --fee=0.3iris --commit
```

## iriscli asset token transfer

Transfer the ownership of a token

```bash
iriscli asset token transfer [symbol] [flags]
```

**Flags:**

| Name | Type   | Required | Default | Description           |
| ---- | ------ | -------- | ------- | --------------------- |
| --to | string | Yes      |         | The new owner address |

### Transfer Token Owner

```bash
iriscli asset token transfer kitty --to=<new-owner-address> --from=<key-name> --chain-id=irishub --fee=0.3iris --commit
```

## iriscli asset token mint

Mint tokens to a specified address

```bash
iriscli asset token mint [symbol] [flags]
```

**Flags:**

| Name     | Type   | Required | Default | Description                                           |
| -------- | ------ | -------- | ------- | ----------------------------------------------------- |
| --to     | string |          |         | Address to which the token will be minted, default to the owner address |
| --amount | uint64 | Yes      | 0       | Amount of the tokens to be minted                         |

### Mint Token

```bash
iriscli asset token mint kitty --amount=1000000 --from=<key-name> --chain-id=irishub --fee=0.3iris
```

## iriscli asset token token

Query a token by symbol

```bash
iriscli asset token token [symbol] [flags]
```

### Query a token

```bash
iriscli asset token token kitty
```

## iriscli asset token tokens

Query tokens by the owner which is optional

```bash
iriscli asset token tokens [owner] [flags]
```

### Query all tokens

```bash
iriscli asset token tokens
```

### Query tokens with the specified owner

```bash
iriscli asset token tokens <owner>
```

## iriscli asset token fee

Query the token related fees, including token issuance and minting

```bash
iriscli asset token fee [symbol] [flags]
```

### query fees of issuing and minting a token

```bash
iriscli asset token fee kitty
```
