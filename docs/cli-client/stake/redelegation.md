# iriscli stake redelegation

## Description

Query a redelegation record based on delegator and source validator address and destination validator address

## Usage

```
iriscli stake redelegation --address-validator-source=<source-validator-address> --address-validator-dest=<destination-validator-address> --address-delegator=<address-delegator> <flags>
```

Print help messages:
```
iriscli stake redelegation --help
```

## Unique Flags

| Name, shorthand            | Default | Description                               | Required |
| -------------------------- | ------- | ----------------------------------------- | -------- | 
| --address-delegator        |         | Bech address of the delegator             | true     |
| --address-validator-dest   |         | Bech address of the destination validator | true     |
| --address-validator-source |         | Bech address of the source validator      | true     |

## Examples

Query a redelegation record
```
iriscli stake redelegation --address-validator-source=SourceValidatorAddress --address-validator-dest=DestinationValidatorAddress --address-delegator=DelegatorAddress
```

After that, you will get specified redelegation's info based on delegator and source validator address and destination validator address

```bash
Redelegation
Delegator: iaa10s0arq9khpl0cfzng3qgxcxq0ny6hmc9gtd2ft
Source Validator: iva1dayujdfnxjggd5ydlvvgkerp2supknth9a8qhh
Destination Validator: iva1h27xdw6t9l5jgvun76qdu45kgrx9lqedpg3ecs
Creation height: 1130
Min time to unbond (unix): 2018-11-16 07:22:48.740311064 +0000 UTC
Source shares: 0.1000000000
Destination shares: 0.1000000000
```
