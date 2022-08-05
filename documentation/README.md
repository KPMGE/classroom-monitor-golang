# Api Documentation

### Setting up google credentials
This projects relies on the google apis. So, you need to go to google cloud platform and 
create an application there. After that, you can go ahead and download a *.json* file with 
your app's credentials. Just name that file *credentials.json* and place it at the project's
root directory.

### How to set up docker container
First of all, you should have your google oauth credentials. In order to get them, 
go to the google's documentation. Once you've made the app, download your credentials and 
name your file *credentials.json*. Then, run the following command: 

```bash
sudo docker run -it \
  --name classroom-go \
  -p 3333:3333 \
  -v $(pwd)/credentials.json:/api/credentials.json \
  -p 5003:5003 \
  kevincarvalhodejesus/classroom-go-api
```

After a few seconds, you should have a container running and a message with a link for authentication 
should appear, click on it and allow the application to access your google classroom data.

### How to test the api? 
You can test the api directly calling the routes using some sort of a client. If you just want to test it 
locally, you can use something like *insomnia* for example.


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
