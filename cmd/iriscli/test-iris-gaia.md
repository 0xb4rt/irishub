# IBC Test

**Use local `cosmos-sdk` for test**

```bash
# go.mod
replace github.com/cosmos/cosmos-sdk => /path-to-your-local/cosmos-sdk
```

**Environment setup**

```bash
cd ~ && mkdir ibc-testnets && cd ibc-testnets
iris testnet -o ibc-iris --v 1 --chain-id chain-iris --node-dir-prefix n
gaiad testnet -o ibc-gaia --v 1 --chain-id chain-gaia --node-dir-prefix n
```

**Set configuration**

```bash
sed -i '' 's/"leveldb"/"goleveldb"/g' ibc-iris/n0/iris/config/config.toml
sed -i '' 's/"leveldb"/"goleveldb"/g' ibc-gaia/n0/gaiad/config/config.toml
sed -i '' 's#"tcp://0.0.0.0:26656"#"tcp://0.0.0.0:26556"#g' ibc-gaia/n0/gaiad/config/config.toml
sed -i '' 's#"tcp://0.0.0.0:26657"#"tcp://0.0.0.0:26557"#g' ibc-gaia/n0/gaiad/config/config.toml
sed -i '' 's#"localhost:6060"#"localhost:6061"#g' ibc-gaia/n0/gaiad/config/config.toml
sed -i '' 's#"tcp://127.0.0.1:26658"#"tcp://127.0.0.1:26558"#g' ibc-gaia/n0/gaiad/config/config.toml

iriscli config --home ibc-iris/n0/iriscli/ chain-id chain-iris
gaiacli config --home ibc-gaia/n0/gaiacli/ chain-id chain-gaia
iriscli config --home ibc-iris/n0/iriscli/ output json
gaiacli config --home ibc-gaia/n0/gaiacli/ output json
iriscli config --home ibc-iris/n0/iriscli/ node http://localhost:26657
gaiacli config --home ibc-gaia/n0/gaiacli/ node http://localhost:26557
```

**View Keys**

```bash
iriscli --home ibc-iris/n0/iriscli keys list | jq '.[].address'
gaiacli --home ibc-gaia/n0/gaiacli keys list | jq '.[].address'
```

**Start**

```bash
# run in background
nohup iris --home ibc-iris/n0/iris start >ibc-iris.log &
nohup gaiad --home ibc-gaia/n0/gaiad start >ibc-gaia.log &

# run in terminal
iris --home ibc-iris/n0/iris start
gaiad --home ibc-gaia/n0/gaiad start
```

**Create client**

create client on chain-iris

```bash
# view consensus-state of chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client self-consensus-state -o json | jq
# export consensus_state.json from chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client self-consensus-state -o json >ibc-iris/n0/consensus_state.json
# create client on chain-iris
iriscli --home ibc-iris/n0/iriscli tx ibc client create client-to-gaia ibc-iris/n0/consensus_state.json \
  --from n0 -y -o text --broadcast-mode=block
```

create client on chain-gaia

```bash
# view consensus-state of chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client self-consensus-state -o json | jq
# export consensus_state.json from chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client self-consensus-state -o json >ibc-gaia/n0/consensus_state.json
# create client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client create client-to-iris ibc-gaia/n0/consensus_state.json \
  --from n0 -y -o text --broadcast-mode=block
```

**Query client**

query client state

```bash
# query client state on chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client state client-to-gaia | jq
# query client state on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client state client-to-iris | jq
```

query client consensus-state

```bash
# query client consensus-state on chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client consensus-state client-to-gaia | jq
# query client consensus-state on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client consensus-state client-to-iris | jq
```

query client path

```bash
# query client path of chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client path | jq
# query client path of chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client path | jq
```

**Update client**

update chain-iris

```bash
# query header of chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client header -o json | jq
# export header of chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client header -o json >ibc-iris/n0/header.json
# update client on chain-iris
iriscli --home ibc-iris/n0/iriscli tx ibc client update client-to-gaia ibc-iris/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
```

update chain-gaia

```bash
# query header of chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json | jq
# export header of chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
# update client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
```

**Create connection**

`open-init` on chain-iris

```bash
# export prefix.json
gaiacli --home ibc-gaia/n0/gaiacli q ibc client path -o json >ibc-iris/n0/prefix.json
# view prefix.json
jq -r '' ibc-iris/n0/prefix.json
# open-init
iriscli --home ibc-iris/n0/iriscli tx ibc connection open-init \
  conn-to-gaia client-to-gaia \
  conn-to-iris client-to-iris \
  ibc-iris/n0/prefix.json \
  --from n0 -y -o text \
  --broadcast-mode=block
```

`open-try` on chain-gaia

```bash
# export prefix.json
iriscli --home ibc-iris/n0/iriscli q ibc client path -o json >ibc-gaia/n0/prefix.json
# export header.json from chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
# export proof_init.json from chain-iris with hight in header.json
iriscli --home ibc-iris/n0/iriscli q ibc connection proof conn-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  -o json >ibc-gaia/n0/conn_proof_init.json
# view proof_init.json
jq -r '' ibc-gaia/n0/conn_proof_init.json
# update client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# query client consense state
gaiacli --home ibc-gaia/n0/gaiacli q ibc client consensus-state client-to-iris | jq
# open-try
gaiacli --home ibc-gaia/n0/gaiacli tx ibc connection open-try \
  conn-to-iris client-to-iris \
  conn-to-gaia client-to-gaia \
  ibc-gaia/n0/prefix.json \
  1.0.0 \
  ibc-gaia/n0/conn_proof_init.json \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  --from n0 -y -o text \
  --broadcast-mode=block
```

`open-ack` on chain-iris

```bash
# export header.json from chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client header -o json >ibc-iris/n0/header.json
# export proof_try.json from chain-gaia with hight in header.json
gaiacli --home ibc-gaia/n0/gaiacli q ibc connection proof conn-to-iris \
  $(jq -r '.value.SignedHeader.header.height' ibc-iris/n0/header.json) \
  -o json >ibc-iris/n0/conn_proof_try.json
# view proof_try.json
jq -r '' ibc-iris/n0/conn_proof_try.json
# update client on chain-iris
iriscli --home ibc-iris/n0/iriscli tx ibc client update client-to-gaia ibc-iris/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# query client consense state
iriscli --home ibc-iris/n0/iriscli q ibc client consensus-state client-to-gaia | jq
# open-ack
iriscli --home ibc-iris/n0/iriscli tx ibc connection open-ack \
  conn-to-gaia \
  ibc-iris/n0/conn_proof_try.json \
  $(jq -r '.value.SignedHeader.header.height' ibc-iris/n0/header.json) \
  1.0.0 \
  --from n0 -y -o text \
  --broadcast-mode=block
```

`open-confirm` on chain-gaia

```bash
# export header.json from chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
# export proof_ack.json from chain-iris with hight in header.json
iriscli --home ibc-iris/n0/iriscli q ibc connection proof conn-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  -o json >ibc-gaia/n0/conn_proof_ack.json
# view proof_ack.json
jq -r '' ibc-gaia/n0/conn_proof_ack.json
# update client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# open-confirm
gaiacli --home ibc-gaia/n0/gaiacli tx ibc connection open-confirm \
  conn-to-iris \
  ibc-gaia/n0/conn_proof_ack.json \
  --from n0 -y -o text \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  --broadcast-mode=block
```

**Query connection**

query connection

```bash
# query connection on chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc connection end conn-to-gaia | jq
# query connection on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc connection end conn-to-iris | jq
```

query connection proof

```bash
# query connection proof with height in header.json on chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc connection proof conn-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-iris/n0/header.json) | jq
# query connection proof with height in header.json on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc connection proof conn-to-iris \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) | jq
```

query connections of a client

```bash
# query connections of a client on chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc connection client client-to-gaia | jq
# query connections of a client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc connection client client-to-iris | jq
```

**Create channel**

`open-init` on chain-iris

```bash
# open-init
iriscli --home ibc-iris/n0/iriscli tx ibc channel open-init \
  port-to-gaia chann-to-gaia \
  port-to-iris chann-to-iris \
  conn-to-gaia \
  --from n0 -y -o text \
  --broadcast-mode=block
```

`open-try` on chain-gaia

```bash
# export header.json from chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
# export proof_init.json from chain-iris with hight in header.json
iriscli --home ibc-iris/n0/iriscli q ibc channel proof port-to-gaia chann-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  -o json >ibc-gaia/n0/chann_proof_init.json
# view proof_init.json
jq -r '' ibc-gaia/n0/chann_proof_init.json
# update client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# query client consense state
gaiacli --home ibc-gaia/n0/gaiacli q ibc client consensus-state client-to-iris | jq
# open-try
gaiacli --home ibc-gaia/n0/gaiacli tx ibc channel open-try \
  port-to-iris chann-to-iris \
  port-to-gaia chann-to-gaia \
  conn-to-iris \
  ibc-gaia/n0/chann_proof_init.json \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  --from n0 -y -o text \
  --broadcast-mode=block
```

`open-ack` on chain-iris

```bash
# export header.json from chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc client header -o json >ibc-iris/n0/header.json
# export proof_try.json from chain-gaia with hight in header.json
gaiacli --home ibc-gaia/n0/gaiacli q ibc channel proof port-to-iris chann-to-iris \
  $(jq -r '.value.SignedHeader.header.height' ibc-iris/n0/header.json) \
  -o json >ibc-iris/n0/chann_proof_try.json
# view proof_try.json
jq -r '' ibc-iris/n0/chann_proof_try.json
# update client on chain-iris
iriscli --home ibc-iris/n0/iriscli tx ibc client update client-to-gaia ibc-iris/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# query client consense state
iriscli --home ibc-iris/n0/iriscli q ibc client consensus-state client-to-gaia | jq
# open-ack
iriscli --home ibc-iris/n0/iriscli tx ibc channel open-ack \
  port-to-gaia chann-to-gaia \
  ibc-iris/n0/chann_proof_try.json \
  $(jq -r '.value.SignedHeader.header.height' ibc-iris/n0/header.json) \
  --from n0 -y -o text \
  --broadcast-mode=block
```

`open-confirm` on chain-gaia

```bash
# export header.json from chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
# export proof_ack.json from chain-iris with hight in header.json
iriscli --home ibc-iris/n0/iriscli q ibc channel proof port-to-gaia chann-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  -o json >ibc-gaia/n0/chann_proof_ack.json
# view proof_ack.json
jq -r '' ibc-gaia/n0/chann_proof_ack.json
# update client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# query client consense state
gaiacli --home ibc-gaia/n0/gaiacli q ibc client consensus-state client-to-iris | jq
# open-confirm
gaiacli --home ibc-gaia/n0/gaiacli tx ibc channel open-confirm \
  port-to-iris chann-to-iris \
  ibc-gaia/n0/chann_proof_ack.json \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  --from n0 -y -o text \
  --broadcast-mode=block
```

**Query channel**

query channel

```bash
# query channel on chain-iris
iriscli --home ibc-iris/n0/iriscli query ibc channel end port-to-gaia chann-to-gaia | jq
# query channel on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli query ibc channel end port-to-iris chann-to-iris | jq
```

query channel proof

```bash
# query channel proof with height in header.json on chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc channel proof port-to-gaia chann-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-iris/n0/header.json) | jq
# query channel proof with height in header.json on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli q ibc channel proof port-to-iris chann-to-iris \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) | jq
```

**Bank transfer from chain-iris to chain-gaia**

```bash
# export transfer result to result.json
iriscli --home ibc-iris/n0/iriscli tx ibcmockbank transfer \
  --src-port port-to-gaia --src-channel chann-to-gaia \
  --denom iris-atto --amount 1 \
  --receiver $(gaiacli --home ibc-gaia/n0/gaiacli keys show n0 | jq -r '.address') \
  --source true \
  --from n0 -y -o json > ibc-iris/n0/result.json
# export packet.json
jq -r '.events[1].attributes[2].value' ibc-iris/n0/result.json >ibc-gaia/n0/packet.json
```

**Bank receive**

> use "tx ibc channel send-packet" instead when router-module completed in the future

```bash
# export header.json from chain-iris
iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
# export proof.json from chain-gaia with hight in header.json
iriscli --home ibc-iris/n0/iriscli q ibc channel proof port-to-gaia chann-to-gaia \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  -o json >ibc-gaia/n0/proof.json
# view proof.json
jq -r '' ibc-gaia/n0/proof.json
# update client on chain-gaia
gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json \
  --from n0 -y -o text --broadcast-mode=block
# receive packet
gaiacli --home ibc-gaia/n0/gaiacli tx ibcmockbank recv-packet \
  ibc-gaia/n0/packet.json ibc-gaia/n0/proof.json \
  $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
  --from n0 -y -o text \
  --broadcast-mode=block
```

**Query Account**

```bash
# view sender account
iriscli --home ibc-iris/n0/iriscli q account -o text \
  $(iriscli --home ibc-iris/n0/iriscli keys show n0 | jq -r '.address')
# view receiver account
gaiacli --home ibc-gaia/n0/gaiacli q account -o text \
  $(gaiacli --home ibc-gaia/n0/gaiacli keys show n0 | jq -r '.address')
```
