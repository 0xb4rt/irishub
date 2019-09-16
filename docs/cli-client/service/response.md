# iriscli service response 

## Description

Query a service response

## Usage

```
iriscli service response <flags>
```

## Flags

| Name, shorthand       | Default | Description                                                    | Required |
| --------------------- | ------- | -------------------------------------------------------------- | -------- |
| --request-chain-id    |         | The ID of the blockchain that the service invocation initiated | true     |
| --request-id          |         | The ID of the service invocation                               | true     |

## Examples

### Query a service response

```shell
iriscli service response --request-chain-id=<request_chain_id> --request-id=<request-id>
```
> You can figure out the `request-id` in the return of [service call](call.md)

After that, you will get the response by specified parameters.

```json
{
  "type": "iris-hub/service/SvcResponse",
  "value": {
    "req_chain_id": "test",
    "request_height": "535",
    "request_intra_tx_counter": 0,
    "expiration_height": "635",
    "provider": "iaa1f02ext9duk7h3rx9zm7av0pnlegxve8npm2k6m",
    "consumer": "iaa1f02ext9duk7h3rx9zm7av0pnlegxve8npm2k6m",
    "output": "q80=",
    "error_msg": null
  }
}
```

