-- First query
SELECT
  titles.title, employees.first_name, employees.last_name, salaries.salary
FROM titles
INNER JOIN employees ON
	titles.emp_no=employees.emp_no
INNER JOIN salaries ON
	employees.emp_no=salaries.emp_no
WHERE titles.title LIKE "%Manager%";

-- Second query
SELECT
	e.first_name, e.last_name, e.hire_date, DATEDIFF(CURDATE(), e.hire_date) as at_company,
	t.title,
	d_emp.dept_no, d.dept_name
FROM
	employees e
INNER JOIN titles t on e.emp_no = t.emp_no
INNER JOIN dept_emp d_emp on t.emp_no = d_emp.emp_no
INNER JOIN departments d on d_emp.dept_no = d.dept_no;

-- Third query
SELECT
	d.dept_no,
	COUNT(*),
	SUM(s.salary)
FROM
	departments d
INNER JOIN dept_emp de on d.dept_no = de.dept_no
INNER JOIN salaries s on de.emp_no = s.emp_no
GROUP BY d.dept_no;
