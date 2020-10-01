# ginApi
ginApi

### user表
```
CREATE TABLE user (
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `age` INT NOT NULL,
    `add_time` DATE,
    PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```