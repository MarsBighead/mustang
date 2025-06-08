-- https://leetcode.cn/problems/replace-employee-id-with-the-unique-identifier/description/?envType=study-plan-v2&envId=sql-free-50
-- 1378. 使用唯一标识码替换员工ID
select u.unique_id, e.name from
Employees e left join EmployeeUNI u
on e.id=u.id

-- https://leetcode.cn/problems/product-sales-analysis-i/description/?envType=study-plan-v2&envId=sql-free-50
-- 1068. 产品销售分析 I
select p.product_name, s.year, s.price
from Sales s join Product p on s.product_id=p.product_id

-- https://leetcode.cn/problems/customer-who-visited-but-did-not-make-any-transactions/description/?envType=study-plan-v2&envId=sql-free-50
-- 1581. 进店却未进行过交易的顾客
select v.customer_id,count(customer_id) count_no_trans
from Visits v 
where v.visit_id not in(
    select distinct t.visit_id f
    from  Transactions t)
group by customer_id

select v.customer_id,count(customer_id) count_no_trans
from Visits v left join  Transactions t
    on v.visit_id = t.visit_id 
where t.visit_id is null
group by v.customer_id

-- https://leetcode.cn/problems/rising-temperature/description/?envType=study-plan-v2&envId=sql-free-50
-- 197. 上升的温度
/*
JOIN 和LEFT JOIN是不同的，
1. JOIN在PostgreSQL中通常指的是 INNER JOIN ，它返回两个表中满足连接条件的行。只有当两个表中都有匹配的情况下，结果集才会包含该行。
2. LEFT JOIN返回左表中的所有行，以及右表中与左表中的行匹配的行。如果没有匹配的行，右表中的列将包含NULL值。
*/
-- Write your PostgreSQL query statement below
select b.id
from  Weather a join Weather b
on a.recordDate=b.recordDate - interval '1 day'
and b.Temperature>a.Temperature

-- https://leetcode.cn/problems/average-time-of-process-per-machine/description/?envType=study-plan-v2&envId=sql-free-50
-- 1661. 每台机器的进程平均运行时间
-- PostgreSQL中数据类型转换需要使用CAST
select a.machine_id, 
    round(cast(avg(b.timestamp-a.timestamp) as numeric),3) processing_time
from Activity a join  Activity b
on a.machine_id = b.machine_id 
and a.process_id = b.process_id
and a.activity_type='start'
and b.activity_type='end'
group by a.machine_id

-- 会快些
select a.machine_id, 
    round(cast(avg(b.timestamp-a.timestamp) as numeric),3) processing_time
from Activity a join  Activity b
on a.machine_id = b.machine_id 
and a.process_id = b.process_id
where a.activity_type='start' and b.activity_type='end'
group by a.machine_id
-- MySQL version
select a.machine_id, 
    round(sum(b.timestamp-a.timestamp)/count(a.machine_id)::numeric,3) processing_time
from Activity a join  Activity b
on a.machine_id = b.machine_id 
and a.process_id = b.process_id
and a.activity_type='start'
and b.activity_type='end'
group by a.machine_id

--https://leetcode.cn/problems/employee-bonus/description/?envType=study-plan-v2&envId=sql-free-50
-- 577. 员工奖金
select e.name,b.bonus
from Employee e left join Bonus b
on e.empId=b.empId
where b.bonus<1000 or b.bonus is null

-- https://leetcode.cn/problems/students-and-examinations/description/?envType=study-plan-v2&envId=sql-free-50
-- 1280. 学生们参加各科测试的次数
-- Write your PostgreSQL query statement below
select st.student_id, st.student_name,s.subject_name,
case 
    when e.attended_exams is null then 0
    else e.attended_exams
end attended_exams
from
   Subjects s CROSS JOIN  Students st
   left join (
    SELECT student_id, subject_name, COUNT(subject_name) AS attended_exams
    FROM Examinations
    GROUP BY student_id, subject_name
) e 
on st.student_id = e.student_id and  s.subject_name = e.subject_name
order by st.student_id, st.student_name,s.subject_name

-- 交叉连接和全连接的区别？
-- Write your PostgreSQL query statement below
select st.student_id, st.student_name,s.subject_name,
COUNT(e.subject_name) AS attended_exams
from 
   Students st CROSS JOIN  Subjects s 
   left join Examinations e 
on st.student_id = e.student_id and s.subject_name = e.subject_name
GROUP by st.student_id, st.student_name,s.subject_name
order BY st.student_id, s.subject_name

-- https://leetcode.cn/problems/managers-with-at-least-5-direct-reports/description/?envType=study-plan-v2&envId=sql-free-50
-- 570. 至少有5名直接下属的经理
select a.name 
from
    Employee a left join(
    select  managerId
        from Employee
    where managerId is not Null
    group by managerId
    having count(managerId)>=5) b
    on a.id=b.managerId
where b.managerid is not null

-- in作为条件查询会比连接+条件查询结果慢一些
select a.name from
    Employee a 
where id in (
    select  managerId
    from Employee
    where managerId is not Null
    group by managerId
    having count(managerId)>=5) 