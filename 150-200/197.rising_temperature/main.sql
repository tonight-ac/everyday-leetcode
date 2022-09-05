-- 初始化
DROP TABLE IF EXISTS `Weather`;
CREATE TABLE `Weather` (
`id` INT,
`recordDate` DATE,
`temperature` INT,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `Weather` VALUES
(1, '2015-01-01', 10),
(2, '2015-01-02', 25),
(3, '2015-01-03', 20),
(4, '2015-01-04', 30);

-- 连接
SELECT t1.id FROM Weather AS t1
INNER JOIN
Weather AS t2
ON DATEDIFF(t1.recordDate, t2.recordDate) = 1 AND t1.temperature > t2.temperature;