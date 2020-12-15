# Cumulative Open OnCalls Timeline
SELECT timeline.date AS date, p.pname AS priority, count(DISTINCT jiraissue.ID) AS count
FROM jiraissue
         LEFT JOIN issuestatus i
                   ON jiraissue.issuestatus = i.ID
         INNER JOIN priority p ON jiraissue.PRIORITY = p.ID
         LEFT JOIN (SELECT changegroup.issueid AS issue_id, changegroup.CREATED AS closed_time
                    FROM jiraissue
                             LEFT JOIN changegroup ON jiraissue.ID = changegroup.issueid
                             LEFT JOIN changeitem ON changeitem.groupid = changegroup.ID
                    WHERE jiraissue.PROJECT = '11101'
                      AND FIELD = 'status'
                      AND (lower(NEWSTRING))
                        IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled')
                      AND (lower(OLDSTRING))
                        NOT IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled')) issue_closed
                   ON jiraissue.id = issue_closed.issue_id
   , (SELECT DISTINCT date
      FROM (SELECT DISTINCT DATE(CREATED) AS date
            FROM jiraissue
            UNION
            DISTINCT
            SELECT date(UPDATED) AS date
            FROM jiraissue) AS jdjd
      WHERE weekday(date) = 0
      ORDER BY date) AS timeline
WHERE PROJECT = '11101'
  AND TIMESTAMPDIFF(YEAR, timeline.date, CURTIME()) <= 1
  AND date(CREATED) <= timeline.date
  AND (lower(i.pname)
           NOT IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled') OR
       closed_time > timeline.date)
GROUP BY timeline.date, p.pname;

# Cumulative Closed OnCalls Timeline
SELECT timeline.date AS date, p.pname AS priority, count(DISTINCT jiraissue.ID) AS count
FROM jiraissue
         LEFT JOIN issuestatus i
                   ON jiraissue.issuestatus = i.ID
         INNER JOIN priority p ON jiraissue.PRIORITY = p.ID
         INNER JOIN (SELECT changegroup.issueid AS issue_id, changegroup.CREATED AS closed_time
                     FROM jiraissue
                              LEFT JOIN changegroup ON jiraissue.ID = changegroup.issueid
                              LEFT JOIN changeitem ON changeitem.groupid = changegroup.ID
                     WHERE jiraissue.PROJECT = '11101'
                       AND FIELD = 'status'
                       AND (lower(NEWSTRING))
                         IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled')
                       AND (lower(OLDSTRING))
                         NOT IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled')) AS oncall_closed
                    ON jiraissue.ID = oncall_closed.issue_id
   , (SELECT DISTINCT date
      FROM (SELECT DISTINCT DATE(CREATED) AS date
            FROM jiraissue
            UNION
            DISTINCT
            SELECT date(UPDATED) AS date
            FROM jiraissue) AS jdjd
      WHERE weekday(date) = 0
      ORDER BY date) AS timeline
WHERE PROJECT = '11101'
  AND TIMESTAMPDIFF(YEAR, timeline.date, CURTIME()) <= 1
  AND closed_time < timeline.date
GROUP BY timeline.date, p.pname
ORDER BY date;

# New OnCalls Every Week
SELECT timeline.date AS date, p.pname AS priority, count(DISTINCT jiraissue.ID) AS count
FROM jiraissue
         LEFT JOIN issuestatus i
                   ON jiraissue.issuestatus = i.ID
         INNER JOIN priority p ON jiraissue.PRIORITY = p.ID
   , (SELECT DISTINCT date
      FROM (SELECT DISTINCT DATE(CREATED) AS date
            FROM jiraissue
            UNION
            DISTINCT
            SELECT date(UPDATED) AS date
            FROM jiraissue) AS jdjd
      WHERE weekday(date) = 0
      ORDER BY date) AS timeline
WHERE PROJECT = '11101'
          AND TIMESTAMPDIFF(YEAR, timeline.date, CURTIME()) <= 1
          AND CREATED < timeline.date && CREATED >= timeline.date - INTERVAL 7 DAY
GROUP BY timeline.date, p.pname
ORDER BY date;

# Closed OnCalls Every Week
SELECT timeline.date AS date, p.pname AS priority, count(DISTINCT jiraissue.ID) AS count
FROM jiraissue
         LEFT JOIN issuestatus i
                   ON jiraissue.issuestatus = i.ID
         JOIN (SELECT changegroup.issueid, changegroup.CREATED, changeitem.NEWSTRING
               FROM changegroup
                        JOIN changeitem ON changegroup.id = changeitem.groupid,
                    (SELECT changegroup.issueid, max(CREATED) AS maxcreated
                     FROM changegroup
                              JOIN changeitem ON changegroup.id = changeitem.groupid
                     WHERE changeitem.FIELD = 'status'
                     GROUP BY changegroup.issueid
                     ORDER BY issueid) lastchangeigem
               WHERE changegroup.CREATED = lastchangeigem.maxcreated
                 AND changegroup.issueid = lastchangeigem.issueid
                 AND changeitem.FIELD = 'status') changeitem
              ON jiraissue.ID = changeitem.issueid
         INNER JOIN priority p ON jiraissue.PRIORITY = p.ID
   , (SELECT DISTINCT DATE_FORMAT(date, '%Y-%m-%d') AS date
      FROM (SELECT date(UPDATED) AS date
            FROM jiraissue) AS jdjd
      WHERE weekday(date) = 0
      ORDER BY date) AS timeline
WHERE PROJECT = '11101'
  AND TIMESTAMPDIFF(YEAR, timeline.date, CURTIME()) <= 1
  AND DATE_FORMAT(DATE_SUB(changeitem.CREATED, INTERVAL DAYOFWEEK(changeitem.CREATED) - 2 DAY), '%Y-%m-%d') =
      timeline.date
  AND changeitem.NEWSTRING IN
      ('Job Closed', 'Resolved', 'Closed', 'Canceled', 'Finished', 'CAN\'\'T REPRODUCE', 'WON\'\'T FIX')
  AND lower(i.pname)
    IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled')
GROUP BY timeline.date, p.pname
ORDER BY date;

# Open Oncalls by Priority
SELECT p.pname Priority, count(jiraissue.ID) count
FROM jiraissue
         JOIN priority p ON p.ID = jiraissue.PRIORITY AND jiraissue.PROJECT = "11101"
         JOIN issuestatus i ON i.ID = jiraissue.issuestatus AND
                               lower(i.pname) NOT IN
                               ("duplicated", "resolved", "can't reproduce", "invalid", "canceled")
GROUP BY p.SEQUENCE, p.pname
ORDER BY p.SEQUENCE;

# Open OnCalls by Assignee
SELECT j.ASSIGNEE User, count(j.ID) count
FROM jiraissue j
         JOIN priority p ON p.ID = j.PRIORITY AND j.PROJECT = "11101"
         JOIN issuestatus i ON i.ID = j.issuestatus AND
                               lower(i.pname) NOT IN
                               ("duplicated", "resolved", "can't reproduce", "invalid", "canceled")
GROUP BY j.ASSIGNEE
ORDER BY count(j.ID) DESC;

# Closed OnCalls by Assignee
SELECT j.ASSIGNEE User, count(j.ID) IssueCount
FROM jiraissue j
         JOIN priority p ON p.ID = j.PRIORITY AND j.PROJECT = "11101"
         JOIN issuestatus i ON i.ID = j.issuestatus AND
                               lower(i.pname) IN ("duplicated", "resolved", "can't reproduce", "invalid", "canceled")
GROUP BY j.ASSIGNEE
ORDER BY count(j.ID) DESC;

# OnCalls Opened Last 7 Days
SELECT j.SUMMARY                             summary,
       p.pname                               priority,
       j.ASSIGNEE                            assignee,
       timestampdiff(HOUR, j.CREATED, now()) live_hour,
       j.issuenum
FROM jiraissue j
         JOIN priority p ON p.ID = j.PRIORITY AND j.PROJECT = "11101"
         JOIN issuestatus i ON i.ID = j.issuestatus AND
                               lower(i.pname) NOT IN
                               ("duplicated", "resolved", "can't reproduce", "invalid", "canceled")
WHERE j.CREATED > date_sub(now(), INTERVAL 7 DAY)
ORDER BY live_hour;

# OnCalls Closed Last 7 Days
SELECT jiraissue.SUMMARY                                          summary,
       p.pname                                                    priority,
       jiraissue.ASSIGNEE                                         assignee,
       timestampdiff(HOUR, jiraissue.CREATED, changeitem.CREATED) live_hour,
       jiraissue.issuenum
FROM jiraissue
         LEFT JOIN issuestatus i
                   ON jiraissue.issuestatus = i.ID
         JOIN (SELECT changegroup.issueid, changegroup.CREATED, changeitem.NEWSTRING
               FROM changegroup
                        JOIN changeitem ON changegroup.id = changeitem.groupid,
                    (SELECT changegroup.issueid, max(CREATED) AS maxcreated
                     FROM changegroup
                              JOIN changeitem ON changegroup.id = changeitem.groupid
                     WHERE changeitem.FIELD = 'status'
                     GROUP BY changegroup.issueid
                     ORDER BY issueid) lastchangeigem
               WHERE changegroup.CREATED = lastchangeigem.maxcreated
                 AND changegroup.issueid = lastchangeigem.issueid
                 AND changeitem.FIELD = 'status') changeitem
              ON jiraissue.ID = changeitem.issueid
         INNER JOIN priority p ON jiraissue.PRIORITY = p.ID
WHERE PROJECT = '11101'
  AND changeitem.CREATED > date_sub(now(), INTERVAL 7 DAY)
  AND changeitem.NEWSTRING IN
      ('Job Closed', 'Resolved', 'Closed', 'Canceled', 'Finished', 'CAN\'\'T REPRODUCE', 'WON\'\'T FIX')
  AND lower(i.pname)
    IN ('duplicated', 'resolved', 'can\'t reproduce', 'invalid', 'canceled')

# OnCalls Opened Last 7 Days
SELECT j.SUMMARY                             summary,
       p.pname                               priority,
       j.ASSIGNEE                            assignee,
       timestampdiff(HOUR, j.CREATED, now()) live_hour,
       j.issuenum
FROM jiraissue j
         JOIN priority p ON p.ID = j.PRIORITY AND j.PROJECT = "11101"
         JOIN issuestatus i ON i.ID = j.issuestatus AND
                               lower(i.pname) NOT IN
                               ("duplicated", "resolved", "can't reproduce", "invalid", "canceled")
WHERE j.CREATED > date_sub(now(), INTERVAL 30 DAY)
ORDER BY priority;