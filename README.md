# wtz

_What timezone?_

A small command line tool that shows times in timezones you're interested in:

The current timezone of your computer:
```
$ wtz
+-----------+
| MELBOURNE |
+-----------+
| 00:00     |
| 01:00     |
| 02:00     |
| 03:00     |
[...]
```

The current timezone of your computer and other timezones you're interested in.
```
$ wtz --tz Europe/Madrid,America/New_York,America/Argentina/Buenos_Aires,UTC
+-----------+--------+----------+--------------+-------+
| MELBOURNE | MADRID | NEW YORK | BUENOS AIRES |  UTC  |
+-----------+--------+----------+--------------+-------+
| 00:00     | 14:00  | 08:00    | 10:00        | 13:00 |
| 01:00     | 15:00  | 09:00    | 11:00        | 14:00 |
| 02:00     | 16:00  | 10:00    | 12:00        | 15:00 |
| 03:00     | 17:00  | 11:00    | 13:00        | 16:00 |
| 04:00     | 18:00  | 12:00    | 14:00        | 17:00 |
| 05:00     | 19:00  | 13:00    | 15:00        | 18:00 |
| 06:00     | 20:00  | 14:00    | 16:00        | 19:00 |
| 07:00     | 21:00  | 15:00    | 17:00        | 20:00 |
| 08:00     | 22:00  | 16:00    | 18:00        | 21:00 |
| 09:00     | 23:00  | 17:00    | 19:00        | 22:00 |
| 10:00     | 00:00  | 18:00    | 20:00        | 23:00 |
| 11:00     | 01:00  | 19:00    | 21:00        | 00:00 |
| 12:00     | 02:00  | 20:00    | 22:00        | 01:00 |
| 13:00     | 03:00  | 21:00    | 23:00        | 02:00 |
| 14:00     | 04:00  | 22:00    | 00:00        | 03:00 |
| 15:00     | 05:00  | 23:00    | 01:00        | 04:00 |
| 16:00     | 06:00  | 00:00    | 02:00        | 05:00 |
| 17:00     | 07:00  | 01:00    | 03:00        | 06:00 |
| 18:00     | 08:00  | 02:00    | 04:00        | 07:00 |
| 19:00     | 09:00  | 03:00    | 05:00        | 08:00 |
| 20:00     | 10:00  | 04:00    | 06:00        | 09:00 |
| 21:00     | 11:00  | 05:00    | 07:00        | 10:00 |
| 22:00     | 12:00  | 06:00    | 08:00        | 11:00 |
| 23:00     | 13:00  | 07:00    | 09:00        | 12:00 |
+-----------+--------+----------+--------------+-------+
```

Inspired by [wtftz.sh](https://gist.github.com/mricon/3860883).