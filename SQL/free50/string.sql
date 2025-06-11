
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

-- https://leetcode.cn/problems/delete-duplicate-emails/description/?envType=study-plan-v2&envId=sql-free-50
-- 196. 删除重复的电子邮箱
-- Write your PostgreSQL query statement below
delete from Person p1
using Person p2
where p1.email=p2.email and p1.id>p2.id

-- https://leetcode.cn/problems/patients-with-a-condition/description/?envType=study-plan-v2&envId=sql-free-50
-- 1527. 患某种疾病的患者
select * from Patients
where conditions ~ '^DIAB1|\sDIAB1'

-- https://leetcode.cn/problems/group-sold-products-by-the-date/?envType=study-plan-v2&envId=sql-free-50
-- 1484. 按日期分组销售产品
-- PostgreSQL version
/*
PostgreSQL没有MySQL中的GROUP_CONCAT函数，可以用STRING_AGG实现类似功能
*/
SELECT 
    sell_date,
    COUNT(DISTINCT(product)) AS num_sold, 
    STRING_AGG(DISTINCT product, ',') AS products
FROM 
    Activities
GROUP BY 
    sell_date
ORDER BY 
    sell_date ASC

-- MySQL version
SELECT 
    sell_date,
    COUNT(DISTINCT(product)) AS num_sold, 
    GROUP_CONCAT(DISTINCT product ORDER BY product SEPARATOR ',') AS products
FROM 
    Activities
GROUP BY 
    sell_date
ORDER BY 
    sell_date ASC

-- 1327. 列出指定时间段内所有的下单产品
-- https://leetcode.cn/problems/list-the-products-ordered-in-a-period/description/?envType=study-plan-v2&envId=sql-free-50
-- Write your PostgreSQL query statement below
select p.product_name, sum(o.unit) unit
from Orders o left join Products p using(product_id)
where to_char(o.order_date, 'YYYY-MM') ='2020-02' 
group by 1
having sum(o.unit)>=100