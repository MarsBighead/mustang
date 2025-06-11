-- 620. 有趣的电影
-- https://leetcode.cn/problems/not-boring-movies/description/?envType=study-plan-v2&envId=sql-free-50
select id, movie, description, rating
from cinema
where description!='boring' and id%2=1
order by rating desc

- mod()为聚合函数
select id, movie, description, rating
from cinema
where description!='boring' and mod(id,2)=1
order by rating desc

-- https://leetcode.cn/problems/average-selling-price/description/?envType=study-plan-v2&envId=sql-free-50
-- 1251. 平均售价
select p.product_id, 
(case when sum(u.units) is null then 0
else round(sum(u.units*p.price)::numeric/sum(u.units), 2) end)average_price
from Prices p left join UnitsSold u 
on u.product_id=p.product_id
and (u.purchase_date between p.start_date
and p.end_date)
group by p.product_id

-- 最稳定，但是sum(u.units)存在为0可能，导致除数为0风险
select p.product_id, 
COALESCE(round(sum(u.units*p.price)::numeric/sum(u.units), 2), 0) average_price
from Prices p left join UnitsSold u 
on u.product_id=p.product_id
and (u.purchase_date between p.start_date
and p.end_date)
group by p.product_id

select p.product_id, 
(case when 
sum(COALESCE(u.units,0))=0 then 0
else 
round(sum(COALESCE(u.units,0)*p.price)::numeric/
    sum(COALESCE(u.units,0)),2) end) average_price
from Prices p left join UnitsSold u 
on u.product_id=p.product_id
and (u.purchase_date between p.start_date
and p.end_date)
group by p.product_id

-- 1075. 项目员工 I
-- https://leetcode.cn/problems/project-employees-i/description/?envType=study-plan-v2&envId=sql-free-50
select p.project_id, round(avg(e.experience_years)::numeric,2) average_years
from Project p left join Employee e
on p.employee_id = e.employee_id
group by p.project_id

-- 元素获取时进行类型转换，会比聚合函数中进行类型转换速度快
select p.project_id, round(avg(e.experience_years*1.00)::numeric,2) average_years
from Project p left join Employee e
on p.employee_id = e.employee_id
group by p.project_id

-- 1633. 各赛事的用户注册率
-- https://leetcode.cn/problems/percentage-of-users-attended-a-contest/description/?envType=study-plan-v2&envId=sql-free-50
select contest_id, 
    round(count(user_id)::numeric*100
          /(select count(user_id) as n from Users),2)
    as percentage
from Register 
group by  1
order by 2 desc, 1 asc


-- https://leetcode.cn/problems/queries-quality-and-percentage/description/?envType=study-plan-v2&envId=sql-free-50
-- 1211. 查询结果的质量和占比
-- Write your PostgreSQL query statement below
SELECT 
    query_name, 
    round(avg(rating::numeric/position),2) quality,
    ROUND(SUM(CASE WHEN rating < 3 then 100.00 else 0.00 end) / COUNT(rating), 2) poor_query_percentage
FROM Queries
GROUP BY query_name

-- https://leetcode.cn/problems/monthly-transactions-i/submissions/635946300/?envType=study-plan-v2&envId=sql-free-50
-- 1193. 每月交易 I
select to_char(trans_date, 'YYYY-MM') as "month",
country, count(id) as trans_count,
sum(case when state='approved' then 1 else 0 end) approved_count,
sum(amount) trans_total_amount,
sum(case when state='approved' then amount else 0 end) approved_total_amount
from Transactions
group by 1,2

select to_char(trans_date, 'YYYY-MM') as "month",
country, count(id) as trans_count,
count(id) filter (where state='approved') as approved_count,
sum(amount) trans_total_amount,
COALESCE(sum(amount) filter (where state='approved'), 0) approved_total_amount
from Transactions
group by 1,2


-- 1174. 即时食物配送 II
-- https://leetcode.cn/studyplan/sql-free-50/
select round((count(customer_id) filter(where order_date=customer_pref_delivery_date))*100.0/count(delivery_id),2)
as immediate_percentage
from Delivery
where (customer_id, order_date) in (
     select customer_id, min(order_date)
    from delivery
    group by customer_id
)

-- 550. 游戏玩法分析 IV
-- https://leetcode.cn/problems/game-play-analysis-iv/description/?envType=study-plan-v2&envId=sql-free-50
select round((count(a.player_id) filter(where a.event_date is not null))::numeric
    /(select count(distinct c.player_id) from Activity c), 2) fraction
from (
    select player_id, MIN(event_date)+INTERVAL '1 DAY' as event_date
        from Activity
    group by 1) t left join Activity a 
using (player_id, event_date)