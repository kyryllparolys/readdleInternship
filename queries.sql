select
  titles.title, employees.first_name, employees.last_name, salaries.salary
from titles
inner join employees on
	titles.emp_no=employees.emp_no
inner join salaries on
	employees.emp_no=salaries.emp_no
where titles.title like "%Manager%";

