
SELECT s1.score
FROM
    Scores s1,
    Scores s2,
    Scores s3
WHERE
        s1.id=s2.id - 1
  AND s2.id=s3.id-1
  AND s1.score=s2.score
  AND s2.score=s3.score
