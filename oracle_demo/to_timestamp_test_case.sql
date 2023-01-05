-- TO_TIMESTAMP

---- 日期模型
-------- a.将 字符串类型参数 转换为 时间戳类型
-----------------------------------------------------
-----------------------------------------------------
-- - 结果中会复制标点符号和引用文本。
Select TO_TIMESTAMP( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS') From dual;
Select TO_TIMESTAMP( '2023-10-29','YYYY-MM-DD') From dual;

-- / 同上
Select TO_TIMESTAMP('2023////10//29','YYYY////MM//DD') From dual;

-- , 同上
Select TO_TIMESTAMP('2023,,,,,10//29','YYYY,,,,,MM//DD') From dual;

-- . 同上
Select TO_TIMESTAMP('2023....10..29','YYYY....MM..DD') From dual;

-- ; 同上
Select TO_TIMESTAMP('2023....10..29','YYYY;;;;MM..DD') From dual;

-- : 同上
Select TO_TIMESTAMP('2023....10..29','YYYY::::MM..DD') From dual;

-- "text"
Select TO_TIMESTAMP('2023abcd10..29','YYYY"abcd"MM,,DD') From dual;
SELECT TO_TIMESTAMP('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_TIMESTAMP('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;

SELECT TO_TIMESTAMP('2008xxxx05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_TIMESTAMP('2008"xxxx"05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_TIMESTAMP('2008xxxx05,,20','YYYYabcdMM,,DD') FROM DUAL;

-- 年 月 日 时 分 秒 毫秒 微妙 纳秒
SELECT TO_TIMESTAMP('2008 05 20','YYYY MM DD') FROM DUAL;
SELECT TO_TIMESTAMP('2008-05-20','YYYY-MM-DD') FROM DUAL;
select TO_TIMESTAMP('09:26:50','HH:MI:SS') from dual;
select TO_TIMESTAMP('23--26-50','HH24--MI-SS') from dual;
select TO_TIMESTAMP('2023-01-04 09:26:50','YYYY-MM-DD HH:MI:SS') from dual;
select TO_TIMESTAMP('2023-1月-04 09:26:50','YYYY-MON-DD HH:MI:SS') from dual;
select TO_TIMESTAMP('2023-1月-04 09:26:50','FMYYYY-MON-DD HH:MI:SS') from dual;
select TO_TIMESTAMP('2023-1月-04 09:26:50','FXYYYY-MON-DD HH:MI:SS') from dual;
select TO_TIMESTAMP('2023-01-04 09:26:50.231456897','YYYY-MM-DD HH:MI:SS.FF9') from dual;