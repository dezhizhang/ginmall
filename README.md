# ginApi
ginApi

### user表
```
CREATE TABLE users (
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `age` INT NOT NULL,
    PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
### role角色表
```
CREATE TABLE roles(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    status INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    add_time INT NOT NULL,
    PRIMARY KEY (id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
