select concat(Name,'(',Substring(Occupation,1,1),')') as Name
from occupations
order by Name asc;

select concat('There are a total of ',count(occupation),' ',lower(occupation),'s.') as total
from occupations
group by occupation
order by count(occupation) asc,occupation asc


set @r1=0, @r2=0, @r3=0, @r4=0;
select min(Doctor), min(Professor), min(Singer), min(Actor)
from(
  select case when Occupation='Doctor' then (@r1:=@r1+1)
            when Occupation='Professor' then (@r2:=@r2+1)
            when Occupation='Singer' then (@r3:=@r3+1)
            when Occupation='Actor' then (@r4:=@r4+1) end as RowNumber,
    case when Occupation='Doctor' then Name end as Doctor,
    case when Occupation='Professor' then Name end as Professor,
    case when Occupation='Singer' then Name end as Singer,
    case when Occupation='Actor' then Name end as Actor
  from OCCUPATIONS
  order by Name
) Temp
group by RowNumber;


/*
Binary Tree Nodes

https://www.hackerrank.com/challenges/binary-search-tree-1/problem
*/
Select x.N,
CASE WHEN x.N NOT IN (select distinct P from BST where P IS NOT NULL) THEN 'Leaf'
     WHEN x.P IS NULL THEN 'Root'
     ELSE 'Inner'
END AS NodeType
From BST x 
order by x.N

select c.company_code, c.founder, 
    count(distinct l.lead_manager_code),
    count(distinct s.senior_manager_code), 
    count(distinct m.manager_code),
    count(distinct e.employee_code) 
from Company c, Lead_Manager l, Senior_Manager s, Manager m, Employee e 
where c.company_code = l.company_code 
    and l.lead_manager_code=s.lead_manager_code 
    and s.senior_manager_code=m.senior_manager_code 
    and m.manager_code=e.manager_code 
group by c.company_code order by c.company_code;


/*
Contest Leaderboard

https://www.hackerrank.com/challenges/contest-leaderboard/problem
*/

SELECT h.hacker_id, h.name, SUM(score) FROM (
    SELECT hacker_id, challenge_id, MAX(score) AS score FROM SUBMISSIONS
    GROUP BY hacker_id, challenge_id
)t 
JOIN Hackers h on t.hacker_id = h.hacker_id
GROUP BY h.hacker_id, h.name
HAVING SUM(score) > 0
ORDER BY SUM(score) desc, h.hacker_ids




select c.hacker_id, h.name ,count(c.hacker_id) as c_count
from Hackers as h
    inner join Challenges as c on c.hacker_id = h.hacker_id
group by c.hacker_id, h.name
having 
    c_count = 
        (SELECT MAX(t1.cnt)
        from (SELECT COUNT(hacker_id) as cnt
             from Challenges
             group by hacker_id
             order by hacker_id) t1)
    or c_count in 
        (select t.cnt
         from (select count(*) as cnt 
               from Challenges
               group by hacker_id) t
         group by t.cnt
         having count(t.cnt) = 1)
order by c_count DESC, c.hacker_id;



select  NAME
from city
where CountryCode in (
    select code
    from country
    where CONTINENT='Africa'
)

/*
Average Population of Each Continent

https://www.hackerrank.com/challenges/average-population-of-each-continent/problem
*/
select  co.Continent, FLOOR(AVG(ci.population))
from  COUNTRY co,
CITY ci 
where co.Code= ci.CountryCode
group by co.Continent

/*
The Report

https://www.hackerrank.com/challenges/the-report/problem
*/
select (case when t.grade <8 THEN NULL ELSE t.name END) name, t.grade, t.marks
from
(select s.name, g.grade, s.marks
from students s, grades g
where s.marks between g.min_Mark and g.Max_Mark) t
order by t.grade desc, t.name,t.marks;


select round(long_w , 4) from station
where LAT_N=(select max(LAT_N) 
    from station
    where LAT_N<137.2345
);

/*
Weather Observation Station 18

https://www.hackerrank.com/challenges/weather-observation-station-18
*/
SELECT ROUND(abs((MAX(LONG_W) - MIN(LONG_W)))+ABS((MAX(LAT_N)-MIN(LAT_N))),4) FROM STATION

SELECT ROUND(sqrt(power((MAX(LONG_W) - MIN(LONG_W)),2)+power((MAX(LAT_N)-MIN(LAT_N)),2)),4) 
FROM STATION


SELECT REPEAT('* ', @NUMBER := @NUMBER - 1) FROM information_schema.tables, (SELECT @NUMBER:=21) t LIMIT 20 ;

SELECT REPEAT('* ', @NUMBER := @NUMBER + 1) FROM information_schema.tables, (SELECT @NUMBER:=0) t LIMIT 20 ;


/*
Print Prime Numbers

https://www.hackerrank.com/challenges/print-prime-numbers/problem
*/
SELECT GROUP_CONCAT(NUMB SEPARATOR '&') FROM (     
    SELECT @num:=@num+1 as NUMB 
    FROM information_schema.tables t1,  
         information_schema.tables t2, 
         (SELECT @num:=1) tmp 
    ) tempNum 
WHERE NUMB<=1000 
AND NOT EXISTS( 
    SELECT * FROM 
        (SELECT @nu:=@nu+1 as NUMA 
         FROM information_schema.tables t1, 
              information_schema.tables t2, 
              (SELECT @nu:=1) tmp1    
         LIMIT 1000 ) tatata 
    WHERE FLOOR(NUMB/NUMA)=(NUMB/NUMA) AND NUMA<NUMB AND NUMA>1 
);


/*
Top Competitors

https://www.hackerrank.com/challenges/full-score/problem
*/
select h.hacker_id, h.name
from submissions s
    inner join challenges c on s.challenge_id = c.challenge_id
    inner join difficulty d on c.difficulty_level = d.difficulty_level 
    inner join hackers h on s.hacker_id = h.hacker_id
where s.score = d.score 
    and c.difficulty_level = d.difficulty_level
group by h.hacker_id, h.name
having count(s.hacker_id) > 1
order by count(s.hacker_id) desc, 
    s.hacker_id;

-- psql -U usgmtr -d verify -c "Copy (Select * From \"VcUpdates\") To STDOUT With CSV HEADER DELIMITER ',';" > VcUpdates.csv