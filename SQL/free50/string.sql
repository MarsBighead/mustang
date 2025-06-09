
-- 1517. 查找拥有有效邮箱的用户
-- https://leetcode.cn/problems/find-users-with-valid-e-mails/description/?envType=study-plan-v2&envId=sql-free-50
-- Write your PostgreSQL query statement below
- 使用 ~ 和 ~*（不区分大小写）操作符
SELECT user_id, name, mail
FROM Users
-- 请注意，我们还转义了`@`字符，因为它在某些正则表达式中具有特殊意义
WHERE mail ~ '^[a-zA-Z][a-zA-Z0-9_.-]*\@leetcode\.com$'

-- https://leetcode.cn/problems/fix-names-in-a-table/description/?envType=study-plan-v2&envId=sql-free-50
-- 1667. 修复表中的名字
/*
-- initcap会按照空格将所有单词首字母大写，与题目要求只有第一个字母大写，有差异
select user_id, initcap(name) as name
from Users
order by user_id
*/

select user_id, upper(substring(name,1,1))|| lower(substring(name from 2)) name from Users
order by user_id
