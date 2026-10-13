SELECT e.name AS Employee
FROM Employee e
INNER JOIN Employee m
    ON e.managerId = m.idgit 
WHERE e.salary > m.salary;