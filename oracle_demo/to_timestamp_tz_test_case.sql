-- TO_TIMESTAMP_TZ

---- 日期模型
-------- a.将 字符串类型参数 转换为 时间戳类型
-----------------------------------------------------
-----------------------------------------------------
-- TO_TIMESTAMP_TZ

-- - 结果中会复制标点符号和引用文本。
Select TO_TIMESTAMP_TZ( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS') From dual;
Select TO_TIMESTAMP_TZ( '2023-10-29','YYYY-MM-DD') From dual;

-- / 同上
Select TO_TIMESTAMP_TZ('2023////10//29','YYYY////MM//DD') From dual;

-- , 同上
Select TO_TIMESTAMP_TZ('2023,,,,,10//29','YYYY,,,,,MM//DD') From dual;

-- . 同上
Select TO_TIMESTAMP_TZ('2023....10..29','YYYY....MM..DD') From dual;

-- ; 同上
Select TO_TIMESTAMP_TZ('2023....10..29','YYYY;;;;MM..DD') From dual;

-- : 同上
Select TO_TIMESTAMP_TZ('2023....10..29','YYYY::::MM..DD') From dual;

-- "text"
Select TO_TIMESTAMP_TZ('2023abcd10..29','YYYY"abcd"MM,,DD') From dual;
SELECT TO_TIMESTAMP_TZ('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_TIMESTAMP_TZ('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;

SELECT TO_TIMESTAMP_TZ('2008xxxx05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_TIMESTAMP_TZ('2008"xxxx"05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_TIMESTAMP_TZ('2008xxxx05,,20','YYYYabcdMM,,DD') FROM DUAL;

-- 年 月 日 时 分 秒 毫秒 微妙 纳秒 时区
-- 时区目前已知2种表达方式:
--   1. TZH:TZM 如 +08:00 , -09:03
--   2. TZR 如 +08:00 , -09:03 , Asia/Shanghai ,CST ,UTC
SELECT TO_TIMESTAMP_TZ('2008 05 20','YYYY MM DD') FROM DUAL;
SELECT TO_TIMESTAMP_TZ('2008-05-20','YYYY-MM-DD') FROM DUAL;
select TO_TIMESTAMP_TZ('09:26:50','HH:MI:SS') from dual;
select TO_TIMESTAMP_TZ('23--26-50','HH24--MI-SS') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50','YYYY-MM-DD HH:MI:SS') from dual;
select TO_TIMESTAMP_TZ('2023-1月-04 09:26:50','YYYY-MON-DD HH:MI:SS') from dual;
select TO_TIMESTAMP_TZ('2023-1月-04 09:26:50','FMYYYY-MON-DD HH:MI:SS') from dual;
select TO_TIMESTAMP_TZ('2023-1月-04 09:26:50','FXYYYY-MON-DD HH:MI:SS') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897','YYYY-MM-DD HH:MI:SS.FF9') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 09','YYYY-MM-DD HH:MI:SS.FF9 TZH') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 09:02','YYYY-MM-DD HH:MI:SS.FF9 TZH:TZM') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 +09:02','YYYY-MM-DD HH:MI:SS.FF9 TZH:TZM') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 Asia/Shanghai','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 Asia/Urumqi','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 +09:03','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 CST','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 GMT','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 UTC','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 EDT','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;
select TO_TIMESTAMP_TZ('2023-01-04 09:26:50.231456897 PST','YYYY-MM-DD HH:MI:SS.FF9 TZR') from dual;

