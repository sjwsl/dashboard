INSERT INTO team_bug_jail(time, team_id, DI, in_jail)
SELECT CURTIME(),
       t.id,
       SUM(weight) AS DI,
       IF(SUM(weight) >= t.size * 10
              OR (SUM(weight) >= t.size * 5 AND
                  EXISTS(SELECT team_bug_jail.team_id FROM team_bug_jail WHERE t.id = team_bug_jail.team_id)),
          TRUE,
          FALSE)
FROM issue
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label l ON il.label_id = l.id
         INNER JOIN label_severity_weight lsw ON l.id = lsw.label_id
         INNER JOIN team_issue ti ON issue.id = ti.issue_id
         INNER JOIN team t ON ti.team_id = t.id
WHERE closed = 0
GROUP BY t.id, t.size
;