-- TO_DATE

---- 日期模型
-------- a.将 字符串类型参数 转换为 日期类型
-----------------------------------------------------
-----------------------------------------------------
-- - 结果中会复制标点符号和引用文本。
Select TO_DATE( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS') From dual;
Select TO_DATE( '2023-10-29','YYYY-MM-DD') From dual;

-- / 同上
Select TO_DATE('2023////10//29','YYYY////MM//DD') From dual;

-- , 同上
Select TO_DATE('2023,,,,,10//29','YYYY,,,,,MM//DD') From dual;

-- . 同上
Select TO_DATE('2023....10..29','YYYY....MM..DD') From dual;

-- ; 同上
Select TO_DATE('2023....10..29','YYYY;;;;MM..DD') From dual;

-- : 同上
Select TO_DATE('2023....10..29','YYYY::::MM..DD') From dual;

-- "text"
Select TO_DATE('2023abcd10..29','YYYY"abcd"MM,,DD') From dual;
SELECT TO_DATE('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;

SELECT TO_DATE('2008xxxx05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008"xxxx"05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_DATE('2008xxxx05,,20','YYYYabcdMM,,DD') FROM DUAL;

-- 年 月 日 时 分 秒
SELECT TO_DATE('2008 05 20','YYYY MM DD') FROM DUAL;
SELECT TO_DATE('2008-05-20','YYYY-MM-DD') FROM DUAL;
select TO_DATE('09:26:50','HH:MI:SS') from dual;
select TO_DATE('23--26-50','HH24--MI-SS') from dual;
select TO_DATE('2023-01-04 09:26:50','YYYY-MM-DD HH:MI:SS') from dual;
select TO_DATE('2023-1月-04 09:26:50','YYYY-MON-DD HH:MI:SS') from dual;
select TO_DATE('2023-1月-04 09:26:50','FMYYYY-MON-DD HH:MI:SS') from dual;
select TO_DATE('2023-1月-04 09:26:50','FXYYYY-MON-DD HH:MI:SS') from dual;
