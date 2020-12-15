# Searchable Open Bugs
SELECT issue.url,
       title,
       t.name                        AS team,
       u.login                       AS assignee,
       r.repo_name                   AS repo,
       LOWER(SUBSTR(label_name, 10)) AS severity
FROM issue
         LEFT JOIN team_issue ti ON issue.id = ti.issue_id
         LEFT JOIN team t ON ti.team_id = t.id
         LEFT JOIN issue_label il ON issue.id = il.issue_id
         LEFT JOIN label l ON il.label_id = l.id
         INNER JOIN label_severity_weight lsw ON l.id = lsw.label_id
         LEFT JOIN user_issue ui ON issue.id = ui.issue_id
         LEFT JOIN user u ON ui.user_id = u.id
         LEFT JOIN repository r ON issue.repository_id = r.id
WHERE closed = FALSE [[AND t.name = {{team}}]]
[[AND u.login = {{user}}]]
[[AND r.repo_name = {{repo}}]]
[[AND LOWER(SUBSTR(label_name, 10)) = {{severity}}]]
;