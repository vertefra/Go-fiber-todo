import { createApp } from 'vue'

import App from './App.vue'
import LoginPage from './components/LoginPage.vue'

const app = createApp(App)

app.component('login-page', LoginPage)

app.mount('#app')
