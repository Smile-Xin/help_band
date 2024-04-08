<template>
  <div>
    <div class="d-flex justify-center pa-3 ma-1 text-h4 font-weight-bold">{{ artInfo.title }}</div>
    <div class="d-flex justify-center align-center">
      <div class="d-flex mx-10 justify-center">
        <v-icon class="mr-1" color="indigo" small>
          {{
          'mdi-calendar-month'
          }}
        </v-icon>
        <span>{{ artInfo.UpdatedAt | dateformat('YYYY-MM-DD') }}</span>
      </div>
      <div class="d-flex mr-10 justify-center">
        <v-icon class="mr-1" color="pink" small>{{ 'mdi-comment' }}</v-icon>
        <span>{{ total }}</span>
      </div>
      <div class="d-flex mr-10 justify-center">
        <v-icon class="mr-1" color="green" small>{{ 'mdi-eye' }}</v-icon>
        <span>{{ artInfo.read_count }}</span>
      </div>
      <div class="d-flex mr-10 justify-center">
        <span>作者：{{ artInfo.User.user_name }}</span>
      </div>
    </div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert
      class="ma-4"
      elevation="1"
      color="indigo"
      dark
      border="left"
      outlined
    >{{ artInfo.desc }}</v-alert>

    <div v-html="artInfo.content" class="content ma-5 pa-3 text-justify"></div>

    <v-divider class="ma-5"></v-divider>

    <!-- // 评论区 -->
    <!-- <v-sheet class="ma-3 pa-3">
      <div>
        <v-list
          outlined
          class="ma-3 pa-3"
          v-for="item in commentList"
          :key="item.ID"
          v-show="item.status === 1"
        >
          <template>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>
                  {{ item.username }}
                  {{
                  item.CreatedAt | dateformat('YYYY-MM-DD')
                  }}
                </v-list-item-title>
                <v-list-item-subtitle class="mr-3">
                  {{
                  item.content
                  }}
                </v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </template>
        </v-list>
      </div>
      
      <div>
         <v-pagination
         v-model="queryParam.pagenum"
      :length="total"
    ></v-pagination>
      </div>

      <div>
        <template>
          <v-card flat>
            <v-alert v-if="!headers.user_name" class="ma-3" dense outlined type="error">你还未登录，请登录后留言</v-alert>
            <div v-if="headers.user_name">
              <v-textarea class="mx-3" outlined v-model="comment.content"></v-textarea>
              <v-btn class="ml-3 mb-1" dark color="indigo" small @click="pushComment()">确定</v-btn>
            </div>
          </v-card>
        </template>
      </div>
    </v-sheet> -->
  </div>
</template>
<script>
import MarkdownIt from "markdown-it";
export default {
  props: ['id'],
  data() {
    return {
      artInfo: {},
      commentList: [],
      comment: {
        content: ''
      },
      total: 0,
      headers: {
        user_name: '',
        user_id: 0
      },
      queryParam: {
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created() {
    this.getArtInfo()
    // this.getCommentList()
    this.headers = {
      // 登陆后上传的用户信息
      user_name: window.sessionStorage.getItem('user_name'),
      user_id: window.sessionStorage.getItem('user_id'),
    }
  },
  watch:{
    queryParam:{
      async handler(page){
        const { data: res } = await this.$http.get(`commentfront/${this.id}`, {
        params: {
          pagesize: page.pagesize,
          pagenum: page.pagenum
        }
      })
      console.log(res.total)
      this.commentList = res.data
      this.total = Math.ceil(res.total / page.pagesize)
      },
      deep:true,
      immediate:true
    }
  },
  methods: {
    // 查询文章
    async getArtInfo() {
      const { data: res } = await this.$http.get(`article/query`, {params:{id: this.id}})
      const mk = new MarkdownIt();
      console.log("res.data", res.data)
      this.artInfo = res.data;
      this.artInfo.content = mk.render(res.data.content);
      window.sessionStorage.setItem('title', this.artInfo.title)
    },
    // 获取评论
    async getCommentList() {
      const { data: res } = await this.$http.get(`commentfront/${this.id}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      console.log(res.total)
      this.commentList = res.data
      this.total = Math.ceil(res.total / this.queryParam.pagesize)
    },
    // 发送评论
    async pushComment() {
      const { data: res } = await this.$http.post('addcomment', {
        article_id: parseInt(this.id),
        content: this.comment.content,
        user_id: parseInt(this.headers.user_id),
        username: this.headers.username
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.$message.success('评论成功，待审核后显示')
      this.comment.content = '';
    }
  }
}
</script>
<style scoped>
.content >>> div,
img,
span {
  width: auto;
  max-width: 100%;
}

.content >>> pre,
code {
  margin: 10px;
  padding: 14px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
  background-color: rgba(27, 31, 35, 0.05);
  border-left-width: 0.5rem;
  border-left-style: solid;
  border-color: #303f9f;
}
</style>
