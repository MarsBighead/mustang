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