create table ASSIGNED_ISSUE_NUM_TIMELINE
(
    DATETIME datetime not null,
    REPO_ID int not null,
    ASSIGNED_ISSUE_NUM int not null,
    primary key (DATETIME, REPO_ID)
);

create table COVERAGE_TIMELINE
(
    TIME datetime not null,
    REPO varchar(20) not null,
    COVERAGE float not null
);

create table LABEL
(
    ID int auto_increment
        primary key,
    NAME varchar(100) not null,
    constraint LABEL_NAME_uindex
        unique (NAME)
);

create table LABEL_SEVERITY_WEIGHT
(
    LABEL_ID int not null,
    WEIGHT float not null
);

create table REPOSITORY
(
    ID int not null
        primary key,
    OWNER varchar(100) not null,
    REPO_NAME varchar(100) not null
);

create table ISSUE
(
    ID int not null
        primary key,
    NUMBER int not null,
    REPOSITORY_ID int not null,
    CLOSED tinyint(1) not null,
    CLOSED_AT datetime null,
    CREATED_AT datetime null,
    TITLE varchar(1000) not null,
    constraint ISSUE_NUMBER_REPOSITORY_ID_uindex
        unique (NUMBER, REPOSITORY_ID),
    constraint REPOSITORY_ID_fk
        foreign key (REPOSITORY_ID) references REPOSITORY (ID)
            on update cascade on delete cascade
);

create table COMMENT
(
    ID int auto_increment,
    ISSUE_ID int not null,
    BODY longtext not null,
    primary key (ID, ISSUE_ID),
    constraint COMMENT_ISSUE_NUMBER_fk
        foreign key (ISSUE_ID) references ISSUE (ID)
            on update cascade on delete cascade
);

create table COMMENT_VERSION
(
    COMMENT_ID int not null
        primary key,
    VERSIONS varchar(1000) null,
    constraint COMMENT_VERSION_COMMENT_ID_fk
        foreign key (COMMENT_ID) references COMMENT (ID)
            on update cascade on delete cascade
);

create table LABEL_ISSUE_RELATIONSHIP
(
    LABEL_ID int not null,
    ISSUE_ID int not null,
    primary key (LABEL_ID, ISSUE_ID),
    constraint ISSUE_LABEL_RELATIONSHIP_ISSUE_ID_fk
        foreign key (ISSUE_ID) references ISSUE (ID)
            on update cascade on delete cascade,
    constraint ISSUE_LABEL_RELATIONSHIP_LABEL_ID_fk
        foreign key (LABEL_ID) references LABEL (ID)
            on update cascade on delete cascade
);

create table REPO_VERSION
(
    TAG varchar(100) not null,
    REPO_ID int not null,
    primary key (TAG, REPO_ID)
);

create table SIG
(
    REPO_NAME varchar(100) null,
    LABEL_NAME varchar(100) null
);

create table TEAM
(
    ID int auto_increment
        primary key,
    NAME varchar(20) not null,
    SIZE int null
);

create table TEAM_LABEL_RELATIONSHIP
(
    TEAM_ID int not null,
    REPOSITORY_ID int not null,
    LABEL_ID int not null,
    primary key (LABEL_ID, TEAM_ID),
    constraint TEAM_LABEL_RELATIONSHIP_LABEL_ID_fk
        foreign key (LABEL_ID) references LABEL (ID),
    constraint TEAM_LABEL_RELATIONSHIP_TEAM_ID_fk
        foreign key (TEAM_ID) references TEAM (ID)
);

create table USER
(
    ID int auto_increment
        primary key,
    LOGIN_NAME varchar(100) not null,
    EMAIL varchar(100) not null,
    constraint USER_LOGIN_NAME_uindex
        unique (LOGIN_NAME)
);

create table ASSIGNEE
(
    USER_ID int not null,
    ISSUE_ID int not null,
    primary key (USER_ID, ISSUE_ID),
    constraint ASSIGNEE_ISSUE_ID_fk
        foreign key (ISSUE_ID) references ISSUE (ID)
            on update cascade on delete cascade,
    constraint ASSIGNEE_USER_ID_fk
        foreign key (USER_ID) references USER (ID)
            on update cascade on delete cascade
);

