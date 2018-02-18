import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import locale from "element-ui/lib/locale/lang/en";

Vue.config.productionTip = false;
Vue.use(ElementUI, { locale });

/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  render: h => h(App)
});
