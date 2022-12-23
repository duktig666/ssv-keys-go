# ssv-keys-go
**ETH SSV Key's Golang implementation**



Library and CLI to work with the ETH keystore file:

1. Parse the private key using the keystore password
2. Use the private key to get shares for operators
3. Build the payload for the transaction



For SSV, please see：[doc.ssv.network](https://docs.ssv.network/learn/introduction)

For the TypeScript version, please refer to the official：[ssv-keys](https://github.com/bloxapp/ssv-keys)

If you want to pledge ETH by SSV on ChainupCloud, please see：[ChainUpCloud Ethereum2](https://cloud.chainup.com/app/eth2.0)



## Run

Generate the CLI command tool: ssv-key

```shell
./cli-init.sh
```



`ssv-key --help`

```shell
% ./ssv-key --help

https://github.com/duktig666/ssv-keys-go

Usage:
   [flags]
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  shares      ssv keystore shares
  version     Get version info

Flags:
  -c, --config string   config file (default is config/config.yaml) (default "config/config.yaml")
  -h, --help            help for this command

Use " [command] --help" for more information about a command.
```



Execute the command of the chip result of the keystore

```shell
./ssv-key shares --keystore <YOUR KEYSTORE-m_XXX.json> --password <YOUR PASSWORD>  --output ./temp/output/shares.json
```

> ./ssv-key shares --keystore ./temp/input/keystore-m_12381_3600_0_0_0-1657004059.json --password xxx  --output ./temp/output/shares2.json



Enter four Operatorid and OperatorKey in order

```shell
set operators and operator-ids. count:4
……
```



## TODO

- [ ] Optimize the parameters of CLI
- [ ] Input and output file and folder optimization



## Authors

- [duktig666](https://github.com/duktig666)



## Support

- [ChainUp](https://www.chainup.com/)



## License

Apache-2.0 license