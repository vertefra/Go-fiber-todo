# todo-vue-client

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


###Orientation in Vuejs
App.vue is the main file that will hold the configuration of the application.
Navigate to the main.js file and this file is what mounts .mount('#app') to the index.html in the public folder. index.html has an id tag with app that that is where our application connects to the face of the web to manipulate the DOM directly.
Do note in Vue CLI created projects we do not touch the index.js in public folder because we have template scripts in our .vue files that will handle the html.

One of the things I like most about Vue is that it is component focused. Everything is modular and packed up into different files that contain only information relative to it.

In naming component files in Vue we use PascalCase.
