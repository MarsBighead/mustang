
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