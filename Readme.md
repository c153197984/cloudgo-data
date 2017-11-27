# Cloudgo-Data

A practice of operating SQL database through go.

## Usage

Run the server by

    $ go run main.go

Try the following

To post a new user

    $ curl -v -d "username=Alice&department=sdvx" localhost:8080/service/user

To query a user by ID

    $ curl -v localhost:8080/service/user?uid=1

To query all users

    $ curl -v localhost:8080/service/user?uid=

## Implementation

### By `database/sql`

The main idea is to copy the idea of `jdbc`(Java Database Connectivity).

For the entity `User`, we have three modules, which are `user_entity.go`, `user_dao.go` and `user_service.go` in `entities/`.

`user_entity` provides the definition of a user.

`user_dao` encapsulate the operations between entity and database such as insertion and query. The functions cannot be directly used.

`user_service` exports the services that are used to work with users in the database. It takes the functions in `user_dao` and make the operations atomic if necessary.

To use the connectivity, just call the services provided by `user_service`.

### By `xorm`

`xorm` is an ORM framework for go, and it greatly supports various databases. You can click [here] to view more details.

Due to the framework, there is no need to create any modules. The only things we need to do is to call the functions of the engine.

For example, to insert a new user, do it as below

    _, err := mySQLEngine.Table("users").Insert(user)
    if err != nil {
        panic(err)
    }

ORM doesn't just implement an automatic DAO. DAO is just an encapsulation of one or multiple SQL statements to a certain data. ORM is more an encapsulation of general SQL statements themselves. It makes us easily do SQL operations to any data we want to work with.

### Comparison

You simply can toggle comments in `main.go` to switch between `database/sql` and `xorm`. The key point is to compare the two ideas.

1. Development efficiency

    `xorm` >>> `database/sql`

    There is no doubt that by applying `xorm`, it is much easier to operate the database. The functions we implemented by five .go files are simply solved by a single line through `xorm`.

2. Program structure

    `xorm` ? `database/sql`

    One thing is clear that the idea of `jdbc` makes the program structure extremely tidy, but it is hard to compare it with `xorm`, because we don't know the details of structure in `xorm`.

3. Performance

    `xorm` <<< `database/sql`

    The more we abstract, the worse the performance will be. Afterall, we can do an experiment.

    Do the ab test as below to both implementations.

        $ ab -n 10000 -c 100 http://localhost:8080/service/user?uid=

    The result of `database/sql` is

        This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
        Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
        Licensed to The Apache Software Foundation, http://www.apache.org/

        Benchmarking localhost (be patient)
        Completed 1000 requests
        Completed 2000 requests
        Completed 3000 requests
        Completed 4000 requests
        Completed 5000 requests
        Completed 6000 requests
        Completed 7000 requests
        Completed 8000 requests
        Completed 9000 requests
        Completed 10000 requests
        Finished 10000 requests

        Server Software:
        Server Hostname:        localhost
        Server Port:            8080

        Document Path:          /service/user?uid=
        Document Length:        461 bytes

        Concurrency Level:      100
        Time taken for tests:   2.330 seconds
        Complete requests:      10000
        Failed requests:        0
        Total transferred:      5850000 bytes
        HTML transferred:       4610000 bytes
        Requests per second:    4292.45 [#/sec] (mean)
        Time per request:       23.297 [ms] (mean)
        Time per request:       0.233 [ms] (mean, across all concurrent requests)
        Transfer rate:          2452.23 [Kbytes/sec] received

        Connection Times (ms)
                        min  mean[+/-sd] median   max
        Connect:        0    1   1.0      0       8
        Processing:     0   22  15.9     22     150
        Waiting:        0   22  15.9     21     149
        Total:          0   23  16.0     23     152

        Percentage of the requests served within a certain time (ms)
        50%     23
        66%     27
        75%     30
        80%     31
        90%     36
        95%     42
        98%     54
        99%     84
        100%    152 (longest request)

    The result of `xorm` is

        This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
        Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
        Licensed to The Apache Software Foundation, http://www.apache.org/

        Benchmarking localhost (be patient)
        Completed 1000 requests
        Completed 2000 requests
        Completed 3000 requests
        Completed 4000 requests
        Completed 5000 requests
        Completed 6000 requests
        Completed 7000 requests
        Completed 8000 requests
        Completed 9000 requests
        Completed 10000 requests
        Finished 10000 requests


        Server Software:
        Server Hostname:        localhost
        Server Port:            8080

        Document Path:          /service/user?uid=
        Document Length:        481 bytes

        Concurrency Level:      100
        Time taken for tests:   2.594 seconds
        Complete requests:      10000
        Failed requests:        0
        Total transferred:      6050000 bytes
        HTML transferred:       4810000 bytes
        Requests per second:    3854.58 [#/sec] (mean)
        Time per request:       25.943 [ms] (mean)
        Time per request:       0.259 [ms] (mean, across all concurrent requests)
        Transfer rate:          2277.36 [Kbytes/sec] received

        Connection Times (ms)
                      min  mean[+/-sd] median   max
        Connect:        0    1   1.0      0      10
        Processing:     1   25  11.7     25     219
        Waiting:        1   24  11.7     24     219
        Total:          1   26  11.8     26     219
        WARNING: The median and mean for the initial connection time are not within a normal deviation
                These results are probably not that reliable.

        Percentage of the requests served within a certain time (ms)
          50%     26
          66%     30
          75%     34
          80%     36
          90%     42
          95%     46
          98%     51
          99%     54
         100%    219 (longest request)

    The difference of total time is nearly 2 seconds, which may be a great number in computer world.

[here]: http://xorm.io/
