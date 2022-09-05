-- 初始化
DROP TABLE IF EXISTS `Employee`;
CREATE TABLE `Employee` (
`id` INT,
`name` VARCHAR(256),
`salary` INT,
`managerId` INT,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `Employee` VALUES
(1, 'Joe', 70000, 3),
(2, 'Henry', 80000, 4),
(3, 'Sam', 60000, NULL),
(4, 'Max', 90000, NULL);

-- 添加 t1.managerId IS NOT NULL 条件可以更快

-- 子查询
SELECT `name` AS Employee FROM Employee AS t1
WHERE managerId IS NOT NULL
AND
EXISTS (SELECT * FROM Employee AS t2 WHERE t1.managerId = t2.id AND t1.salary > t2.salary);

-- 连接
SELECT t1.`name` AS Employee FROM Employee AS t1
INNER JOIN
Employee AS t2
ON t1.managerId IS NOT NULL AND t1.managerId = t2.id AND t1.salary > t2.salary;

-- 笛卡尔积
SElECT t1.`name` AS Employee FROM Employee AS t1, Employee AS t2
WHERE t1.managerId IS NOT NULL AND t1.managerId = t2.id AND t1.salary > t2.salary;