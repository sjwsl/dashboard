# Team Jail
INSERT INTO team_bug_jail(time, team_id, DI, in_jail)
SELECT CURTIME(),
       t.id,
       SUM(weight) AS DI,
       IF(SUM(weight) >= t.size * 10
              OR (SUM(weight) >= t.size * 5 AND
                 EXISTS(SELECT team_bug_jail.team_id FROM team_bug_jail WHERE t.id = team_bug_jail.team_id AND in_jail = TRUE)),
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

# Person Jail
INSERT INTO user_bug_jail(time, user_id, in_jail, di, critical)
SELECT curtime()   AS time,
       u.id        AS user,
       iF(EXISTS(SELECT issue.id
                 FROM issue
                          INNER JOIN issue_label i ON issue.id = i.issue_id
                          INNER JOIN label_severity_weight l ON i.label_id = l.label_id
                          INNER JOIN user_issue ui2 ON issue.id = ui2.issue_id
                          INNER JOIN user u2 ON ui2.user_id = u2.id
                 WHERE l.label_name = 'severity/critical'
                   AND closed = 0
                   AND u2.login = u.login
                   AND TIMESTAMPDIFF(HOUR
                           , created_at
                           , curtime())
                     > 48
              ) OR SUM(weight) >= 15
              OR ((SUM(weight) >= 7.5 AND
                   EXISTS(SELECT user_bug_jail.user_id
                          FROM user_bug_jail
                          WHERE u.id = user_bug_jail.user_id
                            AND in_jail = TRUE))),
          TRUE, FALSE
           )       AS in_jail,
       SUM(weight) AS di,
       (SELECT count(*)
        FROM issue
                 INNER JOIN issue_label i ON issue.id = i.issue_id
                 INNER JOIN label_severity_weight l ON i.label_id = l.label_id
                 INNER JOIN user_issue ui2 ON issue.id = ui2.issue_id
                 INNER JOIN user u2 ON ui2.user_id = u2.id
        WHERE l.label_name = 'severity/critical'
          AND closed = 0
          AND u2.login = u.login
          AND TIMESTAMPDIFF(HOUR
                  , created_at
                  , curtime())
            > 48)  AS critical
FROM issue
         INNER JOIN user_issue ui
                    ON issue.id = ui.issue_id
         INNER JOIN user u ON ui.user_id = u.id
         INNER JOIN issue_label il ON issue.id = il.issue_id
         INNER JOIN label_severity_weight lsw ON il.label_id = lsw.label_id
WHERE closed = 0
GROUP BY u.login
ORDER BY DI
    DESC;
;