from diagrams import Diagram, Cluster
from diagrams.aws.compute import EC2
from diagrams.aws.database import RDS
from diagrams.aws.network import ELB
from diagrams.aws.network import CF
from diagrams.onprem.database import Postgresql
from diagrams.onprem.inmemory import Redis
from diagrams.onprem.queue import Rabbitmq
from diagrams.aws.compute import ElasticContainerServiceService

with Diagram("URL Shortener - MVP", show=False):
    cdn = CF("CDN")
    lb = ELB("LB")
        
    queue = Cluster("Task Queue")
    with queue:
        queue_master = Rabbitmq("Master")
        queue_master - [Rabbitmq("Failover")]
    
    
    with Cluster("Application"):
        # API - Main application (Handles users, URLs, redirection, etc.)
        api = ElasticContainerServiceService("API Service")

        # Worker nodes (cleanup db, send verification email, handle signup request etc.)
        worker = ElasticContainerServiceService("Task workers")

        # redirection svc - only responds to frontend traffic for redirection.
        redirector = ElasticContainerServiceService("Redirection service")


    db_cluster = Cluster("Database Cluster")
    with db_cluster:
        db_master = Postgresql("Master DB")
        db_master - [Postgresql("Replica 1"),Postgresql("Replica 2")]

    with Cluster("Redis Cluster"):
        cache_primary = Redis("Master")
        cache_failover = Redis("Failover")
        cache_primary - [cache_failover]

    cdn >> lb
    lb >> api
    lb >> redirector
    redirector >> cache_primary
    redirector >> db_master
    api >> db_master
    api >> queue_master
    api >> cache_primary
    
    worker >> queue_master
    worker >> db_master