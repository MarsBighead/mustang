
-- 2356. 每位教师所教授的科目种类的数量
-- https://leetcode.cn/problems/number-of-unique-subjects-taught-by-each-teacher/description/?envType=study-plan-v2&envId=sql-free-50
select teacher_id, count(subject_id) cnt
from (select distinct teacher_id, subject_id 
from Teacher
) t
group by t.teacher_id

select
teacher_id,
count(distinct subject_id) as cnt
from Teacher
group by teacher_id

-- https://leetcode.cn/problems/user-activity-for-the-past-30-days-i/description/?envType=study-plan-v2&envId=sql-free-50
-- 1141. 查询近30天活跃用户数
select activity_date as day, count(distinct(user_id)) active_users
from Activity
where activity_date>'2019-07-27'::date-interval '30 day' and activity_date<='2019-07-27'::date
group by activity_date

-- 596. 超过 5 名学生的课
-- https://leetcode.cn/problems/classes-with-at-least-5-students/description/?envType=study-plan-v2&envId=sql-free-50
select class from Courses
group by class
having count(*)>=5

-- https://leetcode.cn/problems/find-followers-count/description/?envType=study-plan-v2&envId=sql-free-50
-- 1729. 求关注者的数量
select user_id, count(follower_id) as followers_count
from Followers
group by 1
order by 1

-- 619. 只出现一次的最大数字
-- https://leetcode.cn/problems/biggest-single-number/description/?envType=study-plan-v2&envId=sql-free-50
select max(num) num 
from (select num
        from MyNumbers
    group by num
    having count(num)=1) as t

-- https://leetcode.cn/problems/customers-who-bought-all-products/description/?envType=study-plan-v2&envId=sql-free-50
-- 1045. 买下所有产品的客户
-- Customer表可能包含重复的行
select c.customer_id
from  Customer c
group by c.customer_id
having count(distinct c.product_key)=(select count(*) from Product)
