-- 初始化
DROP TABLE IF EXISTS `Customers`;
CREATE TABLE `Customers` (
`Id` INT,
`Name` VARCHAR(256),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `Customers` VALUES
(1, 'Joe'),
(2, 'Henry'),
(3, 'Sam'),
(4, 'Max');

DROP TABLE IF EXISTS `Orders`;
CREATE TABLE `Orders` (
`Id` INT,
`CustomerId` INT,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `Orders` VALUES
(1, 3),
(2, 1);

-- 子查询
SELECT t1.Name AS Customers FROM Customers AS t1
WHERE t1.Id NOT IN (SELECT CustomerId FROM Orders);

