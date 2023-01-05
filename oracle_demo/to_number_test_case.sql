-- TO_NUMBER
---- 数值模型
------ 十进制
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型
------ 科学计数
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型
------ 乘积计数
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型
------ 十六进制
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型
------ 罗马数字
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型
------ 最小文本 十进制
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型
------ 最小文本 科学计数
-------- a.将 字符串类型参数 转换为 数值类型
-------- b.将 数值类型参数 转换为 数值类型

-----------------------------------------------------
-----------------------------------------------------

---- 数值模型
------ 十进制
-------- a.将 字符串类型参数 转换为 数值类型

-- ,(comma) 半角逗号 装饰作用
-- 不能以逗号开头
-- 不能出现在数字或点号的右边
-- G
-- 正面案例
SELECT TO_NUMBER('34,50','999,99') from dual;
SELECT TO_NUMBER('34,50','99999') from dual;
SELECT TO_NUMBER('34,50','99999,') from dual;
SELECT TO_NUMBER('34,50',',99999') from dual;
SELECT TO_NUMBER('34,50','99999') from dual;
SELECT TO_NUMBER('3450','99999') from dual;
select TO_NUMBER('12,4,548-', '99G9G999S') from dual;
select TO_NUMBER('12,4,548', '99G9G999') from dual;
select TO_NUMBER('12,4,5.48', '99G9G9D99') from dual;

-- 反面案例
SELECT TO_NUMBER('3450','999,99') from dual;
SELECT TO_NUMBER('3450','99') from dual;
SELECT TO_NUMBER('3450','9G9') from dual;
select TO_NUMBER('12,4,548-', '99G9G999S') from dual;
select TO_NUMBER('12,4,5.4,8', '99G9G9D9G9') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 无

-- 反面案例
SELECT TO_NUMBER(34,50,'999,99') from dual;
SELECT TO_NUMBER(34,50,'999G99') from dual;
SELECT TO_NUMBER(34,50,'99999') from dual;
SELECT TO_NUMBER(34,50,'99999,') from dual;
SELECT TO_NUMBER(34,50,',99999') from dual;
SELECT TO_NUMBER(34,50,'99999') from dual;
SELECT TO_NUMBER(3450,'99999') from dual;
select TO_NUMBER(-124548, 'S999999') from dual;
SELECT TO_NUMBER(3450,'999,99') from dual;
SELECT TO_NUMBER(3450,'99') from dual;
SELECT TO_NUMBER(3450,'9G9') from dual;
select TO_NUMBER(-12,4,548, '99G9G999S') from dual;
select TO_NUMBER(-12,4,548, 'S99G9G999S') from dual;
select TO_NUMBER(12,4,5.4,8, '99G9G9D9G9') from dual;
select TO_NUMBER(-12,4,548, 'S99G9G999') from dual;
select TO_NUMBER(124548, '99G9G999') from dual;
select TO_NUMBER(1245.48, '99G9G9D99') from dual;

-- .(period) 半角句号,点号，小数点
-- 只能有1个点号
-- G 分组符号
-- 不能出现在数字的右侧
-- 不能和句号.共同出现
-- NUMERIC_CHARACTER

-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_NUMBER('23.54', '99.99') from dual;
SELECT TO_NUMBER('23.54', '99D99') from dual;
select TO_NUMBER('12,454.8-', '99G999D9S') from dual;
select TO_NUMBER('12,4,54.8-', '99G9G99D9S') from dual;
select TO_NUMBER('12,4,548-', '99G9G999S') from dual;
select TO_NUMBER('12,4,548', '99G9G999') from dual;


-- 反面案例
SELECT TO_NUMBER('23.5.4', '99.9.9') from dual;
SELECT TO_NUMBER('23.54', '99.9.9') from dual;
SELECT TO_NUMBER('23.5.4', '99D9D9') from dual;
SELECT TO_NUMBER('23.54', '99D9D9') from dual;


-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_NUMBER(23.54, '99.99') from dual;
SELECT TO_NUMBER(23.54, '99D99') from dual;

-- 反面案例
SELECT TO_NUMBER(23.54, '99D99') from dual;


-- 0 以0开头或以0结尾
-- 如果格式模型0对应的数字是0，并且是开头或结尾的0，则占位为0，如果不是0，则保持不变
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例

SELECT TO_NUMBER('1','0') FROM DUAL;
SELECT TO_NUMBER(1,'0') FROM DUAL;
SELECT TO_NUMBER('1322526','0099000') FROM DUAL;
SELECT TO_NUMBER('1322526','0099000') FROM DUAL;
SELECT TO_NUMBER('1322526','9999099') FROM DUAL;
SELECT TO_NUMBER('1322526','99990999999') FROM DUAL;

-- 反面案例
SELECT TO_NUMBER('1322526','999090909999999009') FROM DUAL;
SELECT TO_NUMBER('1322526','99909090999009') FROM DUAL;
SELECT TO_NUMBER('1322526','000000000000000') FROM DUAL;
SELECT TO_NUMBER('1322526','90009000009') FROM DUAL;
SELECT TO_NUMBER('1322526','900000000009') FROM DUAL;
SELECT TO_NUMBER('1322526','00000000009') FROM DUAL;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_NUMBER(1322526,'0099000') FROM DUAL;
SELECT TO_NUMBER(1322526,'0099000') FROM DUAL;
SELECT TO_NUMBER(1322526,'9999099') FROM DUAL;
SELECT TO_NUMBER(1322526,'99990999999') FROM DUAL;

-- 反面案例
SELECT TO_NUMBER(1322526,'999090909999999009') FROM DUAL;
SELECT TO_NUMBER(1322526,'99909090999009') FROM DUAL;
SELECT TO_NUMBER(1322526,'000000000000000') FROM DUAL;
SELECT TO_NUMBER(1322526,'90009000009') FROM DUAL;
SELECT TO_NUMBER(1322526,'900000000009') FROM DUAL;
SELECT TO_NUMBER(1322526,'00000000009') FROM DUAL;

-- 9 替换数字，开头的0替换为空格，0除外
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例

SELECT TO_NUMBER('1322526','99990999999') FROM DUAL;
SELECT TO_NUMBER('1322526','9999999999') FROM DUAL;

-- 反面案例
SELECT TO_NUMBER('1322526','9999') FROM DUAL;
SELECT TO_NUMBER('1322526','99999,99999') FROM DUAL;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_NUMBER(1322526,'99990999999') FROM DUAL;

-- 反面案例
SELECT TO_NUMBER(1322526,'9999') FROM DUAL;


-- $ 美元符号 返回的值以$符号开头
-- B 如果格式对应的输出是0，则返回空白
-- L 本地货币符号 参数 NLS_CURRENCY
-- C 返回指定ISO货币符号 参数 NLS_ISO_CURRENCY
-- U 欧元货币符号 参数 NLS_DUAL_CURRENCY
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER('$12345.678', '$999999.999') from dual;
select TO_NUMBER('123.45', 'B999.99') from dual;
select TO_NUMBER('¥123.45', 'S999.99') from dual;

-- 反面案例
select TO_NUMBER('$12345.678', '$$999999.999') from dual;
select TO_NUMBER('$12345.678', '$999999.999') from dual;
select TO_NUMBER('12345.678', '$99.999') from dual;
select TO_NUMBER('$12345.678', '$99.999') from dual;
select TO_NUMBER('$12345.678', '9$9.999') from dual;
select TO_NUMBER('$12345.678', '99.999$') from dual;
select TO_NUMBER('$12345.678', '99.999') from dual;
select TO_NUMBER('￥123.45', 'B999.99') from dual;

select TO_NUMBER('123.45', 'L999.99') from dual;
select TO_NUMBER('356', 'C999') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 无

-- 反面案例
select TO_NUMBER(1258, 'U9999') from dual;


-- MI 返回以负号结尾的负数 或者 返回以空白结尾的正数
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_NUMBER('485-', '9999MI')	from dual;

-- 反面案例
SELECT TO_NUMBER('-485', '9999MI')	from dual;
SELECT TO_NUMBER('-485', '9999MI')	from dual;
SELECT TO_NUMBER('-485', 'MI9999')	from dual;
SELECT TO_NUMBER('-485', '99MI99')	from dual;
SELECT TO_NUMBER('-485', '99MI99')	from dual;
SELECT TO_NUMBER('485', '9999MIMI')	from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_NUMBER(485, 'FM999MI') from dual;

-- 反面案例
SELECT TO_NUMBER(485, '999MI') from dual;
SELECT TO_NUMBER(-485, '999MI') from dual;
SELECT TO_NUMBER(-485, '999MI')	from dual;
SELECT TO_NUMBER(485, '999MI')	from dual;

-- PR 尖括号包裹 ,如果是正数则返回，负数用尖括号包裹
-- PR格式只能出现在最后
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER('485', '999PR') from dual;

-- 反面案例
select TO_NUMBER('485', 'PR999') from dual;
select TO_NUMBER('-485', '999PR') from dual;
select TO_NUMBER('-485', 'PR999') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER(485, '999PR') from dual;

-- 反面案例
select TO_NUMBER(485, 'PR999') from dual;
select TO_NUMBER(-485, '999PR') from dual;
select TO_NUMBER(-485, 'PR999') from dual;

-- S
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER('-1234567890', '9999999999S') from dual;
select TO_NUMBER('+1234567890', '9999999999S') from dual;
select TO_NUMBER('258-', '999S') from dual;

-- 反面案例
select TO_NUMBER('-258', '9S99') from dual;
select TO_NUMBER('-258', '9SS99') from dual;
select TO_NUMBER('-258', 'S9SS99') from dual;
select TO_NUMBER('1', 'S') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER(-258, 'S999') from dual;
select TO_NUMBER(258, 'S999') from dual;
select TO_NUMBER(-258, '999S') from dual;
select TO_NUMBER(258, '999S') from dual;
-- 反面案例
select TO_NUMBER(-258, '9S99') from dual;
select TO_NUMBER(-258, '9SS99') from dual;
select TO_NUMBER(-258, 'S9SS99') from dual;
select TO_NUMBER(1, 'S') from dual;

-------- a.将 字符串类型参数 转换为 数值类型
-- 组合正面案例
select TO_NUMBER('-258', 'FMSB999') from dual;
select TO_NUMBER('-$258', 'FMS$999') from dual;
select TO_NUMBER('$258-', 'FM$999MI') from dual;

select TO_NUMBER('$2,5,8.36-', 'FM$9,9,9.99MI') from dual;
select TO_NUMBER('$2,5,8.36-', '$9,9,9.99MI') from dual;
select TO_NUMBER('$2,5,8.36-', '$9,9,9.99MI') from dual;
select TO_NUMBER('$2,5,8.36', '$9,9,9.99') from dual;
select TO_NUMBER('$25,8.36', '$9099.99') from dual;
select TO_NUMBER('$2,5,8.36', '$9,99.99') from dual; -- 与预期有差异
select TO_NUMBER('$2,5,8.36', '$9,9,9.99') from dual;

-- 组合反面案例
select TO_NUMBER('$2,5,8.36-', '$9,9,9.99') from dual;
select TO_NUMBER('$2,5,8.36-', '$9,99.99') from dual;
select TO_NUMBER('$2,5,8.36-', '$999.99') from dual;
select TO_NUMBER('$2,5,8.36-', '$99,9.99') from dual;
select TO_NUMBER('$2,58.36-', '$99,9.99') from dual;
select TO_NUMBER('$25,8.36', '$9,99.99') from dual;
select TO_NUMBER('$2,5,8.36', '$9,9,9.99') from dual;
select TO_NUMBER('$258-', 'FM$999PRMI') from dual;
select TO_NUMBER('$258-', 'FMS$999PRMI') from dual;
select TO_NUMBER('$258-', 'S$999PRMI') from dual;
select TO_NUMBER('¥258-', 'SL999PRMI') from dual;
select TO_NUMBER('¥258-', 'FMSL999PRMI') from dual;
select TO_NUMBER('¥258-', 'FMSL999MIPR') from dual;
select TO_NUMBER('¥258-', 'FMSL999MI') from dual;
select TO_NUMBER('¥258-', 'FMSL999PR') from dual;
select TO_NUMBER('258-', 'FMSL999PR') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER(-258, 'S999') from dual;

-- 反面案例
select TO_NUMBER(+258, 'S999') from dual;
select TO_NUMBER(-258, '999PRS') from dual;


------ 科学计数
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER('1', '9999999EEEE') from dual;
select TO_NUMBER('1e2', '9999999EEEE') from dual;
select TO_NUMBER('-1e2', '9999999EEEE') from dual;
select TO_NUMBER('-1e+2', '9999999EEEE') from dual;
select TO_NUMBER('-1e-2', '9999999EEEE') from dual;


-- 反面案例
select TO_NUMBER('+123.45e2', '9999999EEEE') from dual;
select TO_NUMBER('+123.456', '9.9EEEE') from dual;
select TO_NUMBER('+123.456e2', '99999999.99EEEE') from dual;
select TO_NUMBER('+123.456', '9.999EEEE') from dual;
select TO_NUMBER('+123.456', '9.9999EEEE') from dual;
select TO_NUMBER('+123.456', '9.99999EEEE') from dual;
select TO_NUMBER('+1E+123', '9.9EEEE') from dual;
select TO_NUMBER('+123.456', 'FM9.9EEEE') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 反面案例
select TO_NUMBER(1e+2, '9999999EEEE') from dual;
select TO_NUMBER(-1e-2, '9999999EEEE') from dual;
select TO_NUMBER(-1e-2, '9999999EEEE') from dual;

------ 乘积计数
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
select TO_NUMBER('1', '9V9') from dual;
select TO_NUMBER('1', '99V99') from dual;
select TO_NUMBER('1', '9,9V9') from dual;
select TO_NUMBER('1', '99V9') from dual;
select TO_NUMBER('$12', '$99V9') from dual;


-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
select TO_NUMBER(12, '99V9') from dual;
select TO_NUMBER(12, '99V999') from dual;
select TO_NUMBER(12.45, '99V9') from dual;
select TO_NUMBER(1, '9.9V9') from dual;
select TO_NUMBER(1, '9,9V9') from dual;
select TO_NUMBER(1, '9V9') from dual;
select TO_NUMBER(1, '99V99') from dual;
select TO_NUMBER(12, 'FM99V9') from dual;
select TO_NUMBER(12, 'FM99V9MI') from dual;
select TO_NUMBER(12, 'FM99V9S') from dual;
select TO_NUMBER(12, 'FM99V9PR') from dual;
select TO_NUMBER(12, '99V9PR') from dual;
select TO_NUMBER(12, 'S99V9') from dual;
select TO_NUMBER(12, 'S99V9MI') from dual;
select TO_NUMBER(12, '99V9MI') from dual;
select TO_NUMBER(12, 'L99V9') from dual;

------ 十六进制

-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER('17', 'XXXX') from dual;
select TO_NUMBER('017', 'XXXX') from dual;
select TO_NUMBER('17', '0XXXX') from dual;
select TO_NUMBER('0017', 'FMXXXX') from dual;
select TO_NUMBER('17', 'FMXXXX') from dual;
select TO_NUMBER('00017', 'XXXX') from dual;

-- 反面案例
select TO_NUMBER('12', 'LXXXX') from dual;
select TO_NUMBER('$12', 'S$XXXX') from dual;
select TO_NUMBER('1', 'FMLBUSXXXX') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_NUMBER(17, 'XXXX') from dual;
select TO_NUMBER(017, 'XXXX') from dual;
select TO_NUMBER(17, '0XXXX') from dual;
select TO_NUMBER(0017, 'FMXXXX') from dual;
select TO_NUMBER(17, 'FMXXXX') from dual;
select TO_NUMBER(00017, 'XXXX') from dual;

-- 反面案例
select TO_NUMBER(7, 'XXXX9') from dual;
select TO_NUMBER(6, '9XXXX') from dual;
select TO_NUMBER(6, 'X9') from dual;
select TO_NUMBER(6, '9X') from dual;
select TO_NUMBER(6, 'LX') from dual;
select TO_NUMBER(6, 'XS') from dual;
select TO_NUMBER(12, 'FMX') from dual;
select TO_NUMBER(12, 'FMXXXXMI') from dual;
select TO_NUMBER(12, 'FMXXXXS') from dual;
select TO_NUMBER(12, 'XXXXPR') from dual;
select TO_NUMBER(12, 'SXXXX') from dual;
select TO_NUMBER(12, 'XXXXMI') from dual;
select TO_NUMBER(1, 'SX') from dual;
select TO_NUMBER(1, 'FMLBUSXXXXMIPR') from dual;

------ 罗马数字
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
select TO_NUMBER('12', 'LRN') from dual;
select TO_NUMBER('$12', 'S$RN') from dual;
select TO_NUMBER('1', 'RN') from dual;
select TO_NUMBER('11', 'RN') from dual;
select TO_NUMBER('1', 'FMLBUSRN') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
select TO_NUMBER(11, 'RN') from dual;
select TO_NUMBER(1, 'RN') from dual;
select TO_NUMBER(14825, 'RN') from dual;
select TO_NUMBER(1485, 'rn') from dual;
select TO_NUMBER(1485, 'rn') from dual;
select TO_NUMBER(1485, '99999RN') from dual;
select TO_NUMBER(1485, 'LRN') from dual;
select TO_NUMBER(7, 'RN9') from dual;
select TO_NUMBER(6, '9RN') from dual;
select TO_NUMBER(6, 'RN9') from dual;
select TO_NUMBER(6, '9RN') from dual;
select TO_NUMBER(6, 'LRN') from dual;
select TO_NUMBER(6, 'rnS') from dual;
select TO_NUMBER(12, 'FMRN') from dual;
select TO_NUMBER(12, 'FMrn') from dual;
select TO_NUMBER(12, 'FMRNS') from dual;
select TO_NUMBER(12, 'RNPR') from dual;
select TO_NUMBER(12, 'SRN') from dual;
select TO_NUMBER(12, 'RNMI') from dual;
select TO_NUMBER(1, 'SRN') from dual;
select TO_NUMBER(1, 'FMLBUSRNMIPR') from dual;

------ 最小文本 十进制
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
SELECT TO_NUMBER('1234', 'TM') FROM DUAL;
SELECT TO_NUMBER('1234', 'TM9') FROM DUAL;
select TO_NUMBER('12', 'LTM') from dual;
select TO_NUMBER('$12', 'S$TM') from dual;
select TO_NUMBER('1', 'FMLBUSTM') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
SELECT TO_NUMBER(123, 'TM') FROM DUAL;
SELECT TO_NUMBER(1234, 'TM9') FROM DUAL;
SELECT TO_NUMBER(1234, 'TM8') FROM DUAL;
SELECT TO_NUMBER(1234, 'TM99') FROM DUAL;
SELECT TO_NUMBER(1234, '9TM9') FROM DUAL;
SELECT TO_NUMBER(1234, 'LTM9') FROM DUAL;
SELECT TO_NUMBER(1234, '0TM9') FROM DUAL;
SELECT TO_NUMBER(1234, '$TM9') FROM DUAL;
SELECT TO_NUMBER(1234, 'UTM9') FROM DUAL;
select TO_NUMBER(12, 'FMTM') from dual;
select TO_NUMBER(12, 'FMtm') from dual;
select TO_NUMBER(12, 'FMTM9') from dual;
select TO_NUMBER(12, 'FMtM9') from dual;
select TO_NUMBER(12, 'FMTMS') from dual;
select TO_NUMBER(12, 'TMPR') from dual;
select TO_NUMBER(12, 'STM') from dual;
select TO_NUMBER(12, 'TMMI') from dual;
select TO_NUMBER(1, 'STM') from dual;
select TO_NUMBER(1, 'TMTM') from dual;
select TO_NUMBER(1, 'FMLBUSTMMIPR') from dual;

------ 最小文本 科学计数
-------- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
SELECT TO_NUMBER('1234', 'tme') FROM DUAL;
SELECT TO_NUMBER('1234', 'tmetme') FROM DUAL;
select TO_NUMBER('12', 'LTME') from dual;
select TO_NUMBER('$12', 'S$TME') from dual;
select TO_NUMBER('1', 'FMLBUSTME') from dual;

-------- b.将 数值类型参数 转换为 数值类型
-- 正面案例
-- 不支持 to_number

-- 反面案例
SELECT TO_NUMBER(1234, 'TME') FROM DUAL;
SELECT TO_NUMBER(1234, 'TMETME') FROM DUAL;
SELECT TO_NUMBER(1234, 'TMETM') FROM DUAL;
SELECT TO_NUMBER(1234, 'tmE') FROM DUAL;
SELECT TO_NUMBER(1234, 'TMe') FROM DUAL;
SELECT TO_NUMBER(1234, 'tme') FROM DUAL;
SELECT TO_NUMBER(1234, 'TM9e') FROM DUAL;
SELECT TO_NUMBER(1234, 'tm9e') FROM DUAL;
SELECT TO_NUMBER(1234, 'TM8') FROM DUAL;
SELECT TO_NUMBER(1234, 'TM99') FROM DUAL;
SELECT TO_NUMBER(1234, '9TM9') FROM DUAL;
SELECT TO_NUMBER(1234, 'LTM9') FROM DUAL;
SELECT TO_NUMBER(1234, '0TM9') FROM DUAL;
SELECT TO_NUMBER(1234, '$TM9') FROM DUAL;
SELECT TO_NUMBER(1234, 'UTM9') FROM DUAL;
select TO_NUMBER(12, 'FMTME') from dual;
select TO_NUMBER(12, 'FMtme') from dual;
select TO_NUMBER(12, 'FMTM9') from dual;
select TO_NUMBER(12, 'FMTME') from dual;
select TO_NUMBER(12, 'FMTMES') from dual;
select TO_NUMBER(12, 'TMEPR') from dual;
select TO_NUMBER(12, 'STME') from dual;
select TO_NUMBER(1, 'FMLBUSTMEMIPR') from dual;
select TO_NUMBER(12, 'TMEMI') from dual;
select TO_NUMBER(1, 'STME') from dual;
select TO_NUMBER(1, 'STMETME') from dual;
