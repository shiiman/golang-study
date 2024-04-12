create database tbls_app;

use tbls_app;

create table users (
	id int not null auto_increment,
	name varchar(100) not null,
	created_at datetime not null,
	primary key(id)
);

create table posts (
	id int not null auto_increment,
	user_id int not null,
	primary key(id),
	foreign key (user_id) references users(id) on delete cascade
);

create table favorites (
	id int not null auto_increment,
	user_id int not null,
	post_id int not null,
	primary key(id),
	foreign key (user_id) references users(id) on delete cascade,
	foreign key (post_id) references posts(id) on delete cascade
);

create table topics (
	id int not null auto_increment,
	topic_key int,
	body text,
	is_visible bool,
	deleted_at datetime not null,
	primary key(id),
	unique topic_key_idx(topic_key),
	index deleted_at_idx(deleted_at)
);
