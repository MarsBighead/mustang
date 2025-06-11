
-- https://leetcode.cn/problems/recyclable-and-low-fat-products/description/?envType=study-plan-v2&envId=sql-free-50
-- 1757. 可回收且低脂的产品
select product_id from Products where low_fats='Y' and recyclable='Y'


-- https://leetcode.cn/problems/find-customer-referee/description/?envType=study-plan-v2&envId=sql-free-50
-- 584. 寻找用户推荐人
select "name" from Customer
where referee_id!=2 or referee_id is null 

-- https://leetcode.cn/problems/big-countries/description/?envType=study-plan-v2&envId=sql-free-50
-- 595. 大的国家
select "name", area, "population" from  world
where area>=3000000 or "population">= 25000000

-- https://leetcode.cn/problems/article-views-i/description/?envType=study-plan-v2&envId=sql-free-50
-- 1148. 文章浏览 I
select distinct author_id as id from Views
where author_id=viewer_id
order by author_id

-- https://leetcode.cn/problems/invalid-tweets/description/?envType=study-plan-v2&envId=sql-free-50
-- 1683. 无效的推文
select tweet_id from Tweets
where length(content)>15


-- Write your PostgreSQL query statement below
-- 1934. 确认率
-- https://leetcode.cn/problems/confirmation-rate/description/?envType=study-plan-v2&envId=sql-free-50
select s.user_id, (case 
  when a.count is not null then round(cast(t.count as numeric)/a.count,2)
  else 0 end) confirmation_rate
from Signups s left join
    (select user_id, count(user_id) count
        from Confirmations 
    where action='confirmed'
    group by user_id) t
on s.user_id=t.user_id left join 
    (select user_id, 
    count(user_id) count
    from Confirmations
    group by user_id) a
on t.user_id=a.user_id

/*对 case when...else end的运用，以及NUM类型转数字的方法*/
select s.user_id, (case 
    when count(t.user_id) = 0 then 0
    else round(SUM(
        CASE WHEN t.action = 'confirmed' THEN 1 ELSE 0 END)::numeric 
        /count(t.user_id),2)
    end) confirmation_rate
from Signups s left join Confirmations t
on s.user_id=t.user_id
group by s.user_id

-- 高级查询和连接
-- https://leetcode.cn/problems/primary-department-for-each-employee/description/?envType=study-plan-v2&envId=sql-free-50
-- 1789. 员工的直属部门
WITH t AS (
    SELECT
        employee_id,
        department_id,
        primary_flag,
        COUNT(employee_id) OVER(PARTITION BY employee_id) AS count_over
        FROM Employee)
SELECT employee_id, department_id
FROM t
where count_over=1 or primary_flag='Y'


-- 1084. 销售分析 III
-- https://leetcode.cn/problems/sales-analysis-iii/description/?envType=study-plan-v2&envId=sql-free-50
select s.product_id, p.product_name
from Product p left join Sales s
USING(product_id)
group by s.product_id,p.product_name
having Min(s.sale_date) >= '2019-01-01' and Max(s.sale_date) <= '2019-03-31';

-- 610. 判断三角形
-- https://leetcode.cn/problems/triangle-judgement/description/?envType=study-plan-v2&envId=sql-free-50
-- Write your PostgreSQL query statement below

select x,y,z,
case when x+y>z and x+z>y and y+z>x then 'Yes'
else 'No' end triangle 
from Triangle


-- https://leetcode.cn/problems/consecutive-numbers/description/?envType=study-plan-v2&envId=sql-free-50
-- 180. 连续出现的数字
-- 第2种可能更快
select distinct l1.Num as ConsecutiveNums
FROM
    Logs l1 left join
    Logs l2 using (Num) left join
    Logs l3 using (Num)
where l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1

select distinct l1.Num as ConsecutiveNums
FROM
    Logs l1 
    join Logs l2 on l2.id = l1.id+1 and l2.Num=l1.Num
    join Logs l3 on l3.id = l2.id+1 and l3.Num=l2.Num


-- 1164. 指定日期的产品价格
-- https://leetcode.cn/studyplan/sql-free-50/
-- Write your PostgreSQL query statement below
select p1.product_id, COALESCE(p2.new_price, 10) price
from (
    select distinct product_id
    from products
    ) as p1
    left join (
        select product_id, new_price 
        from Products
        where (product_id, change_date) in (
            select product_id, max(change_date) change_date
            from Products
            where change_date<='2019-08-16'
            group by product_id)
        ) p2 using(product_id) 
order by p1.product_id 

-- 减少一层子查询
-- Write your PostgreSQL query statement below
select p1.product_id, COALESCE(p2.new_price,10) price
from (select distinct product_id from Products) p1
    left join (
        SELECT 
        product_id, 
        new_price,
        RANK() OVER (
            PARTITION BY product_id 
            ORDER BY change_date DESC
        ) AS rn
    FROM Products
    WHERE change_date <= '2019-08-16'
    ) p2 
    on p1.product_id =p2.product_id and  p2.rn=1

-- 1907. 按分类统计薪水
-- https://leetcode.cn/problems/count-salary-categories/description/?envType=study-plan-v2&envId=sql-free-50
-- Write your PostgreSQL query statement below
-- 条件在Where里面执行，速度更快
select 'Low Salary' AS category,
COALESCE(count(account_id) filter(where income < 20000 ),0) as accounts_count
from Accounts
union 
select 'Average Salary' AS category,
COALESCE(count(account_id) filter(where income >=20000 and income<=50000 ),0) as accounts_count
from Accounts
union
select 'High Salary' AS category,
COALESCE(count(account_id) filter(where income>50000 ),0) as accounts_count
from Accounts


-- Write your PostgreSQL query statement below
SELECT
    'Low Salary' AS category, 
    COUNT(*) AS accounts_count
FROM
    Accounts
WHERE
    income < 20000
UNION
SELECT
    'Average Salary' AS category,
    COUNT(*) AS accounts_count
FROM
    Accounts
WHERE
    20000 <= income AND income <= 50000
UNION
SELECT
    'High Salary' AS category,
    COUNT(*) AS accounts_count
FROM
    Accounts
WHERE
    income > 50000;

-- https://leetcode.cn/problems/last-person-to-fit-in-the-bus/description/?envType=study-plan-v2&envId=sql-free-50 
-- 1204. 最后一个能进入巴士的人
/*
与MySQL不同的是，PostgreSQL 可以通过SUM...OVER实现累加
*/
select 
person_name
from (
    SELECT person_name, SUM(weight) OVER (ORDER BY turn) AS weight
    FROM Queue
) where weight<=1000
order by weight desc 
limit 1

-- MySQL version
SELECT person_name, @pre := @pre + weight AS weight
FROM Queue, (SELECT @pre := 0) tmp
ORDER BY turn


-- 简单查询，dense rank排序
-- 178. 分数排名
-- https://leetcode.cn/problems/rank-scores/description/?envType=problem-list-v2&envId=database
select score, dense_rank() over (order by score desc) rank
from Scores