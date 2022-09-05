-- 初始化
DROP TABLE IF EXISTS `Person`;
CREATE TABLE `Person` (
`Id` INT,
`Email` VARCHAR(256),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `Person` VALUES
(1, 'a@b.com'),
(2, 'c@d.com'),
(3, 'a@b.com');

-- 聚合
SELECT Email FROM Person
GROUP BY Email
HAVING COUNT(*) > 1;