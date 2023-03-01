CREATE TABLE users (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `created_at` datetime default current_timestamp,
    PRIMARY KEY (`id`)
);
