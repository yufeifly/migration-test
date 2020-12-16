# intro
This is the validator of container migration.

# content
- There exists test code to run tests, and validation code to verify the migration code is correct.
- It supports testing migration of multiple redis services. In other words, you can migrate multiple 
redis services at the same time concurrently.  

# usage
## test migration
`validate migration test -p pc -n 2` 
> -p or --platform means run on the server, or my own pc. 
> -n or --number means the number of services to be migrated concurrently.

## verify migration
`validate migration verify --addr 127.0.0.1:8888 --range key0:key9999`
> --addr means the target addr to get kv pairs to validate.
>-- range means the range of keys