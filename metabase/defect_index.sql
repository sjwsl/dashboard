# TiDB DI Timeline
SELECT sum, week, time_label
FROM ((SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Created" AS time_label
       FROM repository
                JOIN issue i ON repository.id = i.repository_id
                JOIN week_line wl ON i.created_week = wl.week
                JOIN issue_label il ON i.id = il.issue_id
                JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                JOIN issue_label il2 ON i.id = il2.issue_id
                JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
       WHERE repo_name = 'tidb'
       GROUP BY wl.week
       ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Closed" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id AND
                                   i.closed = 1
                   JOIN week_line wl ON i.closed_week = wl.week
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'tidb'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "All" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id
                   JOIN week_line wl ON i.created_at < wl.week AND (!i.closed OR i.closed_at > wl.week)
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'tidb'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )) AS U1
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
ORDER BY week;

# TiKV DI Timeline
SELECT sum, week, time_label
FROM ((SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Created" AS time_label
       FROM repository
                JOIN issue i ON repository.id = i.repository_id
                JOIN week_line wl ON i.created_week = wl.week
                JOIN issue_label il ON i.id = il.issue_id
                JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                JOIN issue_label il2 ON i.id = il2.issue_id
                JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
       WHERE repo_name = 'tikv'
       GROUP BY wl.week
       ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Closed" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id AND
                                   i.closed = 1
                   JOIN week_line wl ON i.closed_week = wl.week
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'tikv'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "All" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id
                   JOIN week_line wl ON i.created_at < wl.week AND (!i.closed OR i.closed_at > wl.week)
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'tikv'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )) AS U1
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
ORDER BY week;

# PD DI Timeline
SELECT sum, week, time_label
FROM ((SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Created" AS time_label
       FROM repository
                JOIN issue i ON repository.id = i.repository_id
                JOIN week_line wl ON i.created_week = wl.week
                JOIN issue_label il ON i.id = il.issue_id
                JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                JOIN issue_label il2 ON i.id = il2.issue_id
                JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
       WHERE repo_name = 'pd'
       GROUP BY wl.week
       ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Closed" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id AND
                                   i.closed = 1
                   JOIN week_line wl ON i.closed_week = wl.week
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'pd'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "All" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id
                   JOIN week_line wl ON i.created_at < wl.week AND (!i.closed OR i.closed_at > wl.week)
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'pd'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )) AS U1
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
ORDER BY week;

# DM DI Timeline
SELECT sum, week, time_label
FROM ((SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Created" AS time_label
       FROM repository
                JOIN issue i ON repository.id = i.repository_id
                JOIN week_line wl ON i.created_week = wl.week
                JOIN issue_label il ON i.id = il.issue_id
                JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                JOIN issue_label il2 ON i.id = il2.issue_id
                JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
       WHERE repo_name = 'dm'
       GROUP BY wl.week
       ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Closed" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id AND
                                   i.closed = 1
                   JOIN week_line wl ON i.closed_week = wl.week
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'dm'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "All" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id
                   JOIN week_line wl ON i.created_at < wl.week AND (!i.closed OR i.closed_at > wl.week)
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'dm'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )) AS U1
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
ORDER BY week;

# TiDB-lightning DI Timeline
SELECT sum, week, time_label
FROM ((SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Created" AS time_label
       FROM repository
                JOIN issue i ON repository.id = i.repository_id
                JOIN week_line wl ON i.created_week = wl.week
                JOIN issue_label il ON i.id = il.issue_id
                JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                JOIN issue_label il2 ON i.id = il2.issue_id
                JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
       WHERE repo_name = 'tidb-lightning'
       GROUP BY wl.week
       ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Closed" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id AND
                                   i.closed = 1
                   JOIN week_line wl ON i.closed_week = wl.week
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'tidb-lightning'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "All" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id
                   JOIN week_line wl ON i.created_at < wl.week AND (!i.closed OR i.closed_at > wl.week)
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'tidb-lightning'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )) AS U1
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
ORDER BY week;

# BR DI Timeline
SELECT sum, week, time_label
FROM ((SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Created" AS time_label
       FROM repository
                JOIN issue i ON repository.id = i.repository_id
                JOIN week_line wl ON i.created_week = wl.week
                JOIN issue_label il ON i.id = il.issue_id
                JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                JOIN issue_label il2 ON i.id = il2.issue_id
                JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
       WHERE repo_name = 'br'
       GROUP BY wl.week
       ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "Closed" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id AND
                                   i.closed = 1
                   JOIN week_line wl ON i.closed_week = wl.week
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'br'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )
      UNION
      (
          SELECT sum(lsw.weight) AS sum, wl.week, repo_name, "All" AS time_label
          FROM repository
                   JOIN issue i ON repository.id = i.repository_id
                   JOIN week_line wl ON i.created_at < wl.week AND (!i.closed OR i.closed_at > wl.week)
                   JOIN issue_label il ON i.id = il.issue_id
                   JOIN label l ON l.id = il.label_id AND l.name = "type/bug"
                   JOIN issue_label il2 ON i.id = il2.issue_id
                   JOIN label_severity_weight lsw ON il2.label_id = lsw.label_id
          WHERE repo_name = 'br'
          GROUP BY wl.week
          ORDER BY wl.week DESC
      )) AS U1
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
ORDER BY week;

# team_di_timeline
SELECT week AS time, t.name AS team, SUM(weight) AS DI
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label l ON il.label_id = l.id
         INNER JOIN label_severity_weight lsw ON l.id = lsw.label_id
         INNER JOIN team_issue ti ON issue.id = ti.issue_id
         INNER JOIN team t ON ti.team_id = t.id
         INNER JOIN week_line wl ON created_at < wl.week AND (!closed OR closed_week > week)
WHERE TIMESTAMPDIFF(YEAR, week, CURTIME()) <= 1
GROUP BY week, team
ORDER BY DI
    DESC;