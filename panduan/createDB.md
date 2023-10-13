<!-- create title -->

# Database

```
create database go_bio
```

<!-- create scope code -->
<!-- create commentar -->

## Table users (sudah auto migrate)

```sql
create table users (
    email varchar(100) not null,
    nama VARCHAR(50) NOT NULL,
    no_hp VARCHAR(15) DEFAULT NULL,
    alamat TEXT DEFAULT NULL,
    ktp VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (email) USING BTREE
)
```
