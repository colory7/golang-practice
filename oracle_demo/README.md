# MySQL 兼容函数说明


# Oracle 兼容函数说明

参考自 Oracle® Database SQL Language Reference 11g Release 2 (11.2)

## Oracle Format Model相关函数的测试
### Number format

Number format 相关函数 
```
to_char
to_number
to_binary_float
to_binary_double
```

测试用例
```
SET lc_numeric = 'C';

-- ,(comma) 半角逗号 装饰作用
-- 不能以逗号开头
-- 不能出现在数字或点号的右边
SELECT to_number('34,50','999,99') from dual;
SELECT to_number('34,50','99999') from dual;
SELECT to_number('34,50','99999,') from dual;
SELECT to_number('34,50',',99999') from dual;

-- .(period) 半角句号,点号，小数点
-- 只能有1个点号
SELECT to_char('32.78','9999') from dual;
SELECT to_number('0', '99.99') from dual;
select to_char('0', '99.99') from dual;
select to_char('+0.1', '99.99') from dual;
select to_char('-0.2', '99.99') from dual;
select to_char('+123.456', '999.999') from dual;
select to_char('-123.456', '999.999') from dual;
SELECT to_char('32.78','9999') from dual;

-- $ 美元符号 返回的值以$符号开头
select to_number('$12345.678', '$$999999.999') from dual;
SELECT TO_CHAR('12345.678', '$99,990.99') FROM DUAL;

-- 0 以0开头或以0结尾
-- 如果格式模型0对应的数字是0，并且是开头或结尾的0，则占位为0，如果不是0，则保持不变
SELECT TO_NUMBER('1','0') FROM DUAL;
SELECT TO_NUMBER(1,'0') FROM DUAL;


SELECT to_char('0','99.99') from dual;
SELECT to_char('00032','99.99') from dual;
SELECT length(to_char('0','99.99')) from dual;
SELECT length(to_char('00032','99.99')) from dual;
SELECT to_char('+0.1','90.99') from dual;
SELECT to_char('+0.1','90.00') from dual;
SELECT to_char('+123.456','999.009') from dual;
select to_char('0', '90.99') from dual;
select to_char('+0.1', '90.99') from dual;
select to_char('-0.2', '90.99') from dual;


-- 9 替换数字，开头的0替换为空格，0除外
select to_char('0', '9999') from dual;
select to_char('032', '9999') from dual;
select to_char('1', '9999') from dual;
select to_char('032', '9999') from dual;
select to_char('+32', '9999') from dual;
select to_char('-32', '9999') from dual;
select to_char('0.32', '9999') from dual;

-- B 如果格式对应的输出是0，则返回空白
select to_char('0.001', 'B9999') from dual;
select to_char('0.001', 'B9999.99') from dual;
select to_char('0.001', 'B9999.999') from dual;

select to_char('0', 'B9999') from dual;
select to_char('1', 'B9999') from dual;
select to_char('0', 'B90.99') from dual;
select to_char('0.134', 'B90.99') from dual;
select to_char('0.135', 'B90.99') from dual;
select to_char('0.134', 'B99.99') from dual;
select to_char('0.135', 'B99.99') from dual;

-- C 返回指定ISO货币符号
-- 参数 NLS_ISO_CURRENCY
select to_char('356', 'C999') from dual;


-- D 返回指定的十进制字符，默认是.
-- 只能有1个十进制字符
-- 参数 NLS_NUMERIC_CHARACTER
select to_number('12,454.8-', '99G999D9S') from dual;

-- EEEE 科学计数
select to_char('+123.456', '9.9EEEE') from dual;
select to_char('+123.456', '9.99EEEE') from dual;
select to_char('+123.456', '9.999EEEE') from dual;
select to_char('+123.456', '9.9999EEEE') from dual;
select to_char('+123.456', '9.99999EEEE') from dual;
select to_char('+1E+123', '9.9EEEE') from dual;
select to_char('+123.456', 'FM9.9EEEE') from dual;

-- G 分组符号
-- 不能出现在数字的右侧
-- 不能和句号.共同出现
-- NUMERIC_CHARACTER
select to_char(1485, '9G999') from dual;
select to_char(1485.3, '99G99.9') from dual;
select to_char(1485.3, '99G99D9') from dual;
select to_char(1485, '99G99') from dual;

-- L 本地货币符号
-- NLS_CURRENCY
select to_char('+123.45', 'L999.99') from dual;
SELECT to_char('1234.56','L99,999.99') from dual;

-- MI 返回以负号结尾的负数 或者 返回以空白结尾的正数
SELECT to_char(-485, '999MI')	from dual;
SELECT to_char(485, '999MI')	from dual;
SELECT to_char(485, 'FM999MI') from dual;

-- PR 尖括号包裹 ,如果是正数则返回，负数用尖括号包裹
-- PR格式只能出现在最后
select to_char(485, '999PR') from dual;
select to_char(-485, '999PR') from dual;

-- RN 大写罗马数字 范围 1-3999
select to_char(14825, 'RN') from dual;

-- rn 小写罗马数字 范围 1-3999
select to_char(1485, 'rn') from dual;

-- S
/*
S 在开头 数值前附加正负号
S 在结尾 数值后附加正负号
*/
-- 只能出现在开头或结尾

select to_char(-258, 'S999') from dual;
select to_char(258, 'S999') from dual;
select to_char(-258, '999S') from dual;
select to_char(258, '999S') from dual;

select to_char('-1234567890', '9999999999S') from dual;
select to_char('+1234567890', '9999999999S') from dual;

-- TM 文本最小数值
-- TM后只能跟随1个9或者1个E(e)
-- 大于64个字符用科学计数法输出
SELECT TO_CHAR(1234, 'TM') FROM DUAL;
SELECT TO_CHAR(1234, 'TM9') FROM DUAL;
SELECT TO_CHAR(1234, 'TME') FROM DUAL;
SELECT TO_CHAR(1234, 'TMe') FROM DUAL;
SELECT TO_CHAR(1234, 'TM9e') FROM DUAL;

-- U 欧元货币符号
-- NLS_DUAL_CURRENCY
select to_char(1258, 'U9999') from dual;

-- V
/*
V之前的9对应原数值m
V之后的9的个数对应n
输出: m乘以10的n次方
*/
-- 可能会四舍五入
select to_char(12, '99V999') from dual;
select to_char(12.45, '99V9') from dual;


-- X 10进制转换为16进制数
-- 浮点数会被四舍五入为整数
-- 只接受整数或0，负数则报错
-- 以0或FM开头，否则会返回空白
select to_char('17', 'XXXX') from dual;
select to_char('017', 'XXXX') from dual;
select to_char('17', '0XXXX') from dual;
select to_char('0017', 'FMXXXX') from dual;
select to_char('17', 'FMXXXX') from dual;
select to_char('00017', 'XXXX') from dual;

格式化修饰符

-- FM
-- 让输出更紧凑，删除前后空白
-- 适用于 to_char to_date* to_number
select to_char(123.456,'9.9EEEE') from dual;
select to_char(+123.456,'FM9.9EEEE') from dual;
select to_char('+123.456', 'FM999.009') from dual;
select to_char('+123.45', 'FM999.009') from dual;
select to_char('+0123.0', 'FM9999.99') from dual;
select to_char('+123.45', 'FML999.99') from dual;
select to_char(+123.456,'999.009') from dual;
select to_char(+123.456,'FM999.009') from dual;

select to_char(+123.456,'L999.99') from dual;
select to_char(+123.456,'FML999.99') from dual;

select to_char(+123.456000,'L999.999999') from dual;
select to_char(+123.456000,'FXL999.999999') from dual;

-- FX
-- 不适用于 数字类型 或 to_char函数???
-- 适用于 to_date
```


## Date format
Date format 相关函数
```
to_date
to_timestamp
to_timestamp_tz
to_char
```

测试用例
```
-- - 结果中会复制标点符号和引用文本。
SELECT TO_DATE('2008-05-20','YYYY-MM-DD') FROM DUAL;
SELECT TO_DATE('2008----05--20','YYYY----MM--DD') FROM DUAL;

-- / 同上
SELECT TO_DATE('2008////05//20','YYYY////MM//DD') FROM DUAL;
SELECT TO_DATE('2008////05//20','YYYY----MM//DD') FROM DUAL;
-- , 同上
SELECT TO_DATE('2008,,,,,05,,20','YYYY,,,,,MM//DD') FROM DUAL;
SELECT TO_DATE('2008,,,,,05,,20','YYYY,,,,,MM,,DD') FROM DUAL;

-- . 同上
SELECT TO_DATE('2008....05,,20','YYYY....MM..DD') FROM DUAL;
SELECT TO_DATE('2008....05,,20','YYYY....MM,,DD') FROM DUAL;

-- ; 同上
SELECT TO_DATE('2008;;;;05,,20','YYYY;;;;MM..DD') FROM DUAL;
SELECT TO_DATE('2008;;;;05,,20','YYYY;;;;MM,,DD') FROM DUAL;

-- : 同上
SELECT TO_DATE('2008::::05,,20','YYYY::::MM..DD') FROM DUAL;
SELECT TO_DATE('2008::::05,,20','YYYY::::MM,,DD') FROM DUAL;

-- "text"
SELECT TO_DATE('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008xxxx05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008"xxxx"05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008xxxx05,,20','YYYYabcdMM,,DD') FROM DUAL;

-- AD 公元后
Select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','AD yyyy-mm-dd hh:mm:ss') From dual;

-- A.D.  同上
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','A.D. yyyy-mm-dd hh:mm:ss') From dual;

-- AM 上午
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','AM yyyy-mm-dd hh:mm:ss') From dual;

-- A.M. 同上
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','A.M. yyyy-mm-dd hh:mm:ss') From dual;

-- BC 公元前
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','BC yyyy-mm-dd hh:mm:ss') From dual;

-- B.C. 同上
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','B.C. yyyy-mm-dd hh:mm:ss') From dual;

-- CC 世纪
/*
如果最后2位数字是01-99，则返回世纪数位前两位数+1
如果最后2位数字是00，则返回世纪数位前两位数
*/
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','CC') from dual;

-- SCC 同上
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','SCC') from dual;
SELECT to_char(date'2000-01-01','SCC') from dual;
SELECT to_char(date'2001-01-01','SCC') from dual;

-- D 星期的第几天
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','D') from dual;

-- DAY 天的名称
SELECT to_char(current_timestamp, 'Day, DD HH12:MI:SS') FROM DUAL;
SELECT to_char(current_timestamp, 'FMDay, FMDD HH12:MI:SS') FROM DUAL;

-- DD 月的第几天
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','dd')       from dual;

-- DDD 年的第几天
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','ddd')       from dual;

-- DL 长日期格式
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','DL') from dual;

-- DS 短日期格式
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','DS') from dual;

-- DY 天的缩写
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','DY') from dual;

-- E 时代名称缩写 中华民国，泰国佛时代
-- EE 完全时代名称

-- FF [1..9] 秒数，最多精确到纳秒，9位小数
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF1') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF2') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF3') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF4') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF5') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF6') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF7') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF8') from DUAL;
SELECT TO_CHAR(SYSTIMESTAMP, 'SS.FF9') from DUAL;

-- FM 填充模式
-- 如果返回的值的开头是0，则替换为空白
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789', 'fmDDTH') from dual;
select TO_DATE('0207','MM/YY') from dual;
select TO_DATE('0207','MM/YY') from dual;
select TO_DATE('02#07','MM/YY') from dual;
select TO_DATE('02#07','FMMM/YY') from dual;
select TO_DATE('02#07','FXMM/YY') from dual; -- error

SELECT TO_CHAR(SYSDATE, 'DDTH') || ' of ' ||
         TO_CHAR(SYSDATE, 'DD') || ', ' ||
         TO_CHAR(SYSDATE, 'YYYY') "Ides"
    FROM DUAL;

SELECT TO_CHAR(SYSDATE, 'fmDDTH') || ' of ' ||
         TO_CHAR(SYSDATE, 'fmDD') || ', ' ||
         TO_CHAR(SYSDATE, 'YYYY') "Ides"
    FROM DUAL;

-- FX 固定模式
/*
1.字符及对应的格式必须严格一一对应
2.不允许有多余的空格
3.数值参与格式需要完全对应(或通过fm参数去掉前置0)
*/
SELECT TO_DATE('15/ JAN /1998', 'DD-MON-YYYY') FROM DUAL; -- Match
SELECT TO_DATE(' 15! JAN % /1998', 'DD-MON-YYYY') FROM DUAL; -- Error
SELECT TO_DATE('15-JAN-1998', 'FXDD-MON-YYYY') FROM DUAL; -- Match
SELECT TO_DATE('15/JAN/1998', 'FXDD-MON-YYYY') FROM DUAL; -- Error


SELECT TO_DATE('01-01-1998' ,'FXDD-MM-YYYY') FROM DUAL; -- Match
SELECT TO_DATE('1-1-1998' ,'FXFMDD-MM-YYYY') FROM DUAL; -- Match
SELECT TO_DATE('1-01-1998' ,'FXDD-MON-YYYY') FROM DUAL; -- Error

select TO_DATE('02#07','FXMM#YY') from dual;
select TO_DATE('02#07','FXMM/YY') from dual;

-- HH
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','hh')   from dual;

-- HH12
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','hh12')   from dual;

-- HH24
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','hh24')   from dual;


-- IW 年的第几个星期 ISO8601
/*
日历星期从周一开始
第1个日历星期包含1月4号
第1个日历星期可能包含12月29，30，31
最后一个日历星期可能包含1月1，2，3
*/
select to_char(TIMESTAMP '2023-01-01 09:26:50.124','IW') from dual;
select to_char(TIMESTAMP '2023-01-02 09:26:50.124','IW') from dual;
select to_char(TIMESTAMP '2023-01-03 09:26:50.124','IW') from dual;
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','IW') from dual;

-- IYYY 4位数字的年，包含日历星期 ISO8601
select to_char(TIMESTAMP '2023-01-01 09:26:50.124','IYYY') from dual;
select to_char(TIMESTAMP '2023-01-02 09:26:50.124','IYYY') from dual;
select to_char(TIMESTAMP '2023-01-03 09:26:50.124','IYYY') from dual;
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','IYYY') from dual;

-- IYY 最后3个数字的年，包含日历星期 ISO8601
select to_char(TIMESTAMP '2023-01-01 09:26:50.124','IYY') from dual;
select to_char(TIMESTAMP '2023-01-02 09:26:50.124','IYY') from dual;
select to_char(TIMESTAMP '2023-01-03 09:26:50.124','IYY') from dual;
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','IYY') from dual;

-- IY 最后2个数字的年，包含日历星期 ISO8601
select to_char(TIMESTAMP '2023-01-01 09:26:50.124','IY') from dual;
select to_char(TIMESTAMP '2023-01-02 09:26:50.124','IY') from dual;
select to_char(TIMESTAMP '2023-01-03 09:26:50.124','IY') from dual;
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','IY') from dual;

-- I 1个数字的年，包含日历星期 ISO8601
select to_char(TIMESTAMP '2023-01-01 09:26:50.124','I') from dual;
select to_char(TIMESTAMP '2023-01-02 09:26:50.124','I') from dual;
select to_char(TIMESTAMP '2023-01-03 09:26:50.124','I') from dual;
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','I') from dual;

-- J  julian calendar 从1, 4712 BC开始，必须是整数
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','J')   from dual;


-- MI 分钟 0-59
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','mi') from dual;


-- MM 月 01-12
SELECT TO_DATE('2008 05 20','YYYY MM DD') FROM DUAL;
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','MM') from dual;


-- MON 月的缩写，同MONTH
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','MON') from dual;

-- MONTH 月的名称
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','MONTH') from dual;

-- PM
Select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY-MM-DD PM hh24:mi:ss') From dual;

-- P.M. 同上
Select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY-MM-DD P.M. hh24:mi:ss') From dual;

-- Q 季节 1-4
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','Q') from dual;

-- RM 罗马数字的月
select to_char(TIMESTAMP '2023-01-04 09:26:50.124','RM') from dual;

-- RR 用2位数字保存20世纪的日期在21世纪的日期里
-- 系统时间 1950-1999
-- 指定时间 50-99  98
SELECT TO_CHAR(TO_DATE('27-10-98', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
-- 输出 1998

-- 系统时间 1950-1999
-- 指定时间 00-49  17
SELECT TO_CHAR(TO_DATE('27-10-17', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
-- 输出 2017

-- 系统时间 2000 - 2049
-- 指定时间 50-99  98
SELECT TO_CHAR(TO_DATE('27-10-98', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
-- 输出 1998

-- 系统时间 2000 - 2049
-- 指定时间 0-49  17
SELECT TO_CHAR(TO_DATE('27-10-17', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
-- 输出 2017

-- RRRR
/*
如果RRRR对应的数字有2位，则与RR相同
a.年的后两位为指定年数的2位
b.年的前两位
    如果指定的年数在0-49，则取当前系统的年的前两位
    如果指定的年数在49-99，则取当前系统的年的前两位-1
*/
SELECT TO_CHAR(TO_DATE('27-10-1998', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
SELECT TO_CHAR(TO_DATE('27-10-2017', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
SELECT TO_CHAR(TO_DATE('27-10-98', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;
SELECT TO_CHAR(TO_DATE('27-10-17', 'DD-MM-RR'), 'YYYY') "Year" FROM DUAL;

-- SS 秒 0-59
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','ss')      from dual;
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','SS') FROM DUAL;


-- SSSSS 0-86399
-- 午夜过后的秒
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.476589 US/Pacific PDT','SSSSS') FROM DUAL;
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.76589 US/Pacific PDT','SSSSS') FROM DUAL;

-- TS
SELECT to_char(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TS') FROM DUAL;

-- TZD
SELECT to_char(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TZD') FROM DUAL;

-- TZH
SELECT to_char(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TZH') FROM DUAL;

-- TZM
SELECT to_char(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TZM') FROM DUAL;

-- TZR
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','TZR') FROM DUAL;

-- WW 年的第几个星期 1-53
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','WW') FROM DUAL;

-- W 月的第几周取值范围1-5。星期从1开始，7结束。
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','WW') FROM DUAL;

-- X 本地基数字符 小数点
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','HH:MI:SSXFF') FROM DUAL;

-- Y,YYY 逗号分隔的固定格式表示的年
SELECT to_char(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','Y,YYY') FROM DUAL;

-- YEAR 年,单词拼写方式
SELECT to_char(date'2000-01-01','YEAR') FROM DUAL;
SELECT to_char(date'-2000-01-01','YEAR') FROM DUAL;
SELECT to_char(date'2000-01-01','AD YEAR') FROM DUAL;
SELECT to_char(date'-2000-01-01','AD YEAR') FROM DUAL;


-- SYEAR 年,单词拼写方式。对年添加标记，公元前则加负号-，公元后则是空格
SELECT to_char(date'2000-01-01','SYEAR') FROM DUAL;
SELECT to_char(date'-2000-01-01','SYEAR') FROM DUAL;
SELECT to_char(date'2000-01-01','AD SYEAR') FROM DUAL;
SELECT to_char(date'-2000-01-01','AD SYEAR') FROM DUAL;

-- YYYY 4位数字的年
SELECT TO_CHAR(TO_DATE('27-OCT-98', 'DD-MON-RR'), 'YYYY') "Year" FROM DUAL;
SELECT TO_CHAR(TO_DATE('27-OCT-17', 'DD-MON-RR'), 'YYYY') "Year" FROM DUAL;
select to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','yyyy')    from dual;

-- SYYYY 4位数字的年。对年添加标记，公元前则加负号-，公元后则是空格
SELECT to_char(TIMESTAMP '1987-10-29 01:30:56.321654789','SYYYY') FROM DUAL;

-- YYY 最后3位数字的年
SELECT to_char(TIMESTAMP '1987-10-29 01:30:56.321654789','YYY') FROM DUAL;


-- YY 最后2位数字的年
SELECT to_char(TIMESTAMP '1987-10-29 01:30:56.321654789','YY') FROM DUAL;
SELECT TO_CHAR(TO_DATE('0207','MM/YY'), 'MM/YY') FROM DUAL;
SELECT TO_CHAR (TO_DATE('02#07','MM/YY'), 'MM/YY') FROM DUAL;


-- Y 最后1位数字的年
SELECT TO_CHAR(TO_DATE('027','MM/Y'), 'MM/Y') FROM DUAL;
SELECT to_char(TIMESTAMP '1987-10-29 01:30:56.321654789 ','Y') FROM DUAL;

-- 日期格式后缀
-- TH 序号,英文序数词描述
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','DDTH') FROM DUAL;


-- SP 拼写,英文拼写方式描述
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','DDSP') FROM DUAL;

-- SPTH 英文拼写方式描述+英文序数词描述
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789','DDTHSP') FROM DUAL;

-- THSP 英文序数词描述+英文拼写方式描述
SELECT to_char(TIMESTAMP '2023-10-29 01:30:56.321654789 ','DDSPTH') FROM DUAL;

```

### 函数按照功能划分

#### 输入 数值类型参数 ,输出 数值类型 
    to_number
    to_char

#### 输入 字符串类型参数 ,输出 数值类型
    to_number
    to_char

#### 输入 时间类型参数 ,输出 字符串类型
    to_char

#### 输入 字符串类型参数 ,输出 字符串类型
    to_char

#### 输入 字符串类型参数 ,输出 时间类型
    to_date
    to_timestamp
    to_timestamp_tz

## 函数实现分析
### to_number 函数
#### 输出模式
    十进制
        左符号 + - < 空
        货币符号
        数值模型 9 0 . ,
        输出模式 十进制 空
        右符号 + - > 空
    科学计数
        左符号 + - < 空
        货币符号
        数值模型 9 0 .  不包含 逗号
        输出模式 EEEE
        右符号 + - > 空
    乘积
        左符号 + - < 空
        货币符号
        数值模型 9 0 , 不包含 小数点
        输出模式 V
        右符号 + - > 空
    十六进制
        X
        参数 大于0的正整数, 浮点数mod取整
        输出模式 十六进制
    罗马数字
        参数 大于0的正整数, 浮点数mod取整
        输出模式 RN
    最小文本 十进制
        TM TM9
    最小文本 科学计数
        TME TMe


### to_binary_float
### to_binary_double
### to_date 
### to_timestamp 
### to_timestamp_tz 
### to_char 
