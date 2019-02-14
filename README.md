# HBASE-GO-APP

## create table

```sh
docker exec -ti hbase sh -c "./hbase/bin/hbase shell"
```

### Example Usage

```sh
create 'table1', 'columnfamily1'
put 'table1', 'row1', 'columnfamily1:column1', 'value1'
get 'table1', 'row1'
```