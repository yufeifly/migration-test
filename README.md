# intro
This is the validator of container migration project.
It consists of three parts: tester, validator and breeder.

# features
- It supports testing migration of multiple redis services. In other words, you can migrate multiple 
redis services at the same time concurrently.  

# usage
## test migration
`validate migration test -p pc -n 2` 
> This command starts the migration process.
>-p or --platform means run on the server, or my own pc. 
> -n or --number means the number of services to be migrated concurrently.

## verify migration
`validate migration verify --addr 127.0.0.1:8888 --range key0:key9999`
> This command verifies that the migration is right or not. 
> --addr means the target addr to get kv pairs to validate.
> --range means the range of keys.

## feed redis server on random kv pairs
`validator redis breed --redis-server 127.0.0.1:39954 --range key{0:9999}`
> This command feeds the redis service some random data.
> --redis-server means the address of service.
> --range means the range of the key.