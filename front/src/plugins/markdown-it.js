import Vue from "vue";
import MarkdownIt from "markdown-it";

Vue.prototype.$mark = MarkdownIt

export default new MarkdownIt({})