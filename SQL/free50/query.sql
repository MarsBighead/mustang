
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