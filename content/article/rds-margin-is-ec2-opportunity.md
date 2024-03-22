
I was writing an article on the topic [less spoken costs of managed databases](https://www.infoq.com/articles/managed-relational-databases-costs/) and one question that the editor asked is how much margin RDS makes compared to running a Postgres instance on EC2? That intrigued me and honestly, I never did the math so far. I have been using AWS-managed databases since 2016 and I thought what's a better time to do an analysis of Cloud cost than today ?!

I started by comparing instance types between EC2 and RDS. As far as I understand, these are the same VMs (EC2) that are powering RDS. Anyone with the knowledge, correct me if I am wrong.

![ec2-instance-vs-rds-instance-cost-per-month](/images/ec2-vs-rds-cost-per-month.png)

When I saw more than a 100% margin, I wasn't sure whether I was making any mathematical errors. My instant reaction was perhaps there is some difference in other systems resources. I looked at network performance which is up to 12.5 Gigabit, it's 4 vCPUs with 16 GiB memory. Same-to-same. My understanding is that the newer generation of instances are generally cheaper. However, that's where the margins are also higher.

So far I analyzed only the instance type. Next, I looked into the storage cost (EBS). These are compared by IOPS, throughput, and no snapshots enabled. 

![ec2-vs-rds-storage-cost-per-month](/images/ec2-vs-rds-storage-cost.png)

Looks like gp3 is an anomaly. I used the AWS cost calculator and it had a few additional notes on that. 

![gp3-ec2-detail](/images/gp3-ec2-detail.png)

![gp3-rds-detail](/images/gp3-rds-detail.png)

Certainly, there is operation value adds that RDS provides - automated backup, and failover in the case of AZ (though it’s not an apple-apple comparison). So there is some value added for the markup. However, features like performance insights, and RDS proxy have to be paid separately.

**Open questions:**

[1]. What is the difference between a vanilla EC2 instance and an RDS instance?

[2]. Is there any free tier in GP3 volume type for RDS?

[3]. What's the markup in GCP and Azure?

**References:**

https://instances.vantage.sh/rds/?cost_duration=monthly

https://instances.vantage.sh/?cost_duration=monthly

https://calculator.aws/#/createCalculator/RDSPostgreSQL

https://rbranson.medium.com/rds-pricing-has-more-than-doubled-ef8c3b7e5218
