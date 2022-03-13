## Steps
### Setup mongo
```bash
docker run --name mongo -p 27017:27017 -d mongo

docker run -it --link mongo:mongo --rm mongo sh \
  -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/store"'

db.transactions.insert([
{
    "ccnum" : "4444333322221111",
    "date" : "2019-01-05",
    "amount" : 100.12,
    "cvv" : "1234",
    "exp" : "09/2020"
},
{
    "ccnum" : "4444123456789012",
    "date" : "2019-01-07",
    "amount" : 2400.18,
    "cvv" : "5544",
    "exp" : "02/2021"
},
{
    "ccnum" : "4465122334455667",
    "date" : "2019-01-29",
    "amount" : 1450.87,
    "cvv" : "9876",
    "exp" : "06/2020"
}
]);
```

### Setup MySQL
```bash
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql

docker run -it --link mysql:mysql --rm mysql sh -c \
'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" \
-uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'

create database store;
use store;
create table transactions(ccnum varchar(32), date date, amount float(7,2),cvv char(4), exp date);
insert into transactions(ccnum, date, amount, cvv, exp) values ('4444333322221111', '2019-01-05', 100.12, '1234', '2020-09-01');
insert into transactions(ccnum, date, amount, cvv, exp) values ('4444123456789012', '2019-01-07', 2400.18, '5544', '2021-02-01');
insert into transactions(ccnum, date, amount, cvv, exp) values ('4465122334455667', '2019-01-29', 1450.87, '9876', '2019-06-01');
```

### Setup Postgresql
```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres

docker run -it --rm --link postgres:postgres postgres psql -h postgres -U postgres

create database store;
\connect store
create table transactions(ccnum varchar(32), date date, amount money, cvv char(4), exp date);
insert into transactions(ccnum, date, amount, cvv, exp) values ('4444333322221111', '2019-01-05', 100.12, '1234', '2020-09-01');
insert into transactions(ccnum, date, amount, cvv, exp) values ('4444123456789012', '2019-01-07', 2400.18, '5544', '2021-02-01');
insert into transactions(ccnum, date, amount, cvv, exp) values ('4465122334455667', '2019-01-29', 1450.87, '9876', '2019-06-01');
```
