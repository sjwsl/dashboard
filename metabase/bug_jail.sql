# Team Jail
WITH team_max_time AS (SELECT max(time) max_time, team_id FROM team_bug_jail GROUP BY team_id)
SELECT DISTINCT t.name                               AS team,
                team_bug_jail.di                     AS DI,
                team_bug_jail.di - t.size * 5        AS `DI to free`,
                team_bug_jail.time + INTERVAL 8 HOUR AS `update time`
FROM team_bug_jail
         JOIN team_max_time
              ON team_max_time.team_id = team_bug_jail.team_id AND team_max_time.max_time = team_bug_jail.time
         JOIN team t ON team_max_time.team_id = t.id
WHERE in_jail = TRUE
ORDER BY di DESC;

# Free Teams
WITH team_max_time AS (SELECT max(time) max_time, team_id FROM team_bug_jail GROUP BY team_id)
SELECT DISTINCT t.name                                                AS team,
                team_bug_jail.di,
                IF(team_bug_jail.DI >= t.size * 5, 'warning', 'good') AS `status`,
                team_bug_jail.time + INTERVAL 8 HOUR                  AS `update time`
FROM team_bug_jail
         JOIN team_max_time
              ON team_max_time.team_id = team_bug_jail.team_id AND team_max_time.max_time = team_bug_jail.time
         JOIN team t
              ON team_max_time.team_id = t.id
WHERE in_jail = FALSE
ORDER BY di
    DESC;

# User Jail
WITH user_max_time AS (SELECT max(time) max_time, user_id FROM user_bug_jail GROUP BY user_id)
SELECT DISTINCT u.login                              AS user,
                user_bug_jail.di                     AS DI,
                user_bug_jail.time + INTERVAL 8 HOUR AS `update time`,
                critical
FROM user_bug_jail
         INNER JOIN user u ON user_bug_jail.user_id = u.id
         INNER JOIN user_max_time
                    ON user_max_time.max_time = user_bug_jail.time AND user_max_time.user_id = user_bug_jail.user_id
WHERE in_jail = TRUE
ORDER BY di
    DESC;

# Free Users
WITH user_max_time AS (SELECT max(time) max_time, user_id FROM user_bug_jail GROUP BY user_id)
SELECT DISTINCT IF(user_bug_jail.di >= 7.5, 'warning', 'good') AS status,
                u.login                                        AS user,
                user_bug_jail.di                               AS DI,
                user_bug_jail.time + INTERVAL 8 HOUR           AS `update time`
FROM user_bug_jail
         INNER JOIN user u ON user_bug_jail.user_id = u.id
         INNER JOIN user_max_time
                    ON user_max_time.max_time = user_bug_jail.time AND user_max_time.user_id = user_bug_jail.user_id
WHERE in_jail = FALSE
ORDER BY di
    DESC;

# Critical Bugs More Than 48 Hours
SELECT issue.url                                  AS issue,
       title,
       TIMESTAMPDIFF(HOUR, created_at, curtime()) AS `open hours`,
       GROUP_CONCAT(u.login)                      AS assignees
FROM issue
         INNER JOIN issue_label il
                    ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE lsw.label_name = 'severity/critical'
  AND closed = 0
  AND TIMESTAMPDIFF(HOUR, created_at, curtime()) > 48
GROUP BY url, title, created_at;

