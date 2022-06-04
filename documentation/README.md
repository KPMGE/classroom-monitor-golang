# Api Documentation

### Setting up google credentials
This projects relies on the google apis. So, you need to go to google cloud platform and 
create an application there. After that, you can go ahead and download a __.json__ file with 
your app's credentials. Just name that file __credentials.json__ and place it at the project's
root directory.

### Setting up .env file
For this project, it's required the use of environment variables. To do that, 
create a __.env__ file at the project's root directory and add the course id of 
your classroom as follows: 

```bash
COURSE_ID=<your course id>
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
