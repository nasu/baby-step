## Usage

```
gcloud spanner databases ddl update [DATABASE] --instance=[INSTANCE] --project=[PROJECT] --ddl="`cat sample.ddl.sql`"
for sql in sample.singers.sql sample.albums.sql; do gcloud spanner databases execute-sql [DATABASE] --instance=[INSTANCE] --project=[PROJECT] --sql="`cat $sql`"; done
```
