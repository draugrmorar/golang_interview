-- 584. Find Customer Referee https://leetcode.com/problems/find-customer-referee/description
SELECT name
FROM customer
WHERE referee_id != 2
OR referee_id is null;

-- 596. Classes With at Least 5 Students https://leetcode.com/problems/classes-with-at-least-5-students/description
SELECT class
FROM courses
GROUP BY class
HAVING count(student) >= 5;

-- 1683. Invalid Tweets https://leetcode.com/problems/invalid-tweets/description
SELECT tweet_id
FROM tweets
WHERE LENGTH(content) > 15;

-- 577. Employee Bonus https://leetcode.com/problems/employee-bonus/description
SELECT e.name,
       b.bonus
FROM employee e
LEFT JOIN bonus b
    ON e.empId = b.empId
WHERE b.bonus is null
OR b.bonus < 1000;

-- 1731. The Number of Employees Which Report to Each Employee https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/description
SELECT  e.reports_to AS employee_id,
        em.name AS name,
        COUNT(*) AS reports_count,
        ROUND(AVG(e.age)) AS average_age
FROM employees e
JOIN employees em
    ON e.reports_to = em.employee_id
GROUP BY e.reports_to
ORDER BY employee_id ASC;

-- 1729. Find Followers Count https://leetcode.com/problems/find-followers-count/description
SELECT user_id,
       count(*) AS followers_count
FROM followers
GROUP BY user_id
ORDER BY user_id ASC;

-- 1934. Confirmation Rate https://leetcode.com/problems/confirmation-rate/description
SELECT s.user_id AS user_id,
       round(avg(if(c.action="confirmed",1,0)),2) as confirmation_rate
FROM signups s
LEFT JOIN confirmations c
    ON s.user_id = c.user_id
GROUP BY s.user_id;

-- 1280. Students and Examinations https://leetcode.com/problems/students-and-examinations/description
SELECT s.student_id,
       s.student_name,
       sub.subject_name,
       COUNT(e.student_id) AS attended_exams
FROM students s
CROSS JOIN subjects sub
LEFT JOIN examinations e
    ON s.student_id = e.student_id
    AND sub.subject_name = e.subject_name
GROUP BY student_id, subject_name
ORDER BY student_id ASC, subject_name ASC

-- 2356. Number of Unique Subjects Taught by Each Teacher https://leetcode.com/problems/number-of-unique-subjects-taught-by-each-teacher/description/
SELECT teacher_id,
       count(distinct(subject_id)) AS cnt
FROM teacher
GROUP BY teacher_id;

-- 1757. Recyclable and Low Fat Products https://leetcode.com/problems/recyclable-and-low-fat-products/description/
SELECT product_id
FROM products
WHERE low_fats = "Y"
  AND recyclable = "Y";

-- 1693. Daily Leads and Partners https://leetcode.com/problems/daily-leads-and-partners/description/
SELECT date_id,
       make_name,
       COUNT(DISTINCT(lead_id)) AS unique_leads,
       COUNT(DISTINCT(partner_id)) AS unique_partners
FROM dailysales
GROUP BY date_id, make_name;

-- 1741. Find Total Time Spent by Each Employee https://leetcode.com/problems/find-total-time-spent-by-each-employee/description/
SELECT  event_day AS day,
        emp_id,
        (sum(out_time - in_time)) AS total_time
FROM employees
GROUP BY day, emp_id;

-- 1068. Product Sales Analysis I https://leetcode.com/problems/product-sales-analysis-i/description/
SELECT p.product_name,
       s.year,
       s.price
FROM product p
RIGHT JOIN sales s
    ON p.product_id = s.product_id;

-- 3475. DNA Pattern Recognition  https://leetcode.com/problems/dna-pattern-recognition/description/
SELECT sample_id,
       dna_sequence,
       species,
       if(dna_sequence LIKE 'ATG%',1,0) AS has_start,
       if(dna_sequence LIKE '%TAA' OR dna_sequence LIKE '%TAG' OR dna_sequence LIKE '%TGA',1,0) AS has_stop,
       if(dna_sequence LIKE '%ATAT%',1,0) AS has_atat,
       if(dna_sequence LIKE '%GGG%',1,0) AS has_ggg
FROM samples
ORDER BY sample_id ASC;

-- 1795. Rearrange Products Table https://leetcode.com/problems/rearrange-products-table/description/
SELECT product_id, 'store1' AS store, store1 AS price
FROM products
WHERE store1 IS NOT NULL
    UNION ALL
SELECT product_id, 'store2' AS store, store2 AS price
FROM products
WHERE store2 IS NOT NULL
    UNION ALL
SELECT product_id, 'store3' AS store, store3 AS price
FROM products
WHERE store3 IS NOT NULL;

-- 1393. Capital Gain/Loss https://leetcode.com/problems/capital-gainloss/description/
SELECT stock_name,
       sum(if (operation = "BUY", 0 - price, price)) AS capital_gain_loss
FROM    stocks
GROUP BY stock_name;

-- 608. Tree Node https://leetcode.com/problems/tree-node
SELECT id, if(p_id is null, "Root", if(id NOT IN (SELECT p_id FROM tree WHERE p_id IS NOT NULL), "Leaf", "Inner")) AS type
FROM    tree

SELECT id,
       CASE WHEN p_id is null THEN "Root"
            WHEN id NOT IN (SELECT p_id FROM tree WHERE p_id IS NOT null) THEN "Leaf"
            ELSE "Inner"
           END AS type
FROM    tree;

-- 3497. Analyze Subscription Conversion https://leetcode.com/problems/analyze-subscription-conversion/description/
SELECT user_id,
       ROUND(AVG(IF(activity_type = 'free_trial', activity_duration, NULL)), 2) AS trial_avg_duration,
       ROUND(AVG(IF(activity_type = 'paid', activity_duration, NULL)),2) AS paid_avg_duration
FROM useractivity
WHERE user_id IN (
    SELECT user_id
    FROM useractivity
    GROUP BY user_id
    HAVING COUNT(DISTINCT CASE WHEN activity_type='free_trial' THEN 1 END)>0
       AND COUNT(DISTINCT CASE WHEN activity_type='paid' THEN 1 END)>0
)
GROUP BY user_id
ORDER BY user_id ASC;

SELECT * FROM
    (SELECT user_id,
            ROUND(AVG(IF(activity_type = 'free_trial', activity_duration, NULL)), 2) AS trial_avg_duration,
            ROUND(AVG(IF(activity_type = 'paid', activity_duration, NULL)),2) AS paid_avg_duration
     FROM useractivity
     GROUP BY user_id
    ) as a
WHERE trial_avg_duration IS NOT NULL
  AND paid_avg_duration IS NOT NULL
ORDER BY user_id ASC;


-- 626. Exchange Seats https://leetcode.com/problems/exchange-seats/description/
SELECT
    CASE
        WHEN id % 2 = 0 THEN (id-1)
        WHEN id % 2 = 1 AND id+1 IN (SELECT id FROM seat) THEN (id+1)
        ELSE id
    END AS id,
        student
FROM seat
ORDER BY id ASC;
