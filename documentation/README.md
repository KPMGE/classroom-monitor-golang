# Api Documentation

### Setting up google credentials
This projects relies on the google apis. So, you need to go to google cloud platform and 
create an application there. After that, you can go ahead and download a __.json__ file with 
your app's credentials. Just name that file __credentials.json__ and place it at the project's
root directory.

### How to set up docker container
You can run this application with docker, first run you've got to create build your image,
go to the root directory and run the command: 

```bash
sudo docker build -t classroom-go-api .
```

Then, you can run a container using: 

```bash
sudo docker run -it -p3333:3333 --name classroom-go classroom-go-api
```

### How to generate api route docs.
In order to generate your api route docs, run the following command:

```bash
npx insomnia-documenter --config ./documentation.json
```

Then, you can run a server with your brand new generated docs with:

```bash
npx serve
```

Now it's as simple as opening a new tab on your favourite browser and accessing the link:
> http://localhost:3000
