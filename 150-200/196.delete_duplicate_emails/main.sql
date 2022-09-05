-- 初始化
DROP TABLE IF EXISTS `Person`;
CREATE TABLE `Person` (
`id` INT,
`email` VARCHAR(256),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `Person` VALUES
(1, 'john@example.com'),
(2, 'bob@example.com'),
(3, 'john@example.com');

-- SubQuery
DELETE FROM Person
WHERE id NOT IN
      (SELECT t1.id FROM
          (SELECT MIN(t2.id) AS id FROM Person AS t2 GROUP BY t2.email) as t1);