
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