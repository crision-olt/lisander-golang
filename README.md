# lisander-golang-backend
Lisander golang backend

    /register       POST -- Register a new user.
    /login          POST -- Get token.

    /user           GET -- get info profile.
    /user           PUT -- change info profile.
    /user/avatar    POST -- upload a avatar for user.
    /user/avatar    GET -- get avatar of user.
    /user/banner    POST -- upload a banner for user.
    /user/banner    GET -- get banner of user.

    /toots          POST -- create a new toot.
    /toots          GET -- read toots.
    /toots          DELETE -- delete toot.

    /relation       POST -- create a relation.
    /relation       DELETE -- delete a relation.
    /relation       GET -- check relation.

    /users          GET -- list users.
    /users/toots    GET -- list toots from users.


DOCKER
    
    HELPERS
        docker ps **show running dockers**

    MONGO
        docker pull mongo:latest **get mongo image for docker**
        docker run -d -p 27017:27017 -v ~/mongodb:/data/db --name lisander mongo:latest **Run docker**
        docker exec -it lisander bash  **Entry bash mongo**
    
    GOLANG
        docker run -p 8080:8080 lisander-golang-backend . **Start golang docker**
        docker exec -it practical_aryabhata bash **Entry bash golang**
