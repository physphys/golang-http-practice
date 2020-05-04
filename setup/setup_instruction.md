```bash
$ createuser todo_owner
$ createdb todo -O todo_owner
$ psql -h localhost -U todo_owner -d todo -f ./setup/base_ddl.sql
$ psql -h localhost -U todo_owner -d todo -f ./setup/dummy_data.sql
```
