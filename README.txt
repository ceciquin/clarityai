1- create a postgres DB for the shared cache (I'll have 2 tables with a 1-1 relation, one is for "Security" and the other one is for "Score")

    a. Why postgres?  postgres is an open source db which is mantained for an active community and
    the most significant characteristic is that can be run in multiple platform and operation system which make it
    very flexible and portable. This is important to avod any conflict between other systems or future changes.
    b. I created a specific user and a user role for the db, this user will have the following acces for security porpuse
    and following the principle of "least access": 
        - SELECT access for Security table.
        - SELECT, INSERT, UPDATE for Score table.

2- create a redis connection for the shared cache (I'll have 2 tables with a 1-1 relation, one is for "Securities" and the other one is for "scores")

    a. Why redis?  redis has lot of functionalities that make it very performant ( eg: several data types ).
    This is relevant because will bring the posibility to have a fast shared cache wich is very important 
    if we need to share financial securities information.

3- create 2 APIs using Golang 

    a. Golang is a programming language that is very performant when we need to integrate with third parties or 
    other internal services. It's very easy to mantain and easy to follow. 
    b. y have one internal api called ´scores_api.go´and a public REST api called ´security_api.go´

4- Opentelemetry: I added opentelemetry modules in the project to work with telemetry data which will help me to introduce
an observability culture within the development cycle. This will be very important to notice poor performances in the services
and also to use it in any incident or downside. Opentelemetry it's easy to instrument, is open source and is part of the cloud native org
and has a very vivid community.

5- Jaeger: I also added Jaeger to visualize the telemetry data (the Makefile it's configured to run jaeger easily)

6- Tests : I believe that Tests are very important, in this project I'd like to add some unit test but for next iteration.

7- CI/CD using github action: for another iteration I'd add github action as CI/CD engine to push a continous integration coding culture
and mantain a stable source code.

8- aplied kubernetes: I created the templates for 4 deploy configuration k8s file, I decided to use deployment object because 
is very flexible and easy to mantain and follow.

9- run system in AWS : I was thinking to use EKS cluster in AWS in order to create the infraestructure to run my k8s components. 
For this  need to configure the cluster, add an configure a node group, set up aws cli and the kuctl utility and then run the deployments.
I can also use loadbalancer in the future if I need them. 
