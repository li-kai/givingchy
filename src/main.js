import Vue from "vue";
import App from "./App.vue";
import router from "./router";
/* eslint-disable no-unused-vars */
import {
  Pagination,
  Dialog,
  Autocomplete,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Menu,
  Submenu,
  MenuItem,
  MenuItemGroup,
  Input,
  InputNumber,
  Radio,
  RadioGroup,
  RadioButton,
  Checkbox,
  CheckboxButton,
  CheckboxGroup,
  Switch,
  Select,
  Option,
  OptionGroup,
  Button,
  ButtonGroup,
  Table,
  TableColumn,
  DatePicker,
  TimeSelect,
  TimePicker,
  Popover,
  Tooltip,
  Breadcrumb,
  BreadcrumbItem,
  Form,
  FormItem,
  Tabs,
  TabPane,
  Tag,
  Tree,
  Alert,
  Slider,
  Icon,
  Row,
  Col,
  Upload,
  Progress,
  Badge,
  Card,
  Rate,
  Steps,
  Step,
  Carousel,
  CarouselItem,
  Collapse,
  CollapseItem,
  Cascader,
  ColorPicker,
  Transfer,
  Container,
  Header,
  Aside,
  Main,
  Footer,
  Loading,
  MessageBox,
  Message,
  Notification
} from "element-ui";
/* eslint-enable no-unused-vars */
import lang from "element-ui/lib/locale/lang/en";
import locale from "element-ui/lib/locale";

// configure language
locale.use(lang);
Vue.config.productionTip = false;

[
// Pagination,
// Dialog,
// Autocomplete,
// Dropdown,
// DropdownMenu,
// DropdownItem,
Menu,
// Submenu,
MenuItem,
// MenuItemGroup,
Input,
InputNumber,
// Radio,
// RadioGroup,
// RadioButton,
// Checkbox,
// CheckboxButton,
// CheckboxGroup,
// Switch,
// Select,
// Option,
// OptionGroup,
Button,
// ButtonGroup,
// Table,
// TableColumn,
DatePicker,
// TimeSelect,
// TimePicker,
Popover,
// Tooltip,
// Breadcrumb,
// BreadcrumbItem,
Form,
FormItem,
// Tabs,
// TabPane,
Tag,
// Tree,
// Alert,
Slider,
Icon,
Row,
Col,
// Upload,
Progress,
Badge,
Card,
// Rate,
// Steps,
// Step,
// Carousel,
// CarouselItem,
// Collapse,
// CollapseItem,
// Cascader,
// ColorPicker,
// Transfer,
Container,
Header,
// Aside,
Main,
// Footer,
].forEach((comp) => Vue.use(comp));

// Vue.use(Loading.directive);

// Vue.prototype.$loading = Loading.service;
// Vue.prototype.$msgbox = MessageBox;
// Vue.prototype.$alert = MessageBox.alert;
// Vue.prototype.$confirm = MessageBox.confirm;
// Vue.prototype.$prompt = MessageBox.prompt;
Vue.prototype.$notify = Notification;
// Vue.prototype.$message = Message;

/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  render: h => h(App)
});
