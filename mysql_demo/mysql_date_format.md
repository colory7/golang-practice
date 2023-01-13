
# MySQL函数
## 参考
MySQL 8.0 Reference Manual Including MySQL NDB Cluster 8.0

## 函数
### date_format
#### 日期模式说明
%U 00-53 数值跨年，则初始化位00,最多累加到53，第1周有可能是小于7天
%V 01-53 数值不跨年 则最多累加到53; 跨年，则最多累加到52
%X 年份跟随V

%u 00-53 数值跨年，则初始化位01，最多累加到52，第1周有可能是小于7天
%v 01-52 数值不跨年,最多累加到52,之后从01开始计算
%x 年份 跟随 v
-- 小写 周一是第一天
-- 大写 周日是第一天

#### 测试用例
##### case 1
```mysql
SELECT DATE_FORMAT('2019-12-28 00:00:00','%u %v %x - %U %V %X'); -- 周六
-- 52 52 2019 - 51 51 2019
SELECT DATE_FORMAT('2019-12-29 00:00:00','%u %v %x - %U %V %X'); -- 周日
-- 52 52 2019 - 52 52 2019
SELECT DATE_FORMAT('2019-12-30 00:00:00','%u %v %x - %U %V %X'); -- 周一
-- 53 01 2020 - 52 52 2019
SELECT DATE_FORMAT('2019-12-31 00:00:00','%u %v %x - %U %V %X'); -- 周二
-- 53 01 2020 - 52 52 2019

SELECT DATE_FORMAT('2020-01-01 00:00:00','%u %v %x - %U %V %X'); -- 周三
-- 01 01 2020 - 00 52 2019
SELECT DATE_FORMAT('2020-01-02 00:00:00','%u %v %x - %U %V %X'); -- 周四
-- 01 01 2020 - 00 52 2019
SELECT DATE_FORMAT('2020-01-03 00:00:00','%u %v %x - %U %V %X'); -- 周五
-- 01 01 2020 - 00 52 2019
SELECT DATE_FORMAT('2020-01-04 00:00:00','%u %v %x - %U %V %X'); --	周六
-- 01 01 2020 - 00 52 2019
SELECT DATE_FORMAT('2020-01-05 00:00:00','%u %v %x - %U %V %X'); -- 周日
-- 01 01 2020 - 01 01 2020
SELECT DATE_FORMAT('2020-01-06 00:00:00','%u %v %x - %U %V %X'); -- 周一
-- 02 02 2020 - 01 01 2020
```

##### case 2
```mysql
SELECT DATE_FORMAT('1998-12-27 00:00:00','%u %v %x - %U %V %X'); -- 周日
-- 52 52 1998 - 52 52 1998
SELECT DATE_FORMAT('1998-12-28 00:00:00','%u %v %x - %U %V %X'); -- 周一
-- 53 53 1998 - 52 52 1998
SELECT DATE_FORMAT('1998-12-29 00:00:00','%u %v %x - %U %V %X'); -- 周二
-- 53 53 1998 - 52 52 1998
SELECT DATE_FORMAT('1998-12-30 00:00:00','%u %v %x - %U %V %X'); -- 周三
-- 53 53 1998 - 52 52 1998
SELECT DATE_FORMAT('1998-12-31 00:00:00','%u %v %x - %U %V %X'); -- 周四
-- 53 53 1998 - 52 52 1998

SELECT DATE_FORMAT('1999-01-01 00:00:00','%u %v %x - %U %V %X'); -- 周五
-- 00 53 1998 - 00 52 1998
SELECT DATE_FORMAT('1999-01-02 00:00:00','%u %v %x - %U %V %X'); -- 周六
-- 00 53 1998 - 00 52 1998
SELECT DATE_FORMAT('1999-01-03 00:00:00','%u %v %x - %U %V %X'); -- 周日
-- 00 53 1998 - 01 01 1999
SELECT DATE_FORMAT('1999-01-04 00:00:00','%u %v %x - %U %V %X'); -- 周一
-- 01 01 1999 - 01 01 1999
SELECT DATE_FORMAT('1999-01-05 00:00:00','%u %v %x - %U %V %X'); -- 周二
-- 01 01 1999 - 01 01 1999
SELECT DATE_FORMAT('1999-01-06 00:00:00','%u %v %x - %U %V %X'); -- 周三
-- 01 01 1999 - 01 01 1999
```

##### case 3
```mysql
SELECT DATE_FORMAT('2000-12-22 00:00:00','%u %v %x - %U %V %X %W %w');
-- 51 51 2000 - 51 51 2000 Friday 5
SELECT DATE_FORMAT('2000-12-23 00:00:00','%u %v %x - %U %V %X %W %w');
-- 51 51 2000 - 51 51 2000 Saturday 6
SELECT DATE_FORMAT('2000-12-24 00:00:00','%u %v %x - %U %V %X %W %w');
-- 51 51 2000 - 52 52 2000 Sunday 0
SELECT DATE_FORMAT('2000-12-25 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 52 52 2000 Monday 1
SELECT DATE_FORMAT('2000-12-26 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 52 52 2000 Tuesday 2
SELECT DATE_FORMAT('2000-12-27 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 52 52 2000
SELECT DATE_FORMAT('2000-12-28 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 52 52 2000
SELECT DATE_FORMAT('2000-12-29 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 52 52 2000
SELECT DATE_FORMAT('2000-12-30 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 52 52 2000
SELECT DATE_FORMAT('2000-12-31 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2000 - 53 53 2000

SELECT DATE_FORMAT('2001-01-01 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 00 53 2000
SELECT DATE_FORMAT('2001-01-02 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 00 53 2000
SELECT DATE_FORMAT('2001-01-03 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 00 53 2000
SELECT DATE_FORMAT('2001-01-04 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 00 53 2000
SELECT DATE_FORMAT('2001-01-05 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 00 53 2000
SELECT DATE_FORMAT('2001-01-06 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 00 53 2000
SELECT DATE_FORMAT('2001-01-07 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 01 01 2001
SELECT DATE_FORMAT('2001-01-08 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2001 - 01 01 2001
```

##### case 4
```mysql
SELECT DATE_FORMAT('2001-12-27 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2001 - 51 51 2001 Thursday 4
SELECT DATE_FORMAT('2001-12-28 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2001 - 51 51 2001 Friday 5
SELECT DATE_FORMAT('2001-12-29 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2001 - 51 51 2001 Saturday 6
SELECT DATE_FORMAT('2001-12-30 00:00:00','%u %v %x - %U %V %X %W %w');
-- 52 52 2001 - 52 52 2001 Sunday 0
SELECT DATE_FORMAT('2001-12-31 00:00:00','%u %v %x - %U %V %X %W %w');
-- 53 01 2002 - 52 52 2001 Monday 1

SELECT DATE_FORMAT('2002-01-01 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2002 - 00 52 2001 Tuesday 2
SELECT DATE_FORMAT('2002-01-02 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2002 - 00 52 2001 Wednesday 3
SELECT DATE_FORMAT('2002-01-03 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2002 - 00 52 2001 Thursday 4
SELECT DATE_FORMAT('2002-01-04 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2002 - 00 52 2001 Friday 5
SELECT DATE_FORMAT('2002-01-05 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2002 - 00 52 2001 Saturday 6
SELECT DATE_FORMAT('2002-01-06 00:00:00','%u %v %x - %U %V %X %W %w');
-- 01 01 2002 - 01 01 2002 Sunday 0
SELECT DATE_FORMAT('2002-01-07 00:00:00','%u %v %x - %U %V %X %W %w');
-- 02 02 2002 - 01 01 2002 Monday 1
SELECT DATE_FORMAT('2002-01-08 00:00:00','%u %v %x - %U %V %X %W %w');
-- 02 02 2002 - 01 01 2002 Tuesday 2
```