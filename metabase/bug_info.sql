# Open Bugs by Repository
SELECT repo_name, count(*) AS count
FROM issue
         JOIN repository r ON r.id = issue.repository_id
         JOIN issue_label il ON issue.id = il.issue_id
         JOIN label l ON l.id = il.label_id
WHERE closed = 0
  AND l.name = 'type/bug'
GROUP BY repo_name
ORDER BY repo_name;

# Open Bugs No SIG by Repository
SELECT repo_name, count(*) AS count
FROM issue
         JOIN repository r ON r.id = issue.repository_id
         JOIN issue_label il ON issue.id = il.issue_id
         JOIN label l ON l.id = il.label_id
WHERE closed = 0
  AND l.name = 'type/bug'
  AND issue.id NOT IN (
    SELECT issue.id
    FROM issue
             JOIN issue_label i ON issue.id = i.issue_id
             JOIN label_sig ls ON i.label_id = ls.label_id
)
GROUP BY repo_name
ORDER BY repo_name;

# Open Bugs No Assignee by Repository
SELECT repo_name, count(*) AS count
FROM issue
         JOIN repository ON issue.repository_id = repository.id
         JOIN issue_label il ON issue.id = il.issue_id
         JOIN label l ON l.id = il.label_id
WHERE closed = 0
  AND l.name = 'type/bug'
  AND issue.id NOT IN (
    SELECT issue.id
    FROM issue
             JOIN user_issue ui ON issue.id = ui.issue_id
             JOIN user u ON u.id = ui.user_id
)
GROUP BY repo_name
ORDER BY repo_name;

# Open Bugs No Severity by Repository
SELECT repo_name, count(*) AS count
FROM issue
         JOIN repository r ON r.id = issue.repository_id
         JOIN issue_label il ON issue.id = il.issue_id
         JOIN label l ON l.id = il.label_id
WHERE closed = 0
  AND l.name = 'type/bug'
  AND issue.id NOT IN (
    SELECT issue.id
    FROM issue
             JOIN issue_label i ON issue.id = i.issue_id
             JOIN label_severity_weight lsw ON i.label_id = lsw.label_id
)
GROUP BY repo_name
ORDER BY repo_name;

# Open Bugs Need More Info
SELECT issue.url                             AS issue,
       r.repo_name                           AS repo,
       group_concat(DISTINCT u.login)        AS assignees,
       group_concat(DISTINCT ls.label_name)  AS sig,
       group_concat(DISTINCT lsw.label_name) AS severity,
       issue.closed_at
FROM issue
         JOIN repository r ON issue.repository_id = r.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON u.id = ui.user_id
         JOIN issue_label il ON issue.id = il.issue_id
         LEFT JOIN label_sig ls ON il.label_id = ls.label_id
         LEFT JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
WHERE issue.id IN (
    SELECT issue.id
    FROM issue
             JOIN issue_label i ON issue.id = i.issue_id
             JOIN label l ON i.label_id = l.id AND l.name = 'need-more-info'
)
GROUP BY issue.id, r.repo_name, issue.closed_at
ORDER BY issue.closed_at DESC;

# Open Critical Bugs
SELECT issue.url                        AS issue,
       title,
       GROUP_CONCAT(DISTINCT u.login)   AS assignees,
       GROUP_CONCAT(DISTINCT team.name) AS teams
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team ON ti.team_id = team.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE closed = 0
  AND lsw.label_name = 'severity/critical'
GROUP BY issue.url, title;

# Open Critical Bugs Created Last 7 Days
SELECT issue.url                        AS issue,
       title,
       GROUP_CONCAT(DISTINCT u.login)   AS assignees,
       GROUP_CONCAT(DISTINCT team.name) AS teams,
       created_at                       AS created
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team ON ti.team_id = team.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE closed = 0
  AND lsw.label_name = 'severity/critical'
  AND TIMESTAMPDIFF(DAY, created_at, CURTIME()) <= 7
GROUP BY issue.url, title, created_at
ORDER BY created_at DESC;

# Critical Bugs Closed Last 7 Days
SELECT issue.url                        AS issue,
       title,
       GROUP_CONCAT(DISTINCT u.login)   AS assignees,
       GROUP_CONCAT(DISTINCT team.name) AS teams,
       closed_at                        AS closed
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team ON ti.team_id = team.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE closed = 1
  AND lsw.label_name = 'severity/critical'
  AND TIMESTAMPDIFF(DAY, closed_at, CURTIME()) <= 7
GROUP BY issue.url, title, closed_at
ORDER BY closed_at DESC;

# Open Major Bugs
SELECT issue.url                        AS issue,
       title,
       GROUP_CONCAT(DISTINCT u.login)   AS assignees,
       GROUP_CONCAT(DISTINCT team.name) AS teams
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team ON ti.team_id = team.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE closed = 0
  AND lsw.label_name = 'severity/major'
GROUP BY issue.url, title;

# Open Major Bugs Created Last 7 Days
SELECT issue.url                        AS issue,
       title,
       GROUP_CONCAT(DISTINCT u.login)   AS assignees,
       GROUP_CONCAT(DISTINCT team.name) AS teams,
       created_at                       AS created
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team ON ti.team_id = team.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE closed = 0
  AND lsw.label_name = 'severity/major'
  AND TIMESTAMPDIFF(DAY, created_at, CURTIME()) <= 7
GROUP BY issue.url, title, created_at
ORDER BY created_at DESC;

# Major Bugs Closed Last 7 Days
SELECT issue.url                        AS issue,
       title,
       GROUP_CONCAT(DISTINCT u.login)   AS assignees,
       GROUP_CONCAT(DISTINCT team.name) AS teams,
       closed_at                        AS closed
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team ON ti.team_id = team.id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
WHERE closed = 1
  AND lsw.label_name = 'severity/major'
  AND TIMESTAMPDIFF(DAY, closed_at, CURTIME()) <= 7
GROUP BY issue.url, title, closed_at
ORDER BY closed_at DESC;

# Open Bugs No SIG TimeLine
SELECT wl.week, repository.repo_name, count(*)
FROM repository
         JOIN issue i ON repository.id = i.repository_id
         JOIN week_line wl ON i.created_at < wl.week AND (i.closed_at > wl.week OR i.closed_at IS NULL)
         JOIN issue_label il ON i.id = il.issue_id
         JOIN label l ON l.id = il.label_id AND l.name = 'type/bug'
WHERE TIMESTAMPDIFF(YEAR, wl.week, CURTIME()) <= 1
  AND i.id NOT IN (
    SELECT issue.id
    FROM issue
             JOIN issue_label il2 ON issue.id = il2.issue_id
             JOIN label_sig ls ON il2.label_id = ls.label_id
)
GROUP BY wl.week, repository.repo_name
ORDER BY wl.week DESC;

# Open Bugs by Team Timeline
SELECT wl.week, team.name AS team, COUNT(DISTINCT issue.id) AS count
FROM issue
         INNER JOIN team_issue ON issue.id = team_issue.issue_id
         LEFT JOIN team ON team_issue.team_id = team.id
         INNER JOIN issue_label ON issue.id = issue_label.issue_id
         LEFT JOIN label ON issue_label.label_id = label.id,
     week_line wl
WHERE TIMESTAMPDIFF(YEAR, wl.week, CURTIME()) <= 1
  AND label.name = 'type/bug'
  AND issue.created_at <= wl.week
  AND (issue.closed_at
           > wl.week
    OR issue.closed = FALSE)
GROUP BY wl.week, team.name;

# Open Bugs by Severity TimeLine
SELECT wl.week, count(*), lsw.label_name
FROM issue i
         JOIN repository r ON i.repository_id = r.id
         JOIN issue_label il ON i.id = il.issue_id
         JOIN label l ON il.label_id = l.id AND l.name = 'type/bug'
         JOIN issue_label il2 ON i.id = il2.issue_id
         JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
         JOIN week_line wl ON i.created_at < wl.week AND (i.closed_at > wl.week OR i.closed_at IS NULL)
WHERE TIMESTAMPDIFF(YEAR, wl.week, CURTIME()) <= 1
GROUP BY wl.week, lsw.label_name
ORDER BY wl.week DESC;

# Open Bugs by Repository TimeLine
SELECT wl.week, repository.repo_name, count(*)
FROM repository
         JOIN issue i ON repository.id = i.repository_id
         JOIN week_line wl ON i.created_at < wl.week AND (i.closed_at > wl.week OR i.closed_at IS NULL)
         JOIN issue_label il ON i.id = il.issue_id
         JOIN label l ON l.id = il.label_id AND l.name = 'type/bug'
WHERE TIMESTAMPDIFF(YEAR, wl.week, CURTIME()) <= 1
GROUP BY wl.week, repository.repo_name
ORDER BY wl.week DESC;

# Open Bugs by SIG TimeLine
SELECT wl.week, ls.label_name, count(*)
FROM repository
         JOIN issue i ON repository.id = i.repository_id
         JOIN week_line wl ON i.created_at < wl.week AND (i.closed_at > wl.week OR i.closed_at IS NULL)
         JOIN issue_label il ON i.id = il.issue_id
         JOIN label l ON l.id = il.label_id AND l.name = 'type/bug'
         JOIN issue_label il2 ON i.id = il2.issue_id
         JOIN label_sig ls ON il2.label_id = ls.label_id
WHERE TIMESTAMPDIFF(YEAR, wl.week, CURTIME()) <= 1
GROUP BY wl.week, ls.label_name
ORDER BY wl.week DESC;