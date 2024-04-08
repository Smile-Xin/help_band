<template>
  <div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-5" elevation="1" color="" dark>{{ taskInfo.title }}</v-alert>
    <v-col>
      <v-card-text class="d-flex align-center">
        <div class=" align-center">
          发布日期：
          <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
          <span>{{ taskInfo.CreatedAt | dateformat('YYYY-MM-DD') }}</span>
        </div>
      </v-card-text>
      <v-card-text class="d-flex align-center">
        <div class="align-center">
          发布者：
          <v-icon class="mr-1" small>{{ 'mdi-account-question' }}</v-icon>
          <span>{{ taskInfo.demander_name }}</span>
        </div>
      </v-card-text>
      <v-subheader>简介：{{ taskInfo.briefing }}</v-subheader>

    </v-col>
    <v-card-text class="d-flex align-center">
      <div class="align-center">
        <p>详细内容：</p>
      </div>
    </v-card-text>
    <div v-html="taskInfo.content" class="content ma-5 pa-3 text-justify"></div>
    <v-divider class="ma-5"></v-divider>

    <v-card-text id="accept" class="d-flex">
      <div class="align-center">
        <el-button type="primary" v-if="this.taskInfo.status == 0" @click="accept()">接受任务<i class="el-icon-s-claim el-icon--right"></i></el-button>
      </div>
    </v-card-text>

  </div>
</template>
<script>
export default {
  // 由路由传递过来的任务id
  props: ['taskId'],
  data() {
    return {
      taskInfo: {},
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
    this.getTaskInfo()
    this.headers = {
      // 登陆后上传的用户信息
      user_name: window.sessionStorage.getItem('user_name'),
      user_id: window.sessionStorage.getItem('user_id'),
    }
  },
  watch: {
  },
  methods: {
    // 接受任务
    accept() {
      if (!this.headers.user_id) {
        return this.$message.error("请先登录");
      }
      if (this.headers.user_id == this.taskInfo.demander_id) {
        return this.$message.error("不能接受自己发布的任务");
      }
      this.$confirm.confirm("是否接受此任务", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(async () => {
        // 接受任务
        console.log("accept")
        this.taskInfo.status = 1;
        this.taskInfo.recipient_id = parseInt(this.headers.user_id);
        this.taskInfo.recipient_name = this.headers.user_name;
        console.log(this.taskInfo)
        const res = await this.$http.post(`task/EditTask`, this.taskInfo);
        if (res.data.state != 200) {
          return this.$message.error(res.data.message);
        } else {
          this.$message({
            type: "success",
            message: "成功接受!",
          });
        }
        this.$router.push({ path: "/" });
      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消",
        });
      });
    },
    // 查询任务
    async getTaskInfo() {
      const { data: res } = await this.$http.get(`task/QueryTaskByID/${this.taskId}`)
      this.taskInfo = res.data;
      console.log(this.taskInfo)
      window.sessionStorage.setItem('title', this.taskInfo.title)
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
#accept {
  display: flex;
  justify-content: center;
}

.content>>>div,
img,
span {
  width: auto;
  max-width: 100%;
}

.content>>>pre,
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
