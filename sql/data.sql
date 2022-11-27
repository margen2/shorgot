insert into  users ( name, nick, email, pw)
values 
("margen1", "margen1", "margen1@gmail.com", "$2a$10$.8O4ZDxZsr9jFtL8vlpYJeVq3Gy.LpOqazgwRyFCCPAwk6l33pwTO"),
("margen2", "margen2", "margen2@gmail.com", "$2a$10$.8O4ZDxZsr9jFtL8vlpYJeVq3Gy.LpOqazgwRyFCCPAwk6l33pwTO"),
("margen3", "margen3", "margen3@gmail.com", "$2a$10$.8O4ZDxZsr9jFtL8vlpYJeVq3Gy.LpOqazgwRyFCCPAwk6l33pwTO");

insert into followers ( user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("User 1's post","This is a post made by user 1", 1),
("User 2's post","This is a post made by user 2", 2),
("User 3's post","This is a post made by user 2", 3);