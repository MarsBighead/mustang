
-- 1978. 上级经理已离职的公司员工
-- https://leetcode.cn/problems/employees-whose-manager-left-the-company/description/?envType=study-plan-v2&envId=sql-free-50
select e.employee_id
from Employees e left join Employees a
on e.manager_id = a.employee_id
where  e.manager_id is not null 
 and a.employee_id is null and e.salary<30000
order by 1

-- 子查询未必比连接快
select employee_id
from Employees 
where salary<30000 and
 manager_id is not null and
 manager_id not in(
    select  employee_id from Employees )
order by 1

-- https://leetcode.cn/problems/exchange-seats/?envType=study-plan-v2&envId=sql-free-50
-- 626. 换座位
SELECT
	CASE
		WHEN ID = MAX_ID
		AND MAX_ID % 2 = 1 THEN ID
		WHEN ID < MAX_ID
		AND ID % 2 = 1 THEN ID + 1
		ELSE ID -1
	END IDX,
	ID,
    STUDENT
FROM
	(
		SELECT ID, STUDENT, MAX(ID) OVER () AS MAX_ID
		FROM Seat
	)
ORDER BY 1

-- 将偶数id减2，然后进行排序
-- 通过rank函数重新为id赋值
SELECT
	RANK() OVER (
		ORDER BY
			CASE
				WHEN ID % 2 = 0 THEN ID -2
				ELSE ID
			END
	) AS ID,
	STUDENT
FROM Seat

-- 1341. 电影评分
-- https://leetcode.cn/problems/movie-rating/description/?envType=study-plan-v2&envId=sql-free-50
-- 查找在 February 2020 平均评分最高 的电影名称。如果出现平局，返回字典序较小的电影名称。
select title results
from Movies m left join (
    select movie_id, avg(rating) avg_rating,
        max(avg(rating)) over() max_rating 
    from MovieRating
    where to_char(created_at, 'YYYY-MM')='2020-02'
    group by movie_id
)  t using(movie_id)
where t.avg_rating=t.max_rating 
order by 1 asc
limit 1

-- Write your PostgreSQL query statement below
(select u.name as results 
from Users u left join
(select user_id, count(*) cnt, max(count(*)) over() max_cnt
from MovieRating
group by user_id) t0 using(user_id)
where t0.cnt=t0.max_cnt
order by 1 asc
limit 1)
union all
(select title results
from Movies m left join (
    select movie_id, avg(rating) avg_rating,
        max(avg(rating)) over() max_rating 
    from MovieRating
    where to_char(created_at, 'YYYY-MM')='2020-02'
    group by movie_id
)  t using(movie_id)
where t.avg_rating=t.max_rating 
order by 1 asc
limit 1)

