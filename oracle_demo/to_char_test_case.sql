-- TO_CHAR
---- 函数作用
------ a.将 字符串类型参数 转换为 数值类型
------ b.将 数值类型参数 转换为 数值类型
---- 数值模型
------ 十进制
------ 科学计数
------ 乘积计数
------ 十六进制
------ 罗马数字
------ 最小文本 十进制
------ 最小文本 科学计数

---- 日期模型
-------- a.将 日期类型参数 转换为 字符串类型
-------- b.将 时间戳类型参数 转换为 字符串类型
-------- c.将 带时区时间戳类型参数 转换为 字符串类型
-----------------------------------------------------
-----------------------------------------------------

---- 数值模型
-------------------------  十进制  -------------------------
---- a.将 字符串类型参数 转换为 数值类型

-- ,(comma) 半角逗号 装饰作用
-- 不能以逗号开头
-- 不能出现在数字或点号的右边
-- G
-- 正面案例
SELECT TO_CHAR('3450','999,99') from dual;
SELECT TO_CHAR('3450','99999') from dual;
SELECT TO_CHAR('3450','99999,') from dual;
SELECT TO_CHAR('3450','99999') from dual;
SELECT TO_CHAR('3450','99999') from dual;
select TO_CHAR('124548', '99G9G999') from dual;
select TO_CHAR('1245.48', '99G9G9D99') from dual;

-- 反面案例
select TO_CHAR('124548-', '99G9G999S') from dual;
select TO_CHAR('12,4,548-', '99G9G999S') from dual;
select TO_CHAR('12,4,548', '99G9G999') from dual;
select TO_CHAR('12,4,5.48', '99G9G9D99') from dual;
SELECT TO_CHAR('34,50','99999') from dual;
SELECT TO_CHAR('3450',',99999') from dual;
SELECT TO_CHAR('34,50','99999,') from dual;
SELECT TO_CHAR('34,50',',99999') from dual;
SELECT TO_CHAR('34,50','999,99') from dual;
SELECT TO_CHAR('34,50','99999') from dual;
SELECT TO_CHAR('3450','999,99') from dual;
SELECT TO_CHAR('3450','99') from dual;
SELECT TO_CHAR('3450','9G9') from dual;
select TO_CHAR('12,4,548-', '99G9G999S') from dual;
select TO_CHAR('12,4,5.4,8', '99G9G9D9G9') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(3450,'999,99') from dual;
SELECT TO_CHAR(3450,'999G99') from dual;
SELECT TO_CHAR(3450,'99G9G99') from dual;
SELECT TO_CHAR(3450,'9,9,9,99') from dual;

-- 反面案例
SELECT TO_CHAR(3450,'9,9G9G99') from dual;
SELECT TO_CHAR(34,50,'999,99') from dual;
SELECT TO_CHAR(34,50,'999G99') from dual;
SELECT TO_CHAR(34,50,'99G9G99') from dual;
SELECT TO_CHAR(34,50,'99999') from dual;
SELECT TO_CHAR(34,50,'99999,') from dual;
SELECT TO_CHAR(34,50,',99999') from dual;
SELECT TO_CHAR(34,50,'99999') from dual;
SELECT TO_CHAR(3450,'99999') from dual;
select TO_CHAR(-124548, 'S999999') from dual;
SELECT TO_CHAR(3450,'999,99') from dual;
SELECT TO_CHAR(3450,'99') from dual;
SELECT TO_CHAR(3450,'9G9') from dual;
select TO_CHAR(-12,4,548, '99G9G999S') from dual;
select TO_CHAR(-12,4,548, 'S99G9G999S') from dual;
select TO_CHAR(12,4,5.4,8, '99G9G9D9G9') from dual;
select TO_CHAR(-12,4,548, 'S99G9G999') from dual;
select TO_CHAR(124548, '99G9G999') from dual;
select TO_CHAR(1245.48, '99G9G9D99') from dual;

-- .(period) 半角句号,点号，小数点
-- 只能有1个点号
-- G 分组符号
-- 不能出现在数字的右侧
-- 不能和句号.共同出现
-- NUMERIC_CHARACTER

---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR('23.54', '99.99') from dual;
SELECT TO_CHAR('23.54', '99D99') from dual;
select TO_CHAR('12454.8', '99G999D9S') from dual;
select TO_CHAR('12454.8', '99G9G99D9S') from dual;
select TO_CHAR('124548', '99G9G999S') from dual;
select TO_CHAR('124548', '99G9G999') from dual;


-- 反面案例
select TO_CHAR('12,454.8-', '99G999D9S') from dual;
select TO_CHAR('12,4,54.8-', '99G9G99D9S') from dual;
select TO_CHAR('12,4,548-', '99G9G999S') from dual;
select TO_CHAR('12,4,548', '99G9G999') from dual;
SELECT TO_CHAR('23.5.4', '99.9.9') from dual;
SELECT TO_CHAR('23.54', '99.9.9') from dual;
SELECT TO_CHAR('23.5.4', '99D9D9') from dual;
SELECT TO_CHAR('23.54', '99D9D9') from dual;


---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(23.54, '99.99') from dual;
SELECT TO_CHAR(23.54, '99D99') from dual;
SELECT TO_CHAR(23.54, '99D99') from dual;
SELECT TO_CHAR(.54, 'D99') from dual;


-- 反面案例
SELECT TO_CHAR(23.54, '99D99') from dual;
SELECT TO_CHAR(23.54, '99D9D9') from dual;
SELECT TO_CHAR(2.54, 'D999') from dual;

-- 0 以0开头或以0结尾
-- 如果格式模型0对应的数字是0，并且是开头或结尾的0，则占位为0，如果不是0，则保持不变
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例

SELECT TO_CHAR('1','0') FROM DUAL;
SELECT TO_CHAR('1322526','0099000') FROM DUAL;
SELECT TO_CHAR('1322526','0099000') FROM DUAL;
SELECT TO_CHAR('1322526','9999099') FROM DUAL;
SELECT TO_CHAR('1322526','99990999999') FROM DUAL;
SELECT TO_CHAR('1322526','999090909999999009') FROM DUAL;
SELECT TO_CHAR('1322526','99909090999009') FROM DUAL;
SELECT TO_CHAR('1322526','000000000000000') FROM DUAL;
SELECT TO_CHAR('1322526','90009000009') FROM DUAL;
SELECT TO_CHAR('1322526','900000000009') FROM DUAL;
SELECT TO_CHAR('1322526','00000000009') FROM DUAL;

-- 反面案例
SELECT TO_CHAR('1322526','00') FROM DUAL;
SELECT TO_CHAR('1322526','99') FROM DUAL;
SELECT TO_CHAR('1322526','99.9') FROM DUAL;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(1,'0') FROM DUAL;
SELECT TO_CHAR(1322526,'0099000') FROM DUAL;
SELECT TO_CHAR(1322526,'0099000') FROM DUAL;
SELECT TO_CHAR(1322526,'9999099') FROM DUAL;
SELECT TO_CHAR(1322526,'99990999999') FROM DUAL;

-- 反面案例
SELECT TO_CHAR(1322526,'999090909999999009') FROM DUAL;
SELECT TO_CHAR(1322526,'99909090999009') FROM DUAL;
SELECT TO_CHAR(1322526,'000000000000000') FROM DUAL;
SELECT TO_CHAR(1322526,'90009000009') FROM DUAL;
SELECT TO_CHAR(1322526,'900000000009') FROM DUAL;
SELECT TO_CHAR(1322526,'00000000009') FROM DUAL;

-- 9 替换数字，开头的0替换为空格，0除外
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR('1322526','99990999999') FROM DUAL;
SELECT TO_CHAR('1322526','9999999999') FROM DUAL;
SELECT TO_CHAR('1322526','99999,99999') FROM DUAL;

-- 反面案例
SELECT TO_CHAR('1322526','9999') FROM DUAL;
SELECT TO_CHAR('1322526','966') FROM DUAL;
SELECT TO_CHAR('1322526','66') FROM DUAL;
SELECT TO_CHAR('1322526','6') FROM DUAL;
SELECT TO_CHAR('1322526','666666666') FROM DUAL;


---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(1322526,'99990999999') FROM DUAL;

-- 反面案例
SELECT TO_CHAR(1322526,'9999') FROM DUAL;


-- $ 美元符号 返回的值以$符号开头
-- B 如果格式对应的输出是0，则返回空白
-- L 本地货币符号 参数 NLS_CURRENCY
-- C 返回指定ISO货币符号 参数 NLS_ISO_CURRENCY
-- U 欧元货币符号 参数 NLS_DUAL_CURRENCY
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('123.45', 'B999.99') from dual;
select TO_CHAR('12345.678', '$999999.999') from dual;
select TO_CHAR('123.45', 'L999.99') from dual;
select TO_CHAR('123.45', 'U999.99') from dual;
select TO_CHAR('123.45', 'U99999.990') from dual;
select TO_CHAR('356', 'C999') from dual;

-- 反面案例
select TO_CHAR('¥123.45', 'S999.99') from dual;
select TO_CHAR('$12345.678', '$999999.999') from dual;
select TO_CHAR('$12345.678', '$$999999.999') from dual;
select TO_CHAR('$12345.678', '$999999.999') from dual;
select TO_CHAR('$12345.678', '$99.999') from dual;
select TO_CHAR('$12345.678', '9$9.999') from dual;
select TO_CHAR('$12345.678', '99.999$') from dual;
select TO_CHAR('$12345.678', '99.999') from dual;
select TO_CHAR('￥123.45', 'B999.99') from dual;


---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(123.45, 'B999.99') from dual;
select TO_CHAR(12345.678, '$999999.999') from dual;
select TO_CHAR(123.45, 'L999.99') from dual;
select TO_CHAR(123.45, 'U999.99') from dual;
select TO_CHAR(123.45, 'U99999.990') from dual;
select TO_CHAR(356, 'C999') from dual;
select TO_CHAR(1258, '9999C') from dual;
select TO_CHAR(1258, '999999C99') from dual;
select TO_CHAR(1258, '9999U') from dual;
select TO_CHAR(1258, '9999B') from dual;
select TO_CHAR(1258, '9999L') from dual;
select TO_CHAR(1258, '999B9L') from dual;
select TO_CHAR(1258, '99B99') from dual;
select TO_CHAR(1258.345, '9999L99') from dual;
select TO_CHAR(1258.235, '9999U99') from dual;
select TO_CHAR(1258.235, '9999C9999') from dual;
select TO_CHAR(1258.235, '9999B9999') from dual;
select TO_CHAR(1258.235, '9999$9999') from dual;

-- 反面案例
select TO_CHAR(1258, '99U99U') from dual;
select TO_CHAR(1258, '99B99B') from dual;
select TO_CHAR(1258, '99B9B9L') from dual;
select TO_CHAR(1258, '9999CL') from dual;
select TO_CHAR(1258, '9999C99C99') from dual;

-- MI 返回以负号结尾的负数 或者 返回以空白结尾的正数
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR('485', '9999MI')	from dual;
SELECT TO_CHAR('-485', '9999MI')	from dual;

-- 反面案例
SELECT TO_CHAR('485-', '9999MI')	from dual;
SELECT TO_CHAR('-485', '9999MI')	from dual;
SELECT TO_CHAR('-485', '9999MI')	from dual;
SELECT TO_CHAR('-485', 'MI9999')	from dual;
SELECT TO_CHAR('-485', '99MI99')	from dual;
SELECT TO_CHAR('-485', '99MI99')	from dual;
SELECT TO_CHAR('485', '9999MIMI')	from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(485, 'FM999MI') from dual;

-- 反面案例
SELECT TO_CHAR(485, '999MI') from dual;
SELECT TO_CHAR(-485, '999MI') from dual;
SELECT TO_CHAR(-485, '999MI')	from dual;
SELECT TO_CHAR(485, '999MI')	from dual;

-- PR 尖括号包裹 ,如果是正数则返回，负数用尖括号包裹
-- PR格式只能出现在最后
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('485', '999PR') from dual;
select TO_CHAR('-485', '999PR') from dual;

-- 反面案例
select TO_CHAR('485', 'PR999') from dual;
select TO_CHAR('-485', 'PR999') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(485, '999PR') from dual;
select TO_CHAR(-485, '999PR') from dual;

-- 反面案例
select TO_CHAR(485, 'PR999') from dual;
select TO_CHAR(-485, 'PR999') from dual;

-- S
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('-1234567890', '9999999999S') from dual;
select TO_CHAR('+1234567890', '9999999999S') from dual;


-- 反面案例
select TO_CHAR('258-', '999S') from dual;
select TO_CHAR('-258', '9S99') from dual;
select TO_CHAR('-258', '9SS99') from dual;
select TO_CHAR('-258', 'S9SS99') from dual;
select TO_CHAR('1', 'S') from dual;
select TO_CHAR('1', 'SS') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(-258, 'S999') from dual;
select TO_CHAR(258, 'S999') from dual;
select TO_CHAR(-258, '999S') from dual;
select TO_CHAR(258, '999S') from dual;

-- 反面案例
select TO_CHAR(-258, '9S99') from dual;
select TO_CHAR(-258, '9SS99') from dual;
select TO_CHAR(-258, 'S9SS99') from dual;
select TO_CHAR(1, 'S') from dual;

---- a.将 字符串类型参数 转换为 数值类型
-- 组合正面案例
select TO_CHAR('-258', 'FMSB999') from dual;
select TO_CHAR('258', 'FMSL999') from dual;
select TO_CHAR('258', 'FML999S') from dual;


-- 组合反面案例
select TO_CHAR('258', 'FML999SPR') from dual;
select TO_CHAR('$258-', 'FM$999MI') from dual;
select TO_CHAR('$2,5,8.36-', '$9,9,9.99') from dual;
select TO_CHAR('$2,5,8.36-', '$9,99.99') from dual;
select TO_CHAR('$2,5,8.36-', '$999.99') from dual;
select TO_CHAR('$2,5,8.36-', '$99,9.99') from dual;
select TO_CHAR('$2,58.36-', '$99,9.99') from dual;
select TO_CHAR('$25,8.36', '$9,99.99') from dual;
select TO_CHAR('$2,5,8.36', '$9,9,9.99') from dual;
select TO_CHAR('$258-', 'FM$999PRMI') from dual;
select TO_CHAR('$258-', 'FMS$999PRMI') from dual;
select TO_CHAR('$258-', 'S$999PRMI') from dual;
select TO_CHAR('¥258-', 'SL999PRMI') from dual;
select TO_CHAR('¥258-', 'FMSL999PRMI') from dual;
select TO_CHAR('¥258-', 'FMSL999MIPR') from dual;
select TO_CHAR('¥258-', 'FMSL999MI') from dual;
select TO_CHAR('¥258-', 'FMSL999PR') from dual;
select TO_CHAR('258-', 'FMSL999PR') from dual;
select TO_CHAR('-$258', 'FMS$999') from dual;
select TO_CHAR('$2,5,8.36-', 'FM$9,9,9.99MI') from dual;
select TO_CHAR('$2,5,8.36-', '$9,9,9.99MI') from dual;
select TO_CHAR('$2,5,8.36-', '$9,9,9.99MI') from dual;
select TO_CHAR('$2,5,8.36', '$9,9,9.99') from dual;
select TO_CHAR('$25,8.36', '$9099.99') from dual;
select TO_CHAR('$2,5,8.36', '$9,99.99') from dual; -- 与预期有差异
select TO_CHAR('$2,5,8.36', '$9,9,9.99') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(-258, 'S999') from dual;

-- 反面案例
select TO_CHAR(+258, 'S999') from dual;
select TO_CHAR(-258, '999PRS') from dual;


-------------------------  科学计数  -------------------------
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('1', '9999999EEEE') from dual;
select TO_CHAR('1e2', '9999999EEEE') from dual;
select TO_CHAR('-1e2', '9999999EEEE') from dual;
select TO_CHAR('-1e+2', '9999999EEEE') from dual;
select TO_CHAR('-1e-2', '9999999EEEE') from dual;


-- 反面案例
select TO_CHAR('+123.45e2', '9999999EEEE') from dual;
select TO_CHAR('+123.456', '9.9EEEE') from dual;
select TO_CHAR('+123.456e2', '99999999.99EEEE') from dual;
select TO_CHAR('+123.456', '9.999EEEE') from dual;
select TO_CHAR('+123.456', '9.9999EEEE') from dual;
select TO_CHAR('+123.456', '9.99999EEEE') from dual;
select TO_CHAR('+1E+123', '9.9EEEE') from dual;
select TO_CHAR('+123.456', 'FM9.9EEEE') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(1e+2, '9999999EEEE') from dual;
select TO_CHAR(-1e-2, '9999999EEEE') from dual;

-- 反面案例
select TO_CHAR(-1e-2, '9999999EEEE') from dual;

-------------------------  乘积计数  -------------------------
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('1', '9V9') from dual;
select TO_CHAR('1', '99V99') from dual;
select TO_CHAR('1', '9,9V9') from dual;
select TO_CHAR('1', '99V9') from dual;

-- 反面案例
select TO_CHAR('$12', '$99V9') from dual;


---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(12, '99V9') from dual;
select TO_CHAR(12, '99V999') from dual;
select TO_CHAR(12.45, '99V9') from dual;
select TO_CHAR(1, '9,9V9') from dual;
select TO_CHAR(1, '9V9') from dual;
select TO_CHAR(1, '99V99') from dual;
select TO_CHAR(12, 'FM99V9') from dual;
select TO_CHAR(12, 'FM99V9MI') from dual;
select TO_CHAR(12, 'FM99V9S') from dual;
select TO_CHAR(12, 'FM99V9PR') from dual;
select TO_CHAR(12, '99V9PR') from dual;
select TO_CHAR(12, 'S99V9') from dual;
select TO_CHAR(12, '99V9MI') from dual;
select TO_CHAR(12, 'L99V9') from dual;

-- 反面案例
select TO_CHAR(12, 'S99V9MI') from dual;
select TO_CHAR(1, '9.9V9') from dual;

-------------------------  十六进制  -------------------------

---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('17', 'XXXX') from dual;
select TO_CHAR('017', 'XXXX') from dual;
select TO_CHAR('17', '0XXXX') from dual;
select TO_CHAR('0017', 'FMXXXX') from dual;
select TO_CHAR('17', 'FMXXXX') from dual;
select TO_CHAR('00017', 'XXXX') from dual;

-- 反面案例
select TO_CHAR('12', 'LXXXX') from dual;
select TO_CHAR('$12', 'S$XXXX') from dual;
select TO_CHAR('1', 'FMLBUSXXXX') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(17, 'XXXX') from dual;
select TO_CHAR(017, 'XXXX') from dual;
select TO_CHAR(17, '0XXXX') from dual;
select TO_CHAR(0017, 'FMXXXX') from dual;
select TO_CHAR(17, 'FMXXXX') from dual;
select TO_CHAR(00017, 'XXXX') from dual;

-- 反面案例
select TO_CHAR(7, 'XXXX9') from dual;
select TO_CHAR(6, '9XXXX') from dual;
select TO_CHAR(6, 'X9') from dual;
select TO_CHAR(6, '9X') from dual;
select TO_CHAR(6, 'LX') from dual;
select TO_CHAR(6, 'XS') from dual;
select TO_CHAR(12, 'FMX') from dual;
select TO_CHAR(12, 'FMXXXXMI') from dual;
select TO_CHAR(12, 'FMXXXXS') from dual;
select TO_CHAR(12, 'XXXXPR') from dual;
select TO_CHAR(12, 'SXXXX') from dual;
select TO_CHAR(12, 'XXXXMI') from dual;
select TO_CHAR(1, 'SX') from dual;
select TO_CHAR(1, 'FMLBUSXXXXMIPR') from dual;

-------------------------  罗马数字  -------------------------
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR('1', 'RN') from dual;
select TO_CHAR('11', 'RN') from dual;

-- 反面案例
select TO_CHAR('12', 'LRN') from dual;
select TO_CHAR('$12', 'S$RN') from dual;
select TO_CHAR('1', 'FMLBUSRN') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
select TO_CHAR(11, 'RN') from dual;
select TO_CHAR(1, 'RN') from dual;
select TO_CHAR(14825, 'RN') from dual;
select TO_CHAR(1485, 'rn') from dual;
select TO_CHAR(12, 'FMRN') from dual;
select TO_CHAR(12, 'FMrn') from dual;
select TO_CHAR(12, 'SRN') from dual;
select TO_CHAR(1, 'SRN') from dual;


-- 反面案例
select TO_CHAR(-12, 'rn') from dual;
select TO_CHAR(1485, '99999RN') from dual;
select TO_CHAR(1485, 'LRN') from dual;
select TO_CHAR(7, 'RN9') from dual;
select TO_CHAR(6, '9RN') from dual;
select TO_CHAR(6, 'RN9') from dual;
select TO_CHAR(6, '9RN') from dual;
select TO_CHAR(6, 'LRN') from dual;
select TO_CHAR(6, 'rnS') from dual;
select TO_CHAR(12, 'FMRNS') from dual;
select TO_CHAR(12, 'RNPR') from dual;
select TO_CHAR(12, 'RNMI') from dual;
select TO_CHAR(1, 'RNPR') from dual;
select TO_CHAR(1, 'RNMI') from dual;
select TO_CHAR(-1, 'RNMI') from dual;
select TO_CHAR(1, 'FMLBUSRNMIPR') from dual;

-------------------------  最小文本 十进制  -------------------------
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR('1234', 'TM') FROM DUAL;
SELECT TO_CHAR('1234', 'TM9') FROM DUAL;

-- 反面案例
select TO_CHAR('12', 'LTM') from dual;
select TO_CHAR('$12', 'S$TM') from dual;
select TO_CHAR('1', 'FMLBUSTM') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(123, 'TM') FROM DUAL;
SELECT TO_CHAR(1234, 'TM9') FROM DUAL;
select TO_CHAR(12, 'STM') from dual;


-- 反面案例
select TO_CHAR(12, 'BTM') from dual;
SELECT TO_CHAR(1234, 'TM8') FROM DUAL;
SELECT TO_CHAR(1234, 'TM99') FROM DUAL;
SELECT TO_CHAR(1234, '9TM9') FROM DUAL;
SELECT TO_CHAR(1234, 'LTM9') FROM DUAL;
SELECT TO_CHAR(1234, '0TM9') FROM DUAL;
SELECT TO_CHAR(1234, '$TM9') FROM DUAL;
SELECT TO_CHAR(1234, 'UTM9') FROM DUAL;
select TO_CHAR(12, 'FMTM') from dual;
select TO_CHAR(12, 'FMtm') from dual;
select TO_CHAR(12, 'FMTM9') from dual;
select TO_CHAR(12, 'FMtM9') from dual;
select TO_CHAR(12, 'FMTMS') from dual;
select TO_CHAR(12, 'TMPR') from dual;
select TO_CHAR(12, 'TMMI') from dual;
select TO_CHAR(1, 'STM') from dual;
select TO_CHAR(1, 'TMTM') from dual;
select TO_CHAR(1, 'FMLBUSTMMIPR') from dual;

-------------------------  最小文本 科学计数  -------------------------
---- a.将 字符串类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR('1234', 'tme') FROM DUAL;
select TO_CHAR('12', 'STME') from dual;
select TO_CHAR('12', 'FMSTME') from dual;


-- 反面案例
select TO_CHAR('12', 'TMES') from dual;
SELECT TO_CHAR('1234', 'tmetme') FROM DUAL;
select TO_CHAR('12', 'LTME') from dual;
select TO_CHAR('$12', 'S$TME') from dual;
select TO_CHAR('1', 'FMLBUSTME') from dual;

---- b.将 数值类型参数 转换为 数值类型
-- 正面案例
SELECT TO_CHAR(1234, 'TME') FROM DUAL;
SELECT TO_CHAR(1234, 'tmE') FROM DUAL;
SELECT TO_CHAR(1234, 'TMe') FROM DUAL;
SELECT TO_CHAR(1234, 'tme') FROM DUAL;
select TO_CHAR(12, 'STME') from dual;
select TO_CHAR(1, 'STME') from dual;

-- 反面案例
SELECT TO_CHAR(1234, 'TMETME') FROM DUAL;
SELECT TO_CHAR(1234, 'TMETM') FROM DUAL;
SELECT TO_CHAR(1234, 'TM9e') FROM DUAL;
SELECT TO_CHAR(1234, 'tm9e') FROM DUAL;
SELECT TO_CHAR(1234, 'TM8') FROM DUAL;
SELECT TO_CHAR(1234, 'TM99') FROM DUAL;
SELECT TO_CHAR(1234, '9TM9') FROM DUAL;
SELECT TO_CHAR(1234, 'LTM9') FROM DUAL;
SELECT TO_CHAR(1234, '0TM9') FROM DUAL;
SELECT TO_CHAR(1234, '$TM9') FROM DUAL;
SELECT TO_CHAR(1234, 'UTM9') FROM DUAL;
select TO_CHAR(12, 'FMTME') from dual;
select TO_CHAR(12, 'FMtme') from dual;
select TO_CHAR(12, 'FMTM9') from dual;
select TO_CHAR(12, 'FMTME') from dual;
select TO_CHAR(12, 'FMTMES') from dual;
select TO_CHAR(12, 'TMEPR') from dual;
select TO_CHAR(1, 'FMLBUSTMEMIPR') from dual;
select TO_CHAR(12, 'TMEMI') from dual;
select TO_CHAR(1, 'STMETME') from dual;



---- c.将 日期类型参数 转换为 字符串类型
---- d.将 时间戳类型参数 转换为 字符串类型
---- e.将 带时区时间戳类型参数 转换为 字符串类型

-- - 结果中会复制标点符号和引用文本。
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY-YYYY') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY-YYYY') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY-YYYY') From dual;

-- / 同上
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY////MM//DD') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY////MM//DD') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY////MM//DD') From dual;

-- , 同上
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY,,,,,MM//DD') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY,,,,,MM//DD') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY,,,,,MM//DD') From dual;

-- . 同上
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY....MM..DD') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY....MM..DD') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY....MM..DD') From dual;

-- ; 同上
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY;;;;MM..DD') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY;;;;MM..DD') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY;;;;MM..DD') From dual;

-- : 同上
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY::::MM..DD') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY::::MM..DD') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY::::MM..DD') From dual;

-- "text"
Select TO_CHAR(to_date( '2023-10-29 01:30:56','YYYY-MM-DD HH:MI:SS'),'YYYY"abcd"MM,,DD') From dual;
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY"abcd"MM,,DD') From dual;
Select TO_CHAR(to_timestamp_tz('2023-10-29 01:30:56.321654789', 'YYYY-MM-DD HH24:MI:SS.FF9'),'YYYY"abcd"MM,,DD') From dual;

SELECT TO_CHAR('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_CHAR('2008abcd05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_CHAR('2008xxxx05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_CHAR('2008"xxxx"05,,20','YYYY"abcd"MM,,DD') FROM DUAL;
SELECT TO_CHAR('2008xxxx05,,20','YYYYabcdMM,,DD') FROM DUAL;

-- AD 公元后
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','AD yyyy-mm-dd hh:mm:ss') From dual;

-- A.D.  同上
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','A.D. yyyy-mm-dd hh:mm:ss') From dual;

-- AM 上午
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','AM yyyy-mm-dd hh:mm:ss') From dual;

-- A.M. 同上
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','A.M. yyyy-mm-dd hh:mm:ss') From dual;

-- BC 公元前
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','BC yyyy-mm-dd hh:mm:ss') From dual;

-- B.C. 同上
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','B.C. yyyy-mm-dd hh:mm:ss') From dual;

-- CC 世纪
/*
如果最后2位数字是01-99，则返回世纪数位前两位数+1
如果最后2位数字是00，则返回世纪数位前两位数
*/
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','CC') from dual;

-- SCC 同上
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','SCC') from dual;
SELECT TO_CHAR(date'2000-01-01','SCC') from dual;
SELECT TO_CHAR(date'2001-01-01','SCC') from dual;

-- D 星期的第几天
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','D') from dual;

-- DAY 天的名称
SELECT TO_CHAR(current_timestamp, 'Day, DD HH12:MI:SS') FROM DUAL;
SELECT TO_CHAR(current_timestamp, 'FMDay, FMDD HH12:MI:SS') FROM DUAL;

-- DD 月的第几天
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','dd')       from dual;

-- DDD 年的第几天
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','ddd')       from dual;

-- DL 长日期格式
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','DL') from dual;

-- DS 短日期格式
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','DS') from dual;

-- DY 天的缩写
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','DY') from dual;

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
select TO_CHAR('0207','MM/YY') from dual;
select TO_CHAR('0207','MM/YY') from dual;
select TO_CHAR('02#07','MM/YY') from dual;
select TO_CHAR('02#07','FMMM/YY') from dual;
select TO_CHAR('02#07','FXMM/YY') from dual; -- error

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
SELECT TO_CHAR('15/ JAN /1998', 'DD-MON-YYYY') FROM DUAL; -- Match
SELECT TO_CHAR(' 15! JAN % /1998', 'DD-MON-YYYY') FROM DUAL; -- Error
SELECT TO_CHAR('15-JAN-1998', 'FXDD-MON-YYYY') FROM DUAL; -- Match
SELECT TO_CHAR('15/JAN/1998', 'FXDD-MON-YYYY') FROM DUAL; -- Error


SELECT TO_CHAR('01-01-1998' ,'FXDD-MM-YYYY') FROM DUAL; -- Match
SELECT TO_CHAR('1-1-1998' ,'FXFMDD-MM-YYYY') FROM DUAL; -- Match
SELECT TO_CHAR('1-01-1998' ,'FXDD-MON-YYYY') FROM DUAL; -- Error

select TO_CHAR('02#07','FXMM#YY') from dual;
select TO_CHAR('02#07','FXMM/YY') from dual;

-- HH
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','hh')   from dual;

-- HH12
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','hh12')   from dual;

-- HH24
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','hh24')   from dual;


-- IW 年的第几个星期 ISO8601
/*
日历星期从周一开始
第1个日历星期包含1月4号
第1个日历星期可能包含12月29，30，31
最后一个日历星期可能包含1月1，2，3
*/
select TO_CHAR(TIMESTAMP '2023-01-01 09:26:50.124','IW') from dual;
select TO_CHAR(TIMESTAMP '2023-01-02 09:26:50.124','IW') from dual;
select TO_CHAR(TIMESTAMP '2023-01-03 09:26:50.124','IW') from dual;
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','IW') from dual;

-- IYYY 4位数字的年，包含日历星期 ISO8601
select TO_CHAR(TIMESTAMP '2023-01-01 09:26:50.124','IYYY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-02 09:26:50.124','IYYY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-03 09:26:50.124','IYYY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','IYYY') from dual;

-- IYY 最后3个数字的年，包含日历星期 ISO8601
select TO_CHAR(TIMESTAMP '2023-01-01 09:26:50.124','IYY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-02 09:26:50.124','IYY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-03 09:26:50.124','IYY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','IYY') from dual;

-- IY 最后2个数字的年，包含日历星期 ISO8601
select TO_CHAR(TIMESTAMP '2023-01-01 09:26:50.124','IY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-02 09:26:50.124','IY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-03 09:26:50.124','IY') from dual;
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','IY') from dual;

-- I 1个数字的年，包含日历星期 ISO8601
select TO_CHAR(TIMESTAMP '2023-01-01 09:26:50.124','I') from dual;
select TO_CHAR(TIMESTAMP '2023-01-02 09:26:50.124','I') from dual;
select TO_CHAR(TIMESTAMP '2023-01-03 09:26:50.124','I') from dual;
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','I') from dual;

-- J  julian calendar 从1, 4712 BC开始，必须是整数
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','J')   from dual;


-- MI 分钟 0-59
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','mi') from dual;


-- MM 月 01-12
SELECT TO_CHAR('2008 05 20','YYYY MM DD') FROM DUAL;
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','MM') from dual;


-- MON 月的缩写，同MONTH
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','MON') from dual;

-- MONTH 月的名称
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','MONTH') from dual;

-- PM
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY-MM-DD PM hh24:mi:ss') From dual;

-- P.M. 同上
Select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','YYYY-MM-DD P.M. hh24:mi:ss') From dual;

-- Q 季节 1-4
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','Q') from dual;

-- RM 罗马数字的月
select TO_CHAR(TIMESTAMP '2023-01-04 09:26:50.124','RM') from dual;

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
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','ss')      from dual;
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','SS') FROM DUAL;


-- SSSSS 0-86399
-- 午夜过后的秒
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.476589 US/Pacific PDT','SSSSS') FROM DUAL;
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.76589 US/Pacific PDT','SSSSS') FROM DUAL;

-- TS
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TS') FROM DUAL;

-- TZD
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TZD') FROM DUAL;

-- TZH
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TZH') FROM DUAL;

-- TZM
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:00 US/Pacific PDT','TZM') FROM DUAL;

-- TZR
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','TZR') FROM DUAL;

-- WW 年的第几个星期 1-53
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','WW') FROM DUAL;

-- W 月的第几周取值范围1-5。星期从1开始，7结束。
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','WW') FROM DUAL;

-- X 本地基数字符 小数点
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','HH:MI:SSXFF') FROM DUAL;

-- Y,YYY 逗号分隔的固定格式表示的年
SELECT TO_CHAR(TIMESTAMP '1999-10-29 01:30:56.321654789 US/Pacific PDT','Y,YYY') FROM DUAL;

-- YEAR 年,单词拼写方式
SELECT TO_CHAR(date'2000-01-01','YEAR') FROM DUAL;
SELECT TO_CHAR(date'-2000-01-01','YEAR') FROM DUAL;
SELECT TO_CHAR(date'2000-01-01','AD YEAR') FROM DUAL;
SELECT TO_CHAR(date'-2000-01-01','AD YEAR') FROM DUAL;


-- SYEAR 年,单词拼写方式。对年添加标记，公元前则加负号-，公元后则是空格
SELECT TO_CHAR(date'2000-01-01','SYEAR') FROM DUAL;
SELECT TO_CHAR(date'-2000-01-01','SYEAR') FROM DUAL;
SELECT TO_CHAR(date'2000-01-01','AD SYEAR') FROM DUAL;
SELECT TO_CHAR(date'-2000-01-01','AD SYEAR') FROM DUAL;

-- YYYY 4位数字的年
SELECT TO_CHAR(TO_DATE('27-OCT-98', 'DD-MON-RR'), 'YYYY') "Year" FROM DUAL;
SELECT TO_CHAR(TO_DATE('27-OCT-17', 'DD-MON-RR'), 'YYYY') "Year" FROM DUAL;
select TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','yyyy')    from dual;

-- SYYYY 4位数字的年。对年添加标记，公元前则加负号-，公元后则是空格
SELECT TO_CHAR(TIMESTAMP '1987-10-29 01:30:56.321654789','SYYYY') FROM DUAL;

-- YYY 最后3位数字的年
SELECT TO_CHAR(TIMESTAMP '1987-10-29 01:30:56.321654789','YYY') FROM DUAL;


-- YY 最后2位数字的年
SELECT TO_CHAR(TIMESTAMP '1987-10-29 01:30:56.321654789','YY') FROM DUAL;
SELECT TO_CHAR(TO_DATE('0207','MM/YY'), 'MM/YY') FROM DUAL;
SELECT TO_CHAR (TO_DATE('02#07','MM/YY'), 'MM/YY') FROM DUAL;


-- Y 最后1位数字的年
SELECT TO_CHAR(TO_DATE('027','MM/Y'), 'MM/Y') FROM DUAL;
SELECT TO_CHAR(TIMESTAMP '1987-10-29 01:30:56.321654789 ','Y') FROM DUAL;

-- 日期格式后缀
-- TH 序号,英文序数词描述
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','DDTH') FROM DUAL;


-- SP 拼写,英文拼写方式描述
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','DDSP') FROM DUAL;

-- SPTH 英文拼写方式描述+英文序数词描述
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789','DDTHSP') FROM DUAL;

-- THSP 英文序数词描述+英文拼写方式描述
SELECT TO_CHAR(TIMESTAMP '2023-10-29 01:30:56.321654789 ','DDSPTH') FROM DUAL;





